package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func (t Task) Do() error {
	err := t()
	if err != nil {
		return err
	}
	return err
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return ErrErrorsLimitExceeded
	}
	errorChan := make(chan error, 50)
	var wg sync.WaitGroup
	i := 1
	k := 1
	for _, t := range tasks {
		if len(errorChan) >= m {
			return ErrErrorsLimitExceeded
		}
		t := t
		if i <= n {
			errorChan = runGoroutine(&wg, t, errorChan)
			i++
			k++
		} else {
			wg.Wait()
			errorChan = runGoroutine(&wg, t, errorChan)
			i = 1
			if len(errorChan) >= m {
				return ErrErrorsLimitExceeded
			}
		}
	}
	wg.Wait()
	close(errorChan)

	return nil
}

func runGoroutine(wg *sync.WaitGroup, t Task, errorChan chan error) chan error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := t()
		if err != nil {
			errorChan <- err
		}
	}()
	return errorChan
}
