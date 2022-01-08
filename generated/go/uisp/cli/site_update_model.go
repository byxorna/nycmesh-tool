// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for SiteUpdate

// register flags to command
func registerModelSiteUpdateFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSiteUpdateDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteUpdateID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteUpdateIdentification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteUpdateNotifications(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteUpdateQos(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSiteUpdateUcrm(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteUpdateDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	if err := registerModelSiteDescription1Flags(depth+1, descriptionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteUpdateID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerSiteUpdateIdentification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelSiteIdentificationDetail1Flags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteUpdateNotifications(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var notificationsFlagName string
	if cmdPrefix == "" {
		notificationsFlagName = "notifications"
	} else {
		notificationsFlagName = fmt.Sprintf("%v.notifications", cmdPrefix)
	}

	if err := registerModelSiteNotificationsFlags(depth+1, notificationsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteUpdateQos(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var qosFlagName string
	if cmdPrefix == "" {
		qosFlagName = "qos"
	} else {
		qosFlagName = fmt.Sprintf("%v.qos", cmdPrefix)
	}

	if err := registerModelSiteTrafficShaping1Flags(depth+1, qosFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSiteUpdateUcrm(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ucrmFlagName string
	if cmdPrefix == "" {
		ucrmFlagName = "ucrm"
	} else {
		ucrmFlagName = fmt.Sprintf("%v.ucrm", cmdPrefix)
	}

	if err := registerModelSiteUcrmDescriptionFlags(depth+1, ucrmFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSiteUpdateFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveSiteUpdateDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, idAdded := retrieveSiteUpdateIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, identificationAdded := retrieveSiteUpdateIdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	err, notificationsAdded := retrieveSiteUpdateNotificationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || notificationsAdded

	err, qosAdded := retrieveSiteUpdateQosFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || qosAdded

	err, ucrmAdded := retrieveSiteUpdateUcrmFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ucrmAdded

	return nil, retAdded
}

func retrieveSiteUpdateDescriptionFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {
		// info: complex object description SiteDescription1 is retrieved outside this Changed() block
	}
	descriptionFlagValue := m.Description
	if swag.IsZero(descriptionFlagValue) {
		descriptionFlagValue = &models.SiteDescription1{}
	}

	err, descriptionAdded := retrieveModelSiteDescription1Flags(depth+1, descriptionFlagValue, descriptionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded
	if descriptionAdded {
		m.Description = descriptionFlagValue
	}

	return nil, retAdded
}

func retrieveSiteUpdateIDFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSiteUpdateIdentificationFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification SiteIdentificationDetail1 is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.SiteIdentificationDetail1{}
	}

	err, identificationAdded := retrieveModelSiteIdentificationDetail1Flags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}

func retrieveSiteUpdateNotificationsFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	notificationsFlagName := fmt.Sprintf("%v.notifications", cmdPrefix)
	if cmd.Flags().Changed(notificationsFlagName) {
		// info: complex object notifications SiteNotifications is retrieved outside this Changed() block
	}
	notificationsFlagValue := m.Notifications
	if swag.IsZero(notificationsFlagValue) {
		notificationsFlagValue = &models.SiteNotifications{}
	}

	err, notificationsAdded := retrieveModelSiteNotificationsFlags(depth+1, notificationsFlagValue, notificationsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || notificationsAdded
	if notificationsAdded {
		m.Notifications = notificationsFlagValue
	}

	return nil, retAdded
}

func retrieveSiteUpdateQosFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	qosFlagName := fmt.Sprintf("%v.qos", cmdPrefix)
	if cmd.Flags().Changed(qosFlagName) {
		// info: complex object qos SiteTrafficShaping1 is retrieved outside this Changed() block
	}
	qosFlagValue := m.Qos
	if swag.IsZero(qosFlagValue) {
		qosFlagValue = &models.SiteTrafficShaping1{}
	}

	err, qosAdded := retrieveModelSiteTrafficShaping1Flags(depth+1, qosFlagValue, qosFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || qosAdded
	if qosAdded {
		m.Qos = qosFlagValue
	}

	return nil, retAdded
}

func retrieveSiteUpdateUcrmFlags(depth int, m *models.SiteUpdate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ucrmFlagName := fmt.Sprintf("%v.ucrm", cmdPrefix)
	if cmd.Flags().Changed(ucrmFlagName) {
		// info: complex object ucrm SiteUcrmDescription is retrieved outside this Changed() block
	}
	ucrmFlagValue := m.Ucrm
	if swag.IsZero(ucrmFlagValue) {
		ucrmFlagValue = &models.SiteUcrmDescription{}
	}

	err, ucrmAdded := retrieveModelSiteUcrmDescriptionFlags(depth+1, ucrmFlagValue, ucrmFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ucrmAdded
	if ucrmAdded {
		m.Ucrm = ucrmFlagValue
	}

	return nil, retAdded
}
