// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for NetworkStatistics

// register flags to command
func registerModelNetworkStatisticsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkStatisticsAllClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsAllSites(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsDataLinksUtilizationScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsDownlinkUtilizationMean(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsFirmwareUpToDateness(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsGateways(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsInterval(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsIspScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsLinkScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsLiveClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsLiveSites(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsNetworkDesignScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsNetworkHealth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsOutages(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsPeriod(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsSignalScore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkStatisticsUplinkUtilizationMean(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkStatisticsAllClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: allClients AllClients array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsAllSites(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: allSites AllSites array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsDataLinksUtilizationScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: dataLinksUtilizationScore DataLinksUtilizationScore array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsDownlinkUtilizationMean(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: downlinkUtilizationMean DownlinkUtilizationMean array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsFirmwareUpToDateness(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: firmwareUpToDateness FirmwareUpToDateness array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsGateways(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: gateways Gateways map type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsInterval(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var intervalFlagName string
	if cmdPrefix == "" {
		intervalFlagName = "interval"
	} else {
		intervalFlagName = fmt.Sprintf("%v.interval", cmdPrefix)
	}

	if err := registerModelNetworkStatisticIntervalFlags(depth+1, intervalFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkStatisticsIspScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ispScore IspScore array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsLinkScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: linkScore LinkScore1 array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsLiveClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: liveClients LiveClients array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsLiveSites(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: liveSites LiveSites array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsNetworkDesignScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: networkDesignScore NetworkDesignScore array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsNetworkHealth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: networkHealth NetworkHealth array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsOutages(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: outages Outages array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsPeriod(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

	var periodFlagDefault int64

	_ = cmd.PersistentFlags().Int64(periodFlagName, periodFlagDefault, periodDescription)

	return nil
}

func registerNetworkStatisticsSignalScore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: signalScore SignalScore array type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkStatisticsUplinkUtilizationMean(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: uplinkUtilizationMean UplinkUtilizationMean array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkStatisticsFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, allClientsAdded := retrieveNetworkStatisticsAllClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allClientsAdded

	err, allSitesAdded := retrieveNetworkStatisticsAllSitesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allSitesAdded

	err, dataLinksUtilizationScoreAdded := retrieveNetworkStatisticsDataLinksUtilizationScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataLinksUtilizationScoreAdded

	err, downlinkUtilizationMeanAdded := retrieveNetworkStatisticsDownlinkUtilizationMeanFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || downlinkUtilizationMeanAdded

	err, firmwareUpToDatenessAdded := retrieveNetworkStatisticsFirmwareUpToDatenessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firmwareUpToDatenessAdded

	err, gatewaysAdded := retrieveNetworkStatisticsGatewaysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gatewaysAdded

	err, intervalAdded := retrieveNetworkStatisticsIntervalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || intervalAdded

	err, ispScoreAdded := retrieveNetworkStatisticsIspScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ispScoreAdded

	err, linkScoreAdded := retrieveNetworkStatisticsLinkScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linkScoreAdded

	err, liveClientsAdded := retrieveNetworkStatisticsLiveClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || liveClientsAdded

	err, liveSitesAdded := retrieveNetworkStatisticsLiveSitesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || liveSitesAdded

	err, networkDesignScoreAdded := retrieveNetworkStatisticsNetworkDesignScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkDesignScoreAdded

	err, networkHealthAdded := retrieveNetworkStatisticsNetworkHealthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkHealthAdded

	err, outagesAdded := retrieveNetworkStatisticsOutagesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || outagesAdded

	err, periodAdded := retrieveNetworkStatisticsPeriodFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || periodAdded

	err, signalScoreAdded := retrieveNetworkStatisticsSignalScoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || signalScoreAdded

	err, uplinkUtilizationMeanAdded := retrieveNetworkStatisticsUplinkUtilizationMeanFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || uplinkUtilizationMeanAdded

	return nil, retAdded
}

func retrieveNetworkStatisticsAllClientsFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allClientsFlagName := fmt.Sprintf("%v.allClients", cmdPrefix)
	if cmd.Flags().Changed(allClientsFlagName) {
		// warning: allClients array type AllClients is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsAllSitesFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allSitesFlagName := fmt.Sprintf("%v.allSites", cmdPrefix)
	if cmd.Flags().Changed(allSitesFlagName) {
		// warning: allSites array type AllSites is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsDataLinksUtilizationScoreFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataLinksUtilizationScoreFlagName := fmt.Sprintf("%v.dataLinksUtilizationScore", cmdPrefix)
	if cmd.Flags().Changed(dataLinksUtilizationScoreFlagName) {
		// warning: dataLinksUtilizationScore array type DataLinksUtilizationScore is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsDownlinkUtilizationMeanFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	downlinkUtilizationMeanFlagName := fmt.Sprintf("%v.downlinkUtilizationMean", cmdPrefix)
	if cmd.Flags().Changed(downlinkUtilizationMeanFlagName) {
		// warning: downlinkUtilizationMean array type DownlinkUtilizationMean is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsFirmwareUpToDatenessFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firmwareUpToDatenessFlagName := fmt.Sprintf("%v.firmwareUpToDateness", cmdPrefix)
	if cmd.Flags().Changed(firmwareUpToDatenessFlagName) {
		// warning: firmwareUpToDateness array type FirmwareUpToDateness is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsGatewaysFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gatewaysFlagName := fmt.Sprintf("%v.gateways", cmdPrefix)
	if cmd.Flags().Changed(gatewaysFlagName) {
		// warning: gateways map type Gateways is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsIntervalFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	intervalFlagName := fmt.Sprintf("%v.interval", cmdPrefix)
	if cmd.Flags().Changed(intervalFlagName) {
		// info: complex object interval NetworkStatisticInterval is retrieved outside this Changed() block
	}
	intervalFlagValue := m.Interval
	if swag.IsZero(intervalFlagValue) {
		intervalFlagValue = &models.NetworkStatisticInterval{}
	}

	err, intervalAdded := retrieveModelNetworkStatisticIntervalFlags(depth+1, intervalFlagValue, intervalFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || intervalAdded
	if intervalAdded {
		m.Interval = intervalFlagValue
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsIspScoreFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ispScoreFlagName := fmt.Sprintf("%v.ispScore", cmdPrefix)
	if cmd.Flags().Changed(ispScoreFlagName) {
		// warning: ispScore array type IspScore is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsLinkScoreFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	linkScoreFlagName := fmt.Sprintf("%v.linkScore", cmdPrefix)
	if cmd.Flags().Changed(linkScoreFlagName) {
		// warning: linkScore array type LinkScore1 is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsLiveClientsFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	liveClientsFlagName := fmt.Sprintf("%v.liveClients", cmdPrefix)
	if cmd.Flags().Changed(liveClientsFlagName) {
		// warning: liveClients array type LiveClients is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsLiveSitesFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	liveSitesFlagName := fmt.Sprintf("%v.liveSites", cmdPrefix)
	if cmd.Flags().Changed(liveSitesFlagName) {
		// warning: liveSites array type LiveSites is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsNetworkDesignScoreFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkDesignScoreFlagName := fmt.Sprintf("%v.networkDesignScore", cmdPrefix)
	if cmd.Flags().Changed(networkDesignScoreFlagName) {
		// warning: networkDesignScore array type NetworkDesignScore is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsNetworkHealthFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkHealthFlagName := fmt.Sprintf("%v.networkHealth", cmdPrefix)
	if cmd.Flags().Changed(networkHealthFlagName) {
		// warning: networkHealth array type NetworkHealth is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsOutagesFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	outagesFlagName := fmt.Sprintf("%v.outages", cmdPrefix)
	if cmd.Flags().Changed(outagesFlagName) {
		// warning: outages array type Outages is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsPeriodFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		periodFlagValue, err := cmd.Flags().GetInt64(periodFlagName)
		if err != nil {
			return err, false
		}
		m.Period = periodFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsSignalScoreFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	signalScoreFlagName := fmt.Sprintf("%v.signalScore", cmdPrefix)
	if cmd.Flags().Changed(signalScoreFlagName) {
		// warning: signalScore array type SignalScore is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNetworkStatisticsUplinkUtilizationMeanFlags(depth int, m *models.NetworkStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	uplinkUtilizationMeanFlagName := fmt.Sprintf("%v.uplinkUtilizationMean", cmdPrefix)
	if cmd.Flags().Changed(uplinkUtilizationMeanFlagName) {
		// warning: uplinkUtilizationMean array type UplinkUtilizationMean is not supported by go-swagger cli yet
	}

	return nil, retAdded
}