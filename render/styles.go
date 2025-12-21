package render

import "github.com/gdamore/tcell/v2"

var (
	StyleHidden = tcell.StyleDefault.
			Foreground(tcell.ColorGray).
			Background(tcell.ColorBlack)

	StyleRevealed = tcell.StyleDefault.
			Foreground(tcell.ColorWhite).
			Background(tcell.ColorBlack)

	StyleMine = tcell.StyleDefault.
			Foreground(tcell.ColorRed).
			Background(tcell.ColorBlack)

	StyleFlag = tcell.StyleDefault.
			Foreground(tcell.ColorRed).
			Background(tcell.ColorBlack)
)

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
		return StyleRevealed
	}
}
