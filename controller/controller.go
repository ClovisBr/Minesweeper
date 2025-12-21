package controller

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/event"
	"github.com/ClovisBr/Minesweeper/view"
)

type Controller struct {
	layout *view.Layout

	cursor *Cursor

	keyboardCursor engine.CellIndex
	mouseCursor    *engine.CellIndex

	lastCursor engine.CellIndex
}

func NewController(rows, cols int) *Controller {
	cur := NewCursor(rows, cols)
	idx := cur.Index()

	return &Controller{
		layout:         view.NewLayout(rows, cols),
		cursor:         cur,
		keyboardCursor: idx,
		lastCursor:     -1,
	}
}

func (c *Controller) currentCursor() engine.CellIndex {
	if c.mouseCursor != nil {
		return *c.mouseCursor
	}
	return c.keyboardCursor
}

func (c *Controller) SyncCursor() bool {
	cur := c.currentCursor()
	if cur == c.lastCursor {
		return false
	}
	c.lastCursor = cur
	return true
}

func (c *Controller) ApplyIntent(i Intent) event.Action {
	switch i {

	// ---------- UI ----------
	case IntentUp:
		if c.cursor.MoveUp() {
			c.keyboardCursor = c.cursor.Index()
			c.mouseCursor = nil

			return event.UIAction{
				Kind: event.UIMoveCursorUp,
			}
		}

	case IntentDown:
		if c.cursor.MoveDown() {
			c.keyboardCursor = c.cursor.Index()
			c.mouseCursor = nil

			return event.UIAction{
				Kind: event.UIMoveCursorDown,
			}
		}

	case IntentLeft:
		if c.cursor.MoveLeft() {
			c.keyboardCursor = c.cursor.Index()
			c.mouseCursor = nil

			return event.UIAction{
				Kind: event.UIMoveCursorLeft,
			}
		}

	case IntentRight:
		if c.cursor.MoveRight() {
			c.keyboardCursor = c.cursor.Index()
			c.mouseCursor = nil

			return event.UIAction{
				Kind: event.UIMoveCursorRight,
			}
		}

	// ---------- Gameplay ----------
	case IntentReveal:
		return event.GameplayAction{
			Time:  now(),
			Kind:  event.ActionReveal,
			Index: c.currentCursor(),
		}

	case IntentToggleFlag:
		return event.GameplayAction{
			Time:  now(),
			Kind:  event.ActionToggleFlag,
			Index: c.currentCursor(),
		}
	}

	return nil
}

func (c *Controller) HandleMouseMove(x, y int) {
	if idx, ok := c.layout.ScreenToCell(x, y); ok {
		c.mouseCursor = &idx
	} else {
		c.mouseCursor = nil
	}
}
