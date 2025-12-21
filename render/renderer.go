package render

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/gdamore/tcell/v2"
)

type Renderer struct {
	screen tcell.Screen
}

const (
	CellW = 3
	CellH = 1
)

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

			ch := ' '
			style := StyleHidden

			switch {
			case cell.Has(engine.FlagFlag):
				ch = 'F'
				style = StyleFlag

			case !cell.Has(engine.FlagReveal):
				ch = '~'
				style = StyleHidden

			case cell.Has(engine.FlagMine):
				ch = '*'
				style = StyleMine

			default:
				n := cell.GetNeighborCount()
				if n > 0 {
					ch = rune('0' + n)
					style = numberStyle(n)
				} else {
					ch = ' '
					style = StyleRevealed
				}
			}

			// curseur logique → sur toute la cellule 3×1
			if r0 == cursorR && c0 == cursorC {
				style = style.Reverse(true)
			}

			x := c0 * CellW
			y := r0

			// padding + contenu centré
			r.screen.SetContent(x, y, ' ', nil, style)
			r.screen.SetContent(x+1, y, ch, nil, style)
			r.screen.SetContent(x+2, y, ' ', nil, style)
		}
	}

	r.screen.Show()
}
