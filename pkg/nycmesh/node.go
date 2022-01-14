package nycmesh

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

// TODO: this function should get cleaned up, we should not be storing
// nodes.json as a single blob. Instead, store each node, and a list of node IDs.
func (c *Client) Nodes(ids ...int) (map[int]Node, error) {
	var body []byte
	var err error

	if !c.diskCache.Has("nodes") {
		body, err := c.do_body("GET", path+"/nodes", map[string]string{})
		if err != nil {
			return nil, fmt.Errorf("Unable to fetch nodes: %w", err)
		}

		err = c.diskCache.Write("nodes", body)
		if err != nil {
			return nil, err
		}
	}

	body, err = c.diskCache.Read("nodes")
	if err != nil {
		return nil, err
	}

	var nodes []Node
	if err = json.Unmarshal(body, &nodes); err != nil {
		return nil, fmt.Errorf("unable to decode nodes: %w\n%s", err, body)
	}

	idFilter := map[int]bool{}
	for _, id := range ids {
		idFilter[id] = true
	}

	m := map[int]Node{}
	for _, node := range nodes {
		if _, matchesFilter := idFilter[node.ID]; !matchesFilter && len(idFilter) > 0 {
			// only return nodes that match provided filter if necessary
			continue
		}
		m[node.ID] = node
	}

	return m, nil
}
