package controller

type Intent int

type Action struct {
	Intent Intent
}

const (
	IntentNone Intent = iota

	IntentUp
	IntentDown
	IntentLeft
	IntentRight

	IntentReveal
	IntentToggleFlag
	IntentQuit
)
