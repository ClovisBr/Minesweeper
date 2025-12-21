package event

import "github.com/ClovisBr/Minesweeper/engine"

type GameState int
type GameplayActionKind int

const (
	GameRunning GameState = iota
	GameWon
	GameLost

	ActionNone GameplayActionKind = iota
	ActionReveal
	ActionToggleFlag
)

type GameplayAction struct {
	Index engine.CellIndex   // index 1D de la cellule
	Time  int32              // timestamp monotonic (ms ou ticks)
	Kind  GameplayActionKind // reveal / flag
}

type GameplayUpdate struct {
	Time  int32
	Cells []CellChange
	State GameState
}

type CellChange struct {
	Index engine.CellIndex
	Mask  engine.Cell
	Value engine.Cell
}
