package render

import "github.com/gdamore/tcell/v2"

var (
	StyleHidden   = tcell.StyleDefault.Foreground(tcell.ColorGray)
	StyleRevealed = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	StyleMine     = tcell.StyleDefault.Foreground(tcell.ColorRed)
	StyleCursor   = tcell.StyleDefault.Reverse(true)
)
