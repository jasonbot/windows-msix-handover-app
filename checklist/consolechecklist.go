package checklist

import (
	_ "embed"
	"fmt"
	"os"
	"weak"
)

var stepMarkers map[StepState]string

func init() {
	stepMarkers = map[StepState]string{
		StepPending:    "[ ]",
		StepInProgress: "[_]",
		StepError:      "[!]",
		StepSuccess:    "[O]",
		StepSkipped:    "[.]",
	}
}

type consolestep struct {
	ChecklistStep
	index  int
	runner weak.Pointer[consolerunner]
}

func (s *consolestep) SetMessage(m string) {
	s.ChecklistStep.Message = m
	if p := s.runner.Value(); p != nil {
		p.View()
	}
}

func (s *consolestep) SetState(state StepState) {
	s.ChecklistStep.State = state
	if p := s.runner.Value(); p != nil {
		p.View()
	}
}

func (s *consolestep) SetProgressPercentage(p int8) {
	s.ChecklistStep.Progress = p
	if p := s.runner.Value(); p != nil {
		p.View()
	}
}

type consolerunner struct {
	title        string
	steps        []*consolestep
	currentindex int
	done         bool
	m            chan any
}

func (c *consolerunner) AddStep(title string) RunStep {
	step := consolestep{
		ChecklistStep: ChecklistStep{
			Title:    title,
			Progress: -1,
		},
		index:  len(c.steps),
		runner: weak.Make(c),
	}
	c.steps = append(c.steps, &step)

	return &step
}

func (c *consolerunner) View() {
	lines := fmt.Sprintf("=== %v ===\n", c.title)

	for i, step := range c.steps {
		lines += fmt.Sprintf("%v %v: %v %v\n", stepMarkers[step.State], i+1, step.Title, step.Message)
		if step.Progress >= 0 {
			lines += fmt.Sprintf("    (%v%%)\n", step.Progress)
		}
	}

	fmt.Print("\033[2J")
	fmt.Print(lines)
}

func (c *consolerunner) Start() {
	c.View()
	<-c.m
}

func (c *consolerunner) Finish() {
	c.done = true
	c.View()
	os.Exit(0)
}

func NewConsoleChecklist(title string) ChecklistRunner {
	return &consolerunner{
		title: title,
		steps: []*consolestep{},
	}
}
