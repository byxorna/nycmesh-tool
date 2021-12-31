// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for SiteDescription1

// register flags to command
func registerModelSiteDescription1Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSiteDescription1Address(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1Contact(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1DeviceCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1DeviceListStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1Elevation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1Endpoints(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1Height(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1IPAddresses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1Location(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1Note(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1RegulatoryDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1SLA(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteDescription1UcrmID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteDescription1Address(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := `Address of the site.`

	var addressFlagName string
	if cmdPrefix == "" {
		addressFlagName = "address"
	} else {
		addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
	}

	var addressFlagDefault string

	_ = cmd.PersistentFlags().String(addressFlagName, addressFlagDefault, addressDescription)

	return nil
}

func registerSiteDescription1Contact(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var contactFlagName string
	if cmdPrefix == "" {
		contactFlagName = "contact"
	} else {
		contactFlagName = fmt.Sprintf("%v.contact", cmdPrefix)
	}

	if err := registerModelSiteDescriptionContactFlags(depth+1, contactFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteDescription1DeviceCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceCountDescription := `Number of devices in this site.`

	var deviceCountFlagName string
	if cmdPrefix == "" {
		deviceCountFlagName = "deviceCount"
	} else {
		deviceCountFlagName = fmt.Sprintf("%v.deviceCount", cmdPrefix)
	}

	var deviceCountFlagDefault float64

	_ = cmd.PersistentFlags().Float64(deviceCountFlagName, deviceCountFlagDefault, deviceCountDescription)

	return nil
}

func registerSiteDescription1DeviceListStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deviceListStatusDescription := `Enum: ["active","disconnected","inactive"]. Deprecated. Use site.identification.status instead.`

	var deviceListStatusFlagName string
	if cmdPrefix == "" {
		deviceListStatusFlagName = "deviceListStatus"
	} else {
		deviceListStatusFlagName = fmt.Sprintf("%v.deviceListStatus", cmdPrefix)
	}

	var deviceListStatusFlagDefault string

	_ = cmd.PersistentFlags().String(deviceListStatusFlagName, deviceListStatusFlagDefault, deviceListStatusDescription)

	if err := cmd.RegisterFlagCompletionFunc(deviceListStatusFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["active","disconnected","inactive"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSiteDescription1Elevation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	elevationDescription := `Site elevation without structure height.`

	var elevationFlagName string
	if cmdPrefix == "" {
		elevationFlagName = "elevation"
	} else {
		elevationFlagName = fmt.Sprintf("%v.elevation", cmdPrefix)
	}

	var elevationFlagDefault float64

	_ = cmd.PersistentFlags().Float64(elevationFlagName, elevationFlagDefault, elevationDescription)

	return nil
}

func registerSiteDescription1Endpoints(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: endpoints SiteEndpointsList array type is not supported by go-swagger cli yet

	return nil
}

func registerSiteDescription1Height(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	heightDescription := `Site structure height.`

	var heightFlagName string
	if cmdPrefix == "" {
		heightFlagName = "height"
	} else {
		heightFlagName = fmt.Sprintf("%v.height", cmdPrefix)
	}

	var heightFlagDefault float64

	_ = cmd.PersistentFlags().Float64(heightFlagName, heightFlagDefault, heightDescription)

	return nil
}

func registerSiteDescription1IPAddresses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ipAddresses IPAddresses array type is not supported by go-swagger cli yet

	return nil
}

func registerSiteDescription1Location(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var locationFlagName string
	if cmdPrefix == "" {
		locationFlagName = "location"
	} else {
		locationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
	}

	if err := registerModelSiteDescriptionLocationFlags(depth+1, locationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteDescription1Note(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	noteDescription := `Any additional site description.`

	var noteFlagName string
	if cmdPrefix == "" {
		noteFlagName = "note"
	} else {
		noteFlagName = fmt.Sprintf("%v.note", cmdPrefix)
	}

	var noteFlagDefault string

	_ = cmd.PersistentFlags().String(noteFlagName, noteFlagDefault, noteDescription)

	return nil
}

func registerSiteDescription1RegulatoryDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regulatoryDomainDescription := `Enum: ["XX","XY","AF","AX","AL","DZ","AS","AD","AO","AI","AQ","AG","AR","AM","AW","AU","AT","AZ","BS","BH","BD","BB","BY","BE","BZ","BJ","BM","BT","BO","BQ","BA","BW","BV","BR","IO","BN","BG","BF","BI","CV","KH","CM","CA","KY","CF","TD","CL","CN","CX","CC","CO","KM","CK","CR","CI","HR","CW","CY","CZ","CD","DK","DJ","DM","DO","EC","EG","SV","GQ","ER","EE","SZ","ET","FK","FO","FJ","FI","FR","GF","PF","TF","GA","GM","GE","DE","GH","GI","GR","GL","GD","GP","GU","GT","GG","GN","GW","GY","HT","HM","HN","HK","HU","IS","IN","ID","IQ","IE","IM","IL","IT","JM","JP","JE","JO","KZ","KE","KI","KW","KG","LA","LV","LB","LS","LR","LY","LI","LT","LU","MO","MK","MG","MW","MY","MV","ML","MT","MH","MQ","MR","MU","YT","MX","FM","MD","MC","MN","ME","MS","MA","MZ","MM","NA","NR","NP","NL","NC","NZ","NI","NE","NG","NU","NF","MP","NO","OM","PK","PW","PA","PG","PY","PE","PH","PN","PL","PT","PR","QA","KR","RS","SC","CG","RE","RO","RU","RW","BL","SH","KN","LC","MF","PM","VC","WS","SM","ST","SA","SN","SL","SG","SX","SK","SI","SB","SO","ZA","GS","SS","ES","LK","PS","SR","SJ","SE","CH","TW","TJ","TZ","TH","TL","TG","TK","TO","TT","TN","TR","TM","TC","TV","UG","UA","AE","GB","US","UY","UZ","VU","VA","VE","VN","VG","VI","WF","EH","YE","ZM","ZW"]. Location of regulation.`

	var regulatoryDomainFlagName string
	if cmdPrefix == "" {
		regulatoryDomainFlagName = "regulatoryDomain"
	} else {
		regulatoryDomainFlagName = fmt.Sprintf("%v.regulatoryDomain", cmdPrefix)
	}

	var regulatoryDomainFlagDefault string

	_ = cmd.PersistentFlags().String(regulatoryDomainFlagName, regulatoryDomainFlagDefault, regulatoryDomainDescription)

	if err := cmd.RegisterFlagCompletionFunc(regulatoryDomainFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["XX","XY","AF","AX","AL","DZ","AS","AD","AO","AI","AQ","AG","AR","AM","AW","AU","AT","AZ","BS","BH","BD","BB","BY","BE","BZ","BJ","BM","BT","BO","BQ","BA","BW","BV","BR","IO","BN","BG","BF","BI","CV","KH","CM","CA","KY","CF","TD","CL","CN","CX","CC","CO","KM","CK","CR","CI","HR","CW","CY","CZ","CD","DK","DJ","DM","DO","EC","EG","SV","GQ","ER","EE","SZ","ET","FK","FO","FJ","FI","FR","GF","PF","TF","GA","GM","GE","DE","GH","GI","GR","GL","GD","GP","GU","GT","GG","GN","GW","GY","HT","HM","HN","HK","HU","IS","IN","ID","IQ","IE","IM","IL","IT","JM","JP","JE","JO","KZ","KE","KI","KW","KG","LA","LV","LB","LS","LR","LY","LI","LT","LU","MO","MK","MG","MW","MY","MV","ML","MT","MH","MQ","MR","MU","YT","MX","FM","MD","MC","MN","ME","MS","MA","MZ","MM","NA","NR","NP","NL","NC","NZ","NI","NE","NG","NU","NF","MP","NO","OM","PK","PW","PA","PG","PY","PE","PH","PN","PL","PT","PR","QA","KR","RS","SC","CG","RE","RO","RU","RW","BL","SH","KN","LC","MF","PM","VC","WS","SM","ST","SA","SN","SL","SG","SX","SK","SI","SB","SO","ZA","GS","SS","ES","LK","PS","SR","SJ","SE","CH","TW","TJ","TZ","TH","TL","TG","TK","TO","TT","TN","TR","TM","TC","TV","UG","UA","AE","GB","US","UY","UZ","VU","VA","VE","VN","VG","VI","WF","EH","YE","ZM","ZW"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerSiteDescription1SLA(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slaDescription := ``

	var slaFlagName string
	if cmdPrefix == "" {
		slaFlagName = "sla"
	} else {
		slaFlagName = fmt.Sprintf("%v.sla", cmdPrefix)
	}

	var slaFlagDefault float64

	_ = cmd.PersistentFlags().Float64(slaFlagName, slaFlagDefault, slaDescription)

	return nil
}

func registerSiteDescription1UcrmID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ucrmIdDescription := `ID of UCRM client bound with this site. Null if no UCRM client is bound to this site.`

	var ucrmIdFlagName string
	if cmdPrefix == "" {
		ucrmIdFlagName = "ucrmId"
	} else {
		ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
	}

	var ucrmIdFlagDefault string

	_ = cmd.PersistentFlags().String(ucrmIdFlagName, ucrmIdFlagDefault, ucrmIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSiteDescription1Flags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveSiteDescription1AddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, contactAdded := retrieveSiteDescription1ContactFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contactAdded

	err, deviceCountAdded := retrieveSiteDescription1DeviceCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceCountAdded

	err, deviceListStatusAdded := retrieveSiteDescription1DeviceListStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceListStatusAdded

	err, elevationAdded := retrieveSiteDescription1ElevationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || elevationAdded

	err, endpointsAdded := retrieveSiteDescription1EndpointsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointsAdded

	err, heightAdded := retrieveSiteDescription1HeightFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || heightAdded

	err, ipAddressesAdded := retrieveSiteDescription1IPAddressesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAddressesAdded

	err, locationAdded := retrieveSiteDescription1LocationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded

	err, noteAdded := retrieveSiteDescription1NoteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noteAdded

	err, regulatoryDomainAdded := retrieveSiteDescription1RegulatoryDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regulatoryDomainAdded

	err, slaAdded := retrieveSiteDescription1SLAFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slaAdded

	err, ucrmIdAdded := retrieveSiteDescription1UcrmIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ucrmIdAdded

	return nil, retAdded
}

func retrieveSiteDescription1AddressFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressFlagName := fmt.Sprintf("%v.address", cmdPrefix)
	if cmd.Flags().Changed(addressFlagName) {

		var addressFlagName string
		if cmdPrefix == "" {
			addressFlagName = "address"
		} else {
			addressFlagName = fmt.Sprintf("%v.address", cmdPrefix)
		}

		addressFlagValue, err := cmd.Flags().GetString(addressFlagName)
		if err != nil {
			return err, false
		}
		m.Address = addressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1ContactFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contactFlagName := fmt.Sprintf("%v.contact", cmdPrefix)
	if cmd.Flags().Changed(contactFlagName) {
		// info: complex object contact SiteDescriptionContact is retrieved outside this Changed() block
	}
	contactFlagValue := m.Contact
	if swag.IsZero(contactFlagValue) {
		contactFlagValue = &models.SiteDescriptionContact{}
	}

	err, contactAdded := retrieveModelSiteDescriptionContactFlags(depth+1, contactFlagValue, contactFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contactAdded
	if contactAdded {
		m.Contact = contactFlagValue
	}

	return nil, retAdded
}

func retrieveSiteDescription1DeviceCountFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceCountFlagName := fmt.Sprintf("%v.deviceCount", cmdPrefix)
	if cmd.Flags().Changed(deviceCountFlagName) {

		var deviceCountFlagName string
		if cmdPrefix == "" {
			deviceCountFlagName = "deviceCount"
		} else {
			deviceCountFlagName = fmt.Sprintf("%v.deviceCount", cmdPrefix)
		}

		deviceCountFlagValue, err := cmd.Flags().GetFloat64(deviceCountFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceCount = deviceCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1DeviceListStatusFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceListStatusFlagName := fmt.Sprintf("%v.deviceListStatus", cmdPrefix)
	if cmd.Flags().Changed(deviceListStatusFlagName) {

		var deviceListStatusFlagName string
		if cmdPrefix == "" {
			deviceListStatusFlagName = "deviceListStatus"
		} else {
			deviceListStatusFlagName = fmt.Sprintf("%v.deviceListStatus", cmdPrefix)
		}

		deviceListStatusFlagValue, err := cmd.Flags().GetString(deviceListStatusFlagName)
		if err != nil {
			return err, false
		}
		m.DeviceListStatus = deviceListStatusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1ElevationFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	elevationFlagName := fmt.Sprintf("%v.elevation", cmdPrefix)
	if cmd.Flags().Changed(elevationFlagName) {

		var elevationFlagName string
		if cmdPrefix == "" {
			elevationFlagName = "elevation"
		} else {
			elevationFlagName = fmt.Sprintf("%v.elevation", cmdPrefix)
		}

		elevationFlagValue, err := cmd.Flags().GetFloat64(elevationFlagName)
		if err != nil {
			return err, false
		}
		m.Elevation = elevationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1EndpointsFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointsFlagName := fmt.Sprintf("%v.endpoints", cmdPrefix)
	if cmd.Flags().Changed(endpointsFlagName) {
		// warning: endpoints array type SiteEndpointsList is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSiteDescription1HeightFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	heightFlagName := fmt.Sprintf("%v.height", cmdPrefix)
	if cmd.Flags().Changed(heightFlagName) {

		var heightFlagName string
		if cmdPrefix == "" {
			heightFlagName = "height"
		} else {
			heightFlagName = fmt.Sprintf("%v.height", cmdPrefix)
		}

		heightFlagValue, err := cmd.Flags().GetFloat64(heightFlagName)
		if err != nil {
			return err, false
		}
		m.Height = heightFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1IPAddressesFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipAddressesFlagName := fmt.Sprintf("%v.ipAddresses", cmdPrefix)
	if cmd.Flags().Changed(ipAddressesFlagName) {
		// warning: ipAddresses array type IPAddresses is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSiteDescription1LocationFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	locationFlagName := fmt.Sprintf("%v.location", cmdPrefix)
	if cmd.Flags().Changed(locationFlagName) {
		// info: complex object location SiteDescriptionLocation is retrieved outside this Changed() block
	}
	locationFlagValue := m.Location
	if swag.IsZero(locationFlagValue) {
		locationFlagValue = &models.SiteDescriptionLocation{}
	}

	err, locationAdded := retrieveModelSiteDescriptionLocationFlags(depth+1, locationFlagValue, locationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded
	if locationAdded {
		m.Location = locationFlagValue
	}

	return nil, retAdded
}

func retrieveSiteDescription1NoteFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	noteFlagName := fmt.Sprintf("%v.note", cmdPrefix)
	if cmd.Flags().Changed(noteFlagName) {

		var noteFlagName string
		if cmdPrefix == "" {
			noteFlagName = "note"
		} else {
			noteFlagName = fmt.Sprintf("%v.note", cmdPrefix)
		}

		noteFlagValue, err := cmd.Flags().GetString(noteFlagName)
		if err != nil {
			return err, false
		}
		m.Note = noteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1RegulatoryDomainFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regulatoryDomainFlagName := fmt.Sprintf("%v.regulatoryDomain", cmdPrefix)
	if cmd.Flags().Changed(regulatoryDomainFlagName) {

		var regulatoryDomainFlagName string
		if cmdPrefix == "" {
			regulatoryDomainFlagName = "regulatoryDomain"
		} else {
			regulatoryDomainFlagName = fmt.Sprintf("%v.regulatoryDomain", cmdPrefix)
		}

		regulatoryDomainFlagValue, err := cmd.Flags().GetString(regulatoryDomainFlagName)
		if err != nil {
			return err, false
		}
		m.RegulatoryDomain = regulatoryDomainFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1SLAFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slaFlagName := fmt.Sprintf("%v.sla", cmdPrefix)
	if cmd.Flags().Changed(slaFlagName) {

		var slaFlagName string
		if cmdPrefix == "" {
			slaFlagName = "sla"
		} else {
			slaFlagName = fmt.Sprintf("%v.sla", cmdPrefix)
		}

		slaFlagValue, err := cmd.Flags().GetFloat64(slaFlagName)
		if err != nil {
			return err, false
		}
		m.SLA = slaFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteDescription1UcrmIDFlags(depth int, m *models.SiteDescription1, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ucrmIdFlagName := fmt.Sprintf("%v.ucrmId", cmdPrefix)
	if cmd.Flags().Changed(ucrmIdFlagName) {

		var ucrmIdFlagName string
		if cmdPrefix == "" {
			ucrmIdFlagName = "ucrmId"
		} else {
			ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
		}

		ucrmIdFlagValue, err := cmd.Flags().GetString(ucrmIdFlagName)
		if err != nil {
			return err, false
		}
		m.UcrmID = ucrmIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
