package engine

import "errors"

type Cell uint16
type CellFlag uint16

const (
	neighborMask Cell = 0b1111

	FlagMine   CellFlag = 0b0000000000010000
	FlagReveal CellFlag = 0b0000000000100000
	FlagFlag   CellFlag = 0b0000000001000000
)

func (c *Cell) Has(flag CellFlag) bool {
	return *c&Cell(flag) != 0
}

func (c *Cell) Set(flag CellFlag) {
	*c |= Cell(flag)
}

func (c *Cell) Clear(flag CellFlag) {
	*c &^= Cell(flag)
}

func (c *Cell) GetNeighborCount() uint8 {
	return uint8(*c & neighborMask)
}

func (c *Cell) SetNeighborCount(n uint8) error {
	if n > 8 {
		return errors.New("Neighbor Count out of range")
	}
	*c &^= neighborMask          // clear
	*c |= Cell(n) & neighborMask // set
	return nil
}
