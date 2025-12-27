package generator

import "github.com/ClovisBr/Minesweeper/engine"

func ComputeNeighbors(g *engine.Grid) error {
	for idx, cell := range g.Cells {
		if !cell.Has(engine.FlagMine) {
			continue
		}

		g.Neighbors(engine.CellIndex(idx), func(n engine.CellIndex) {
			neighbor := g.CellAt(n)
			if neighbor == nil || neighbor.Has(engine.FlagMine) {
				return
			}

			count := neighbor.GetNeighborCount()
			_ = neighbor.SetNeighborCount(count + 1)
		})
	}
	return nil
}
