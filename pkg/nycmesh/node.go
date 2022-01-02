package nycmesh

import (
	"fmt"
)

func (n *Node) GeoURI() string {
	return fmt.Sprintf("geo:%f,%f,%.0f", n.Latitude, n.Longitude, n.AltitudeMeters)
}
