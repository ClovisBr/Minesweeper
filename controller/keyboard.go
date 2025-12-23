package controller

import (
	"github.com/ClovisBr/Minesweeper/config"
	"github.com/gdamore/tcell/v2"
)

type KeyboardMapper struct {
	keys map[rune]Intent
}

func NewKeyboardMapper(cfg config.Keyboard) *KeyboardMapper {
	return &KeyboardMapper{
		keys: map[rune]Intent{
			cfg.Up:         IntentUp,
			cfg.Down:       IntentDown,
			cfg.Left:       IntentLeft,
			cfg.Right:      IntentRight,
			cfg.Reveal:     IntentReveal,
			cfg.ToggleFlag: IntentToggleFlag,
			cfg.Quit:       IntentQuit,
		},
	}
}

func (k *KeyboardMapper) Map(ev *tcell.EventKey) Intent {
	if ev.Key() != tcell.KeyRune {
		return IntentNone
	}

	if i, ok := k.keys[ev.Rune()]; ok {
		return i
	}

	return IntentNone
}

func MapKey(ev *tcell.EventKey) Intent {
	if ev.Key() == tcell.KeyRune {
		switch ev.Rune() {
		case 'k':
			return IntentUp
		case 'j':
			return IntentDown
		case 'h':
			return IntentLeft
		case 'l':
			return IntentRight
		case ' ':
			return IntentReveal
		case 'f':
			return IntentToggleFlag
		case 'q':
			return IntentQuit
		}
	}

	switch ev.Key() {
	case tcell.KeyUp:
		return IntentUp
	case tcell.KeyDown:
		return IntentDown
	case tcell.KeyLeft:
		return IntentLeft
	case tcell.KeyRight:
		return IntentRight
	case tcell.KeyEnter:
		return IntentReveal
	}

	return IntentNone
}
