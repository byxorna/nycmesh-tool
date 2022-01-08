// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/spf13/cobra"
)

// Schema cli for Version

// register flags to command
func registerModelVersionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVersionBuild(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVersionDeployment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVersionTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVersionVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVersionBuild(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	buildDescription := `Required. Build identification.`

	var buildFlagName string
	if cmdPrefix == "" {
		buildFlagName = "build"
	} else {
		buildFlagName = fmt.Sprintf("%v.build", cmdPrefix)
	}

	var buildFlagDefault string

	_ = cmd.PersistentFlags().String(buildFlagName, buildFlagDefault, buildDescription)

	return nil
}

func registerVersionDeployment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deploymentDescription := `Required. How was the UISP deployed. Content of UISP_DEPLOYMENT env variable.`

	var deploymentFlagName string
	if cmdPrefix == "" {
		deploymentFlagName = "deployment"
	} else {
		deploymentFlagName = fmt.Sprintf("%v.deployment", cmdPrefix)
	}

	var deploymentFlagDefault string

	_ = cmd.PersistentFlags().String(deploymentFlagName, deploymentFlagDefault, deploymentDescription)

	return nil
}

func registerVersionTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeDescription := `Required. Time of the build.`

	var timeFlagName string
	if cmdPrefix == "" {
		timeFlagName = "time"
	} else {
		timeFlagName = fmt.Sprintf("%v.time", cmdPrefix)
	}

	var timeFlagDefault string

	_ = cmd.PersistentFlags().String(timeFlagName, timeFlagDefault, timeDescription)

	return nil
}

func registerVersionVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Required. Version of UISP.`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVersionFlags(depth int, m *models.Version, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, buildAdded := retrieveVersionBuildFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || buildAdded

	err, deploymentAdded := retrieveVersionDeploymentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deploymentAdded

	err, timeAdded := retrieveVersionTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeAdded

	err, versionAdded := retrieveVersionVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveVersionBuildFlags(depth int, m *models.Version, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	buildFlagName := fmt.Sprintf("%v.build", cmdPrefix)
	if cmd.Flags().Changed(buildFlagName) {

		var buildFlagName string
		if cmdPrefix == "" {
			buildFlagName = "build"
		} else {
			buildFlagName = fmt.Sprintf("%v.build", cmdPrefix)
		}

		buildFlagValue, err := cmd.Flags().GetString(buildFlagName)
		if err != nil {
			return err, false
		}
		m.Build = &buildFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVersionDeploymentFlags(depth int, m *models.Version, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deploymentFlagName := fmt.Sprintf("%v.deployment", cmdPrefix)
	if cmd.Flags().Changed(deploymentFlagName) {

		var deploymentFlagName string
		if cmdPrefix == "" {
			deploymentFlagName = "deployment"
		} else {
			deploymentFlagName = fmt.Sprintf("%v.deployment", cmdPrefix)
		}

		deploymentFlagValue, err := cmd.Flags().GetString(deploymentFlagName)
		if err != nil {
			return err, false
		}
		m.Deployment = &deploymentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVersionTimeFlags(depth int, m *models.Version, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeFlagName := fmt.Sprintf("%v.time", cmdPrefix)
	if cmd.Flags().Changed(timeFlagName) {

		var timeFlagName string
		if cmdPrefix == "" {
			timeFlagName = "time"
		} else {
			timeFlagName = fmt.Sprintf("%v.time", cmdPrefix)
		}

		timeFlagValue, err := cmd.Flags().GetString(timeFlagName)
		if err != nil {
			return err, false
		}
		m.Time = &timeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVersionVersionFlags(depth int, m *models.Version, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = &versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}