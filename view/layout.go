package view

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/event"
)

const (
	CellW = 3
	CellH = 1
)

type Layout struct {
	Rows int
	Cols int
}

func NewLayout(rows, cols int) *Layout {
	return &Layout{
		Rows: rows,
		Cols: cols,
	}
}

// ScreenToCell convertit des coordonnées écran (x,y)
// vers un index de cellule logique.
func (l *Layout) ScreenToCell(x, y int) (engine.CellIndex, bool) {
	if x < 0 || y < 0 {
		return 0, false
	}

	col := x / CellW
	row := y / CellH

	if col < 0 || col >= l.Cols || row < 0 || row >= l.Rows {
		return 0, false
	}

	return engine.CellIndex(row*l.Cols + col), true
}

func applyCellDiff(cell *engine.Cell, d event.CellChange) {
	*cell &^= d.Mask          // clear bits
	*cell |= d.Value & d.Mask // set new bits
}
