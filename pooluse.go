package perf

import (
	"context"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/sapvs/gopool"
)

func PoolTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	p := gopool.New(gopool.WithNumWorkers(10), gopool.WithLogger(slog.Default()),
		gopool.WithResultBuffer(100), gopool.WithWorkBuffer(100), gopool.WithContext(ctx))

	result, err := p.Start()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func(r chan gopool.Result, w *sync.WaitGroup) {
		for range r {
		}
		log.Println("resutl channel closed")
		w.Done()
	}(result, &wg)

	go func(p *gopool.Pool, w *sync.WaitGroup) {
		for range 1000 {
			p.Submit(&Task{})
		}
		wg.Done()
	}(p, &wg)

	wg.Wait()
}

type Res struct{}

// Get Get func needs to implemented to get
// result back from the processing result of Task
func (r *Res) Get() any {
	return "asn"
}

type Task struct{}

// Do implement this func as the processing block of Task;
// returns the Result
func (t *Task) Do() gopool.Result {
	return &Res{}
}
