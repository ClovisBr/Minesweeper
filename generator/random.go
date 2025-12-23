package generator

import (
	"errors"
	"math/rand"

	"github.com/ClovisBr/Minesweeper/config"
	"github.com/ClovisBr/Minesweeper/engine"
)

func GenerateMines(cfg config.Config) ([]engine.CellIndex, error) {
	total := cfg.Grid.Rows * cfg.Grid.Cols
	mines := cfg.Grid.Mines

	if mines > total {
		return nil, errors.New("more mines than cells")
	}

	rng := rand.New(rand.NewSource(cfg.Grid.Seed))

	indices := make([]engine.CellIndex, total)
	for i := 0; i < total; i++ {
		indices[i] = engine.CellIndex(i)
	}

	// Fisher–Yates partiel : seulement `mines` itérations
	for i := 0; i < mines; i++ {
		j := i + rng.Intn(total-i)
		indices[i], indices[j] = indices[j], indices[i]
	}

	return indices[:mines], nil
}
