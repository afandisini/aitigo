package testingutil

import (
	"sync"
	"time"
)

type Clock interface {
	Now() time.Time
}

type MockClock struct {
	mu  sync.Mutex
	now time.Time
}

func NewMockClock(initial time.Time) *MockClock {
	return &MockClock{now: initial}
}

func (c *MockClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *MockClock) Set(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = t
}
