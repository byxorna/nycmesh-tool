package app

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/authorization"
)

var (
	daemonWatchFailureBackoff = time.Second * 10
)

func (a *App) runLogInformer(ctx context.Context) error {
	wg := sync.WaitGroup{}
	fountain := make(chan LogEvent, 10)

	consumers := []func(context.Context, *sync.WaitGroup, <-chan LogEvent){}
	if a.config.Daemon.DFSEventDetection {
		consumers = append(consumers, a.logConsumerDFSEventDetector)
	}

	a.startLogProducer(ctx, &wg, fountain)
	outputChannels := a.startLogConsumerMultiplexer(ctx, &wg, fountain, len(consumers))
	for i, startCoroutine := range consumers {
		// create the log watchers, giving them their own channel to listen on for logs
		startCoroutine(ctx, &wg, outputChannels[i])
	}

	wg.Wait()
	return nil
}

func (a *App) RunDaemon(daemonCtx context.Context) (errs []error) {
	a.daemonBootupTimestamp = time.Now()

	encodedConfig, err := json.Marshal(a.config)
	if err != nil {
		return []error{err}
	}

	uispUser, err := a.UISPAPI.Authorization.GetUser(authorization.NewGetUserParams(), nil)
	if err != nil {
		log.Printf("unable to fetch current user from UISP API: %s", err.Error())
		return []error{err}
	}

	log.Printf("logged into UISP as %s", uispUser.Payload.Username)
	log.Printf("daemon config: %s", encodedConfig)

	informers := []func(context.Context) error{
		a.runLogInformer,
		a.runOutageInformer,
	}

	errCh := make(chan error)
	wgDone := make(chan bool)

	// TODO: informers should map into a channel of errors
	wg := sync.WaitGroup{}
	for i, informer := range informers {
		i := i // local variable, so we dont lose track of which coroutine is which
		wg.Add(1)
		go func(routine func(context.Context) error) {
			err := routine(daemonCtx)
			if err != nil {
				log.Printf("daemon coroutine %d failed: %s", i, err.Error())
				errCh <- err
			}
			wg.Done()
		}(informer)
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
