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

// PostDevicesUisprsIDRouterOspfAreasReader is a Reader for the PostDevicesUisprsIDRouterOspfAreas structure.
type PostDevicesUisprsIDRouterOspfAreasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDevicesUisprsIDRouterOspfAreasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDevicesUisprsIDRouterOspfAreasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDevicesUisprsIDRouterOspfAreasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDevicesUisprsIDRouterOspfAreasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDevicesUisprsIDRouterOspfAreasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDevicesUisprsIDRouterOspfAreasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDevicesUisprsIDRouterOspfAreasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDevicesUisprsIDRouterOspfAreasOK creates a PostDevicesUisprsIDRouterOspfAreasOK with default headers values
func NewPostDevicesUisprsIDRouterOspfAreasOK() *PostDevicesUisprsIDRouterOspfAreasOK {
	return &PostDevicesUisprsIDRouterOspfAreasOK{}
}

/* PostDevicesUisprsIDRouterOspfAreasOK describes a response with status code 200, with default header values.

Successful
*/
type PostDevicesUisprsIDRouterOspfAreasOK struct {
	Payload *models.Model103
}

func (o *PostDevicesUisprsIDRouterOspfAreasOK) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/router/ospf/areas][%d] postDevicesUisprsIdRouterOspfAreasOK  %+v", 200, o.Payload)
}
func (o *PostDevicesUisprsIDRouterOspfAreasOK) GetPayload() *models.Model103 {
	return o.Payload
}

func (o *PostDevicesUisprsIDRouterOspfAreasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Model103)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDRouterOspfAreasBadRequest creates a PostDevicesUisprsIDRouterOspfAreasBadRequest with default headers values
func NewPostDevicesUisprsIDRouterOspfAreasBadRequest() *PostDevicesUisprsIDRouterOspfAreasBadRequest {
	return &PostDevicesUisprsIDRouterOspfAreasBadRequest{}
}

/* PostDevicesUisprsIDRouterOspfAreasBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDevicesUisprsIDRouterOspfAreasBadRequest struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDRouterOspfAreasBadRequest) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/router/ospf/areas][%d] postDevicesUisprsIdRouterOspfAreasBadRequest  %+v", 400, o.Payload)
}
func (o *PostDevicesUisprsIDRouterOspfAreasBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDRouterOspfAreasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDRouterOspfAreasUnauthorized creates a PostDevicesUisprsIDRouterOspfAreasUnauthorized with default headers values
func NewPostDevicesUisprsIDRouterOspfAreasUnauthorized() *PostDevicesUisprsIDRouterOspfAreasUnauthorized {
	return &PostDevicesUisprsIDRouterOspfAreasUnauthorized{}
}

/* PostDevicesUisprsIDRouterOspfAreasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostDevicesUisprsIDRouterOspfAreasUnauthorized struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDRouterOspfAreasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/router/ospf/areas][%d] postDevicesUisprsIdRouterOspfAreasUnauthorized  %+v", 401, o.Payload)
}
func (o *PostDevicesUisprsIDRouterOspfAreasUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDRouterOspfAreasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDRouterOspfAreasForbidden creates a PostDevicesUisprsIDRouterOspfAreasForbidden with default headers values
func NewPostDevicesUisprsIDRouterOspfAreasForbidden() *PostDevicesUisprsIDRouterOspfAreasForbidden {
	return &PostDevicesUisprsIDRouterOspfAreasForbidden{}
}

/* PostDevicesUisprsIDRouterOspfAreasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostDevicesUisprsIDRouterOspfAreasForbidden struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDRouterOspfAreasForbidden) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/router/ospf/areas][%d] postDevicesUisprsIdRouterOspfAreasForbidden  %+v", 403, o.Payload)
}
func (o *PostDevicesUisprsIDRouterOspfAreasForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDRouterOspfAreasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDRouterOspfAreasNotFound creates a PostDevicesUisprsIDRouterOspfAreasNotFound with default headers values
func NewPostDevicesUisprsIDRouterOspfAreasNotFound() *PostDevicesUisprsIDRouterOspfAreasNotFound {
	return &PostDevicesUisprsIDRouterOspfAreasNotFound{}
}

/* PostDevicesUisprsIDRouterOspfAreasNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostDevicesUisprsIDRouterOspfAreasNotFound struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDRouterOspfAreasNotFound) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/router/ospf/areas][%d] postDevicesUisprsIdRouterOspfAreasNotFound  %+v", 404, o.Payload)
}
func (o *PostDevicesUisprsIDRouterOspfAreasNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDRouterOspfAreasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDevicesUisprsIDRouterOspfAreasInternalServerError creates a PostDevicesUisprsIDRouterOspfAreasInternalServerError with default headers values
func NewPostDevicesUisprsIDRouterOspfAreasInternalServerError() *PostDevicesUisprsIDRouterOspfAreasInternalServerError {
	return &PostDevicesUisprsIDRouterOspfAreasInternalServerError{}
}

/* PostDevicesUisprsIDRouterOspfAreasInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDevicesUisprsIDRouterOspfAreasInternalServerError struct {
	Payload *models.Error
}

func (o *PostDevicesUisprsIDRouterOspfAreasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /devices/uisprs/{id}/router/ospf/areas][%d] postDevicesUisprsIdRouterOspfAreasInternalServerError  %+v", 500, o.Payload)
}
func (o *PostDevicesUisprsIDRouterOspfAreasInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDevicesUisprsIDRouterOspfAreasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}