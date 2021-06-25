package main

import (
	"errors"
	"sync"
)

func Execute(tasks []func() error, E int) error {
	var (
		count int
		wg    sync.WaitGroup
		mu    sync.Mutex
	)

	for _, task := range tasks {
		wg.Add(1)
		go func(task func() error) {
			defer wg.Done()
			if task() != nil {
				mu.Lock()
				defer mu.Unlock()
				count++
			}
		}(task)
	}
	wg.Wait()
	if count > E {
		return errors.New("number of errors is more than E")
	}

	return nil
}

func ExecuteChan(tasks []func() error, E int) error {
	var (
		errCount       int
		performedTasks int
	)
	c := make(chan error)
	performed := make(chan bool)

	for _, task := range tasks {
		go func(task func() error) {
			err := task()
			if err != nil {
				c <- err
			}
			performed <- true
		}(task)
	}

	for {
		select {
		case <-c:
			errCount++
		case <-performed:
			performedTasks++
		default:
			if performedTasks >= len(tasks) {
				if errCount > E {
					return errors.New("number of errors is more than E")
				}
				return nil
			}
		}
	}
}
