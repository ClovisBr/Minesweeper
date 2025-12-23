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
	Time Time
	Kind UIActionKind

	// optionnels selon le Kind
	X int // souris écran
	Y int
}

func (UIAction) isAction() {}

type Viewport struct {
	Row0 int
	Col0 int
	Rows int
	Cols int
}

type UIState struct {
	Time Time

	Cursor   int
	Hover    *int
	Viewport Viewport
}
