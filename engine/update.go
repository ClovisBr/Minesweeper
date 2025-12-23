package engine

// CellChange représente un changement partiel sur une cellule.
// Mask  : bits concernés
// Value : nouvelle valeur pour ces bits
type CellChange struct {
	Index CellIndex
	Mask  Cell
	Value Cell
}

type Update struct {
	Cells []CellChange
	State GameState
}
