package perf_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/s3vt/perf-go"
)

func ExampleCounterMutexGet() {
	counter := perf.CounterMutexGet()
	fmt.Printf("Got counter %d", counter)
}

func ExampleCounterAtomicGet() {
	counter := perf.CounterAtomicGet()
	fmt.Printf("Got counter %d", counter)
}

func ExampleCounterAtomicReset() {
	// Resets and returns last value
	counter := perf.CounterAtomicReset()
	counter = perf.CounterAtomicGet()
	fmt.Printf("Got counter %d", counter)
	//Output : Got Counter 1
}

func ExampleCounterMutexReset() {
	// Resets and returns last value
	counter := perf.CounterMutexReset()
	counter = perf.CounterMutexGet()
	fmt.Printf("Got counter %d", counter)
	//Output : Got Counter 1}
}

var tests = []struct {
	name      string
	getFunc   func() uint64
	resetFunc func() uint64
}{
	{"CounterMutex", perf.CounterMutexGet, perf.CounterMutexReset},
	{"CounterAtomic", perf.CounterAtomicGet, perf.CounterAtomicReset},
}

var numGoRouts = []int{1, 10, 100, 1000, 10000}

func TestCounterGet(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		for _, numGoRout := range numGoRouts {
			t.Run(fmt.Sprintf("%s-%d", tt.name, numGoRout), func(t *testing.T) {
				wg := sync.WaitGroup{}
				wg.Add(numGoRout)
				for i := 0; i < numGoRout; i++ {
					go func(waitgroup *sync.WaitGroup) {
						defer waitgroup.Done()
						tt.getFunc()
					}(&wg)
				}
				wg.Wait()

				if countBeforeReset := tt.resetFunc(); countBeforeReset != uint64(numGoRout) {
					t.Errorf("Inconsistent count after %d updates %d", numGoRout, countBeforeReset)
				}
			})
		}
	}

}

func BenchmarkCounterAtomic(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			perf.CounterAtomicGet()
		}
	})
}

func BenchmarkCounterMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			perf.CounterMutexGet()
		}
	})
}
