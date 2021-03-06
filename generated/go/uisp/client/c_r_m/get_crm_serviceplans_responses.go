// Code generated by go-swagger; DO NOT EDIT.

package c_r_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetCrmServiceplansReader is a Reader for the GetCrmServiceplans structure.
type GetCrmServiceplansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCrmServiceplansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCrmServiceplansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCrmServiceplansUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCrmServiceplansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCrmServiceplansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCrmServiceplansServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCrmServiceplansOK creates a GetCrmServiceplansOK with default headers values
func NewGetCrmServiceplansOK() *GetCrmServiceplansOK {
	return &GetCrmServiceplansOK{}
}

/* GetCrmServiceplansOK describes a response with status code 200, with default header values.

Successful
*/
type GetCrmServiceplansOK struct {
	Payload models.CrmServicePlanListSchema
}

func (o *GetCrmServiceplansOK) Error() string {
	return fmt.Sprintf("[GET /crm/service-plans][%d] getCrmServiceplansOK  %+v", 200, o.Payload)
}
func (o *GetCrmServiceplansOK) GetPayload() models.CrmServicePlanListSchema {
	return o.Payload
}

func (o *GetCrmServiceplansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrmServiceplansUnauthorized creates a GetCrmServiceplansUnauthorized with default headers values
func NewGetCrmServiceplansUnauthorized() *GetCrmServiceplansUnauthorized {
	return &GetCrmServiceplansUnauthorized{}
}

/* GetCrmServiceplansUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCrmServiceplansUnauthorized struct {
	Payload *models.Error
}

func (o *GetCrmServiceplansUnauthorized) Error() string {
	return fmt.Sprintf("[GET /crm/service-plans][%d] getCrmServiceplansUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCrmServiceplansUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCrmServiceplansUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrmServiceplansForbidden creates a GetCrmServiceplansForbidden with default headers values
func NewGetCrmServiceplansForbidden() *GetCrmServiceplansForbidden {
	return &GetCrmServiceplansForbidden{}
}

/* GetCrmServiceplansForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCrmServiceplansForbidden struct {
	Payload *models.Error
}

func (o *GetCrmServiceplansForbidden) Error() string {
	return fmt.Sprintf("[GET /crm/service-plans][%d] getCrmServiceplansForbidden  %+v", 403, o.Payload)
}
func (o *GetCrmServiceplansForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCrmServiceplansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrmServiceplansInternalServerError creates a GetCrmServiceplansInternalServerError with default headers values
func NewGetCrmServiceplansInternalServerError() *GetCrmServiceplansInternalServerError {
	return &GetCrmServiceplansInternalServerError{}
}

/* GetCrmServiceplansInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCrmServiceplansInternalServerError struct {
	Payload *models.Error
}

func (o *GetCrmServiceplansInternalServerError) Error() string {
	return fmt.Sprintf("[GET /crm/service-plans][%d] getCrmServiceplansInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCrmServiceplansInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCrmServiceplansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCrmServiceplansServiceUnavailable creates a GetCrmServiceplansServiceUnavailable with default headers values
func NewGetCrmServiceplansServiceUnavailable() *GetCrmServiceplansServiceUnavailable {
	return &GetCrmServiceplansServiceUnavailable{}
}

/* GetCrmServiceplansServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type GetCrmServiceplansServiceUnavailable struct {
	Payload *models.Error
}

func (o *GetCrmServiceplansServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /crm/service-plans][%d] getCrmServiceplansServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCrmServiceplansServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCrmServiceplansServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
