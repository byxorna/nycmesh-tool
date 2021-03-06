// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetSitesReader is a Reader for the GetSites structure.
type GetSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSitesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSitesOK creates a GetSitesOK with default headers values
func NewGetSitesOK() *GetSitesOK {
	return &GetSitesOK{}
}

/* GetSitesOK describes a response with status code 200, with default header values.

Successful
*/
type GetSitesOK struct {
	Payload models.SitesList
}

func (o *GetSitesOK) Error() string {
	return fmt.Sprintf("[GET /sites][%d] getSitesOK  %+v", 200, o.Payload)
}
func (o *GetSitesOK) GetPayload() models.SitesList {
	return o.Payload
}

func (o *GetSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesBadRequest creates a GetSitesBadRequest with default headers values
func NewGetSitesBadRequest() *GetSitesBadRequest {
	return &GetSitesBadRequest{}
}

/* GetSitesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSitesBadRequest struct {
	Payload *models.Error
}

func (o *GetSitesBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites][%d] getSitesBadRequest  %+v", 400, o.Payload)
}
func (o *GetSitesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesUnauthorized creates a GetSitesUnauthorized with default headers values
func NewGetSitesUnauthorized() *GetSitesUnauthorized {
	return &GetSitesUnauthorized{}
}

/* GetSitesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSitesUnauthorized struct {
	Payload *models.Error
}

func (o *GetSitesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sites][%d] getSitesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSitesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesForbidden creates a GetSitesForbidden with default headers values
func NewGetSitesForbidden() *GetSitesForbidden {
	return &GetSitesForbidden{}
}

/* GetSitesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSitesForbidden struct {
	Payload *models.Error
}

func (o *GetSitesForbidden) Error() string {
	return fmt.Sprintf("[GET /sites][%d] getSitesForbidden  %+v", 403, o.Payload)
}
func (o *GetSitesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesInternalServerError creates a GetSitesInternalServerError with default headers values
func NewGetSitesInternalServerError() *GetSitesInternalServerError {
	return &GetSitesInternalServerError{}
}

/* GetSitesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSitesInternalServerError struct {
	Payload *models.Error
}

func (o *GetSitesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sites][%d] getSitesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSitesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
