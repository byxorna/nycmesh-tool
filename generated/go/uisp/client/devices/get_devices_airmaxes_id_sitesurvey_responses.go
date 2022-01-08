// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetDevicesAirmaxesIDSitesurveyReader is a Reader for the GetDevicesAirmaxesIDSitesurvey structure.
type GetDevicesAirmaxesIDSitesurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesAirmaxesIDSitesurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesAirmaxesIDSitesurveyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesAirmaxesIDSitesurveyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesAirmaxesIDSitesurveyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesAirmaxesIDSitesurveyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesAirmaxesIDSitesurveyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesAirmaxesIDSitesurveyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesAirmaxesIDSitesurveyOK creates a GetDevicesAirmaxesIDSitesurveyOK with default headers values
func NewGetDevicesAirmaxesIDSitesurveyOK() *GetDevicesAirmaxesIDSitesurveyOK {
	return &GetDevicesAirmaxesIDSitesurveyOK{}
}

/* GetDevicesAirmaxesIDSitesurveyOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesAirmaxesIDSitesurveyOK struct {
	Payload models.ListOfSitesAround
}

func (o *GetDevicesAirmaxesIDSitesurveyOK) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/site-survey][%d] getDevicesAirmaxesIdSitesurveyOK  %+v", 200, o.Payload)
}
func (o *GetDevicesAirmaxesIDSitesurveyOK) GetPayload() models.ListOfSitesAround {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSitesurveyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSitesurveyBadRequest creates a GetDevicesAirmaxesIDSitesurveyBadRequest with default headers values
func NewGetDevicesAirmaxesIDSitesurveyBadRequest() *GetDevicesAirmaxesIDSitesurveyBadRequest {
	return &GetDevicesAirmaxesIDSitesurveyBadRequest{}
}

/* GetDevicesAirmaxesIDSitesurveyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesAirmaxesIDSitesurveyBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSitesurveyBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/site-survey][%d] getDevicesAirmaxesIdSitesurveyBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesAirmaxesIDSitesurveyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSitesurveyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSitesurveyUnauthorized creates a GetDevicesAirmaxesIDSitesurveyUnauthorized with default headers values
func NewGetDevicesAirmaxesIDSitesurveyUnauthorized() *GetDevicesAirmaxesIDSitesurveyUnauthorized {
	return &GetDevicesAirmaxesIDSitesurveyUnauthorized{}
}

/* GetDevicesAirmaxesIDSitesurveyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesAirmaxesIDSitesurveyUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSitesurveyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/site-survey][%d] getDevicesAirmaxesIdSitesurveyUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesAirmaxesIDSitesurveyUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSitesurveyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSitesurveyForbidden creates a GetDevicesAirmaxesIDSitesurveyForbidden with default headers values
func NewGetDevicesAirmaxesIDSitesurveyForbidden() *GetDevicesAirmaxesIDSitesurveyForbidden {
	return &GetDevicesAirmaxesIDSitesurveyForbidden{}
}

/* GetDevicesAirmaxesIDSitesurveyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesAirmaxesIDSitesurveyForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSitesurveyForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/site-survey][%d] getDevicesAirmaxesIdSitesurveyForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesAirmaxesIDSitesurveyForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSitesurveyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSitesurveyNotFound creates a GetDevicesAirmaxesIDSitesurveyNotFound with default headers values
func NewGetDevicesAirmaxesIDSitesurveyNotFound() *GetDevicesAirmaxesIDSitesurveyNotFound {
	return &GetDevicesAirmaxesIDSitesurveyNotFound{}
}

/* GetDevicesAirmaxesIDSitesurveyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesAirmaxesIDSitesurveyNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSitesurveyNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/site-survey][%d] getDevicesAirmaxesIdSitesurveyNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesAirmaxesIDSitesurveyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSitesurveyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesAirmaxesIDSitesurveyInternalServerError creates a GetDevicesAirmaxesIDSitesurveyInternalServerError with default headers values
func NewGetDevicesAirmaxesIDSitesurveyInternalServerError() *GetDevicesAirmaxesIDSitesurveyInternalServerError {
	return &GetDevicesAirmaxesIDSitesurveyInternalServerError{}
}

/* GetDevicesAirmaxesIDSitesurveyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesAirmaxesIDSitesurveyInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesAirmaxesIDSitesurveyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/airmaxes/{id}/site-survey][%d] getDevicesAirmaxesIdSitesurveyInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesAirmaxesIDSitesurveyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesAirmaxesIDSitesurveyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}