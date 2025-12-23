package controller

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/event"
	"github.com/ClovisBr/Minesweeper/view"
)

type Controller struct {
	game   *engine.Game
	view   *view.View
	cursor *Cursor
	timer  *Timer
}

func NewController(
	game *engine.Game,
	view *view.View,
	rows, cols int,
) *Controller {
	return &Controller{
		game:   game,
		view:   view,
		cursor: NewCursor(rows, cols),
		timer:  NewTimer(),
	}
}

func (c *Controller) HandleAction(i Intent) *event.GameplayUpdate {
	action := c.ApplyIntent(i)
	if action == nil {
		return nil
	}

	switch a := action.(type) {

	// ---------- UI ----------
	case event.UIAction:
		c.view.ApplyUI(a)
		return nil

	// ---------- Gameplay ----------
	case event.GameplayAction:
		engineUpdate := c.game.Apply(toEngineAction(a))
		eventUpdate := toEventUpdate(engineUpdate, a.Time)
		c.view.ApplyUpdate(eventUpdate)
		return &eventUpdate
	}

	return nil
}

func (c *Controller) ApplyIntent(i Intent) event.Action {
	now := c.timer.Now()

	switch i {

	// ---------- UI (clavier) ----------
	case IntentUp:
		if c.cursor.MoveUp() {
			return event.UIAction{Time: now, Kind: event.UIMoveCursorUp}
		}

	case IntentDown:
		if c.cursor.MoveDown() {
			return event.UIAction{Time: now, Kind: event.UIMoveCursorDown}
		}

	case IntentLeft:
		if c.cursor.MoveLeft() {
			return event.UIAction{Time: now, Kind: event.UIMoveCursorLeft}
		}

	case IntentRight:
		if c.cursor.MoveRight() {
			return event.UIAction{Time: now, Kind: event.UIMoveCursorRight}
		}

	// ---------- Gameplay ----------
	case IntentReveal:
		return event.GameplayAction{
			Time:  now,
			Kind:  event.ActionReveal,
			Index: int(c.view.Cursor),
		}

	case IntentToggleFlag:
		return event.GameplayAction{
			Time:  now,
			Kind:  event.ActionToggleFlag,
			Index: int(c.view.Cursor),
		}
	}

	return nil
}

func (c *Controller) HandleMouseMove(x, y int) {
	now := c.timer.Now()

	c.view.ApplyUI(event.UIAction{
		Time: now,
		Kind: event.UIHover,
		X:    x,
		Y:    y,
	})
}

