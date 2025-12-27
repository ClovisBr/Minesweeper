package engine

type virtualGrid struct {
	base    *Grid
	cells   []Cell
	touched map[CellIndex]struct{}
}

func NewVirtualGrid(g *Grid) *virtualGrid {
	cells := make([]Cell, len(g.Cells))
	copy(cells, g.Cells)

	return &virtualGrid{
		base:    g,
		cells:   cells,
		touched: make(map[CellIndex]struct{}),
	}
}

// Pas nécessaire mais une sécurité en plus.
// Potentiellement remplacer par un panic plus tard.
// On peut théoriquement le supprimer sans trop de soucis.
func (v *virtualGrid) CellAt(idx CellIndex) *Cell {
	if idx < 0 || int(idx) >= len(v.cells) {
		return nil
	}
	return &v.cells[idx]
}

func (v *virtualGrid) Neighbors(idx CellIndex, fn func(CellIndex)) {
	v.base.Neighbors(idx, fn)
}

func (v *virtualGrid) CountFlagsAround(idx CellIndex) int {
	count := 0
	v.Neighbors(idx, func(n CellIndex) {
		if c := v.CellAt(n); c != nil && c.Has(FlagFlag) {
			count++
		}
	})
	return count
}

// mutation contrôlée + traçage
func (v *virtualGrid) Set(idx CellIndex, flag CellFlag) {
	c := &v.cells[idx]
	before := *c
	c.Set(flag)
	if *c != before {
		v.touched[idx] = struct{}{}
	}
}

// point fixe
func (v *virtualGrid) Stabilize() {
	queue := []CellIndex{}
	visited := make(map[CellIndex]struct{})

	for i, c := range v.cells {
		if c.Has(FlagReveal) {
			queue = append(queue, CellIndex(i))
		}
	}

	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]

		if _, seen := visited[idx]; seen {
			continue
		}
		visited[idx] = struct{}{}

		cell := v.CellAt(idx)
		if cell == nil || !cell.Has(FlagReveal) {
			continue
		}

		if v.CountFlagsAround(idx) < int(cell.GetNeighborCount()) {
			continue
		}

		v.Neighbors(idx, func(n CellIndex) {
			c := v.CellAt(n)
			if c == nil || c.Has(FlagReveal) || c.Has(FlagFlag) {
				return
			}
			v.Set(n, FlagReveal)
			queue = append(queue, n)
		})
	}
}

// production finale du diff
func (v *virtualGrid) Changes() []CellChange {
	changes := make([]CellChange, 0, len(v.touched))

	for idx := range v.touched {
		before := v.base.Cells[idx]
		after := v.cells[idx]

		changes = append(changes, CellChange{
			Index: idx,
			Mask:  before ^ after,
			Value: after,
		})
	}

	return changes
}
