package event

type UIActionKind int

const (
	UINone UIActionKind = iota
	UIMoveCursorUp
	UIMoveCursorDown
	UIMoveCursorLeft
	UIMoveCursorRight
	UIHover  // souris
	UIScroll // futur viewport
)

type UIAction struct {
	Time int32
	Kind UIActionKind

	// optionnels selon le Kind
	X int // souris écran
	Y int
}

type Viewport struct {
	Row0 int
	Col0 int
	Rows int
	Cols int
}

type UIState struct {
	Time int32

	Cursor   int
	Hover    *int
	Viewport Viewport
}
