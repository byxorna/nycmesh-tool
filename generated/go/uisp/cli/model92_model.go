// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Model92

// register flags to command
func registerModelModel92Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModel92Account(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel92Interface(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel92Mtu(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel92Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel92Password(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModel92PppoeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModel92Account(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accountDescription := ``

	var accountFlagName string
	if cmdPrefix == "" {
		accountFlagName = "account"
	} else {
		accountFlagName = fmt.Sprintf("%v.account", cmdPrefix)
	}

	var accountFlagDefault string

	_ = cmd.PersistentFlags().String(accountFlagName, accountFlagDefault, accountDescription)

	return nil
}

func registerModel92Interface(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	interfaceDescription := `Required. `

	var interfaceFlagName string
	if cmdPrefix == "" {
		interfaceFlagName = "interface"
	} else {
		interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
	}

	var interfaceFlagDefault string

	_ = cmd.PersistentFlags().String(interfaceFlagName, interfaceFlagDefault, interfaceDescription)

	return nil
}

func registerModel92Mtu(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	mtuDescription := `Required. `

	var mtuFlagName string
	if cmdPrefix == "" {
		mtuFlagName = "mtu"
	} else {
		mtuFlagName = fmt.Sprintf("%v.mtu", cmdPrefix)
	}

	var mtuFlagDefault float64

	_ = cmd.PersistentFlags().Float64(mtuFlagName, mtuFlagDefault, mtuDescription)

	return nil
}

func registerModel92Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerModel92Password(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	passwordDescription := `Required. `

	var passwordFlagName string
	if cmdPrefix == "" {
		passwordFlagName = "password"
	} else {
		passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
	}

	var passwordFlagDefault string

	_ = cmd.PersistentFlags().String(passwordFlagName, passwordFlagDefault, passwordDescription)

	return nil
}

func registerModel92PppoeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pppoeIdDescription := `Required. `

	var pppoeIdFlagName string
	if cmdPrefix == "" {
		pppoeIdFlagName = "pppoeId"
	} else {
		pppoeIdFlagName = fmt.Sprintf("%v.pppoeId", cmdPrefix)
	}

	var pppoeIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pppoeIdFlagName, pppoeIdFlagDefault, pppoeIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModel92Flags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accountAdded := retrieveModel92AccountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accountAdded

	err, interfaceAdded := retrieveModel92InterfaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || interfaceAdded

	err, mtuAdded := retrieveModel92MtuFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mtuAdded

	err, nameAdded := retrieveModel92NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, passwordAdded := retrieveModel92PasswordFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || passwordAdded

	err, pppoeIdAdded := retrieveModel92PppoeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pppoeIdAdded

	return nil, retAdded
}

func retrieveModel92AccountFlags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accountFlagName := fmt.Sprintf("%v.account", cmdPrefix)
	if cmd.Flags().Changed(accountFlagName) {

		var accountFlagName string
		if cmdPrefix == "" {
			accountFlagName = "account"
		} else {
			accountFlagName = fmt.Sprintf("%v.account", cmdPrefix)
		}

		accountFlagValue, err := cmd.Flags().GetString(accountFlagName)
		if err != nil {
			return err, false
		}
		m.Account = accountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel92InterfaceFlags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	interfaceFlagName := fmt.Sprintf("%v.interface", cmdPrefix)
	if cmd.Flags().Changed(interfaceFlagName) {

		var interfaceFlagName string
		if cmdPrefix == "" {
			interfaceFlagName = "interface"
		} else {
			interfaceFlagName = fmt.Sprintf("%v.interface", cmdPrefix)
		}

		interfaceFlagValue, err := cmd.Flags().GetString(interfaceFlagName)
		if err != nil {
			return err, false
		}
		m.Interface = &interfaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel92MtuFlags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mtuFlagName := fmt.Sprintf("%v.mtu", cmdPrefix)
	if cmd.Flags().Changed(mtuFlagName) {

		var mtuFlagName string
		if cmdPrefix == "" {
			mtuFlagName = "mtu"
		} else {
			mtuFlagName = fmt.Sprintf("%v.mtu", cmdPrefix)
		}

		mtuFlagValue, err := cmd.Flags().GetFloat64(mtuFlagName)
		if err != nil {
			return err, false
		}
		m.Mtu = &mtuFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel92NameFlags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel92PasswordFlags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	passwordFlagName := fmt.Sprintf("%v.password", cmdPrefix)
	if cmd.Flags().Changed(passwordFlagName) {

		var passwordFlagName string
		if cmdPrefix == "" {
			passwordFlagName = "password"
		} else {
			passwordFlagName = fmt.Sprintf("%v.password", cmdPrefix)
		}

		passwordFlagValue, err := cmd.Flags().GetString(passwordFlagName)
		if err != nil {
			return err, false
		}
		m.Password = &passwordFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModel92PppoeIDFlags(depth int, m *models.Model92, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pppoeIdFlagName := fmt.Sprintf("%v.pppoeId", cmdPrefix)
	if cmd.Flags().Changed(pppoeIdFlagName) {

		var pppoeIdFlagName string
		if cmdPrefix == "" {
			pppoeIdFlagName = "pppoeId"
		} else {
			pppoeIdFlagName = fmt.Sprintf("%v.pppoeId", cmdPrefix)
		}

		pppoeIdFlagValue, err := cmd.Flags().GetInt64(pppoeIdFlagName)
		if err != nil {
			return err, false
		}
		m.PppoeID = &pppoeIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
