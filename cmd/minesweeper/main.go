package main

import (
	"github.com/ClovisBr/Minesweeper/config"
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/generator"
	"github.com/ClovisBr/Minesweeper/render"
)

func main() {
	cfg := config.Default()

	grid := engine.NewGrid(cfg)

	indices, err := generator.GenerateMines(cfg)
	if err != nil {
		panic(err)
	}

	if err := grid.PlaceMines(indices); err != nil {
		panic(err)
	}

	if err := generator.ComputeNeighbors(grid); err != nil {
		panic(err)
	}

	cursorR, cursorC := 0, 0

	renderer, err := render.New()
	if err != nil {
		panic(err)
	}
	renderer.DrawGrid(grid, cursorR, cursorC)
	defer renderer.Close()

	select {}
}
