package app

import (
	"context"
	"log"
	"regexp"
	"sort"
	"sync"
	"time"
)

var (
	daemonWatchFailureBackoff = time.Second * 10
)

/* DFS Event log example

{
      "device": {
      "authorized": true,
      "category": "wireless",
      "displayName": "nycmesh-sn3-south",
      "firmwareVersion": "8.7.7",
      "id": "62351c24-d13b-4e5f-92d9-a39492e718c0",
      "ip": "10.70.88.116/24",
      "mac": "b4:fb:e4:5a:5c:5c",
      "model": "LAP-120",
      "modelName": "LiteAP AC",
      "name": "nycmesh-sn3-south",
      "platformId": "WA",
      "platformName": "WA",
      "site": {
        "id": "fd468966-adc9-4f28-b72d-4d97891d6734",
        "name": "SN3 - 713",
        "parent": {
          "id": "d0b4cf55-15a7-4e60-ba4f-973ef74b5d5e",
          "name": "Degraw - 2282",
          "parent": null,
          "status": "active",
          "type": "site"
        },
        "status": "active",
        "type": "site"
      },
      "started": "0001-01-01T00:00:00.000Z",
      "type": "airMax",
      "updated": "0001-01-01T00:00:00.000Z"
    },
    "deviceMetadata": {
      "alias": "nycmesh-sn3-south"
    },
    "id": "49205812-6bfe-46b8-8b1a-4b16c0581c12",
    "level": "error",
    "message": "Device nycmesh-sn3-south main radio changed frequency due to DFS detection",
    "site": {
      "id": "fd468966-adc9-4f28-b72d-4d97891d6734",
      "name": "SN3 - 713",
      "parent": {
        "id": "d0b4cf55-15a7-4e60-ba4f-973ef74b5d5e",
        "name": "Degraw - 2282",
        "parent": null,
        "status": "active",
        "type": "site"
      },
      "status": "active",
      "type": "site"
    },
    "tags": [
      "device",
      "device-state"
    ],
    "timestamp": "2022-02-03T23:11:03.900Z"
  }
*/

func (a *App) coroutineLogWatch(ctx context.Context) error {
	wg := sync.WaitGroup{}
	logFountain := make(chan LogEvent, 10)
	logConsumerChannels := map[string]chan LogEvent{}

	initialQueryPeriod := (time.Hour * 24 * 2)

	startLogProducer := func(ctx context.Context, wg *sync.WaitGroup, ch chan<- LogEvent) {
		// start up the log producer, which pushes events into ch
		wg.Add(1)
		ctxDone := ctx.Done()

		go func() {
			for {
				select {
				case <-ctxDone:
					break
				default:
					log.Printf("initial log query window is %s", initialQueryPeriod)
					if err := a.WatchLogs(ctx, time.Now().Add(-initialQueryPeriod), ch); err != nil {
						log.Printf("unable to watch logs: %s", err.Error())
						log.Printf("retrying log watch after %s backoff", daemonWatchFailureBackoff)
						time.Sleep(daemonWatchFailureBackoff)
					}
				}
			}

			close(ch)
			wg.Done()
		}()
	}

	startLogConsumerMultiplexer := func(ctx context.Context, wg *sync.WaitGroup, srcCh <-chan LogEvent) {
		// create a multiplexer, so for each consumer that cares about logs, we make a new channel and send
		// each log to each consumer
		wg.Add(1)
		ctxDone := ctx.Done()

		go func() {
			defer func() {
				for _, ch := range logConsumerChannels {
					close(ch)
				}
				wg.Done()
			}()

			for {
				select {
				case <-ctxDone:
					break
				case evt := <-srcCh:
					// multiplex this log to all our logConsumerChannels
					for _, consumerCh := range logConsumerChannels {
						consumerCh <- evt
					}
				}
			}
		}()
	}

	createDFSEventDetector := func(wg *sync.WaitGroup, logsCh <-chan LogEvent) {
		dfsEventCh := make(chan LogEvent)
		// start up the DFS detection consumer, that filters events for DFS events, and emits them to dfsEventCh
		wg.Add(1)
		ctxDone := ctx.Done()
		go func() {
			defer func() {
				close(dfsEventCh)
				wg.Done()
			}()

			dfsMessageRegex := regexp.MustCompile(`\bchanged frequency due to DFS detection\b`)
			//dfsMessageRegex := regexp.MustCompile(`\bhas been disconnected\b`)
			log.Printf("watching for DFS events with `%v`", dfsMessageRegex)
			for {
				select {
				case <-ctxDone:
					break
				case rawLog := <-logsCh:
					t := rawLog.Tags
					sort.Strings(t)

					if rawLog.Message != nil && dfsMessageRegex.MatchString(*rawLog.Message) { // && sort.SearchStrings(t, "device-state") != len(t) {
						dfsEventCh <- rawLog
					}
				}
			}

		}()

		// read from dfsEventCh, log events as we see them
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctxDone:
					break
				case dfsEvent := <-dfsEventCh:
					log.Printf("DFS event detected at nn:%d on %s: %s", dfsEvent.NN, dfsEvent.Timestamp, *dfsEvent.Message)
				}
			}
		}()
	}

	logWatchers := map[string]func(*sync.WaitGroup, <-chan LogEvent){}
	if a.config.DaemonConfig.DFSEventDetection {
		id := "dfs"
		logConsumerChannels[id] = make(chan LogEvent)
		logWatchers[id] = createDFSEventDetector
	}

	startLogProducer(ctx, &wg, logFountain)
	startLogConsumerMultiplexer(ctx, &wg, logFountain)
	for id, startCoroutine := range logWatchers {
		// create the log watchers, giving them their own channel to listen on for logs
		startCoroutine(&wg, logConsumerChannels[id])
	}

	wg.Wait()
	return nil
}

func (a *App) RunDaemon(daemonCtx context.Context) (errs []error) {
	coroutines := []func(context.Context) error{
		a.coroutineLogWatch,
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
