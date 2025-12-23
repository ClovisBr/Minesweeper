package event

type GameState int
type GameplayActionKind int

const (
	GameRunning GameState = iota
	GameWon
	GameLost
)

const (
	ActionNone GameplayActionKind = iota
	ActionReveal
	ActionToggleFlag
)

type GameplayAction struct {
	Time  Time
	Kind  GameplayActionKind
	Index int // index 1D (transport only)
}

type CellChange struct {
	Index int
	Mask  uint16
	Value uint16
}

type GameplayUpdate struct {
	Time  Time
	Cells []CellChange
	State GameState
}

func (GameplayAction) isAction() {}
