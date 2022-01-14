package meshapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type Device struct {
	ID             int        `json:"id"`
	Latitude       float64    `json:"lat"`
	Longitude      float64    `json:"lng"`
	AltitudeMeters float64    `json:"alt"`
	Azimuth        int        `json:"azimuth"`
	Status         string     `json:"status"`
	Name           string     `json:"name"`
	SSID           string     `json:"ssid"`
	Notes          string     `json:"notes"`
	NodeID         int        `json:"node_id"`
	Created        time.Time  `json:"create_date"`
	Abandoned      *time.Time `json:"abandon_date,omitempty"`
	Type           DeviceType `json:"type"`

	// TODO
}

type DeviceType struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Manufacturer string  `json:"manufacturer"`
	Range        float64 `json:"range"`
	Width        float64 `json:"width"`
}

type DevicesByID []Device

func (s DevicesByID) Len() int {
	return len(s)
}

func (s DevicesByID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s DevicesByID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

// Devices returns the list of Devices from NYCMesh API
// TODO: should devices be cached/fetched as 1 block, or as individual entries?
func (c *Client) Devices() (map[int]*Device, error) {
	var body []byte
	var err error

	var devices []*Device
	if err = json.Unmarshal(body, &devices); err != nil {
		return nil, fmt.Errorf("unable to decode devices: %w\n%s", err, body)
	}

	m := map[int]*Device{}
	for _, dev := range devices {
		m[dev.ID] = dev
	}

	return m, nil
}
