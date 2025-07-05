package perf

import (
	"sync"
	"sync/atomic"
)

var (
	simplecounter uint64 = 0
	atomiccount   uint64 = 0
	mut           sync.Mutex
)

// CounterMutexGet used sync.Mutex to lock over incrementing an int var
// Returns the incremented counter value
func CounterMutexGet() uint64 {
	mut.Lock()
	defer mut.Unlock()
	simplecounter++
	return simplecounter
}

// CounterMutexReset reset counter and returns last value
func CounterMutexReset() uint64 {
	mut.Lock()
	defer mut.Unlock()
	oldCounter := simplecounter
	simplecounter = 0

	return oldCounter
}

// CounterAtomicGet uses atomic api and returns the incremented counter value
func CounterAtomicGet() uint64 {
	return atomic.AddUint64(&atomiccount, 1)
}

// CounterAtomicReset reset counter and returns last value
func CounterAtomicReset() uint64 {
	return atomic.SwapUint64(&atomiccount, 0)
}
