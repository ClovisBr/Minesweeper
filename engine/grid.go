package engine

import "errors"
import "fmt"
import "github.com/ClovisBr/Minesweeper/config"

type CellIndices []int

type Grid struct {
	Rows  int
	Cols  int
	Cells []Cell
}

func NewGrid(cfg config.Config) *Grid {
	return &Grid{
		Rows:  cfg.Rows,
		Cols:  cfg.Cols,
		Cells: make([]Cell, cfg.Rows*cfg.Cols),
	}
}

func (g *Grid) index(r, c int) int {
	return r*g.Cols + c
}

func (g *Grid) Cell(r, c int) *Cell {
	return &g.Cells[g.index(r, c)]
}

func (g *Grid) PlaceMines(indices CellIndices) error {
	total := len(g.Cells)

	if len(indices) > total {
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
