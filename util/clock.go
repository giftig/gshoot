package util

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (c realClock) Now() time.Time {
	return time.Now()
}

type mockClock struct {
	Time time.Time
}

func (c mockClock) Now() time.Time {
	return c.Time
}

func NewClock() realClock {
	return realClock{}
}

func MockClock(t time.Time) mockClock {
	return mockClock{Time: t}
}
