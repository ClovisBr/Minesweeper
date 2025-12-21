package controller

import "github.com/ClovisBr/Minesweeper/engine"

type Cursor struct {
	pos  engine.CellIndex
	rows int
	cols int
}

func NewCursor(rows, cols int) *Cursor {
	return &Cursor{
		pos:  0,
		rows: rows,
		cols: cols,
	}
}

func (c *Cursor) Index() engine.CellIndex {
	return c.pos
}

func (c *Cursor) MoveUp() bool {
	if int(c.pos) >= c.cols {
		c.pos -= engine.CellIndex(c.cols)
		return true
	}
	return false
}

func (c *Cursor) MoveDown() bool {
	if int(c.pos) < (c.rows-1)*c.cols {
		c.pos += engine.CellIndex(c.cols)
		return true
	}
	return false
}

func (c *Cursor) MoveLeft() bool {
	if int(c.pos)%c.cols != 0 {
		c.pos--
		return true
	}
	return false
}

func (c *Cursor) MoveRight() bool {
	if int(c.pos)%c.cols != c.cols-1 {
		c.pos++
		return true
	}
	return false
}
