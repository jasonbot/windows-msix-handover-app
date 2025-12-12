package main

import (
	"log"
	"time"

	"github.com/jasonbot/windows-msix-handover-app/checklist"
)

func handleSteps(steps []checklist.RunStep, cl checklist.ChecklistRunner) {
	defer cl.Finish()
	for i := range steps {
		item := steps[i]
		item.SetState(checklist.StepInProgress)
		item.SetMessage("Hi there")
		for progress := range 100 {
			steps[i].SetProgressPercentage(int8(progress))
			time.Sleep(3 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
		item.SetState(checklist.StepSuccess)
		log.Println("Step", i+1)
	}
	log.Println("---")
}

func main() {
	cl := checklist.NewConsoleChecklist("Hello, world")
	steps := []checklist.RunStep{
		cl.AddStep("Step 1. Ham"),
		cl.AddStep("Step 2. Bacon"),
		cl.AddStep("Step 3. Bacon"),
	}

	go handleSteps(steps, cl)
	cl.Start()
}
