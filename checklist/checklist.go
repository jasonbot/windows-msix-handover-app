package checklist

type StepState string

const (
	StepInProgress StepState = "InProgress"
	StepError      StepState = "Error"
	StepSuccess    StepState = "Success"
	StepSkipped    StepState = "Skipped"
)

type RunStep interface {
	SetStepName(string)
	SetStepMessage(string)
	SetState(StepState)
	SetProgressPercentage(uint8)
}
