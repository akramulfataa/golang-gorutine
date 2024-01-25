package main

import (
	"sync"
	"testing"
)

func TestDatatRaceCondition(t *testing.T) {
	var counter int32
	// implementasi mutex
	var mu sync.RWMutex
	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			counter += int32(i)
			mu.Unlock()
		}(i)
	}
}
