// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
  "github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for NewClientSchema

// register flags to command
func registerModelNewClientSchemaFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNewClientSchemaAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaAddressData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaFirstName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaLastName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaLocation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaNote(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaPhone(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNewClientSchemaServicePlanID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewClientSchemaAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addressDescription := `Address of the Client site.`

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

func registerNewClientSchemaAddressData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var addressDataFlagName string
	if cmdPrefix == "" {
		addressDataFlagName = "addressData"
	} else {
		addressDataFlagName = fmt.Sprintf("%v.addressData", cmdPrefix)
	}

	if err := registerModelAddressDataFlags(depth+1, addressDataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewClientSchemaEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := `Client's e-mail.`

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

func registerNewClientSchemaFirstName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firstNameDescription := `Required. First name of the Client.`

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

func registerNewClientSchemaLastName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastNameDescription := `Required. Last name of the Client.`

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

func registerNewClientSchemaLocation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var locationFlagName string
	if cmdPrefix == "" {
		locationFlagName = "location"
	} else {
		locationFlagName = fmt.Sprintf("%v.location", cmdPrefix)
	}

	if err := registerModelNewClientLocationFlags(depth+1, locationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNewClientSchemaNote(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	noteDescription := `Custom note.`

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

func registerNewClientSchemaPhone(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	phoneDescription := `Client's phone number.`

	var phoneFlagName string
	if cmdPrefix == "" {
		phoneFlagName = "phone"
	} else {
		phoneFlagName = fmt.Sprintf("%v.phone", cmdPrefix)
	}

	var phoneFlagDefault string

	_ = cmd.PersistentFlags().String(phoneFlagName, phoneFlagDefault, phoneDescription)

	return nil
}

func registerNewClientSchemaServicePlanID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	servicePlanIdDescription := `Required. ID of the CRM Service plan that should be assigned to the new Client`

	var servicePlanIdFlagName string
	if cmdPrefix == "" {
		servicePlanIdFlagName = "servicePlanId"
	} else {
		servicePlanIdFlagName = fmt.Sprintf("%v.servicePlanId", cmdPrefix)
	}

	var servicePlanIdFlagDefault string

	_ = cmd.PersistentFlags().String(servicePlanIdFlagName, servicePlanIdFlagDefault, servicePlanIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNewClientSchemaFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressAdded := retrieveNewClientSchemaAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressAdded

	err, addressDataAdded := retrieveNewClientSchemaAddressDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressDataAdded

	err, emailAdded := retrieveNewClientSchemaEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, firstNameAdded := retrieveNewClientSchemaFirstNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firstNameAdded

	err, lastNameAdded := retrieveNewClientSchemaLastNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastNameAdded

	err, locationAdded := retrieveNewClientSchemaLocationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded

	err, noteAdded := retrieveNewClientSchemaNoteFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noteAdded

	err, phoneAdded := retrieveNewClientSchemaPhoneFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || phoneAdded

	err, servicePlanIdAdded := retrieveNewClientSchemaServicePlanIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || servicePlanIdAdded

	return nil, retAdded
}

func retrieveNewClientSchemaAddressFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNewClientSchemaAddressDataFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressDataFlagName := fmt.Sprintf("%v.addressData", cmdPrefix)
	if cmd.Flags().Changed(addressDataFlagName) {
		// info: complex object addressData AddressData is retrieved outside this Changed() block
	}
	addressDataFlagValue := m.AddressData
	if swag.IsZero(addressDataFlagValue) {
		addressDataFlagValue = &models.AddressData{}
	}

	err, addressDataAdded := retrieveModelAddressDataFlags(depth+1, addressDataFlagValue, addressDataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressDataAdded
	if addressDataAdded {
		m.AddressData = addressDataFlagValue
	}

	return nil, retAdded
}

func retrieveNewClientSchemaEmailFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewClientSchemaFirstNameFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.FirstName = &firstNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewClientSchemaLastNameFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.LastName = &lastNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewClientSchemaLocationFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	locationFlagName := fmt.Sprintf("%v.location", cmdPrefix)
	if cmd.Flags().Changed(locationFlagName) {
		// info: complex object location NewClientLocation is retrieved outside this Changed() block
	}
	locationFlagValue := m.Location
	if swag.IsZero(locationFlagValue) {
		locationFlagValue = &models.NewClientLocation{}
	}

	err, locationAdded := retrieveModelNewClientLocationFlags(depth+1, locationFlagValue, locationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || locationAdded
	if locationAdded {
		m.Location = locationFlagValue
	}

	return nil, retAdded
}

func retrieveNewClientSchemaNoteFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNewClientSchemaPhoneFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	phoneFlagName := fmt.Sprintf("%v.phone", cmdPrefix)
	if cmd.Flags().Changed(phoneFlagName) {

		var phoneFlagName string
		if cmdPrefix == "" {
			phoneFlagName = "phone"
		} else {
			phoneFlagName = fmt.Sprintf("%v.phone", cmdPrefix)
		}

		phoneFlagValue, err := cmd.Flags().GetString(phoneFlagName)
		if err != nil {
			return err, false
		}
		m.Phone = phoneFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNewClientSchemaServicePlanIDFlags(depth int, m *models.NewClientSchema, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	servicePlanIdFlagName := fmt.Sprintf("%v.servicePlanId", cmdPrefix)
	if cmd.Flags().Changed(servicePlanIdFlagName) {

		var servicePlanIdFlagName string
		if cmdPrefix == "" {
			servicePlanIdFlagName = "servicePlanId"
		} else {
			servicePlanIdFlagName = fmt.Sprintf("%v.servicePlanId", cmdPrefix)
		}

		servicePlanIdFlagValue, err := cmd.Flags().GetString(servicePlanIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServicePlanID = &servicePlanIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
