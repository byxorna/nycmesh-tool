// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DeviceStatistics

// register flags to command
func registerModelDeviceStatisticsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceStatisticsAirTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsCcq(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsCPU(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsDownlinkCapacity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsErrors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsInterfaces(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsInterval(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsLocalChain0(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsLocalChain1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsOutput(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsPeriod(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsPing(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsPsu(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsPv(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsRAM(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsRemoteChain0(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsRemoteChain1(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsRemoteSignal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsRemoteSignal60g(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsSignal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsSignal60g(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsStations(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsTemperature(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceStatisticsUplinkCapacity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatisticsAirTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: airTime AirTime array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsCcq(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ccq ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	clientsDescription := ``

	var clientsFlagName string
	if cmdPrefix == "" {
		clientsFlagName = "clients"
	} else {
		clientsFlagName = fmt.Sprintf("%v.clients", cmdPrefix)
	}

	var clientsFlagDefault string

	_ = cmd.PersistentFlags().String(clientsFlagName, clientsFlagDefault, clientsDescription)

	return nil
}

func registerDeviceStatisticsCPU(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: cpu ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsDownlinkCapacity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	downlinkCapacityDescription := ``

	var downlinkCapacityFlagName string
	if cmdPrefix == "" {
		downlinkCapacityFlagName = "downlinkCapacity"
	} else {
		downlinkCapacityFlagName = fmt.Sprintf("%v.downlinkCapacity", cmdPrefix)
	}

	var downlinkCapacityFlagDefault string

	_ = cmd.PersistentFlags().String(downlinkCapacityFlagName, downlinkCapacityFlagDefault, downlinkCapacityDescription)

	return nil
}

func registerDeviceStatisticsErrors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: errors ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsInterfaces(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: interfaces Interfaces array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsInterval(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var intervalFlagName string
	if cmdPrefix == "" {
		intervalFlagName = "interval"
	} else {
		intervalFlagName = fmt.Sprintf("%v.interval", cmdPrefix)
	}

	if err := registerModelIntervalFlags(depth+1, intervalFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceStatisticsLocalChain0(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: localChain0 LocalChain0 array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsLocalChain1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: localChain1 LocalChain1 array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsOutput(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	outputDescription := ``

	var outputFlagName string
	if cmdPrefix == "" {
		outputFlagName = "output"
	} else {
		outputFlagName = fmt.Sprintf("%v.output", cmdPrefix)
	}

	var outputFlagDefault string

	_ = cmd.PersistentFlags().String(outputFlagName, outputFlagDefault, outputDescription)

	return nil
}

func registerDeviceStatisticsPeriod(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	periodDescription := ``

	var periodFlagName string
	if cmdPrefix == "" {
		periodFlagName = "period"
	} else {
		periodFlagName = fmt.Sprintf("%v.period", cmdPrefix)
	}

	var periodFlagDefault float64

	_ = cmd.PersistentFlags().Float64(periodFlagName, periodFlagDefault, periodDescription)

	return nil
}

func registerDeviceStatisticsPing(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ping ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsPsu(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	psuDescription := ``

	var psuFlagName string
	if cmdPrefix == "" {
		psuFlagName = "psu"
	} else {
		psuFlagName = fmt.Sprintf("%v.psu", cmdPrefix)
	}

	var psuFlagDefault string

	_ = cmd.PersistentFlags().String(psuFlagName, psuFlagDefault, psuDescription)

	return nil
}

func registerDeviceStatisticsPv(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pvDescription := ``

	var pvFlagName string
	if cmdPrefix == "" {
		pvFlagName = "pv"
	} else {
		pvFlagName = fmt.Sprintf("%v.pv", cmdPrefix)
	}

	var pvFlagDefault string

	_ = cmd.PersistentFlags().String(pvFlagName, pvFlagDefault, pvDescription)

	return nil
}

func registerDeviceStatisticsRAM(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ram ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsRemoteChain0(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: remoteChain0 RemoteChain0 array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsRemoteChain1(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: remoteChain1 RemoteChain1 array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsRemoteSignal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: remoteSignal ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsRemoteSignal60g(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: remoteSignal60g ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsSignal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: signal ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsSignal60g(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: signal60g ListOfCoordinates array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsStations(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: stations Stations1 map type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsTemperature(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: temperature Temperature array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceStatisticsUplinkCapacity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	uplinkCapacityDescription := ``

	var uplinkCapacityFlagName string
	if cmdPrefix == "" {
		uplinkCapacityFlagName = "uplinkCapacity"
	} else {
		uplinkCapacityFlagName = fmt.Sprintf("%v.uplinkCapacity", cmdPrefix)
	}

	var uplinkCapacityFlagDefault string

	_ = cmd.PersistentFlags().String(uplinkCapacityFlagName, uplinkCapacityFlagDefault, uplinkCapacityDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceStatisticsFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, airTimeAdded := retrieveDeviceStatisticsAirTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || airTimeAdded

	err, ccqAdded := retrieveDeviceStatisticsCcqFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ccqAdded

	err, clientsAdded := retrieveDeviceStatisticsClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientsAdded

	err, cpuAdded := retrieveDeviceStatisticsCPUFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuAdded

	err, downlinkCapacityAdded := retrieveDeviceStatisticsDownlinkCapacityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || downlinkCapacityAdded

	err, errorsAdded := retrieveDeviceStatisticsErrorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorsAdded

	err, interfacesAdded := retrieveDeviceStatisticsInterfacesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfacesAdded

	err, intervalAdded := retrieveDeviceStatisticsIntervalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || intervalAdded

	err, localChain0Added := retrieveDeviceStatisticsLocalChain0Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || localChain0Added

	err, localChain1Added := retrieveDeviceStatisticsLocalChain1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || localChain1Added

	err, outputAdded := retrieveDeviceStatisticsOutputFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outputAdded

	err, periodAdded := retrieveDeviceStatisticsPeriodFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || periodAdded

	err, pingAdded := retrieveDeviceStatisticsPingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pingAdded

	err, psuAdded := retrieveDeviceStatisticsPsuFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || psuAdded

	err, pvAdded := retrieveDeviceStatisticsPvFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pvAdded

	err, ramAdded := retrieveDeviceStatisticsRAMFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ramAdded

	err, remoteChain0Added := retrieveDeviceStatisticsRemoteChain0Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || remoteChain0Added

	err, remoteChain1Added := retrieveDeviceStatisticsRemoteChain1Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || remoteChain1Added

	err, remoteSignalAdded := retrieveDeviceStatisticsRemoteSignalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || remoteSignalAdded

	err, remoteSignal60gAdded := retrieveDeviceStatisticsRemoteSignal60gFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || remoteSignal60gAdded

	err, signalAdded := retrieveDeviceStatisticsSignalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || signalAdded

	err, signal60gAdded := retrieveDeviceStatisticsSignal60gFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || signal60gAdded

	err, stationsAdded := retrieveDeviceStatisticsStationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stationsAdded

	err, temperatureAdded := retrieveDeviceStatisticsTemperatureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || temperatureAdded

	err, uplinkCapacityAdded := retrieveDeviceStatisticsUplinkCapacityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uplinkCapacityAdded

	return nil, retAdded
}

func retrieveDeviceStatisticsAirTimeFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	airTimeFlagName := fmt.Sprintf("%v.airTime", cmdPrefix)
	if cmd.Flags().Changed(airTimeFlagName) {
		// warning: airTime array type AirTime is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsCcqFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ccqFlagName := fmt.Sprintf("%v.ccq", cmdPrefix)
	if cmd.Flags().Changed(ccqFlagName) {
		// warning: ccq array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsClientsFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clientsFlagName := fmt.Sprintf("%v.clients", cmdPrefix)
	if cmd.Flags().Changed(clientsFlagName) {

		var clientsFlagName string
		if cmdPrefix == "" {
			clientsFlagName = "clients"
		} else {
			clientsFlagName = fmt.Sprintf("%v.clients", cmdPrefix)
		}

		clientsFlagValue, err := cmd.Flags().GetString(clientsFlagName)
		if err != nil {
			return err, false
		}
		m.Clients = clientsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsCPUFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuFlagName := fmt.Sprintf("%v.cpu", cmdPrefix)
	if cmd.Flags().Changed(cpuFlagName) {
		// warning: cpu array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsDownlinkCapacityFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	downlinkCapacityFlagName := fmt.Sprintf("%v.downlinkCapacity", cmdPrefix)
	if cmd.Flags().Changed(downlinkCapacityFlagName) {

		var downlinkCapacityFlagName string
		if cmdPrefix == "" {
			downlinkCapacityFlagName = "downlinkCapacity"
		} else {
			downlinkCapacityFlagName = fmt.Sprintf("%v.downlinkCapacity", cmdPrefix)
		}

		downlinkCapacityFlagValue, err := cmd.Flags().GetString(downlinkCapacityFlagName)
		if err != nil {
			return err, false
		}
		m.DownlinkCapacity = downlinkCapacityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsErrorsFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorsFlagName := fmt.Sprintf("%v.errors", cmdPrefix)
	if cmd.Flags().Changed(errorsFlagName) {
		// warning: errors array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsInterfacesFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfacesFlagName := fmt.Sprintf("%v.interfaces", cmdPrefix)
	if cmd.Flags().Changed(interfacesFlagName) {
		// warning: interfaces array type Interfaces is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsIntervalFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	intervalFlagName := fmt.Sprintf("%v.interval", cmdPrefix)
	if cmd.Flags().Changed(intervalFlagName) {
		// info: complex object interval Interval is retrieved outside this Changed() block
	}
	intervalFlagValue := m.Interval
	if swag.IsZero(intervalFlagValue) {
		intervalFlagValue = &models.Interval{}
	}

	err, intervalAdded := retrieveModelIntervalFlags(depth+1, intervalFlagValue, intervalFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || intervalAdded
	if intervalAdded {
		m.Interval = intervalFlagValue
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsLocalChain0Flags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	localChain0FlagName := fmt.Sprintf("%v.localChain0", cmdPrefix)
	if cmd.Flags().Changed(localChain0FlagName) {
		// warning: localChain0 array type LocalChain0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsLocalChain1Flags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	localChain1FlagName := fmt.Sprintf("%v.localChain1", cmdPrefix)
	if cmd.Flags().Changed(localChain1FlagName) {
		// warning: localChain1 array type LocalChain1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsOutputFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	outputFlagName := fmt.Sprintf("%v.output", cmdPrefix)
	if cmd.Flags().Changed(outputFlagName) {

		var outputFlagName string
		if cmdPrefix == "" {
			outputFlagName = "output"
		} else {
			outputFlagName = fmt.Sprintf("%v.output", cmdPrefix)
		}

		outputFlagValue, err := cmd.Flags().GetString(outputFlagName)
		if err != nil {
			return err, false
		}
		m.Output = outputFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsPeriodFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	periodFlagName := fmt.Sprintf("%v.period", cmdPrefix)
	if cmd.Flags().Changed(periodFlagName) {

		var periodFlagName string
		if cmdPrefix == "" {
			periodFlagName = "period"
		} else {
			periodFlagName = fmt.Sprintf("%v.period", cmdPrefix)
		}

		periodFlagValue, err := cmd.Flags().GetFloat64(periodFlagName)
		if err != nil {
			return err, false
		}
		m.Period = periodFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsPingFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pingFlagName := fmt.Sprintf("%v.ping", cmdPrefix)
	if cmd.Flags().Changed(pingFlagName) {
		// warning: ping array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsPsuFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	psuFlagName := fmt.Sprintf("%v.psu", cmdPrefix)
	if cmd.Flags().Changed(psuFlagName) {

		var psuFlagName string
		if cmdPrefix == "" {
			psuFlagName = "psu"
		} else {
			psuFlagName = fmt.Sprintf("%v.psu", cmdPrefix)
		}

		psuFlagValue, err := cmd.Flags().GetString(psuFlagName)
		if err != nil {
			return err, false
		}
		m.Psu = psuFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsPvFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pvFlagName := fmt.Sprintf("%v.pv", cmdPrefix)
	if cmd.Flags().Changed(pvFlagName) {

		var pvFlagName string
		if cmdPrefix == "" {
			pvFlagName = "pv"
		} else {
			pvFlagName = fmt.Sprintf("%v.pv", cmdPrefix)
		}

		pvFlagValue, err := cmd.Flags().GetString(pvFlagName)
		if err != nil {
			return err, false
		}
		m.Pv = pvFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsRAMFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ramFlagName := fmt.Sprintf("%v.ram", cmdPrefix)
	if cmd.Flags().Changed(ramFlagName) {
		// warning: ram array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsRemoteChain0Flags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	remoteChain0FlagName := fmt.Sprintf("%v.remoteChain0", cmdPrefix)
	if cmd.Flags().Changed(remoteChain0FlagName) {
		// warning: remoteChain0 array type RemoteChain0 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsRemoteChain1Flags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	remoteChain1FlagName := fmt.Sprintf("%v.remoteChain1", cmdPrefix)
	if cmd.Flags().Changed(remoteChain1FlagName) {
		// warning: remoteChain1 array type RemoteChain1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsRemoteSignalFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	remoteSignalFlagName := fmt.Sprintf("%v.remoteSignal", cmdPrefix)
	if cmd.Flags().Changed(remoteSignalFlagName) {
		// warning: remoteSignal array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsRemoteSignal60gFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	remoteSignal60gFlagName := fmt.Sprintf("%v.remoteSignal60g", cmdPrefix)
	if cmd.Flags().Changed(remoteSignal60gFlagName) {
		// warning: remoteSignal60g array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsSignalFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	signalFlagName := fmt.Sprintf("%v.signal", cmdPrefix)
	if cmd.Flags().Changed(signalFlagName) {
		// warning: signal array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsSignal60gFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	signal60gFlagName := fmt.Sprintf("%v.signal60g", cmdPrefix)
	if cmd.Flags().Changed(signal60gFlagName) {
		// warning: signal60g array type ListOfCoordinates is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsStationsFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stationsFlagName := fmt.Sprintf("%v.stations", cmdPrefix)
	if cmd.Flags().Changed(stationsFlagName) {
		// warning: stations map type Stations1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsTemperatureFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	temperatureFlagName := fmt.Sprintf("%v.temperature", cmdPrefix)
	if cmd.Flags().Changed(temperatureFlagName) {
		// warning: temperature array type Temperature is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceStatisticsUplinkCapacityFlags(depth int, m *models.DeviceStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uplinkCapacityFlagName := fmt.Sprintf("%v.uplinkCapacity", cmdPrefix)
	if cmd.Flags().Changed(uplinkCapacityFlagName) {

		var uplinkCapacityFlagName string
		if cmdPrefix == "" {
			uplinkCapacityFlagName = "uplinkCapacity"
		} else {
			uplinkCapacityFlagName = fmt.Sprintf("%v.uplinkCapacity", cmdPrefix)
		}

		uplinkCapacityFlagValue, err := cmd.Flags().GetString(uplinkCapacityFlagName)
		if err != nil {
			return err, false
		}
		m.UplinkCapacity = uplinkCapacityFlagValue

		retAdded = true
	}

	return nil, retAdded
}
