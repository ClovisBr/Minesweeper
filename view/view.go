package view

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/event"
)

type View struct {
	Layout *Layout
	Grid   []engine.Cell

	Cursor engine.CellIndex
	Hover  *engine.CellIndex
}

func NewView(layout *Layout, initial []engine.Cell) *View {
	cells := make([]engine.Cell, len(initial))
	copy(cells, initial)

	return &View{
		Layout: layout,
		Grid:   cells,
		Cursor: 0,
		Hover:  nil,
	}
}

func (v *View) clampCursor() {
	max := engine.CellIndex(len(v.Grid) - 1)
	if v.Cursor < 0 {
		v.Cursor = 0
	}
	if v.Cursor > max {
		v.Cursor = max
	}
}

func (v *View) ApplyUI(a event.UIAction) {
	switch a.Kind {

	case event.UIMoveCursorUp:
		v.Cursor -= engine.CellIndex(v.Layout.Cols)
		v.Hover = nil

	case event.UIMoveCursorDown:
		v.Cursor += engine.CellIndex(v.Layout.Cols)
		v.Hover = nil

	case event.UIMoveCursorLeft:
		v.Cursor--
		v.Hover = nil

	case event.UIMoveCursorRight:
		v.Cursor++
		v.Hover = nil

	case event.UIHover:
		if idx, ok := v.Layout.ScreenToCell(a.X, a.Y); ok {
			v.Hover = &idx
		} else {
			v.Hover = nil
		}
	}

	v.clampCursor()
}

func (v *View) ApplyUpdate(u event.GameplayUpdate) {
	for _, d := range u.Cells {
		if d.Index < 0 || d.Index >= len(v.Grid) {
			continue
		}
		applyCellDiff(&v.Grid[d.Index], d)
	}
}
