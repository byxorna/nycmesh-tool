package app

import "fmt"

var (
	// TODO: make this mapping dynamic
	DefaultMonitoringChannel = "CSJK7P7C3" // #monitoring-unms
	NNToSlackChannel         = map[int]string{
		//713:  "CAZ4V9E0P",   // #sn3
		//713:  "C02HZLLH85R", // #hub-clay-5712
		5712: "C02HZLLH85R", // #hub-clay-5712
	}
)

func lookupSlackChannelIDFromNN(nn int) (string, error) {
	if ch, ok := NNToSlackChannel[nn]; ok {
		return ch, nil
	}
	return "", fmt.Errorf("implement me: lookup of slack channel for nn:%d", nn)
}
