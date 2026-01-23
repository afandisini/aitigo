package ratelimit

import (
	"context"
	"sync"
	"time"
)

type bucket struct {
	tokens float64
	last   time.Time
}

type InMemoryStore struct {
	mu      sync.Mutex
	buckets map[string]*bucket
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		buckets: make(map[string]*bucket),
	}
}

func (s *InMemoryStore) Allow(_ context.Context, key string, limit Limit) (bool, time.Duration, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	b, ok := s.buckets[key]
	if !ok {
		b = &bucket{tokens: float64(limit.Burst), last: now}
		s.buckets[key] = b
	}

	refillRate := float64(limit.Rate) / limit.Per.Seconds()
	elapsed := now.Sub(b.last).Seconds()
	b.tokens = minFloat(float64(limit.Burst), b.tokens+(elapsed*refillRate))
	b.last = now

	if b.tokens >= 1 {
		b.tokens -= 1
		return true, 0, nil
	}

	needed := (1 - b.tokens) / refillRate
	return false, time.Duration(needed * float64(time.Second)), nil
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
