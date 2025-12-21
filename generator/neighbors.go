package generator

import "github.com/ClovisBr/Minesweeper/engine"

func ComputeNeighbors(g *engine.Grid) error {
	rows := g.Rows
	cols := g.Cols

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cell := g.Cell(r, c)

			if !cell.Has(engine.FlagMine) {
				continue
			}

			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}

					nr := r + dr
					nc := c + dc

					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}

					neighbor := g.Cell(nr, nc)

					if neighbor.Has(engine.FlagMine) {
						continue
					}

					count := neighbor.GetNeighborCount()
					_ = neighbor.SetNeighborCount(count + 1)
				}
			}
		}
	}

	return nil
}
