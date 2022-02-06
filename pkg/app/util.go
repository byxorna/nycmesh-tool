package app

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// TODO: write unit tests for this function
// TODO: make this use other device fields (model?) to help identify
func GetNNFromDeviceName(name string) (int, error) {
	extracter := regexp.MustCompile(`nycmesh-(?:[^\d]+-)?(\d{3,6})\b`)
	if extracter.MatchString(name) {
		res := extracter.FindStringSubmatch(name)
		var match string
		switch len(res) {
		case 2:
			match = res[1]
		case 3:
			match = res[2]
		}
		nn, err := strconv.Atoi(match)
		if err != nil {
			return 0, fmt.Errorf("unable to parse %s to nn %s to int: %w", name, match, err)
		}
		return nn, nil
	}
	return 0, fmt.Errorf("unable to derive nn from %s", name)
}

func GetNNFromUISPDevice(d *models.DeviceStatusOverview) (int, error) {
	return GetNNFromDeviceName(d.Identification.Name)
}
