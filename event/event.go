package event

type Time uint32 // milliseconds since game start

type Action interface {
	isAction()
}

func (UIAction) isAction()       {}
func (GameplayAction) isAction() {}
