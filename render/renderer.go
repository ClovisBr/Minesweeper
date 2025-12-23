package render

import (
	"github.com/ClovisBr/Minesweeper/view"
	"github.com/gdamore/tcell/v2"
)

type Renderer struct {
	screen tcell.Screen
}

func New() (*Renderer, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}

	s.Clear()
	return &Renderer{screen: s}, nil
}

func (r *Renderer) Close() {
	r.screen.Fini()
}

func (r *Renderer) Draw(v *view.View) {
	r.screen.Clear()

	cols := v.Layout.Cols

	for i, cell := range v.Grid {
		row := i / cols
		col := i % cols

		ch, style := cellStyle(cell)

		if int(v.Cursor) == i {
			style = style.Reverse(true)
		}

		if v.Hover != nil && int(*v.Hover) == i {
			style = style.Reverse(true)
		}

		x := col * view.CellW
		y := row * view.CellH

		r.drawCell(x, y, ch, style)
	}

	r.screen.Show()
}

func (r *Renderer) drawCell(x, y int, ch rune, style tcell.Style) {
	r.screen.SetContent(x, y, ' ', nil, style)
	r.screen.SetContent(x+1, y, ch, nil, style)
	r.screen.SetContent(x+2, y, ' ', nil, style)
}

func (r *Renderer) Screen() tcell.Screen {
	return r.screen
}
