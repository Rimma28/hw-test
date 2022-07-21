package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type countError int32

func (c *countError) inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *countError) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return ErrErrorsLimitExceeded
	}
	var wg sync.WaitGroup
	i := 1
	k := 1
	var c countError
	for _, t := range tasks {
		if c.get() >= int32(m) {
			return ErrErrorsLimitExceeded
		}
		t := t
		if i <= n {
			runGoroutine(&wg, t, &c)
			i++
			k++
		} else {
			wg.Wait()
			runGoroutine(&wg, t, &c)
			i = 1
			if c.get() >= int32(m) {
				fmt.Println(444)
				return ErrErrorsLimitExceeded
			}
		}
	}
	wg.Wait()

	return nil
}

func runGoroutine(wg *sync.WaitGroup, t Task, c *countError) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := t()
		if err != nil {
			c.inc()
		}
	}()
	return
}
