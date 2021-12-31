// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model46

// register flags to command
func registerModelModel46Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel46Mac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel46Origin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel46Vendor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel46Mac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := `Required. MAC address`

	var macFlagName string
	if cmdPrefix == "" {
		macFlagName = "mac"
	} else {
		macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
	}

	var macFlagDefault string

	_ = cmd.PersistentFlags().String(macFlagName, macFlagDefault, macDescription)

	return nil
}

func registerModel46Origin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	originDescription := `Required. `

	var originFlagName string
	if cmdPrefix == "" {
		originFlagName = "origin"
	} else {
		originFlagName = fmt.Sprintf("%v.origin", cmdPrefix)
	}

	var originFlagDefault string

	_ = cmd.PersistentFlags().String(originFlagName, originFlagDefault, originDescription)

	return nil
}

func registerModel46Vendor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorDescription := ``

	var vendorFlagName string
	if cmdPrefix == "" {
		vendorFlagName = "vendor"
	} else {
		vendorFlagName = fmt.Sprintf("%v.vendor", cmdPrefix)
	}

	var vendorFlagDefault string

	_ = cmd.PersistentFlags().String(vendorFlagName, vendorFlagDefault, vendorDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel46Flags(depth int, m *models.Model46, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, macAdded := retrieveModel46MacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	err, originAdded := retrieveModel46OriginFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || originAdded

	err, vendorAdded := retrieveModel46VendorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorAdded

	return nil, retAdded
}

func retrieveModel46MacFlags(depth int, m *models.Model46, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macFlagName := fmt.Sprintf("%v.mac", cmdPrefix)
	if cmd.Flags().Changed(macFlagName) {

		var macFlagName string
		if cmdPrefix == "" {
			macFlagName = "mac"
		} else {
			macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
		}

		macFlagValue, err := cmd.Flags().GetString(macFlagName)
		if err != nil {
			return err, false
		}
		m.Mac = &macFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel46OriginFlags(depth int, m *models.Model46, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModel46VendorFlags(depth int, m *models.Model46, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorFlagName := fmt.Sprintf("%v.vendor", cmdPrefix)
	if cmd.Flags().Changed(vendorFlagName) {

		var vendorFlagName string
		if cmdPrefix == "" {
			vendorFlagName = "vendor"
		} else {
			vendorFlagName = fmt.Sprintf("%v.vendor", cmdPrefix)
		}

		vendorFlagValue, err := cmd.Flags().GetString(vendorFlagName)
		if err != nil {
			return err, false
		}
		m.Vendor = vendorFlagValue

		retAdded = true
	}

	return nil, retAdded
}
