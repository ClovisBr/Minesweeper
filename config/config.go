package config

import "time"

type Config struct {
	Grid     Grid
	Controls Controls
}

type Grid struct {
	Rows  int
	Cols  int
	Mines int
	Seed  int64
}

type Controls struct {
	Keyboard Keyboard
	Mouse    Mouse
}

type Keyboard struct {
	Up    rune
	Down  rune
	Left  rune
	Right rune

	Reveal     rune
	ToggleFlag rune
	Quit       rune
}

type Mouse struct {
	Reveal     MouseButton
	ToggleFlag MouseButton
}

type Map struct {
	Controls Controls
}

type MouseButton int

const (
	MouseNone MouseButton = iota
	MouseLeft
	MouseRight
	MouseMiddle
)

func NewMap(c Controls) *Map {
	return &Map{Controls: c}
}

func Default() Config {
	return Config{
		Grid: Grid{
			Rows:  20,
			Cols:  20,
			Mines: 50,
			Seed:  time.Now().UnixNano(),
		},
		Controls: Controls{
			Keyboard: Keyboard{
				Up:         'k',
				Down:       'j',
				Left:       'h',
				Right:      'l',
				Reveal:     ' ',
				ToggleFlag: 'f',
				Quit:       'q',
			},
			Mouse: Mouse{
				Reveal:     MouseLeft,
				ToggleFlag: MouseRight,
			},
		},
	}
}
