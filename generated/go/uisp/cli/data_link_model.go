// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DataLink

// register flags to command
func registerModelDataLinkFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDataLinkCanDelete(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkDevicesDistance(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkFrequency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkFrom(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkOrigin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkSignal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkSitesDistance(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkSsid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkTo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataLinkType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataLinkCanDelete(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canDeleteDescription := `Required. `

	var canDeleteFlagName string
	if cmdPrefix == "" {
		canDeleteFlagName = "canDelete"
	} else {
		canDeleteFlagName = fmt.Sprintf("%v.canDelete", cmdPrefix)
	}

	var canDeleteFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canDeleteFlagName, canDeleteFlagDefault, canDeleteDescription)

	return nil
}

func registerDataLinkDevicesDistance(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	devicesDistanceDescription := `Nullable property in meters.`

	var devicesDistanceFlagName string
	if cmdPrefix == "" {
		devicesDistanceFlagName = "devicesDistance"
	} else {
		devicesDistanceFlagName = fmt.Sprintf("%v.devicesDistance", cmdPrefix)
	}

	var devicesDistanceFlagDefault float64

	_ = cmd.PersistentFlags().Float64(devicesDistanceFlagName, devicesDistanceFlagDefault, devicesDistanceDescription)

	return nil
}

func registerDataLinkFrequency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	frequencyDescription := `Required. `

	var frequencyFlagName string
	if cmdPrefix == "" {
		frequencyFlagName = "frequency"
	} else {
		frequencyFlagName = fmt.Sprintf("%v.frequency", cmdPrefix)
	}

	var frequencyFlagDefault float64

	_ = cmd.PersistentFlags().Float64(frequencyFlagName, frequencyFlagDefault, frequencyDescription)

	return nil
}

func registerDataLinkFrom(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var fromFlagName string
	if cmdPrefix == "" {
		fromFlagName = "from"
	} else {
		fromFlagName = fmt.Sprintf("%v.from", cmdPrefix)
	}

	if err := registerModelFromFlags(depth+1, fromFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataLinkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerDataLinkOrigin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	originDescription := `Enum: ["unms","manual"]. Required. Generated by UISP or manually added`

	var originFlagName string
	if cmdPrefix == "" {
		originFlagName = "origin"
	} else {
		originFlagName = fmt.Sprintf("%v.origin", cmdPrefix)
	}

	var originFlagDefault string

	_ = cmd.PersistentFlags().String(originFlagName, originFlagDefault, originDescription)

	if err := cmd.RegisterFlagCompletionFunc(originFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["unms","manual"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDataLinkSignal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	signalDescription := `Required. `

	var signalFlagName string
	if cmdPrefix == "" {
		signalFlagName = "signal"
	} else {
		signalFlagName = fmt.Sprintf("%v.signal", cmdPrefix)
	}

	var signalFlagDefault float64

	_ = cmd.PersistentFlags().Float64(signalFlagName, signalFlagDefault, signalDescription)

	return nil
}

func registerDataLinkSitesDistance(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sitesDistanceDescription := `Nullable property in meters.`

	var sitesDistanceFlagName string
	if cmdPrefix == "" {
		sitesDistanceFlagName = "sitesDistance"
	} else {
		sitesDistanceFlagName = fmt.Sprintf("%v.sitesDistance", cmdPrefix)
	}

	var sitesDistanceFlagDefault float64

	_ = cmd.PersistentFlags().Float64(sitesDistanceFlagName, sitesDistanceFlagDefault, sitesDistanceDescription)

	return nil
}

func registerDataLinkSsid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ssidDescription := `Required. `

	var ssidFlagName string
	if cmdPrefix == "" {
		ssidFlagName = "ssid"
	} else {
		ssidFlagName = fmt.Sprintf("%v.ssid", cmdPrefix)
	}

	var ssidFlagDefault string

	_ = cmd.PersistentFlags().String(ssidFlagName, ssidFlagDefault, ssidDescription)

	return nil
}

func registerDataLinkState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stateDescription := `Enum: ["active","disabled","disconnected","proposed"]. `

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	if err := cmd.RegisterFlagCompletionFunc(stateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["active","disabled","disconnected","proposed"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerDataLinkTo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var toFlagName string
	if cmdPrefix == "" {
		toFlagName = "to"
	} else {
		toFlagName = fmt.Sprintf("%v.to", cmdPrefix)
	}

	if err := registerModelToFlags(depth+1, toFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataLinkType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["ethernet","pon","wireless"]. `

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
			if err := json.Unmarshal([]byte(`["ethernet","pon","wireless"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDataLinkFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, canDeleteAdded := retrieveDataLinkCanDeleteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canDeleteAdded

	err, devicesDistanceAdded := retrieveDataLinkDevicesDistanceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || devicesDistanceAdded

	err, frequencyAdded := retrieveDataLinkFrequencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frequencyAdded

	err, fromAdded := retrieveDataLinkFromFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fromAdded

	err, idAdded := retrieveDataLinkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, originAdded := retrieveDataLinkOriginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || originAdded

	err, signalAdded := retrieveDataLinkSignalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || signalAdded

	err, sitesDistanceAdded := retrieveDataLinkSitesDistanceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sitesDistanceAdded

	err, ssidAdded := retrieveDataLinkSsidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ssidAdded

	err, stateAdded := retrieveDataLinkStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stateAdded

	err, toAdded := retrieveDataLinkToFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || toAdded

	err, typeAdded := retrieveDataLinkTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveDataLinkCanDeleteFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	canDeleteFlagName := fmt.Sprintf("%v.canDelete", cmdPrefix)
	if cmd.Flags().Changed(canDeleteFlagName) {

		var canDeleteFlagName string
		if cmdPrefix == "" {
			canDeleteFlagName = "canDelete"
		} else {
			canDeleteFlagName = fmt.Sprintf("%v.canDelete", cmdPrefix)
		}

		canDeleteFlagValue, err := cmd.Flags().GetBool(canDeleteFlagName)
		if err != nil {
			return err, false
		}
		m.CanDelete = &canDeleteFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkDevicesDistanceFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	devicesDistanceFlagName := fmt.Sprintf("%v.devicesDistance", cmdPrefix)
	if cmd.Flags().Changed(devicesDistanceFlagName) {

		var devicesDistanceFlagName string
		if cmdPrefix == "" {
			devicesDistanceFlagName = "devicesDistance"
		} else {
			devicesDistanceFlagName = fmt.Sprintf("%v.devicesDistance", cmdPrefix)
		}

		devicesDistanceFlagValue, err := cmd.Flags().GetFloat64(devicesDistanceFlagName)
		if err != nil {
			return err, false
		}
		m.DevicesDistance = devicesDistanceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkFrequencyFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	frequencyFlagName := fmt.Sprintf("%v.frequency", cmdPrefix)
	if cmd.Flags().Changed(frequencyFlagName) {

		var frequencyFlagName string
		if cmdPrefix == "" {
			frequencyFlagName = "frequency"
		} else {
			frequencyFlagName = fmt.Sprintf("%v.frequency", cmdPrefix)
		}

		frequencyFlagValue, err := cmd.Flags().GetFloat64(frequencyFlagName)
		if err != nil {
			return err, false
		}
		m.Frequency = &frequencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkFromFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fromFlagName := fmt.Sprintf("%v.from", cmdPrefix)
	if cmd.Flags().Changed(fromFlagName) {
		// info: complex object from From is retrieved outside this Changed() block
	}
	fromFlagValue := m.From
	if swag.IsZero(fromFlagValue) {
		fromFlagValue = &models.From{}
	}

	err, fromAdded := retrieveModelFromFlags(depth+1, fromFlagValue, fromFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fromAdded
	if fromAdded {
		m.From = fromFlagValue
	}

	return nil, retAdded
}

func retrieveDataLinkIDFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDataLinkOriginFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	originFlagName := fmt.Sprintf("%v.origin", cmdPrefix)
	if cmd.Flags().Changed(originFlagName) {

		var originFlagName string
		if cmdPrefix == "" {
			originFlagName = "origin"
		} else {
			originFlagName = fmt.Sprintf("%v.origin", cmdPrefix)
		}

		originFlagValue, err := cmd.Flags().GetString(originFlagName)
		if err != nil {
			return err, false
		}
		m.Origin = &originFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkSignalFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	signalFlagName := fmt.Sprintf("%v.signal", cmdPrefix)
	if cmd.Flags().Changed(signalFlagName) {

		var signalFlagName string
		if cmdPrefix == "" {
			signalFlagName = "signal"
		} else {
			signalFlagName = fmt.Sprintf("%v.signal", cmdPrefix)
		}

		signalFlagValue, err := cmd.Flags().GetFloat64(signalFlagName)
		if err != nil {
			return err, false
		}
		m.Signal = &signalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkSitesDistanceFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sitesDistanceFlagName := fmt.Sprintf("%v.sitesDistance", cmdPrefix)
	if cmd.Flags().Changed(sitesDistanceFlagName) {

		var sitesDistanceFlagName string
		if cmdPrefix == "" {
			sitesDistanceFlagName = "sitesDistance"
		} else {
			sitesDistanceFlagName = fmt.Sprintf("%v.sitesDistance", cmdPrefix)
		}

		sitesDistanceFlagValue, err := cmd.Flags().GetFloat64(sitesDistanceFlagName)
		if err != nil {
			return err, false
		}
		m.SitesDistance = sitesDistanceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkSsidFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ssidFlagName := fmt.Sprintf("%v.ssid", cmdPrefix)
	if cmd.Flags().Changed(ssidFlagName) {

		var ssidFlagName string
		if cmdPrefix == "" {
			ssidFlagName = "ssid"
		} else {
			ssidFlagName = fmt.Sprintf("%v.ssid", cmdPrefix)
		}

		ssidFlagValue, err := cmd.Flags().GetString(ssidFlagName)
		if err != nil {
			return err, false
		}
		m.Ssid = &ssidFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkStateFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stateFlagName := fmt.Sprintf("%v.state", cmdPrefix)
	if cmd.Flags().Changed(stateFlagName) {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = stateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataLinkToFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	toFlagName := fmt.Sprintf("%v.to", cmdPrefix)
	if cmd.Flags().Changed(toFlagName) {
		// info: complex object to To is retrieved outside this Changed() block
	}
	toFlagValue := m.To
	if swag.IsZero(toFlagValue) {
		toFlagValue = &models.To{}
	}

	err, toAdded := retrieveModelToFlags(depth+1, toFlagValue, toFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || toAdded
	if toAdded {
		m.To = toFlagValue
	}

	return nil, retAdded
}

func retrieveDataLinkTypeFlags(depth int, m *models.DataLink, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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