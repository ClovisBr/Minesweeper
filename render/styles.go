package render

import (
	"github.com/ClovisBr/Minesweeper/engine"
	"github.com/gdamore/tcell/v2"
)

var (
	styleHidden = tcell.StyleDefault.
			Foreground(tcell.ColorGray).
			Background(tcell.ColorBlack)

	styleRevealed = tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorBlack)

	styleMine = tcell.StyleDefault.
			Foreground(tcell.ColorRed).
			Background(tcell.ColorBlack)

	styleFlag = tcell.StyleDefault.
			Foreground(tcell.ColorRed).
			Background(tcell.ColorBlack)
)

func cellStyle(cell engine.Cell) (rune, tcell.Style) {
	switch {
	case cell.Has(engine.FlagFlag):
		return 'F', styleFlag

	case !cell.Has(engine.FlagReveal):
		return '~', styleHidden

	case cell.Has(engine.FlagMine):
		return '*', styleMine

	default:
		n := cell.GetNeighborCount()
		if n > 0 {
			return rune('0' + n), numberStyle(n)
		}
		return ' ', styleRevealed
	}
}

func numberStyle(n uint8) tcell.Style {
	switch n {
	case 1:
		return tcell.StyleDefault.Foreground(tcell.ColorBlue)
	case 2:
		return tcell.StyleDefault.Foreground(tcell.ColorGreen)
	case 3:
		return tcell.StyleDefault.Foreground(tcell.ColorRed)
	case 4:
		return tcell.StyleDefault.Foreground(tcell.ColorPurple)
	case 5:
		return tcell.StyleDefault.Foreground(tcell.ColorMaroon)
	case 6:
		return tcell.StyleDefault.Foreground(tcell.ColorTeal)
	case 7:
		return tcell.StyleDefault.Foreground(tcell.ColorBlack)
	case 8:
		return tcell.StyleDefault.Foreground(tcell.ColorGray)
	default:
		return styleRevealed
	}
}
