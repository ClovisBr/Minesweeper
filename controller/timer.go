package controller

import (
	"time"

	"github.com/ClovisBr/Minesweeper/event"
)

type Timer struct {
	start time.Time
}

func NewTimer() *Timer {
	return &Timer{start: time.Now()}
}

func (t *Timer) Now() event.Time {
	d := time.Since(t.start)
	return event.Time(d.Milliseconds())
}
