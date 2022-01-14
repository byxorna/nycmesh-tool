package meshapi

import (
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
	// TODO: move this cache to pkg/app.Application
	diskCache cache.DiskCache
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

// GeoURI is a utility format helper to create a geo URI string from lat/lon/alt
// TODO: find a better home for this?
func GeoURI(lat float64, long float64, altMeters float64) string {
	return fmt.Sprintf("geo:%f,%f,%.0f", lat, long, altMeters)
}
