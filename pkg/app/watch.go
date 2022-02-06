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

type LogEvent models.Model9 // -_- generated code suxxx, why this confusing name

func (a *App) WatchLogs(ctx context.Context, dstCh chan<- LogEvent) error {
	qPeriod := float64(logQueryPeriod.Milliseconds())

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("cancelled by context")
		case <-time.After(logFetchInterval):
			// fetch more logs, and handle pagination
			fetchLogsOnce := func(dstCh chan<- LogEvent) error {
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

					log.Printf("got pg %.0f items %d, pagination page=%d pages=%d total=%d", params.Page, len(res.Payload.Items), *res.Payload.Pagination.Page, *res.Payload.Pagination.Pages, *res.Payload.Pagination.Total)

					// collect events, and send them to the dstCh
					// TODO: sort them reverse, so we write oldest first?
					for _, l := range res.Payload.Items {
						//log.Printf("sending %+v", LogEvent(*l))
						dstCh <- LogEvent(*l)
					}

					if res.Payload.Pagination.Pages != nil && float64(*res.Payload.Pagination.Pages) == params.Page {
						// all done, lets wrap up
						log.Printf("reached end of pagination")
						return nil
					}

					// set the params for the next page
					params.SetPage(float64(*res.Payload.Pagination.Page + 1))
				}
			}

			if err := fetchLogsOnce(dstCh); err != nil {
				log.Printf("fetching logs failed: %s", err.Error())
			}
			log.Printf("fetching logs CLOSE")

		}
	}
	log.Printf("watchlogs exit")

	return nil
}
