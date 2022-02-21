package app

import "fmt"

var (
	// TODO: make this mapping dynamic
	DefaultMonitoringChannel = "U010B096FPH" // @gabe
	//DefaultMonitoringChannel = "CSJK7P7C3" // #monitoring-unms
	NNToSlackChannel = map[int]string{
		//713:  "CAZ4V9E0P",   // #sn3
		//713:  "C02HZLLH85R", // #hub-clay-5712
		//5712: "C02HZLLH85R", // #hub-clay-5712
	}
)

func lookupSlackChannelIDFromNN(nn int) (string, error) {
	if ch, ok := NNToSlackChannel[nn]; ok {
		return ch, nil
	}
	return "", fmt.Errorf("implement me: lookup of slack channel for nn:%d", nn)
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

func (a *App) updateSlackOutages(outageMap OutageMap, prevOutageMap OutageMap) error {
	newOutageNNs, clearedNNs, prevImpactedNNs, currentImpactedNNs := outageNNGroups(outageMap, prevOutageMap)
	return fmt.Errorf("TODO: updateSlackOutages threads with current state of each nn @gabe %+v %+v %+v %+v", newOutageNNs, clearedNNs, prevImpactedNNs, currentImpactedNNs)
}
