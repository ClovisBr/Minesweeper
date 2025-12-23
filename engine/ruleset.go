package engine

// Ruleset définit un ensemble de règles de gameplay.
// Il est volontairement minimal et stateless.
type Ruleset interface {
	Reveal(g *Grid, idx CellIndex) []CellChange
	ToggleFlag(g *Grid, idx CellIndex) []CellChange
}
