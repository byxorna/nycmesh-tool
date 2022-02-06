package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/logs"
	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

var (
	logFetchErrorBackoff = time.Second * 2
	logFetchInterval     = time.Second * 5
	// timeout for how long a single http request can take
	logFetchTimeout = time.Second * 10
	// how many logs to fetch in a single http request (page size)
	// 1000 logs from UISP is about 1Mb JSON, roughly
	logFetchCount = float64(1000)

	// each time we ask for logs, we ask for an interval (period in uisp api)
	// thats represented in milliseconds, from 1 hour up to ?
	logQueryPeriod = time.Hour * 1
)

type LogEventUISP models.Model9 // -_- generated code suxxx, why this confusing name

type LogEvent struct {
	Time time.Time
	LogEventUISP
}

func (a *App) WatchLogs(ctx context.Context, dstCh chan<- LogEvent) error {
	qPeriod := float64(logQueryPeriod.Milliseconds())
	var latestLogTimestampObserved time.Time

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("cancelled by context")
		case <-time.After(logFetchInterval):
			// fetch more logs, and handle pagination
			fetchLogsOnce := func(latestTimestampSeen *time.Time, dstCh chan<- LogEvent) error {
				mustBeNewerThan := *latestTimestampSeen
				params := logs.NewGetLogsParams().
					WithDefaults().
					WithContext(ctx).
					WithTimeout(logFetchTimeout).
					WithCount(logFetchCount).
					WithPage(float64(1)).
					WithPeriod(&qPeriod)

				for {
					res, err := a.UISPAPI.Logs.GetLogs(params, nil)
					if err != nil {
						log.Printf("fetching logs failed: %s", err.Error())
						time.Sleep(logFetchErrorBackoff)
						continue
					}
					fetchedLogEvents := []LogEvent{}

					//log.Printf("got pg %.0f items %d, pagination page=%d pages=%d total=%d", params.Page, len(res.Payload.Items), *res.Payload.Pagination.Page, *res.Payload.Pagination.Pages, *res.Payload.Pagination.Total)

					// collect events, and send them to the dstCh
					for _, l := range res.Payload.Items {
						if ets, err := time.Parse(time.RFC3339, l.Timestamp.String()); err != nil {
							log.Printf("unable to parse timestamp: %s %s", l.Timestamp, err.Error())
						} else {
							logevent := LogEvent{
								Time:         ets,
								LogEventUISP: LogEventUISP(*l),
							}
							if logevent.Time.After(latestLogTimestampObserved) {
								// record the latest event we have seen
								latestLogTimestampObserved = logevent.Time
							}
							if logevent.Time.After(mustBeNewerThan) {
								fetchedLogEvents = append(fetchedLogEvents, logevent)
							}
							//otherwise, drop the log by not emitting it, it has already been seen
						}
					}

					if res.Payload.Pagination.Pages != nil && float64(*res.Payload.Pagination.Pages) == params.Page {
						// all done with pagination, lets wrap up this fetch

						// first, reverse the events, because they come from UISP newest to oldest
						for i, j := 0, len(fetchedLogEvents)-1; i < j; i, j = i+1, j-1 {
							fetchedLogEvents[i], fetchedLogEvents[j] = fetchedLogEvents[j], fetchedLogEvents[i]
						}

						// dump the log events into our destination channel in the right order
						for _, logevent := range fetchedLogEvents {
							dstCh <- logevent
						}

						return nil
					}

					// set the params for the next page
					params.SetPage(float64(*res.Payload.Pagination.Page + 1))
				}
			}

			if err := fetchLogsOnce(&latestLogTimestampObserved, dstCh); err != nil {
				log.Printf("fetching logs failed: %s", err.Error())
			}

		}
	}
	log.Printf("watchlogs exit")

	return nil
}
