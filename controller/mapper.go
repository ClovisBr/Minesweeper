package controller

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/event"
)

func toEngineAction(a event.GameplayAction) engine.Action {
	switch a.Kind {
	case event.ActionReveal:
		return engine.Action{
			Kind:  engine.ActionReveal,
			Index: engine.CellIndex(a.Index),
		}
	case event.ActionToggleFlag:
		return engine.Action{
			Kind:  engine.ActionToggleFlag,
			Index: engine.CellIndex(a.Index),
		}
	default:
		panic("unknown GameplayActionKind")
	}
}

func toEventUpdate(u engine.Update, t event.Time) event.GameplayUpdate {
	cells := make([]event.CellChange, len(u.Cells))
	for i, c := range u.Cells {
		cells[i] = event.CellChange{
			Index: int(c.Index),
			Mask:  uint16(c.Mask),
			Value: uint16(c.Value),
		}
	}

	return event.GameplayUpdate{
		Time:  t,
		Cells: cells,
		State: event.GameState(u.State),
	}
}
