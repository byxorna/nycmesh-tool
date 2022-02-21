package app

import (
	"context"
	"log"
	"sync"
	"time"
)

func (a *App) startLogProducer(ctx context.Context, wg *sync.WaitGroup, ch chan<- LogEvent) {
	// start up the log producer, which pushes events into ch
	wg.Add(1)
	ctxDone := ctx.Done()

	initialQueryPeriod := (time.Hour * 48)
	log.Printf("bootstrapping by fetching logs %s old", initialQueryPeriod)

	go func() {
		defer func() {
			close(ch)
			wg.Done()
		}()
		for {
			select {
			case <-ctxDone:
				break
			default:
				if err := a.WatchLogs(ctx, time.Now().Add(-initialQueryPeriod), ch); err != nil {
					log.Printf("unable to watch logs: %s", err.Error())
					log.Printf("retrying log watch after %s backoff", daemonWatchFailureBackoff)
					time.Sleep(daemonWatchFailureBackoff)
				}
			}
		}
	}()
}

// TODO: good candidate for genericization
func (a *App) startLogConsumerMultiplexer(ctx context.Context, wg *sync.WaitGroup, srcCh <-chan LogEvent, numChannelsToMultiplexTo int) []chan LogEvent {
	// create a multiplexer, so for each consumer that cares about logs, we make a new channel and send
	// each log to each consumer
	wg.Add(1)
	ctxDone := ctx.Done()

	downstreamChannels := make([]chan LogEvent, numChannelsToMultiplexTo)
	for i := range downstreamChannels {
		downstreamChannels[i] = make(chan LogEvent)
	}

	go func() {
		defer func() {
			for _, ch := range downstreamChannels {
				close(ch)
			}
			wg.Done()
		}()

		for {
			select {
			case <-ctxDone:
				break
			case evt := <-srcCh:
				// multiplex this log to all our downstreamChannels
				for _, consumerCh := range downstreamChannels {
					consumerCh <- evt
				}
			}
		}
	}()

	return downstreamChannels
}
