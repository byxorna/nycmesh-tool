package app

import (
	"context"
	"log"
	"regexp"
	"sync"
	"time"

	"github.com/slack-go/slack"
)

func (a *App) logConsumerDFSEventDetector(ctx context.Context, wg *sync.WaitGroup, logsCh <-chan LogEvent) {
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
				if rawLog.Message != nil && dfsMessageRegex.MatchString(*rawLog.Message) {

					// drop any DFS events that happened before our daemon started running
					if rawLog.Time.Before(a.daemonBootupTimestamp) {
						elapsed := time.Since(rawLog.Time).Round(time.Minute)
						log.Printf("[ignored] DFS event detected before daemon starup at nn:%d (%s, %s ago): %s", rawLog.NN, rawLog.Time.String(), elapsed.String(), *rawLog.Message)
						continue
					}

					dfsEventCh <- rawLog
				}
			}
		}

	}()

	// read from dfsEventCh, handle valid DFS events
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctxDone:
				break
			case dfsEvent := <-dfsEventCh:
				elapsed := time.Since(dfsEvent.Time)

				log.Printf("DFS event detected at nn:%d %s ago: %s (%s)", dfsEvent.NN, elapsed.String(), *dfsEvent.Message, dfsEvent.Time.String())

				if !a.config.Daemon.EnableSlack {
					continue
				}

				slackChannel, err := lookupSlackChannelIDFromNN(dfsEvent.NN)
				if err != nil {
					log.Printf("Unable to resolve slack channel (will use default monitoring channel %s): %s", DefaultMonitoringChannel, err.Error())
					slackChannel = DefaultMonitoringChannel
				}

				log.Printf("notifying slack channel %s of DFS on nn:%d %s", slackChannel, dfsEvent.NN, dfsEvent.Device.Name)
				log.Printf("%+v %s", dfsEvent.LogEventUISP, *dfsEvent.Message)
				go func() {
					// TODO: add buttons to link out to UISP for the device triggering the event

					channelID, timestamp, err := a.Slack.PostMessage(
						slackChannel,
						slack.MsgOptionText(*dfsEvent.Message, false),
						nil,                         //slack.MsgOptionAttachments(attachment)
						slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
					)
					if err != nil {
						log.Printf("Slack error: %s\n", err.Error())
						return
					} else {
						log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
					}
				}()
			}
		}
	}()
}

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
