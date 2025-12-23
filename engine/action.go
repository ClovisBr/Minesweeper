package engine

type ActionKind int

const (
	ActionReveal ActionKind = iota
	ActionToggleFlag
)

type Action struct {
	Kind  ActionKind
	Index CellIndex
}
