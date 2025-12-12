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
		if i == 2 {
			item.SetState(checklist.StepError)
		} else if i == 3 {
			item.SetState(checklist.StepSkipped)
		} else {
			item.SetMessage("Hi there")
			for progress := range 100 {
				steps[i].SetProgressPercentage(int8(progress))
				time.Sleep(3 * time.Millisecond)
			}
			item.SetState(checklist.StepSuccess)
		}
		time.Sleep(500 * time.Millisecond)
		log.Println("Step", i+1)
	}
	log.Println("---")
}

func main() {
	cl := checklist.NewGioChecklist("Hello, world")
	steps := []checklist.RunStep{
		cl.AddStep("Step 1. Ham"),
		cl.AddStep("Step 2. Bacon"),
		cl.AddStep("Step 3. Eggs"),
		cl.AddStep("Step 4. Bread"),
		cl.AddStep("Step 5. Milk"),
	}

	go handleSteps(steps, cl)
	cl.Start()
}
