package rules

import "github.com/ClovisBr/Minesweeper/engine"

type Punitive struct{}

func (Punitive) Reveal(g *engine.Grid, idx engine.CellIndex) []engine.CellChange {
	v := engine.NewVirtualGrid(g)

	c := v.CellAt(idx)
	if c == nil || c.Has(engine.FlagReveal) || c.Has(engine.FlagFlag) {
		return nil
	}

	v.Set(idx, engine.FlagReveal)
	v.Stabilize()

	return v.Changes()
}

func (Punitive) ToggleFlag(g *engine.Grid, idx engine.CellIndex) []engine.CellChange {
	v := engine.NewVirtualGrid(g)

	c := v.CellAt(idx)
	if c == nil || c.Has(engine.FlagReveal) || c.Has(engine.FlagFlag) {
		return nil
	}

	v.Set(idx, engine.FlagFlag)
	v.Stabilize()

	return v.Changes()
}
