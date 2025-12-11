package checklist

import (
	"image/color"
	"log"
	"os"
	"sync"
	"weak"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

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
	steps        []giostep
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

func (g *giorunner) run() error {
	theme := material.NewTheme()
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
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
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
	g.steps = append(g.steps, step)

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
}

func NewGioChecklist(title string) ChecklistRunner {
	return &giorunner{
		title: title,
		steps: []giostep{},
	}
}
