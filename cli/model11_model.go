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

// Schema cli for Model11

// register flags to command
func registerModelModel11Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel11AggregatedTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11Device(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11DeviceMetadata(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11EndTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11InProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11Ongoing(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11Site(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11StartTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel11Type(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel11AggregatedTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	aggregatedTimeDescription := `Required. `

	var aggregatedTimeFlagName string
	if cmdPrefix == "" {
		aggregatedTimeFlagName = "aggregatedTime"
	} else {
		aggregatedTimeFlagName = fmt.Sprintf("%v.aggregatedTime", cmdPrefix)
	}

	var aggregatedTimeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(aggregatedTimeFlagName, aggregatedTimeFlagDefault, aggregatedTimeDescription)

	return nil
}

func registerModel11Device(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var deviceFlagName string
	if cmdPrefix == "" {
		deviceFlagName = "device"
	} else {
		deviceFlagName = fmt.Sprintf("%v.device", cmdPrefix)
	}

	if err := registerModelDeviceOutageFlags(depth+1, deviceFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel11DeviceMetadata(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var deviceMetadataFlagName string
	if cmdPrefix == "" {
		deviceMetadataFlagName = "deviceMetadata"
	} else {
		deviceMetadataFlagName = fmt.Sprintf("%v.deviceMetadata", cmdPrefix)
	}

	if err := registerModelOutageDeviceMetadataFlags(depth+1, deviceMetadataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel11EndTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endTimestampDescription := `Required. `

	var endTimestampFlagName string
	if cmdPrefix == "" {
		endTimestampFlagName = "endTimestamp"
	} else {
		endTimestampFlagName = fmt.Sprintf("%v.endTimestamp", cmdPrefix)
	}

	var endTimestampFlagDefault string

	_ = cmd.PersistentFlags().String(endTimestampFlagName, endTimestampFlagDefault, endTimestampDescription)

	return nil
}

func registerModel11ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerModel11InProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	inProgressDescription := `Required. `

	var inProgressFlagName string
	if cmdPrefix == "" {
		inProgressFlagName = "inProgress"
	} else {
		inProgressFlagName = fmt.Sprintf("%v.inProgress", cmdPrefix)
	}

	var inProgressFlagDefault bool

	_ = cmd.PersistentFlags().Bool(inProgressFlagName, inProgressFlagDefault, inProgressDescription)

	return nil
}

func registerModel11Ongoing(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ongoingDescription := `Required. `

	var ongoingFlagName string
	if cmdPrefix == "" {
		ongoingFlagName = "ongoing"
	} else {
		ongoingFlagName = fmt.Sprintf("%v.ongoing", cmdPrefix)
	}

	var ongoingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ongoingFlagName, ongoingFlagDefault, ongoingDescription)

	return nil
}

func registerModel11Site(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var siteFlagName string
	if cmdPrefix == "" {
		siteFlagName = "site"
	} else {
		siteFlagName = fmt.Sprintf("%v.site", cmdPrefix)
	}

	if err := registerModelSiteIdentificationFlags(depth+1, siteFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel11StartTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startTimestampDescription := `Required. `

	var startTimestampFlagName string
	if cmdPrefix == "" {
		startTimestampFlagName = "startTimestamp"
	} else {
		startTimestampFlagName = fmt.Sprintf("%v.startTimestamp", cmdPrefix)
	}

	var startTimestampFlagDefault string

	_ = cmd.PersistentFlags().String(startTimestampFlagName, startTimestampFlagDefault, startTimestampDescription)

	return nil
}

func registerModel11Type(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["outage","unreachable"]. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["outage","unreachable"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel11Flags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aggregatedTimeAdded := retrieveModel11AggregatedTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aggregatedTimeAdded

	err, deviceAdded := retrieveModel11DeviceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded

	err, deviceMetadataAdded := retrieveModel11DeviceMetadataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceMetadataAdded

	err, endTimestampAdded := retrieveModel11EndTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endTimestampAdded

	err, idAdded := retrieveModel11IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, inProgressAdded := retrieveModel11InProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || inProgressAdded

	err, ongoingAdded := retrieveModel11OngoingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ongoingAdded

	err, siteAdded := retrieveModel11SiteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded

	err, startTimestampAdded := retrieveModel11StartTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startTimestampAdded

	err, typeAdded := retrieveModel11TypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveModel11AggregatedTimeFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	aggregatedTimeFlagName := fmt.Sprintf("%v.aggregatedTime", cmdPrefix)
	if cmd.Flags().Changed(aggregatedTimeFlagName) {

		var aggregatedTimeFlagName string
		if cmdPrefix == "" {
			aggregatedTimeFlagName = "aggregatedTime"
		} else {
			aggregatedTimeFlagName = fmt.Sprintf("%v.aggregatedTime", cmdPrefix)
		}

		aggregatedTimeFlagValue, err := cmd.Flags().GetFloat64(aggregatedTimeFlagName)
		if err != nil {
			return err, false
		}
		m.AggregatedTime = &aggregatedTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel11DeviceFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceFlagName := fmt.Sprintf("%v.device", cmdPrefix)
	if cmd.Flags().Changed(deviceFlagName) {
		// info: complex object device DeviceOutage is retrieved outside this Changed() block
	}
	deviceFlagValue := m.Device
	if swag.IsZero(deviceFlagValue) {
		deviceFlagValue = &models.DeviceOutage{}
	}

	err, deviceAdded := retrieveModelDeviceOutageFlags(depth+1, deviceFlagValue, deviceFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceAdded
	if deviceAdded {
		m.Device = deviceFlagValue
	}

	return nil, retAdded
}

func retrieveModel11DeviceMetadataFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceMetadataFlagName := fmt.Sprintf("%v.deviceMetadata", cmdPrefix)
	if cmd.Flags().Changed(deviceMetadataFlagName) {
		// info: complex object deviceMetadata OutageDeviceMetadata is retrieved outside this Changed() block
	}
	deviceMetadataFlagValue := m.DeviceMetadata
	if swag.IsZero(deviceMetadataFlagValue) {
		deviceMetadataFlagValue = &models.OutageDeviceMetadata{}
	}

	err, deviceMetadataAdded := retrieveModelOutageDeviceMetadataFlags(depth+1, deviceMetadataFlagValue, deviceMetadataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceMetadataAdded
	if deviceMetadataAdded {
		m.DeviceMetadata = deviceMetadataFlagValue
	}

	return nil, retAdded
}

func retrieveModel11EndTimestampFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endTimestampFlagName := fmt.Sprintf("%v.endTimestamp", cmdPrefix)
	if cmd.Flags().Changed(endTimestampFlagName) {

		var endTimestampFlagName string
		if cmdPrefix == "" {
			endTimestampFlagName = "endTimestamp"
		} else {
			endTimestampFlagName = fmt.Sprintf("%v.endTimestamp", cmdPrefix)
		}

		endTimestampFlagValue, err := cmd.Flags().GetString(endTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.EndTimestamp = &endTimestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel11IDFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = &idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel11InProgressFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	inProgressFlagName := fmt.Sprintf("%v.inProgress", cmdPrefix)
	if cmd.Flags().Changed(inProgressFlagName) {

		var inProgressFlagName string
		if cmdPrefix == "" {
			inProgressFlagName = "inProgress"
		} else {
			inProgressFlagName = fmt.Sprintf("%v.inProgress", cmdPrefix)
		}

		inProgressFlagValue, err := cmd.Flags().GetBool(inProgressFlagName)
		if err != nil {
			return err, false
		}
		m.InProgress = &inProgressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel11OngoingFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ongoingFlagName := fmt.Sprintf("%v.ongoing", cmdPrefix)
	if cmd.Flags().Changed(ongoingFlagName) {

		var ongoingFlagName string
		if cmdPrefix == "" {
			ongoingFlagName = "ongoing"
		} else {
			ongoingFlagName = fmt.Sprintf("%v.ongoing", cmdPrefix)
		}

		ongoingFlagValue, err := cmd.Flags().GetBool(ongoingFlagName)
		if err != nil {
			return err, false
		}
		m.Ongoing = &ongoingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel11SiteFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	siteFlagName := fmt.Sprintf("%v.site", cmdPrefix)
	if cmd.Flags().Changed(siteFlagName) {
		// info: complex object site SiteIdentification is retrieved outside this Changed() block
	}
	siteFlagValue := m.Site
	if swag.IsZero(siteFlagValue) {
		siteFlagValue = &models.SiteIdentification{}
	}

	err, siteAdded := retrieveModelSiteIdentificationFlags(depth+1, siteFlagValue, siteFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || siteAdded
	if siteAdded {
		m.Site = siteFlagValue
	}

	return nil, retAdded
}

func retrieveModel11StartTimestampFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startTimestampFlagName := fmt.Sprintf("%v.startTimestamp", cmdPrefix)
	if cmd.Flags().Changed(startTimestampFlagName) {

		var startTimestampFlagName string
		if cmdPrefix == "" {
			startTimestampFlagName = "startTimestamp"
		} else {
			startTimestampFlagName = fmt.Sprintf("%v.startTimestamp", cmdPrefix)
		}

		startTimestampFlagValue, err := cmd.Flags().GetString(startTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.StartTimestamp = &startTimestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel11TypeFlags(depth int, m *models.Model11, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
