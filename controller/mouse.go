package controller

import (
	"github.com/ClovisBr/Minesweeper/config"
	"github.com/gdamore/tcell/v2"
)

type MouseMapper struct {
	buttons map[config.MouseButton]Intent
}

func NewMouseMapper(cfg config.Mouse) *MouseMapper {
	return &MouseMapper{
		buttons: map[config.MouseButton]Intent{
			cfg.Reveal:     IntentReveal,
			cfg.ToggleFlag: IntentToggleFlag,
		},
	}
}

func (m *MouseMapper) Map(ev *tcell.EventMouse) Intent {
	btn := mouseButton(ev.Buttons())

	if i, ok := m.buttons[btn]; ok {
		return i
	}

	return IntentNone
}

func mouseButton(b tcell.ButtonMask) config.MouseButton {
	switch {
	case b&tcell.Button1 != 0:
		return config.MouseLeft
	case b&tcell.Button2 != 0:
		return config.MouseMiddle
	case b&tcell.Button3 != 0:
		return config.MouseRight
	default:
		return config.MouseNone
	}
}
