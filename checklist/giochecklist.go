package checklist

import (
	_ "embed"
	"image"
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
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/inkeliz/giosvg"
	"golang.org/x/sys/windows/registry"
)

//go:embed wordmark.svg
var wordmark []byte

type iconWithColor struct {
	icon  *giosvg.Icon
	color color.NRGBA
}

var wordMark *giosvg.Icon
var StatusSVGs map[StepState]*iconWithColor

func init() {
	if vector, err := giosvg.NewVector(wordmark); err == nil {
		wordMark = giosvg.NewIcon(vector)
	}

	StatusSVGs = map[StepState]*iconWithColor{}

	for _, item := range []struct {
		status StepState
		svg    string
		color  color.NRGBA
	}{
		{
			status: StepSuccess,
			svg:    `<?xml version="1.0" encoding="UTF-8"?><svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor"><path d="M7 12.5L10 15.5L17 8.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path></svg>`,
			color:  color.NRGBA{R: 68, G: 131, B: 97, A: 255},
		},
		{
			status: StepInProgress,
			svg:    `<?xml version="1.0" encoding="UTF-8"?><svg width="24px" height="24px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor" stroke-width="1.5"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 1.25C6.06294 1.25 1.25 6.06294 1.25 12C1.25 17.9371 6.06294 22.75 12 22.75C17.9371 22.75 22.75 17.9371 22.75 12C22.75 6.06294 17.9371 1.25 12 1.25ZM13.0303 7.96967L16.5303 11.4697C16.8232 11.7626 16.8232 12.2374 16.5303 12.5303L13.0303 16.0303C12.7374 16.3232 12.2626 16.3232 11.9697 16.0303C11.6768 15.7374 11.6768 15.2626 11.9697 14.9697L14.1893 12.75H8C7.58579 12.75 7.25 12.4142 7.25 12C7.25 11.5858 7.58579 11.25 8 11.25H14.1893L11.9697 9.03033C11.6768 8.73744 11.6768 8.26256 11.9697 7.96967C12.2626 7.67678 12.7374 7.67678 13.0303 7.96967Z" fill="currentColor"></path></svg>`,
			color:  color.NRGBA{R: 16, G: 95, B: 173, A: 255},
		},
		{
			status: StepError,
			svg:    `<?xml version="1.0" encoding="UTF-8"?><svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor"><path d="M9.17218 14.8284L12.0006 12M14.829 9.17157L12.0006 12M12.0006 12L9.17218 9.17157M12.0006 12L14.829 14.8284" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path></svg>`,
			color:  color.NRGBA{R: 205, G: 60, B: 58, A: 255},
		},
		{
			status: StepSkipped,
			svg:    `<?xml version="1.0" encoding="UTF-8"?><svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor"><path d="M12 8V16M12 16L15.5 12.5M12 16L8.5 12.5" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path><path d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path></svg>`,
			color:  color.NRGBA{R: 115, G: 114, B: 110, A: 255},
		},
		{
			status: StepPending,
			svg:    `<?xml version="1.0" encoding="UTF-8"?><svg width="24px" height="24px" stroke-width="1.5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" color="currentColor"><path d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path></svg>`,
			color:  color.NRGBA{R: 115, G: 114, B: 110, A: 255},
		},
	} {
		if vector, err := giosvg.NewVector([]byte(item.svg)); err == nil {
			icon := giosvg.NewIcon(vector)
			StatusSVGs[item.status] = &iconWithColor{icon, item.color}
		}
	}
}

func isDarkMode() bool {
	personalizeKey := `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`

	if key, err := registry.OpenKey(registry.CURRENT_USER, personalizeKey, registry.ALL_ACCESS); err == nil {
		defer key.Close()
		if val, _, err := key.GetIntegerValue("AppsUseLightTheme"); err == nil {
			return val == 0
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

func (s *giostep) SetProgressPercentage(p int8) {
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
	button       widget.Clickable
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
	retVal := []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Define an large label with an appropriate text:
			title := material.Label(theme, 14, g.title)
			// Change the position of the label.
			title.Alignment = text.Middle
			title.Font.Weight = font.Bold

			return layout.UniformInset(unit.Dp(24)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Middle,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if wordMark != nil {
							gtx.Constraints.Min = image.Pt(1100/8, 317/8)
							gtx.Constraints.Max = image.Pt(1100/4, 317/4)
							return wordMark.Layout(gtx)
						}
						return layout.Spacer{}.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Top: 12}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return title.Layout(gtx)
						})
					}))
			})
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
			layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Top: 4}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						label := material.Label(theme, 11, step.Message)
						return label.Layout(gtx)
					})
				}),
		}

		if step.Progress >= 0 {
			additionalSteps = append(additionalSteps, layout.Rigid(
				func(gtx layout.Context) layout.Dimensions {
					progressPercentage := float32(step.Progress) / float32(100)
					pg := material.ProgressBar(theme, progressPercentage)
					return layout.Inset{Top: 4, Bottom: 4}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return pg.Layout(gtx)
					})
				}))
		}

		x := func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Baseline,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if StatusSVGs[step.State] != nil {
						return layout.UniformInset(4).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							paint.ColorOp{Color: StatusSVGs[step.State].color}.Add(gtx.Ops)
							gtx.Constraints.Max = image.Pt(28, 28)
							return StatusSVGs[step.State].icon.Layout(gtx)
						})
					}
					return material.Label(theme, 18, string(step.State)).Layout(gtx)
				}),
				layout.Flexed(
					1.0,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Bottom: 12}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx, additionalSteps...)
						})
					}),
			)
		}

		r := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return x(gtx)
		})

		retVal = append(retVal, r)
	}

	retVal = append(retVal,
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			labelContent := "Running..."
			if g.done {
				labelContent = "Done. You can close this window."
			}
			l := material.Label(theme, 12, labelContent)
			l.Alignment = text.Middle
			return l.Layout(gtx)
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
			if g.button.Clicked(gtx) {
				log.Println("OK")
			}

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
			Title:    title,
			Progress: -1,
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
