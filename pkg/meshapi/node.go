package meshapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// https://github.com/meshcenter/mesh-api/blob/master/migrations/1608616875404_base.js#L27
type Node struct {
	ID             int        `json:"id"`
	Latitude       float64    `json:"lat"`
	Longitude      float64    `json:"lng"`
	AltitudeMeters float64    `json:"alt"`
	Status         string     `json:"status"`
	Location       string     `json:"location"`
	Name           string     `json:"name"`
	Notes          string     `json:"notes"`
	Created        time.Time  `json:"create_date"`
	Abandoned      *time.Time `json:"abandon_date,omitempty"`
	BuildingID     int        `json:"building_id"`
	Building       string     `json:"building"`
	MemberID       int        `json:"member_id"`
	Devices        []Device   `json:"devices"`
}

type NodesByID []Node

func (s NodesByID) Len() int {
	return len(s)
}

func (s NodesByID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s NodesByID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Client) Nodes() ([]Node, error) {
	var body []byte
	var err error

	body, err = c.do_body("GET", path+"/nodes", map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch nodes: %w", err)
	}

	var nodes []Node
	if err = json.Unmarshal(body, &nodes); err != nil {
		return nil, fmt.Errorf("unable to decode nodes: %w\n%s", err, body)
	}

	return nodes, nil
}
