package main

import (
	"github.com/ClovisBr/Minesweeper/config"
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/generator"
)

func main() {
	cfg := config.Default()
	indices, _ := generator.GenerateMines(cfg)
	grid := engine.NewGrid(cfg)
	grid.PlaceMines(indices)
	grid.PrintCells()

}
