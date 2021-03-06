// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for SsoUserSchema

// register flags to command
func registerModelSsoUserSchemaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSsoUserSchemaEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoUserSchemaFirstName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoUserSchemaID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoUserSchemaImageURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoUserSchemaLastName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSsoUserSchemaUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSsoUserSchemaEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := `Required. `

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerSsoUserSchemaFirstName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firstNameDescription := ``

	var firstNameFlagName string
	if cmdPrefix == "" {
		firstNameFlagName = "firstName"
	} else {
		firstNameFlagName = fmt.Sprintf("%v.firstName", cmdPrefix)
	}

	var firstNameFlagDefault string

	_ = cmd.PersistentFlags().String(firstNameFlagName, firstNameFlagDefault, firstNameDescription)

	return nil
}

func registerSsoUserSchemaID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSsoUserSchemaImageURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imageUrlDescription := ``

	var imageUrlFlagName string
	if cmdPrefix == "" {
		imageUrlFlagName = "imageUrl"
	} else {
		imageUrlFlagName = fmt.Sprintf("%v.imageUrl", cmdPrefix)
	}

	var imageUrlFlagDefault string

	_ = cmd.PersistentFlags().String(imageUrlFlagName, imageUrlFlagDefault, imageUrlDescription)

	return nil
}

func registerSsoUserSchemaLastName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastNameDescription := ``

	var lastNameFlagName string
	if cmdPrefix == "" {
		lastNameFlagName = "lastName"
	} else {
		lastNameFlagName = fmt.Sprintf("%v.lastName", cmdPrefix)
	}

	var lastNameFlagDefault string

	_ = cmd.PersistentFlags().String(lastNameFlagName, lastNameFlagDefault, lastNameDescription)

	return nil
}

func registerSsoUserSchemaUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := `Required. `

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSsoUserSchemaFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveSsoUserSchemaEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, firstNameAdded := retrieveSsoUserSchemaFirstNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firstNameAdded

	err, idAdded := retrieveSsoUserSchemaIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, imageUrlAdded := retrieveSsoUserSchemaImageURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imageUrlAdded

	err, lastNameAdded := retrieveSsoUserSchemaLastNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastNameAdded

	err, usernameAdded := retrieveSsoUserSchemaUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrieveSsoUserSchemaEmailFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = &emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSsoUserSchemaFirstNameFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firstNameFlagName := fmt.Sprintf("%v.firstName", cmdPrefix)
	if cmd.Flags().Changed(firstNameFlagName) {

		var firstNameFlagName string
		if cmdPrefix == "" {
			firstNameFlagName = "firstName"
		} else {
			firstNameFlagName = fmt.Sprintf("%v.firstName", cmdPrefix)
		}

		firstNameFlagValue, err := cmd.Flags().GetString(firstNameFlagName)
		if err != nil {
			return err, false
		}
		m.FirstName = firstNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSsoUserSchemaIDFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSsoUserSchemaImageURLFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imageUrlFlagName := fmt.Sprintf("%v.imageUrl", cmdPrefix)
	if cmd.Flags().Changed(imageUrlFlagName) {

		var imageUrlFlagName string
		if cmdPrefix == "" {
			imageUrlFlagName = "imageUrl"
		} else {
			imageUrlFlagName = fmt.Sprintf("%v.imageUrl", cmdPrefix)
		}

		imageUrlFlagValue, err := cmd.Flags().GetString(imageUrlFlagName)
		if err != nil {
			return err, false
		}
		m.ImageURL = imageUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSsoUserSchemaLastNameFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastNameFlagName := fmt.Sprintf("%v.lastName", cmdPrefix)
	if cmd.Flags().Changed(lastNameFlagName) {

		var lastNameFlagName string
		if cmdPrefix == "" {
			lastNameFlagName = "lastName"
		} else {
			lastNameFlagName = fmt.Sprintf("%v.lastName", cmdPrefix)
		}

		lastNameFlagValue, err := cmd.Flags().GetString(lastNameFlagName)
		if err != nil {
			return err, false
		}
		m.LastName = lastNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSsoUserSchemaUsernameFlags(depth int, m *models.SsoUserSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = &usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
