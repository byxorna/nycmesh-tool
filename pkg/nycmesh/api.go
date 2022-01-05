package nycmesh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/byxorna/nycmesh-tool/pkg/cache"
)

var (
	host           = "api.nycmesh.net"
	path           = "v1"
	scheme         = "https"
	defaultTimeout = time.Second * 5
)

// wrapper around https://github.com/meshcenter/mesh-api
type Client struct {
	http.Client
	diskCache cache.DiskCache
}

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
	MemberID       int        `json:"member_id"`
	Devices        []Device   `json:"devices"`
}

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

func New() (*Client, error) {
	dc, err := cache.NewDiskVCache()
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: *&http.Client{
			Timeout: defaultTimeout,
		},
		diskCache: dc,
	}, nil
}

// TODO: this function should get cleaned up, we should not be storing
// nodes.json as a single blob. Instead, store each node, and a list of node IDs.
func (c *Client) Nodes() (map[int]Node, error) {
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

	m := map[int]Node{}
	for _, node := range nodes {
		m[node.ID] = node
	}

	return m, nil
}

func (c *Client) do_body(method, endpoint string, params map[string]string) ([]byte, error) {
	resp, err := c.do(method, endpoint, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s://%s/%s", scheme, host, endpoint)
	log.Printf("-> %s %s\n", method, baseURL)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.Client.Do(req)
}
