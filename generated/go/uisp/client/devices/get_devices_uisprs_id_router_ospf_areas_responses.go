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

// GetDevicesUisprsIDRouterOspfAreasReader is a Reader for the GetDevicesUisprsIDRouterOspfAreas structure.
type GetDevicesUisprsIDRouterOspfAreasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesUisprsIDRouterOspfAreasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDevicesUisprsIDRouterOspfAreasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDevicesUisprsIDRouterOspfAreasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDevicesUisprsIDRouterOspfAreasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDevicesUisprsIDRouterOspfAreasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDevicesUisprsIDRouterOspfAreasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDevicesUisprsIDRouterOspfAreasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDevicesUisprsIDRouterOspfAreasOK creates a GetDevicesUisprsIDRouterOspfAreasOK with default headers values
func NewGetDevicesUisprsIDRouterOspfAreasOK() *GetDevicesUisprsIDRouterOspfAreasOK {
	return &GetDevicesUisprsIDRouterOspfAreasOK{}
}

/* GetDevicesUisprsIDRouterOspfAreasOK describes a response with status code 200, with default header values.

Successful
*/
type GetDevicesUisprsIDRouterOspfAreasOK struct {
	Payload string
}

func (o *GetDevicesUisprsIDRouterOspfAreasOK) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/router/ospf/areas][%d] getDevicesUisprsIdRouterOspfAreasOK  %+v", 200, o.Payload)
}
func (o *GetDevicesUisprsIDRouterOspfAreasOK) GetPayload() string {
	return o.Payload
}

func (o *GetDevicesUisprsIDRouterOspfAreasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDRouterOspfAreasBadRequest creates a GetDevicesUisprsIDRouterOspfAreasBadRequest with default headers values
func NewGetDevicesUisprsIDRouterOspfAreasBadRequest() *GetDevicesUisprsIDRouterOspfAreasBadRequest {
	return &GetDevicesUisprsIDRouterOspfAreasBadRequest{}
}

/* GetDevicesUisprsIDRouterOspfAreasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDevicesUisprsIDRouterOspfAreasBadRequest struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDRouterOspfAreasBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/router/ospf/areas][%d] getDevicesUisprsIdRouterOspfAreasBadRequest  %+v", 400, o.Payload)
}
func (o *GetDevicesUisprsIDRouterOspfAreasBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDRouterOspfAreasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDRouterOspfAreasUnauthorized creates a GetDevicesUisprsIDRouterOspfAreasUnauthorized with default headers values
func NewGetDevicesUisprsIDRouterOspfAreasUnauthorized() *GetDevicesUisprsIDRouterOspfAreasUnauthorized {
	return &GetDevicesUisprsIDRouterOspfAreasUnauthorized{}
}

/* GetDevicesUisprsIDRouterOspfAreasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDevicesUisprsIDRouterOspfAreasUnauthorized struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDRouterOspfAreasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/router/ospf/areas][%d] getDevicesUisprsIdRouterOspfAreasUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDevicesUisprsIDRouterOspfAreasUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDRouterOspfAreasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDRouterOspfAreasForbidden creates a GetDevicesUisprsIDRouterOspfAreasForbidden with default headers values
func NewGetDevicesUisprsIDRouterOspfAreasForbidden() *GetDevicesUisprsIDRouterOspfAreasForbidden {
	return &GetDevicesUisprsIDRouterOspfAreasForbidden{}
}

/* GetDevicesUisprsIDRouterOspfAreasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDevicesUisprsIDRouterOspfAreasForbidden struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDRouterOspfAreasForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/router/ospf/areas][%d] getDevicesUisprsIdRouterOspfAreasForbidden  %+v", 403, o.Payload)
}
func (o *GetDevicesUisprsIDRouterOspfAreasForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDRouterOspfAreasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDRouterOspfAreasNotFound creates a GetDevicesUisprsIDRouterOspfAreasNotFound with default headers values
func NewGetDevicesUisprsIDRouterOspfAreasNotFound() *GetDevicesUisprsIDRouterOspfAreasNotFound {
	return &GetDevicesUisprsIDRouterOspfAreasNotFound{}
}

/* GetDevicesUisprsIDRouterOspfAreasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDevicesUisprsIDRouterOspfAreasNotFound struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDRouterOspfAreasNotFound) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/router/ospf/areas][%d] getDevicesUisprsIdRouterOspfAreasNotFound  %+v", 404, o.Payload)
}
func (o *GetDevicesUisprsIDRouterOspfAreasNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDRouterOspfAreasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDevicesUisprsIDRouterOspfAreasInternalServerError creates a GetDevicesUisprsIDRouterOspfAreasInternalServerError with default headers values
func NewGetDevicesUisprsIDRouterOspfAreasInternalServerError() *GetDevicesUisprsIDRouterOspfAreasInternalServerError {
	return &GetDevicesUisprsIDRouterOspfAreasInternalServerError{}
}

/* GetDevicesUisprsIDRouterOspfAreasInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDevicesUisprsIDRouterOspfAreasInternalServerError struct {
	Payload *models.Error
}

func (o *GetDevicesUisprsIDRouterOspfAreasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /devices/uisprs/{id}/router/ospf/areas][%d] getDevicesUisprsIdRouterOspfAreasInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDevicesUisprsIDRouterOspfAreasInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDevicesUisprsIDRouterOspfAreasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
