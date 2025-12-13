package generator

import "github.com/ClovisBr/Minesweeper/engine"

type Generator interface {
	Generate(rows, cols, mines int) *engine.Grid
}
