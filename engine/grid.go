package engine

import (
	"errors"
	"fmt"

	"github.com/ClovisBr/Minesweeper/config"
)

type CellIndex int

type Grid struct {
	Rows  int
	Cols  int
	Cells []Cell
}

func NewGrid(cfg config.Config) *Grid {
	return &Grid{
		Rows:  cfg.Grid.Rows,
		Cols:  cfg.Grid.Cols,
		Cells: make([]Cell, cfg.Grid.Rows*cfg.Grid.Cols),
	}
}

func (g *Grid) index(r, c int) CellIndex {
	return CellIndex(r*g.Cols + c)
}

func (g *Grid) Cell(r, c int) *Cell {
	return &g.Cells[g.index(r, c)]
}

func (g *Grid) CellAt(idx CellIndex) *Cell {
	if idx < 0 || int(idx) >= len(g.Cells) {
		return nil
	}
	return &g.Cells[idx]
}

func (g *Grid) PlaceMines(indices []CellIndex) error {
	total := CellIndex(len(g.Cells))

	if CellIndex(len(indices)) > total {
		return errors.New("more mines than cells")
	}

	// reset
	for i := range g.Cells {
		g.Cells[i].Clear(FlagMine)
	}

	for _, idx := range indices {
		if idx < 0 || idx >= total {
			return errors.New("invalid cell index")
		}
		g.Cells[idx].Set(FlagMine)
	}

	return nil
}

func (g *Grid) PrintCells(binary ...bool) {
	showBinary := len(binary) > 0 && binary[0]

	for i, c := range g.Cells {
		if showBinary {
			fmt.Printf("%016b ", uint16(c))
		} else {
			if c.Has(FlagMine) {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}

		if (i+1)%g.Cols == 0 {
			fmt.Println()
		}
	}
}
