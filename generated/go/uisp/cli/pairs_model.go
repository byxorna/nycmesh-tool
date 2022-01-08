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

// Schema cli for Pairs

// register flags to command
func registerModelPairsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPairsPairA(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPairsPairB(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPairsPairC(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPairsPairD(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPairsPairA(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var pairAFlagName string
	if cmdPrefix == "" {
		pairAFlagName = "pairA"
	} else {
		pairAFlagName = fmt.Sprintf("%v.pairA", cmdPrefix)
	}

	if err := registerModelPairAFlags(depth+1, pairAFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPairsPairB(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var pairBFlagName string
	if cmdPrefix == "" {
		pairBFlagName = "pairB"
	} else {
		pairBFlagName = fmt.Sprintf("%v.pairB", cmdPrefix)
	}

	if err := registerModelPairBFlags(depth+1, pairBFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPairsPairC(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var pairCFlagName string
	if cmdPrefix == "" {
		pairCFlagName = "pairC"
	} else {
		pairCFlagName = fmt.Sprintf("%v.pairC", cmdPrefix)
	}

	if err := registerModelPairCFlags(depth+1, pairCFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPairsPairD(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var pairDFlagName string
	if cmdPrefix == "" {
		pairDFlagName = "pairD"
	} else {
		pairDFlagName = fmt.Sprintf("%v.pairD", cmdPrefix)
	}

	if err := registerModelPairDFlags(depth+1, pairDFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPairsFlags(depth int, m *models.Pairs, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, pairAAdded := retrievePairsPairAFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairAAdded

	err, pairBAdded := retrievePairsPairBFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairBAdded

	err, pairCAdded := retrievePairsPairCFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairCAdded

	err, pairDAdded := retrievePairsPairDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairDAdded

	return nil, retAdded
}

func retrievePairsPairAFlags(depth int, m *models.Pairs, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pairAFlagName := fmt.Sprintf("%v.pairA", cmdPrefix)
	if cmd.Flags().Changed(pairAFlagName) {
		// info: complex object pairA PairA is retrieved outside this Changed() block
	}
	pairAFlagValue := m.PairA
	if swag.IsZero(pairAFlagValue) {
		pairAFlagValue = &models.PairA{}
	}

	err, pairAAdded := retrieveModelPairAFlags(depth+1, pairAFlagValue, pairAFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairAAdded
	if pairAAdded {
		m.PairA = pairAFlagValue
	}

	return nil, retAdded
}

func retrievePairsPairBFlags(depth int, m *models.Pairs, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pairBFlagName := fmt.Sprintf("%v.pairB", cmdPrefix)
	if cmd.Flags().Changed(pairBFlagName) {
		// info: complex object pairB PairB is retrieved outside this Changed() block
	}
	pairBFlagValue := m.PairB
	if swag.IsZero(pairBFlagValue) {
		pairBFlagValue = &models.PairB{}
	}

	err, pairBAdded := retrieveModelPairBFlags(depth+1, pairBFlagValue, pairBFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairBAdded
	if pairBAdded {
		m.PairB = pairBFlagValue
	}

	return nil, retAdded
}

func retrievePairsPairCFlags(depth int, m *models.Pairs, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pairCFlagName := fmt.Sprintf("%v.pairC", cmdPrefix)
	if cmd.Flags().Changed(pairCFlagName) {
		// info: complex object pairC PairC is retrieved outside this Changed() block
	}
	pairCFlagValue := m.PairC
	if swag.IsZero(pairCFlagValue) {
		pairCFlagValue = &models.PairC{}
	}

	err, pairCAdded := retrieveModelPairCFlags(depth+1, pairCFlagValue, pairCFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairCAdded
	if pairCAdded {
		m.PairC = pairCFlagValue
	}

	return nil, retAdded
}

func retrievePairsPairDFlags(depth int, m *models.Pairs, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pairDFlagName := fmt.Sprintf("%v.pairD", cmdPrefix)
	if cmd.Flags().Changed(pairDFlagName) {
		// info: complex object pairD PairD is retrieved outside this Changed() block
	}
	pairDFlagValue := m.PairD
	if swag.IsZero(pairDFlagValue) {
		pairDFlagValue = &models.PairD{}
	}

	err, pairDAdded := retrieveModelPairDFlags(depth+1, pairDFlagValue, pairDFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pairDAdded
	if pairDAdded {
		m.PairD = pairDFlagValue
	}

	return nil, retAdded
}