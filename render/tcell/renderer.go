package render

import (
	"github.com/ClovisBr/Minesweeper/engine"
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

	return &Renderer{screen: s}, nil
}

func (r *Renderer) Close() {
	r.screen.Fini()
}

func (r *Renderer) DrawGrid(g *engine.Grid, cursorR, cursorC int) {
	r.screen.Clear()

	for r0 := 0; r0 < g.Rows; r0++ {
		for c0 := 0; c0 < g.Cols; c0++ {
			cell := g.Cell(r0, c0)

			ch := '.'
			style := StyleHidden

			if cell.Has(engine.FlagReveal) {
				style = StyleRevealed
				if cell.Has(engine.FlagMine) {
					ch = '*'
					style = StyleMine
				} else if cell.GetNeighborCount() > 0 {
					ch = rune('0' + cell.GetNeighborCount())
				} else {
					ch = ' '
				}
			}

			if r0 == cursorR && c0 == cursorC {
				style = style.Reverse(true)
			}

			r.screen.SetContent(c0, r0, ch, nil, style)
		}
	}

	r.screen.Show()
}
