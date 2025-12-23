package engine

type GameState int

const (
	GameRunning GameState = iota
	GameWon
	GameLost
)

type Game struct {
	Grid  *Grid
	Rules Ruleset
	State GameState
}

func NewGame(grid *Grid, rules Ruleset) *Game {
	return &Game{
		Grid:  grid,
		Rules: rules,
		State: GameRunning,
	}
}

func (g *Game) Apply(a Action) Update {
	// jeu déjà terminé → on ignore
	if g.State != GameRunning {
		return Update{State: g.State}
	}

	var changes []CellChange

	switch a.Kind {
	case ActionReveal:
		changes = g.Rules.Reveal(g.Grid, a.Index)

	case ActionToggleFlag:
		changes = g.Rules.ToggleFlag(g.Grid, a.Index)
	}

	applyChanges(g.Grid, changes)

	if revealedMine(changes) {
		g.State = GameLost
	}

	return Update{
		Cells: changes,
		State: g.State,
	}
}

func applyChanges(grid *Grid, changes []CellChange) {
	for _, ch := range changes {
		cell := grid.CellAt(ch.Index)
		if cell == nil {
			continue
		}
		*cell &^= ch.Mask
		*cell |= ch.Value & ch.Mask
	}
}

func revealedMine(changes []CellChange) bool {
	for _, ch := range changes {
		// une mine est révélée si on a mis FlagReveal à 1 ET que FlagMine vaut 1 sur la cellule
		if ch.Mask&Cell(FlagReveal) != 0 &&
			ch.Value&Cell(FlagReveal) != 0 &&
			ch.Value&Cell(FlagMine) != 0 {
			return true
		}
	}
	return false
}
