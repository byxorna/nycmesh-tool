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

// PostDevicesEroutersIDRouterOspfAreasReader is a Reader for the PostDevicesEroutersIDRouterOspfAreas structure.
type PostDevicesEroutersIDRouterOspfAreasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesEroutersIDRouterOspfAreasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesEroutersIDRouterOspfAreasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesEroutersIDRouterOspfAreasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesEroutersIDRouterOspfAreasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesEroutersIDRouterOspfAreasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesEroutersIDRouterOspfAreasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesEroutersIDRouterOspfAreasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesEroutersIDRouterOspfAreasOK creates a PostDevicesEroutersIDRouterOspfAreasOK with default headers values
func NewPostDevicesEroutersIDRouterOspfAreasOK() *PostDevicesEroutersIDRouterOspfAreasOK {
	return &PostDevicesEroutersIDRouterOspfAreasOK{}
}

/* PostDevicesEroutersIDRouterOspfAreasOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesEroutersIDRouterOspfAreasOK struct {
	Payload *models.Model101
}

func (o *PostDevicesEroutersIDRouterOspfAreasOK) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/router/ospf/areas][%d] postDevicesEroutersIdRouterOspfAreasOK  %+v", 200, o.Payload)
}
func (o *PostDevicesEroutersIDRouterOspfAreasOK) GetPayload() *models.Model101 {
	return o.Payload
}

func (o *PostDevicesEroutersIDRouterOspfAreasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model101)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDRouterOspfAreasBadRequest creates a PostDevicesEroutersIDRouterOspfAreasBadRequest with default headers values
func NewPostDevicesEroutersIDRouterOspfAreasBadRequest() *PostDevicesEroutersIDRouterOspfAreasBadRequest {
	return &PostDevicesEroutersIDRouterOspfAreasBadRequest{}
}

/* PostDevicesEroutersIDRouterOspfAreasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesEroutersIDRouterOspfAreasBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDRouterOspfAreasBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/router/ospf/areas][%d] postDevicesEroutersIdRouterOspfAreasBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesEroutersIDRouterOspfAreasBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDRouterOspfAreasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDRouterOspfAreasUnauthorized creates a PostDevicesEroutersIDRouterOspfAreasUnauthorized with default headers values
func NewPostDevicesEroutersIDRouterOspfAreasUnauthorized() *PostDevicesEroutersIDRouterOspfAreasUnauthorized {
	return &PostDevicesEroutersIDRouterOspfAreasUnauthorized{}
}

/* PostDevicesEroutersIDRouterOspfAreasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesEroutersIDRouterOspfAreasUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDRouterOspfAreasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/router/ospf/areas][%d] postDevicesEroutersIdRouterOspfAreasUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesEroutersIDRouterOspfAreasUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDRouterOspfAreasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDRouterOspfAreasForbidden creates a PostDevicesEroutersIDRouterOspfAreasForbidden with default headers values
func NewPostDevicesEroutersIDRouterOspfAreasForbidden() *PostDevicesEroutersIDRouterOspfAreasForbidden {
	return &PostDevicesEroutersIDRouterOspfAreasForbidden{}
}

/* PostDevicesEroutersIDRouterOspfAreasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesEroutersIDRouterOspfAreasForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDRouterOspfAreasForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/router/ospf/areas][%d] postDevicesEroutersIdRouterOspfAreasForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesEroutersIDRouterOspfAreasForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDRouterOspfAreasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDRouterOspfAreasNotFound creates a PostDevicesEroutersIDRouterOspfAreasNotFound with default headers values
func NewPostDevicesEroutersIDRouterOspfAreasNotFound() *PostDevicesEroutersIDRouterOspfAreasNotFound {
	return &PostDevicesEroutersIDRouterOspfAreasNotFound{}
}

/* PostDevicesEroutersIDRouterOspfAreasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesEroutersIDRouterOspfAreasNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDRouterOspfAreasNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/router/ospf/areas][%d] postDevicesEroutersIdRouterOspfAreasNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesEroutersIDRouterOspfAreasNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDRouterOspfAreasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesEroutersIDRouterOspfAreasInternalServerError creates a PostDevicesEroutersIDRouterOspfAreasInternalServerError with default headers values
func NewPostDevicesEroutersIDRouterOspfAreasInternalServerError() *PostDevicesEroutersIDRouterOspfAreasInternalServerError {
	return &PostDevicesEroutersIDRouterOspfAreasInternalServerError{}
}

/* PostDevicesEroutersIDRouterOspfAreasInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesEroutersIDRouterOspfAreasInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesEroutersIDRouterOspfAreasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/erouters/{id}/router/ospf/areas][%d] postDevicesEroutersIdRouterOspfAreasInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesEroutersIDRouterOspfAreasInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesEroutersIDRouterOspfAreasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
