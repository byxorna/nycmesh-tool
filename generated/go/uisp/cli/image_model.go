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

// Schema cli for Image

// register flags to command
func registerModelImageFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageDate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageFileName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageFileType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageFullURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageFullURLRelative(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageHeight(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageIdentification(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageOrder(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageThumbURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageThumbURLRelative(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageWidth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageDate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: date Date map type is not supported by go-swagger cli yet

	return nil
}

func registerImageDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Required. `

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerImageFileName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileNameDescription := `Required. `

	var fileNameFlagName string
	if cmdPrefix == "" {
		fileNameFlagName = "fileName"
	} else {
		fileNameFlagName = fmt.Sprintf("%v.fileName", cmdPrefix)
	}

	var fileNameFlagDefault string

	_ = cmd.PersistentFlags().String(fileNameFlagName, fileNameFlagDefault, fileNameDescription)

	return nil
}

func registerImageFileType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileTypeDescription := `Required. `

	var fileTypeFlagName string
	if cmdPrefix == "" {
		fileTypeFlagName = "fileType"
	} else {
		fileTypeFlagName = fmt.Sprintf("%v.fileType", cmdPrefix)
	}

	var fileTypeFlagDefault string

	_ = cmd.PersistentFlags().String(fileTypeFlagName, fileTypeFlagDefault, fileTypeDescription)

	return nil
}

func registerImageFullURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fullUrlDescription := `Required. `

	var fullUrlFlagName string
	if cmdPrefix == "" {
		fullUrlFlagName = "fullUrl"
	} else {
		fullUrlFlagName = fmt.Sprintf("%v.fullUrl", cmdPrefix)
	}

	var fullUrlFlagDefault string

	_ = cmd.PersistentFlags().String(fullUrlFlagName, fullUrlFlagDefault, fullUrlDescription)

	return nil
}

func registerImageFullURLRelative(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fullUrlRelativeDescription := ``

	var fullUrlRelativeFlagName string
	if cmdPrefix == "" {
		fullUrlRelativeFlagName = "fullUrlRelative"
	} else {
		fullUrlRelativeFlagName = fmt.Sprintf("%v.fullUrlRelative", cmdPrefix)
	}

	var fullUrlRelativeFlagDefault string

	_ = cmd.PersistentFlags().String(fullUrlRelativeFlagName, fullUrlRelativeFlagDefault, fullUrlRelativeDescription)

	return nil
}

func registerImageHeight(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	heightDescription := `Required. `

	var heightFlagName string
	if cmdPrefix == "" {
		heightFlagName = "height"
	} else {
		heightFlagName = fmt.Sprintf("%v.height", cmdPrefix)
	}

	var heightFlagDefault float64

	_ = cmd.PersistentFlags().Float64(heightFlagName, heightFlagDefault, heightDescription)

	return nil
}

func registerImageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerImageIdentification(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var identificationFlagName string
	if cmdPrefix == "" {
		identificationFlagName = "identification"
	} else {
		identificationFlagName = fmt.Sprintf("%v.identification", cmdPrefix)
	}

	if err := registerModelImageIdentificationFlags(depth+1, identificationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerImageOrder(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	orderDescription := `Required. `

	var orderFlagName string
	if cmdPrefix == "" {
		orderFlagName = "order"
	} else {
		orderFlagName = fmt.Sprintf("%v.order", cmdPrefix)
	}

	var orderFlagDefault float64

	_ = cmd.PersistentFlags().Float64(orderFlagName, orderFlagDefault, orderDescription)

	return nil
}

func registerImageSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `Required. `

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault float64

	_ = cmd.PersistentFlags().Float64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerImageThumbURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	thumbUrlDescription := `Required. `

	var thumbUrlFlagName string
	if cmdPrefix == "" {
		thumbUrlFlagName = "thumbUrl"
	} else {
		thumbUrlFlagName = fmt.Sprintf("%v.thumbUrl", cmdPrefix)
	}

	var thumbUrlFlagDefault string

	_ = cmd.PersistentFlags().String(thumbUrlFlagName, thumbUrlFlagDefault, thumbUrlDescription)

	return nil
}

func registerImageThumbURLRelative(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	thumbUrlRelativeDescription := ``

	var thumbUrlRelativeFlagName string
	if cmdPrefix == "" {
		thumbUrlRelativeFlagName = "thumbUrlRelative"
	} else {
		thumbUrlRelativeFlagName = fmt.Sprintf("%v.thumbUrlRelative", cmdPrefix)
	}

	var thumbUrlRelativeFlagDefault string

	_ = cmd.PersistentFlags().String(thumbUrlRelativeFlagName, thumbUrlRelativeFlagDefault, thumbUrlRelativeDescription)

	return nil
}

func registerImageWidth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	widthDescription := `Required. `

	var widthFlagName string
	if cmdPrefix == "" {
		widthFlagName = "width"
	} else {
		widthFlagName = fmt.Sprintf("%v.width", cmdPrefix)
	}

	var widthFlagDefault float64

	_ = cmd.PersistentFlags().Float64(widthFlagName, widthFlagDefault, widthDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dateAdded := retrieveImageDateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dateAdded

	err, descriptionAdded := retrieveImageDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, fileNameAdded := retrieveImageFileNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileNameAdded

	err, fileTypeAdded := retrieveImageFileTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileTypeAdded

	err, fullUrlAdded := retrieveImageFullURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fullUrlAdded

	err, fullUrlRelativeAdded := retrieveImageFullURLRelativeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fullUrlRelativeAdded

	err, heightAdded := retrieveImageHeightFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || heightAdded

	err, idAdded := retrieveImageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, identificationAdded := retrieveImageIdentificationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded

	err, nameAdded := retrieveImageNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, orderAdded := retrieveImageOrderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || orderAdded

	err, sizeAdded := retrieveImageSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, thumbUrlAdded := retrieveImageThumbURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || thumbUrlAdded

	err, thumbUrlRelativeAdded := retrieveImageThumbURLRelativeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || thumbUrlRelativeAdded

	err, widthAdded := retrieveImageWidthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || widthAdded

	return nil, retAdded
}

func retrieveImageDateFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dateFlagName := fmt.Sprintf("%v.date", cmdPrefix)
	if cmd.Flags().Changed(dateFlagName) {
		// warning: date map type Date is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveImageDescriptionFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = &descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageFileNameFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileNameFlagName := fmt.Sprintf("%v.fileName", cmdPrefix)
	if cmd.Flags().Changed(fileNameFlagName) {

		var fileNameFlagName string
		if cmdPrefix == "" {
			fileNameFlagName = "fileName"
		} else {
			fileNameFlagName = fmt.Sprintf("%v.fileName", cmdPrefix)
		}

		fileNameFlagValue, err := cmd.Flags().GetString(fileNameFlagName)
		if err != nil {
			return err, false
		}
		m.FileName = &fileNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageFileTypeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileTypeFlagName := fmt.Sprintf("%v.fileType", cmdPrefix)
	if cmd.Flags().Changed(fileTypeFlagName) {

		var fileTypeFlagName string
		if cmdPrefix == "" {
			fileTypeFlagName = "fileType"
		} else {
			fileTypeFlagName = fmt.Sprintf("%v.fileType", cmdPrefix)
		}

		fileTypeFlagValue, err := cmd.Flags().GetString(fileTypeFlagName)
		if err != nil {
			return err, false
		}
		m.FileType = &fileTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageFullURLFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fullUrlFlagName := fmt.Sprintf("%v.fullUrl", cmdPrefix)
	if cmd.Flags().Changed(fullUrlFlagName) {

		var fullUrlFlagName string
		if cmdPrefix == "" {
			fullUrlFlagName = "fullUrl"
		} else {
			fullUrlFlagName = fmt.Sprintf("%v.fullUrl", cmdPrefix)
		}

		fullUrlFlagValue, err := cmd.Flags().GetString(fullUrlFlagName)
		if err != nil {
			return err, false
		}
		m.FullURL = &fullUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageFullURLRelativeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fullUrlRelativeFlagName := fmt.Sprintf("%v.fullUrlRelative", cmdPrefix)
	if cmd.Flags().Changed(fullUrlRelativeFlagName) {

		var fullUrlRelativeFlagName string
		if cmdPrefix == "" {
			fullUrlRelativeFlagName = "fullUrlRelative"
		} else {
			fullUrlRelativeFlagName = fmt.Sprintf("%v.fullUrlRelative", cmdPrefix)
		}

		fullUrlRelativeFlagValue, err := cmd.Flags().GetString(fullUrlRelativeFlagName)
		if err != nil {
			return err, false
		}
		m.FullURLRelative = fullUrlRelativeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageHeightFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	heightFlagName := fmt.Sprintf("%v.height", cmdPrefix)
	if cmd.Flags().Changed(heightFlagName) {

		var heightFlagName string
		if cmdPrefix == "" {
			heightFlagName = "height"
		} else {
			heightFlagName = fmt.Sprintf("%v.height", cmdPrefix)
		}

		heightFlagValue, err := cmd.Flags().GetFloat64(heightFlagName)
		if err != nil {
			return err, false
		}
		m.Height = &heightFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageIDFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageIdentificationFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	identificationFlagName := fmt.Sprintf("%v.identification", cmdPrefix)
	if cmd.Flags().Changed(identificationFlagName) {
		// info: complex object identification ImageIdentification is retrieved outside this Changed() block
	}
	identificationFlagValue := m.Identification
	if swag.IsZero(identificationFlagValue) {
		identificationFlagValue = &models.ImageIdentification{}
	}

	err, identificationAdded := retrieveModelImageIdentificationFlags(depth+1, identificationFlagValue, identificationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || identificationAdded
	if identificationAdded {
		m.Identification = identificationFlagValue
	}

	return nil, retAdded
}

func retrieveImageNameFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageOrderFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	orderFlagName := fmt.Sprintf("%v.order", cmdPrefix)
	if cmd.Flags().Changed(orderFlagName) {

		var orderFlagName string
		if cmdPrefix == "" {
			orderFlagName = "order"
		} else {
			orderFlagName = fmt.Sprintf("%v.order", cmdPrefix)
		}

		orderFlagValue, err := cmd.Flags().GetFloat64(orderFlagName)
		if err != nil {
			return err, false
		}
		m.Order = &orderFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageSizeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeFlagName := fmt.Sprintf("%v.size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetFloat64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = &sizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageThumbURLFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	thumbUrlFlagName := fmt.Sprintf("%v.thumbUrl", cmdPrefix)
	if cmd.Flags().Changed(thumbUrlFlagName) {

		var thumbUrlFlagName string
		if cmdPrefix == "" {
			thumbUrlFlagName = "thumbUrl"
		} else {
			thumbUrlFlagName = fmt.Sprintf("%v.thumbUrl", cmdPrefix)
		}

		thumbUrlFlagValue, err := cmd.Flags().GetString(thumbUrlFlagName)
		if err != nil {
			return err, false
		}
		m.ThumbURL = &thumbUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageThumbURLRelativeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	thumbUrlRelativeFlagName := fmt.Sprintf("%v.thumbUrlRelative", cmdPrefix)
	if cmd.Flags().Changed(thumbUrlRelativeFlagName) {

		var thumbUrlRelativeFlagName string
		if cmdPrefix == "" {
			thumbUrlRelativeFlagName = "thumbUrlRelative"
		} else {
			thumbUrlRelativeFlagName = fmt.Sprintf("%v.thumbUrlRelative", cmdPrefix)
		}

		thumbUrlRelativeFlagValue, err := cmd.Flags().GetString(thumbUrlRelativeFlagName)
		if err != nil {
			return err, false
		}
		m.ThumbURLRelative = thumbUrlRelativeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageWidthFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	widthFlagName := fmt.Sprintf("%v.width", cmdPrefix)
	if cmd.Flags().Changed(widthFlagName) {

		var widthFlagName string
		if cmdPrefix == "" {
			widthFlagName = "width"
		} else {
			widthFlagName = fmt.Sprintf("%v.width", cmdPrefix)
		}

		widthFlagValue, err := cmd.Flags().GetFloat64(widthFlagName)
		if err != nil {
			return err, false
		}
		m.Width = &widthFlagValue

		retAdded = true
	}

	return nil, retAdded
}
