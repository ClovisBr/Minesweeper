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
