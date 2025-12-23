package rules

import "github.com/ClovisBr/Minesweeper/engine"

func floodReveal(g *engine.Grid, start engine.CellIndex) []engine.CellChange {
	rows := g.Rows
	cols := g.Cols

	visited := make(map[engine.CellIndex]struct{})
	queue := []engine.CellIndex{start}

	var changes []engine.CellChange

	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]

		if _, ok := visited[idx]; ok {
			continue
		}
		visited[idx] = struct{}{}

		cell := g.CellAt(idx)
		if cell == nil || cell.Has(engine.FlagReveal) {
			continue
		}

		changes = append(changes, engine.CellChange{
			Index: idx,
			Mask:  engine.Cell(engine.FlagReveal),
			Value: engine.Cell(engine.FlagReveal),
		})

		if cell.GetNeighborCount() > 0 {
			continue
		}

		r := int(idx) / cols
		c := int(idx) % cols

		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}

				nr, nc := r+dr, c+dc
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}

				nidx := engine.CellIndex(nr*cols + nc)
				queue = append(queue, nidx)
			}
		}
	}

	return changes
}
