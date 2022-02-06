package app

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func (a *App) RunDaemon() (errs []error) {
	coroutines := []func() error{
		func() error {
			time.Sleep(10 * time.Second)
			return fmt.Errorf("BOOM")
		},
	}
	errCh := make(chan error)
	wgDone := make(chan bool)

	// TODO: coroutines should map into a channel of errors
	wg := sync.WaitGroup{}
	for i, cr := range coroutines {
		i := i // local variable, so we dont lose track of which coroutine is which
		wg.Add(1)
		go func(routine func() error) {
			err := routine()
			if err != nil {
				log.Printf("daemon coroutine %d failed: %s", i, err.Error())
				errCh <- err
			}
			wg.Done()
		}(cr)
	}

	go func() {
		wg.Wait()
		close(wgDone) // make sure to close the done channel, so we know when to wrap up
	}()

	select {
	case <-wgDone:
		break
	case err := <-errCh:
		errs = append(errs, err)
	}

	return errs
}
