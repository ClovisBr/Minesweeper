package generator

import (
	"errors"
	"math/rand"

	"github.com/ClovisBr/Minesweeper/config"
	"github.com/ClovisBr/Minesweeper/engine"
)

func GenerateMines(cfg config.Config) (engine.CellIndices, error) {
	total := cfg.Rows * cfg.Cols
	mines := cfg.Mines

	if mines < 0 {
		return nil, errors.New("mines cannot be negative")
	}
	if mines > total {
		return nil, errors.New("more mines than cells")
	}

	rng := rand.New(rand.NewSource(cfg.Seed))

	indices := make(engine.CellIndices, total)
	for i := range indices {
		indices[i] = i
	}

	// Fisher–Yates partiel : seulement mines itérations
	for i := 0; i < mines; i++ {
		j := i + rng.Intn(total-i)
		indices[i], indices[j] = indices[j], indices[i]
	}

	return indices[:mines], nil
}
