// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/models"
	"github.com/spf13/cobra"
)

// Schema cli for WebServer

// register flags to command
func registerModelWebServerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebServerEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebServerHTTPPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebServerHTTPSPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebServerEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := ``

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerWebServerHTTPPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	httpPortDescription := `Web UI HTTP port`

	var httpPortFlagName string
	if cmdPrefix == "" {
		httpPortFlagName = "httpPort"
	} else {
		httpPortFlagName = fmt.Sprintf("%v.httpPort", cmdPrefix)
	}

	var httpPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(httpPortFlagName, httpPortFlagDefault, httpPortDescription)

	return nil
}

func registerWebServerHTTPSPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	httpsPortDescription := `Web UI HTTPS port`

	var httpsPortFlagName string
	if cmdPrefix == "" {
		httpsPortFlagName = "httpsPort"
	} else {
		httpsPortFlagName = fmt.Sprintf("%v.httpsPort", cmdPrefix)
	}

	var httpsPortFlagDefault int64

	_ = cmd.PersistentFlags().Int64(httpsPortFlagName, httpsPortFlagDefault, httpsPortDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWebServerFlags(depth int, m *models.WebServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, enabledAdded := retrieveWebServerEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, httpPortAdded := retrieveWebServerHTTPPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || httpPortAdded

	err, httpsPortAdded := retrieveWebServerHTTPSPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || httpsPortAdded

	return nil, retAdded
}

func retrieveWebServerEnabledFlags(depth int, m *models.WebServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebServerHTTPPortFlags(depth int, m *models.WebServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	httpPortFlagName := fmt.Sprintf("%v.httpPort", cmdPrefix)
	if cmd.Flags().Changed(httpPortFlagName) {

		var httpPortFlagName string
		if cmdPrefix == "" {
			httpPortFlagName = "httpPort"
		} else {
			httpPortFlagName = fmt.Sprintf("%v.httpPort", cmdPrefix)
		}

		httpPortFlagValue, err := cmd.Flags().GetInt64(httpPortFlagName)
		if err != nil {
			return err, false
		}
		m.HTTPPort = &httpPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebServerHTTPSPortFlags(depth int, m *models.WebServer, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	httpsPortFlagName := fmt.Sprintf("%v.httpsPort", cmdPrefix)
	if cmd.Flags().Changed(httpsPortFlagName) {

		var httpsPortFlagName string
		if cmdPrefix == "" {
			httpsPortFlagName = "httpsPort"
		} else {
			httpsPortFlagName = fmt.Sprintf("%v.httpsPort", cmdPrefix)
		}

		httpsPortFlagValue, err := cmd.Flags().GetInt64(httpsPortFlagName)
		if err != nil {
			return err, false
		}
		m.HTTPSPort = &httpsPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}
