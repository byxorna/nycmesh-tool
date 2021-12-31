// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostDevicesIDRouterOspfAreasReader is a Reader for the PostDevicesIDRouterOspfAreas structure.
type PostDevicesIDRouterOspfAreasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesIDRouterOspfAreasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesIDRouterOspfAreasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesIDRouterOspfAreasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesIDRouterOspfAreasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesIDRouterOspfAreasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesIDRouterOspfAreasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesIDRouterOspfAreasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesIDRouterOspfAreasOK creates a PostDevicesIDRouterOspfAreasOK with default headers values
func NewPostDevicesIDRouterOspfAreasOK() *PostDevicesIDRouterOspfAreasOK {
	return &PostDevicesIDRouterOspfAreasOK{}
}

/* PostDevicesIDRouterOspfAreasOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesIDRouterOspfAreasOK struct {
	Payload *models.Model99
}

func (o *PostDevicesIDRouterOspfAreasOK) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/router/ospf/areas][%d] postDevicesIdRouterOspfAreasOK  %+v", 200, o.Payload)
}
func (o *PostDevicesIDRouterOspfAreasOK) GetPayload() *models.Model99 {
	return o.Payload
}

func (o *PostDevicesIDRouterOspfAreasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model99)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRouterOspfAreasBadRequest creates a PostDevicesIDRouterOspfAreasBadRequest with default headers values
func NewPostDevicesIDRouterOspfAreasBadRequest() *PostDevicesIDRouterOspfAreasBadRequest {
	return &PostDevicesIDRouterOspfAreasBadRequest{}
}

/* PostDevicesIDRouterOspfAreasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesIDRouterOspfAreasBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesIDRouterOspfAreasBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/router/ospf/areas][%d] postDevicesIdRouterOspfAreasBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesIDRouterOspfAreasBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRouterOspfAreasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRouterOspfAreasUnauthorized creates a PostDevicesIDRouterOspfAreasUnauthorized with default headers values
func NewPostDevicesIDRouterOspfAreasUnauthorized() *PostDevicesIDRouterOspfAreasUnauthorized {
	return &PostDevicesIDRouterOspfAreasUnauthorized{}
}

/* PostDevicesIDRouterOspfAreasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesIDRouterOspfAreasUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesIDRouterOspfAreasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/router/ospf/areas][%d] postDevicesIdRouterOspfAreasUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesIDRouterOspfAreasUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRouterOspfAreasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRouterOspfAreasForbidden creates a PostDevicesIDRouterOspfAreasForbidden with default headers values
func NewPostDevicesIDRouterOspfAreasForbidden() *PostDevicesIDRouterOspfAreasForbidden {
	return &PostDevicesIDRouterOspfAreasForbidden{}
}

/* PostDevicesIDRouterOspfAreasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesIDRouterOspfAreasForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesIDRouterOspfAreasForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/router/ospf/areas][%d] postDevicesIdRouterOspfAreasForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesIDRouterOspfAreasForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRouterOspfAreasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRouterOspfAreasNotFound creates a PostDevicesIDRouterOspfAreasNotFound with default headers values
func NewPostDevicesIDRouterOspfAreasNotFound() *PostDevicesIDRouterOspfAreasNotFound {
	return &PostDevicesIDRouterOspfAreasNotFound{}
}

/* PostDevicesIDRouterOspfAreasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesIDRouterOspfAreasNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesIDRouterOspfAreasNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/router/ospf/areas][%d] postDevicesIdRouterOspfAreasNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesIDRouterOspfAreasNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRouterOspfAreasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesIDRouterOspfAreasInternalServerError creates a PostDevicesIDRouterOspfAreasInternalServerError with default headers values
func NewPostDevicesIDRouterOspfAreasInternalServerError() *PostDevicesIDRouterOspfAreasInternalServerError {
	return &PostDevicesIDRouterOspfAreasInternalServerError{}
}

/* PostDevicesIDRouterOspfAreasInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesIDRouterOspfAreasInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesIDRouterOspfAreasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/{id}/router/ospf/areas][%d] postDevicesIdRouterOspfAreasInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesIDRouterOspfAreasInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesIDRouterOspfAreasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
