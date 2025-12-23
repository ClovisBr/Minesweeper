package main

import (
	"log"

	"github.com/ClovisBr/Minesweeper/config"
	"github.com/ClovisBr/Minesweeper/controller"
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/ClovisBr/Minesweeper/engine/rules"
	"github.com/ClovisBr/Minesweeper/generator"
	"github.com/ClovisBr/Minesweeper/render"
	"github.com/ClovisBr/Minesweeper/view"
	"github.com/gdamore/tcell/v2"
)

func main() {
	// ---------- config ----------
	cfg := config.Default()

	// ---------- engine ----------
	grid := engine.NewGrid(cfg)

	indices, err := generator.GenerateMines(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := grid.PlaceMines(indices); err != nil {
		log.Fatal(err)
	}

	if err := generator.ComputeNeighbors(grid); err != nil {
		log.Fatal(err)
	}

	game := engine.NewGame(grid, rules.Punitive{})

	// ---------- view ----------
	layout := view.NewLayout(cfg.Grid.Rows, cfg.Grid.Cols)
	v := view.NewView(layout, grid.Cells)

	// ---------- controller ----------
	ctrl := controller.NewController(
		game,
		v,
		cfg.Grid.Rows,
		cfg.Grid.Cols,
	)

	// ---------- render ----------
	r, err := render.New()
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// ---------- initial draw ----------
	r.Draw(v)

	// ---------- event loop ----------
	for {
		ev := r.Screen().PollEvent()

		switch e := ev.(type) {

		case *tcell.EventKey:
			if e.Key() == tcell.KeyEscape || e.Rune() == 'q' {
				return
			}

			intent := controller.MapKey(e)
			if intent == controller.IntentNone {
				continue
			}

			update := ctrl.HandleAction(intent)
			if update != nil {
				_ = update // prêt pour replay / réseau
			}

			r.Draw(v)

		case *tcell.EventMouse:
			x, y := e.Position()
			ctrl.HandleMouseMove(x, y)
			r.Draw(v)

		case *tcell.EventResize:
			r.Draw(v)
		}
	}
}

