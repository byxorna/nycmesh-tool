// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Model85

// register flags to command
func registerModelModel85Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel85SMTP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel85To(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel85SMTP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var smtpFlagName string
	if cmdPrefix == "" {
		smtpFlagName = "smtp"
	} else {
		smtpFlagName = fmt.Sprintf("%v.smtp", cmdPrefix)
	}

	if err := registerModelSMTPFlags(depth+1, smtpFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel85To(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	toDescription := `Required. `

	var toFlagName string
	if cmdPrefix == "" {
		toFlagName = "to"
	} else {
		toFlagName = fmt.Sprintf("%v.to", cmdPrefix)
	}

	var toFlagDefault string

	_ = cmd.PersistentFlags().String(toFlagName, toFlagDefault, toDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel85Flags(depth int, m *models.Model85, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, smtpAdded := retrieveModel85SMTPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || smtpAdded

	err, toAdded := retrieveModel85ToFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || toAdded

	return nil, retAdded
}

func retrieveModel85SMTPFlags(depth int, m *models.Model85, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	smtpFlagName := fmt.Sprintf("%v.smtp", cmdPrefix)
	if cmd.Flags().Changed(smtpFlagName) {
		// info: complex object smtp SMTP is retrieved outside this Changed() block
	}
	smtpFlagValue := m.SMTP
	if swag.IsZero(smtpFlagValue) {
		smtpFlagValue = &models.SMTP{}
	}

	err, smtpAdded := retrieveModelSMTPFlags(depth+1, smtpFlagValue, smtpFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || smtpAdded
	if smtpAdded {
		m.SMTP = smtpFlagValue
	}

	return nil, retAdded
}

func retrieveModel85ToFlags(depth int, m *models.Model85, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	toFlagName := fmt.Sprintf("%v.to", cmdPrefix)
	if cmd.Flags().Changed(toFlagName) {

		var toFlagName string
		if cmdPrefix == "" {
			toFlagName = "to"
		} else {
			toFlagName = fmt.Sprintf("%v.to", cmdPrefix)
		}

		toFlagValue, err := cmd.Flags().GetString(toFlagName)
		if err != nil {
			return err, false
		}
		m.To = &toFlagValue

		retAdded = true
	}

	return nil, retAdded
}
