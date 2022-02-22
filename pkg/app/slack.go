package app

import (
	"fmt"
	"strings"
	"time"
)

var (
	// TODO: make this mapping dynamic, use config file for this
	DefaultMonitoringChannel = "U010B096FPH" // @gabe
	//DefaultMonitoringChannel = "CSJK7P7C3" // #monitoring-unms
	NNToSlackChannel = map[int]string{
		//713:  "CAZ4V9E0P",   // #sn3
		//5712: "C02HZLLH85R", // #hub-clay-5712
	}
)

func lookupSlackChannelIDFromNN(nn int) (string, error) {
	if ch, ok := NNToSlackChannel[nn]; ok {
		return ch, nil
	}
	return "", fmt.Errorf("implement me: lookup of slack channel for nn:%d", nn)
}

func (a *App) slackEnsureOutageStatusUpdate(nn int, outageMap OutageMap) error {
	devOutages, err := outageMap.NodeDeviceOutages(nn)
	if err != nil {
		return fmt.Errorf("unable to get device outages for nn:%d: %w", nn, err)
	}

	msg := createMarkdownStatus(nn, baseURL, devOutages) // TODO: should we include site info, or all devices in this site?

	outageThreadID, err := findOrCreateOutageThread(nn)
	if err != nil {
		return fmt.Errorf("unable to create outage thread: %w", err)
	}

	if err := ensureNodeStatusMessageInThread(outageThreadID, msg); err != nil {
		return fmt.Errorf("unable to ensure node status in outage thread: %w", err)
	}

	return nil
}

// given the devices that are currently in outage, build a markdown status message
// useful to show the current state of a NN's offline devices
// TODO: should this show all devices, not just those in outage?
func createMarkdownStatus(nn int, baseURL string, devOutages []Outage) string {
	sb := strings.Builder{}
	var outageBeginEarliest time.Time
	var numOfflineDevs = 0
	var numUnreachableDevs = 0

	items := []string{}

	for _, o := range devOutages {
		if outageBeginEarliest.After(o.Start) {
			outageBeginEarliest = o.Start
		}

		dur := time.Since(o.Start).Round(time.Minute)

		var typeFormat string
		switch o.Type {
		case "outage":
			numOfflineDevs++
			typeFormat = fmt.Sprintf("**%s**", o.Type)
		case "unreachable":
			numUnreachableDevs++
			typeFormat = fmt.Sprintf("_%s_", o.Type)
		default:
			typeFormat = fmt.Sprintf("%s", o.Type)
		}

		items = append(items, fmt.Sprintf("`%s` %s in %s for %s", *o.Device.Name, *o.Device.IP, typeFormat, dur))
	}

	aggrHealth := []string{}

	switch {
	case numOfflineDevs > 0:
		aggrHealth = append(aggrHealth, fmt.Sprintf("%d devices offline", numOfflineDevs))
	case numUnreachableDevs > 0:
		aggrHealth = append(aggrHealth, fmt.Sprintf("%d unreachable", numUnreachableDevs))
	default:
		aggrHealth = []string{"online"}
	}

	if outageBeginEarliest.IsZero() {
		sb.WriteString(fmt.Sprintf("nn:%d %s\n", nn, strings.Join(aggrHealth, ", ")))
	} else {
		sb.WriteString(fmt.Sprintf("nn:%d %s for %s\n", nn, strings.Join(aggrHealth, ", "), outageBeginEarliest.Round(time.Minute)))
	}
	for _, v := range items {
		sb.WriteString(fmt.Sprintf("- %s\n", v))
	}

	return sb.String()
}

func findOrCreateOutageThread(nn int) (string, error) {
	return "", fmt.Errorf("FindOrCreateOutageThread")
}

func ensureNodeStatusMessageInThread(message_id string) error {
	return fmt.Errorf("EnsureNodeStatusMessageInThread")
}

func postThreadUpdate(message_id string, update string) error {
	return fmt.Errorf("PostThreadUpdate")
}

// iterate over all NNs in an outage map, and update its absolute status in the "Outage"
// thread
func (a *App) updateSlackOutageThreadsFromOutageMaps(outageMap OutageMap, prevOutageMap OutageMap) error {
	for _, nn := range allNNs(outageMap, prevOutageMap) {
		if err := a.slackEnsureOutageStatusUpdate(nn, outageMap); err != nil {
			return fmt.Errorf("unable to update slack outage thread: %w", err)
		}
	}

	return nil
}
