package nycmesh

import (
	"fmt"
)

func GeoURI(lat float64, long float64, altMeters float64) string {
	return fmt.Sprintf("geo:%f,%f,%.0f", lat, long, altMeters)
}
