package app

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func (a *App) coroutineDFSWatch(ctx context.Context) error {
	log.Printf("dfs watch: start")
	return fmt.Errorf("implement DFS watching")
	time.Sleep(15 * time.Second)
	log.Printf("dfs watch: end")
	return nil
}

func (a *App) RunDaemon(daemonCtx context.Context) (errs []error) {
	coroutines := []func(context.Context) error{
		// NOTE(gabe): define all tasks that should run within the context of the daemon here
		a.coroutineDFSWatch,
	}

	errCh := make(chan error)
	wgDone := make(chan bool)

	// TODO: coroutines should map into a channel of errors
	wg := sync.WaitGroup{}
	for i, cr := range coroutines {
		i := i // local variable, so we dont lose track of which coroutine is which
		wg.Add(1)
		go func(routine func(context.Context) error) {
			err := routine(daemonCtx)
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
