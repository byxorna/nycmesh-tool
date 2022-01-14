package meshapi

import (
	"fmt"
)

func (c *Client) KML() ([]byte, error) {
	var body []byte
	var err error

	body, err = c.do_body("GET", path+"/kml", map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch kml: %w", err)
	}

	return body, nil
}
