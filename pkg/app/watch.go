package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client/logs"
)

var (
	logFetchErrorBackoff = time.Second * 2
	logFetchInterval     = time.Second * 5
	// timeout for how long a single http request can take
	logFetchTimeout = time.Second * 10
	// how many logs to fetch in a single http request (page size)
	// 1000 logs from UISP is about 1Mb JSON, roughly
	logFetchCount = float64(1000)

	MinLogQueryPeriod = float64(time.Hour.Milliseconds())
	MaxLogQueryPeriod = float64((24 * 365 * time.Hour).Milliseconds())
)

func (a *App) WatchLogs(ctx context.Context, since time.Time, dstCh chan<- LogEvent) error {
	initialQueryPeriod := float64(time.Now().Sub(since).Milliseconds())
	latestLogTimestampObserved := since // this in effect forces us to drop any messages we fetch from before this time

	qPeriod := initialQueryPeriod
	for {
		if qPeriod < MinLogQueryPeriod {
			qPeriod = MinLogQueryPeriod
		}
		if qPeriod > MaxLogQueryPeriod {
			qPeriod = MaxLogQueryPeriod
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("cancelled by context")
		case <-time.After(logFetchInterval):
			// fetch more logs, and handle pagination
			//params.SetPeriod(&MinLogQueryPeriod)
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

							if logevent.Device != nil {
								// try to find NN from device name
								if nn, err := GetNNFromDeviceName(logevent.Device.Name); err == nil && logevent.NN == 0 {
									logevent.NN = nn
								}
							}

							if logevent.Site != nil {
								// try to find NN from device name
								if nn, err := GetNNFromSiteName(*logevent.Site.Name); err == nil && logevent.NN == 0 {
									logevent.NN = nn
								}

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

						// decrement the period to minimum, so subsequent queries will not do as much
						// damage to UISP API
						qPeriod = MinLogQueryPeriod
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
