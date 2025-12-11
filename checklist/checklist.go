package checklist

type StepState string

const (
	StepPending    StepState = ""
	StepInProgress StepState = "InProgress"
	StepError      StepState = "Error"
	StepSuccess    StepState = "Success"
	StepSkipped    StepState = "Skipped"
)

type ChecklistStep struct {
	Title    string
	Message  string
	State    StepState
	Progress *int8
}

type RunStep interface {
	SetMessage(string)
	SetState(StepState)
	SetProgressPercentage(*int8)
}

type ChecklistRunner interface {
	AddStep(string) RunStep
	Start()
	Finish()
}
