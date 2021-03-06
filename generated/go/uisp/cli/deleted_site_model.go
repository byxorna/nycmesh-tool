// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

// Schema cli for DeletedSite

// register flags to command
func registerModelDeletedSiteFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeletedSiteDeletedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeletedSiteID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeletedSiteUcrmID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeletedSiteDeletedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deletedAtDescription := `Required. Time when the site was deleted in ISO format.`

	var deletedAtFlagName string
	if cmdPrefix == "" {
		deletedAtFlagName = "deletedAt"
	} else {
		deletedAtFlagName = fmt.Sprintf("%v.deletedAt", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(deletedAtFlagName, "", deletedAtDescription)

	return nil
}

func registerDeletedSiteID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. Site ID.`

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

func registerDeletedSiteUcrmID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ucrmIdDescription := `Required. ID of UCRM client bound with this site. Null if no UCRM client is bound to this site.`

	var ucrmIdFlagName string
	if cmdPrefix == "" {
		ucrmIdFlagName = "ucrmId"
	} else {
		ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
	}

	var ucrmIdFlagDefault string

	_ = cmd.PersistentFlags().String(ucrmIdFlagName, ucrmIdFlagDefault, ucrmIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeletedSiteFlags(depth int, m *models.DeletedSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deletedAtAdded := retrieveDeletedSiteDeletedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deletedAtAdded

	err, idAdded := retrieveDeletedSiteIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ucrmIdAdded := retrieveDeletedSiteUcrmIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ucrmIdAdded

	return nil, retAdded
}

func retrieveDeletedSiteDeletedAtFlags(depth int, m *models.DeletedSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deletedAtFlagName := fmt.Sprintf("%v.deletedAt", cmdPrefix)
	if cmd.Flags().Changed(deletedAtFlagName) {

		var deletedAtFlagName string
		if cmdPrefix == "" {
			deletedAtFlagName = "deletedAt"
		} else {
			deletedAtFlagName = fmt.Sprintf("%v.deletedAt", cmdPrefix)
		}

		deletedAtFlagValueStr, err := cmd.Flags().GetString(deletedAtFlagName)
		if err != nil {
			return err, false
		}
		var deletedAtFlagValue strfmt.DateTime
		if err := deletedAtFlagValue.UnmarshalText([]byte(deletedAtFlagValueStr)); err != nil {
			return err, false
		}
		m.DeletedAt = &deletedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeletedSiteIDFlags(depth int, m *models.DeletedSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDeletedSiteUcrmIDFlags(depth int, m *models.DeletedSite, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ucrmIdFlagName := fmt.Sprintf("%v.ucrmId", cmdPrefix)
	if cmd.Flags().Changed(ucrmIdFlagName) {

		var ucrmIdFlagName string
		if cmdPrefix == "" {
			ucrmIdFlagName = "ucrmId"
		} else {
			ucrmIdFlagName = fmt.Sprintf("%v.ucrmId", cmdPrefix)
		}

		ucrmIdFlagValue, err := cmd.Flags().GetString(ucrmIdFlagName)
		if err != nil {
			return err, false
		}
		m.UcrmID = &ucrmIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
