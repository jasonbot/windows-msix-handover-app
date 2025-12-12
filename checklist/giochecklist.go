package checklist

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"sync"
	"weak"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"golang.org/x/sys/windows/registry"
)

func isDarkMode() bool {
	personalizeKey := `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`

	if key, err := registry.OpenKey(registry.CURRENT_USER, personalizeKey, registry.ALL_ACCESS); err == nil {
		defer key.Close()
		if val, _, err := key.GetIntegerValue("AppsUseLightTheme"); err == nil {
			return err != nil && val != 0
		} else {
			return false
		}
	}

	return true
}

type giostep struct {
	ChecklistStep
	index  int
	runner weak.Pointer[giorunner]
}

func (s *giostep) SetMessage(m string) {
	s.ChecklistStep.Message = m
	if r := s.runner.Value(); r != nil && r.window != nil {
		r.window.Invalidate()
	}
}

func (s *giostep) SetState(state StepState) {
	s.ChecklistStep.State = state
	if r := s.runner.Value(); r != nil {
		if state != StepPending {
			for i := range s.index {
				if r.steps[i].State == StepPending {
					r.steps[i].State = StepSkipped
				}
			}
		}

		if r.window != nil {
			r.window.Invalidate()
		}
	}
}

func (s *giostep) SetProgressPercentage(p *int8) {
	s.ChecklistStep.Progress = p
	if r := s.runner.Value(); r != nil && r.window != nil {
		r.window.Invalidate()
	}
}

type giorunner struct {
	title        string
	steps        []*giostep
	currentindex int
	window       *app.Window
	m            sync.Mutex
	done         bool
}

func (g *giorunner) main() {
	go func() {
		g.window = new(app.Window)
		g.window.Option(
			app.Title(g.title),
			app.MinSize(800, 600),
			app.MaxSize(800, 600),
		)
		err := g.run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func (g *giorunner) stepListWidgets(theme *material.Theme) []layout.FlexChild {
	// Define an large label with an appropriate text:
	title := material.Label(theme, 14, g.title)
	// Change the position of the label.
	title.Alignment = text.Middle
	title.Font.Weight = font.Bold
	retVal := []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return title.Layout(gtx)
		}),
	}

	for _, step := range g.steps {
		additionalSteps := []layout.FlexChild{
			layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					label := material.Label(theme, 12, step.Title)
					if step.State == StepInProgress {
						label.Font.Weight = font.Bold
					}
					return label.Layout(gtx)
				}),
		}
		if step.Message != "" {
			fmt.Println("Adding progress bar")
			additionalSteps = append(additionalSteps, layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					label := material.Label(theme, 11, step.Message)
					return label.Layout(gtx)
				}))
		}

		if step.Progress != nil {
			fmt.Println("Adding progress bar")
			additionalSteps = append(additionalSteps, layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					progressPercentage := float32(*step.Progress) / float32(100)
					fmt.Println("FFF", progressPercentage)
					pg := material.ProgressBar(theme, progressPercentage)
					return pg.Layout(gtx)
				}))
		}

		retVal = append(retVal, additionalSteps...)
	}

	retVal = append(retVal,
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			text := "Running..."
			if g.done {
				text = "Done"
			}
			return material.Label(theme, 12, text).Layout(gtx)
		}))

	return retVal
}

func (g *giorunner) run() error {
	theme := material.NewTheme()
	if isDarkMode() {
		theme.Palette = material.Palette{
			Bg:         color.NRGBA{R: 0x04, G: 0x04, B: 0x04, A: 0xFF},
			Fg:         color.NRGBA{R: 0xFF, G: 0xFE, B: 0xFC, A: 0xFF},
			ContrastBg: color.NRGBA{R: 0xEF, G: 0xF3, B: 0xF5, A: 0x88},
			ContrastFg: color.NRGBA{R: 0x23, G: 0x83, B: 0xE2, A: 0xFF},
		}
	} else {
		theme.Palette = material.Palette{
			Bg:         color.NRGBA{R: 0xFF, G: 0xFE, B: 0xFC, A: 0xFF},
			Fg:         color.NRGBA{R: 0x04, G: 0x04, B: 0x04, A: 0xFF},
			ContrastBg: color.NRGBA{R: 0xEF, G: 0xF3, B: 0xF5, A: 0x88},
			ContrastFg: color.NRGBA{R: 0xFF, G: 0xFE, B: 0xFC, A: 0xAA},
		}
	}
	var ops op.Ops
	for {
		switch e := g.window.Event().(type) {
		case app.DestroyEvent:
			done := g.done
			if done {
				os.Exit(0)
				return e.Err
			}
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			paint.ColorOp{Color: theme.Palette.Bg}.Add(&ops)
			paint.PaintOp{}.Add(&ops)

			layout.UniformInset(unit.Dp(24)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					g.stepListWidgets(theme)...,
				)
			})

			e.Frame(gtx.Ops)
		}
	}
}

func (g *giorunner) AddStep(title string) RunStep {
	step := giostep{
		ChecklistStep: ChecklistStep{
			Title: title,
		},
		index:  len(g.steps),
		runner: weak.Make(g),
	}
	g.steps = append(g.steps, &step)

	if g.window != nil {
		g.window.Invalidate()
	}
	return &step
}

func (g *giorunner) Start() {
	g.main()
}

func (g *giorunner) Finish() {
	g.m.Lock()
	defer g.m.Unlock()
	g.done = true
	for i := range g.steps {
		if g.steps[i].State == StepPending {
			g.steps[i].State = StepSkipped
		}
	}
	if g.window != nil {
		g.window.Invalidate()
	}
}

func NewGioChecklist(title string) ChecklistRunner {
	return &giorunner{
		title: title,
		steps: []*giostep{},
	}
}
