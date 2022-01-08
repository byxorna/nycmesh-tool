// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for CreateSite

// register flags to command
func registerModelCreateSiteFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCreateSiteAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteContactEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteContactName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteContactPhone(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteElevation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteHeight(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteLocation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteNote(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteParentSiteID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteRegulatoryDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCreateSiteType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateSiteAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerCreateSiteContactEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contactEmailDescription := `Email of the contact person.`

	var contactEmailFlagName string
	if cmdPrefix == "" {
		contactEmailFlagName = "contactEmail"
	} else {
		contactEmailFlagName = fmt.Sprintf("%v.contactEmail", cmdPrefix)
	}

	var contactEmailFlagDefault string

	_ = cmd.PersistentFlags().String(contactEmailFlagName, contactEmailFlagDefault, contactEmailDescription)

	return nil
}

func registerCreateSiteContactName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contactNameDescription := `Name of the contact person.`

	var contactNameFlagName string
	if cmdPrefix == "" {
		contactNameFlagName = "contactName"
	} else {
		contactNameFlagName = fmt.Sprintf("%v.contactName", cmdPrefix)
	}

	var contactNameFlagDefault string

	_ = cmd.PersistentFlags().String(contactNameFlagName, contactNameFlagDefault, contactNameDescription)

	return nil
}

func registerCreateSiteContactPhone(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contactPhoneDescription := `Phone number of the contact person.`

	var contactPhoneFlagName string
	if cmdPrefix == "" {
		contactPhoneFlagName = "contactPhone"
	} else {
		contactPhoneFlagName = fmt.Sprintf("%v.contactPhone", cmdPrefix)
	}

	var contactPhoneFlagDefault string

	_ = cmd.PersistentFlags().String(contactPhoneFlagName, contactPhoneFlagDefault, contactPhoneDescription)

	return nil
}

func registerCreateSiteElevation(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerCreateSiteHeight(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerCreateSiteLocation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var locationFlagName string
	if cmdPrefix == "" {
		locationFlagName = "location"
	} else {
		locationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
	}

	if err := registerModelCreateSiteLocationFlags(depth+1, locationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCreateSiteName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. Name of the site.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerCreateSiteNote(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerCreateSiteParentSiteID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	parentSiteIdDescription := `ID of the parent site (deprecated).`

	var parentSiteIdFlagName string
	if cmdPrefix == "" {
		parentSiteIdFlagName = "parentSiteId"
	} else {
		parentSiteIdFlagName = fmt.Sprintf("%v.parentSiteId", cmdPrefix)
	}

	var parentSiteIdFlagDefault string

	_ = cmd.PersistentFlags().String(parentSiteIdFlagName, parentSiteIdFlagDefault, parentSiteIdDescription)

	return nil
}

func registerCreateSiteRegulatoryDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerCreateSiteType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["site","endpoint","client"]. Type of the site.`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string = "site"

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["site","endpoint","client"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCreateSiteFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveCreateSiteAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, contactEmailAdded := retrieveCreateSiteContactEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contactEmailAdded

	err, contactNameAdded := retrieveCreateSiteContactNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contactNameAdded

	err, contactPhoneAdded := retrieveCreateSiteContactPhoneFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contactPhoneAdded

	err, elevationAdded := retrieveCreateSiteElevationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || elevationAdded

	err, heightAdded := retrieveCreateSiteHeightFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || heightAdded

	err, locationAdded := retrieveCreateSiteLocationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded

	err, nameAdded := retrieveCreateSiteNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, noteAdded := retrieveCreateSiteNoteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noteAdded

	err, parentSiteIdAdded := retrieveCreateSiteParentSiteIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentSiteIdAdded

	err, regulatoryDomainAdded := retrieveCreateSiteRegulatoryDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regulatoryDomainAdded

	err, typeAdded := retrieveCreateSiteTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveCreateSiteAddressFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCreateSiteContactEmailFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contactEmailFlagName := fmt.Sprintf("%v.contactEmail", cmdPrefix)
	if cmd.Flags().Changed(contactEmailFlagName) {

		var contactEmailFlagName string
		if cmdPrefix == "" {
			contactEmailFlagName = "contactEmail"
		} else {
			contactEmailFlagName = fmt.Sprintf("%v.contactEmail", cmdPrefix)
		}

		contactEmailFlagValue, err := cmd.Flags().GetString(contactEmailFlagName)
		if err != nil {
			return err, false
		}
		m.ContactEmail = contactEmailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateSiteContactNameFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contactNameFlagName := fmt.Sprintf("%v.contactName", cmdPrefix)
	if cmd.Flags().Changed(contactNameFlagName) {

		var contactNameFlagName string
		if cmdPrefix == "" {
			contactNameFlagName = "contactName"
		} else {
			contactNameFlagName = fmt.Sprintf("%v.contactName", cmdPrefix)
		}

		contactNameFlagValue, err := cmd.Flags().GetString(contactNameFlagName)
		if err != nil {
			return err, false
		}
		m.ContactName = contactNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateSiteContactPhoneFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contactPhoneFlagName := fmt.Sprintf("%v.contactPhone", cmdPrefix)
	if cmd.Flags().Changed(contactPhoneFlagName) {

		var contactPhoneFlagName string
		if cmdPrefix == "" {
			contactPhoneFlagName = "contactPhone"
		} else {
			contactPhoneFlagName = fmt.Sprintf("%v.contactPhone", cmdPrefix)
		}

		contactPhoneFlagValue, err := cmd.Flags().GetString(contactPhoneFlagName)
		if err != nil {
			return err, false
		}
		m.ContactPhone = contactPhoneFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateSiteElevationFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCreateSiteHeightFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCreateSiteLocationFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	locationFlagName := fmt.Sprintf("%v.location", cmdPrefix)
	if cmd.Flags().Changed(locationFlagName) {
		// info: complex object location CreateSiteLocation is retrieved outside this Changed() block
	}
	locationFlagValue := m.Location
	if swag.IsZero(locationFlagValue) {
		locationFlagValue = &models.CreateSiteLocation{}
	}

	err, locationAdded := retrieveModelCreateSiteLocationFlags(depth+1, locationFlagValue, locationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded
	if locationAdded {
		m.Location = locationFlagValue
	}

	return nil, retAdded
}

func retrieveCreateSiteNameFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateSiteNoteFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCreateSiteParentSiteIDFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentSiteIdFlagName := fmt.Sprintf("%v.parentSiteId", cmdPrefix)
	if cmd.Flags().Changed(parentSiteIdFlagName) {

		var parentSiteIdFlagName string
		if cmdPrefix == "" {
			parentSiteIdFlagName = "parentSiteId"
		} else {
			parentSiteIdFlagName = fmt.Sprintf("%v.parentSiteId", cmdPrefix)
		}

		parentSiteIdFlagValue, err := cmd.Flags().GetString(parentSiteIdFlagName)
		if err != nil {
			return err, false
		}
		m.ParentSiteID = parentSiteIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCreateSiteRegulatoryDomainFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCreateSiteTypeFlags(depth int, m *models.CreateSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}