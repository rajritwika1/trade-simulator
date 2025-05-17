package utils

import "time"

type Timer struct {
	start time.Time
	end   time.Time
}

func (t *Timer) Stop() {
	panic("unimplemented")
}

func NewTimer() *Timer {
	return &Timer{start: time.Now()}
}

func (t *Timer) stop() {
	t.end = time.Now()
}

func (t *Timer) Elapsed() float64 {
	return float64(t.end.Sub(t.start).Microseconds()) / 1000.0
}
