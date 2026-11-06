package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidWorkersCount = errors.New("workers count must be positive")
)

type Task func() error

func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrInvalidWorkersCount
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	if len(tasks) == 0 {
		return nil
	}

	var (
		wg       sync.WaitGroup
		errCount int32
	)

	tasksCh := make(chan Task)

	done := make(chan struct{})
	var stopOnce sync.Once
	stop := func() {
		stopOnce.Do(func() { close(done) })
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for task := range tasksCh {
				select {
				case <-done:
					return
				default:
				}

				if err := task(); err != nil {
					if int(atomic.AddInt32(&errCount, 1)) >= m {
						stop()
					}
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(tasksCh)

		for _, task := range tasks {
			select {
			case <-done:
				return
			case tasksCh <- task:
			}
		}
	}()

	wg.Wait()

	if int(atomic.LoadInt32(&errCount)) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
