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
	log.Printf("dfs watch: start")

	ctxDone := ctx.Done()
	wg := sync.WaitGroup{}
	rawLogsCh := make(chan LogEvent, 10)
	dfsEventCh := make(chan LogEvent)

	// start up the log producer, which pushes events into rawLogsCh
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctxDone:
				break
			default:
				// we start watching logs since 24h ago, so we catch any recent events
				if err := a.WatchLogs(ctx, time.Now().Add(-(time.Hour * 24)), rawLogsCh); err != nil {
					log.Printf("unable to watch logs: %w", err)
					log.Printf("retrying log watch after %s backoff", daemonWatchFailureBackoff)
					time.Sleep(daemonWatchFailureBackoff)
				}
			}
		}
		wg.Done()
	}()

	// start up the DFS detection consumer, that filters events for DFS events, and emits them to dfsEventCh
	wg.Add(1)
	go func() {
		dfsMessageRegex := regexp.MustCompile(`\bchanged frequency due to DFS detection\b`)
		//dfsMessageRegex := regexp.MustCompile(`\bhas been disconnected\b`)
		for {
			select {
			case <-ctxDone:
				break
			case rawLog := <-rawLogsCh:
				t := rawLog.Tags
				sort.Strings(t)

				if rawLog.Message != nil && dfsMessageRegex.MatchString(*rawLog.Message) && sort.SearchStrings(t, "device-state") != len(t) {
					dfsEventCh <- rawLog
				}
			}
		}
		wg.Done()
	}()

	// read from dfsEventCh, log events as we see them
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctxDone:
				break
			case dfsEvent := <-dfsEventCh:
				log.Printf("DFS event detected at nn:%d: %s", dfsEvent.NN, *dfsEvent.Message)
			}
		}
		wg.Done()
	}()

	wg.Wait()
	log.Printf("dfs watch: end")
	return nil
}

func (a *App) RunDaemon(daemonCtx context.Context) (errs []error) {
	coroutines := []func(context.Context) error{
		// NOTE(gabe): define all tasks that should run within the context of the daemon here
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
