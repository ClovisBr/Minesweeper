package rules

import "github.com/ClovisBr/Minesweeper/engine"

type Punitive struct{}

func (Punitive) Reveal(g *engine.Grid, idx engine.CellIndex) []engine.CellChange {
	cell := g.CellAt(idx)
	if cell == nil {
		return nil
	}

	// déjà révélée ou flagguée (flag définitif)
	if cell.Has(engine.FlagReveal) || cell.Has(engine.FlagFlag) {
		return nil
	}

	changes := []engine.CellChange{{
		Index: idx,
		Mask:  engine.Cell(engine.FlagReveal),
		Value: engine.Cell(engine.FlagReveal),
	}}

	// mine révélée -> GameLost
	if cell.Has(engine.FlagMine) {
		return changes
	}

	if cell.GetNeighborCount() == 0 {
		changes = append(changes, floodReveal(g, idx)...)
	}

	return changes
}
func (Punitive) ToggleFlag(g *engine.Grid, idx engine.CellIndex) []engine.CellChange {
	cell := g.CellAt(idx)
	if cell == nil {
		return nil
	}

	// interdit si déjà révélé ou déjà flaggé
	if cell.Has(engine.FlagReveal) || cell.Has(engine.FlagFlag) {
		return nil
	}

	changes := []engine.CellChange{{
		Index: idx,
		Mask:  engine.Cell(engine.FlagFlag),
		Value: engine.Cell(engine.FlagFlag),
	}}

	// si c’est une mine → flag = reveal implicite → défaite
	if cell.Has(engine.FlagMine) {
		changes = append(changes, engine.CellChange{
			Index: idx,
			Mask:  engine.Cell(engine.FlagReveal),
			Value: engine.Cell(engine.FlagReveal),
		})
	}

	return changes
}
