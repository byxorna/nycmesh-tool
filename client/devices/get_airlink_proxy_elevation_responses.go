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

// GetAirlinkProxyElevationReader is a Reader for the GetAirlinkProxyElevation structure.
type GetAirlinkProxyElevationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAirlinkProxyElevationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetAirlinkProxyElevationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAirlinkProxyElevationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAirlinkProxyElevationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAirlinkProxyElevationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetAirlinkProxyElevationBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAirlinkProxyElevationBadRequest creates a GetAirlinkProxyElevationBadRequest with default headers values
func NewGetAirlinkProxyElevationBadRequest() *GetAirlinkProxyElevationBadRequest {
	return &GetAirlinkProxyElevationBadRequest{}
}

/* GetAirlinkProxyElevationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAirlinkProxyElevationBadRequest struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyElevationBadRequest) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/elevation][%d] getAirlinkProxyElevationBadRequest  %+v", 400, o.Payload)
}
func (o *GetAirlinkProxyElevationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyElevationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyElevationUnauthorized creates a GetAirlinkProxyElevationUnauthorized with default headers values
func NewGetAirlinkProxyElevationUnauthorized() *GetAirlinkProxyElevationUnauthorized {
	return &GetAirlinkProxyElevationUnauthorized{}
}

/* GetAirlinkProxyElevationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAirlinkProxyElevationUnauthorized struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyElevationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/elevation][%d] getAirlinkProxyElevationUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAirlinkProxyElevationUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyElevationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyElevationForbidden creates a GetAirlinkProxyElevationForbidden with default headers values
func NewGetAirlinkProxyElevationForbidden() *GetAirlinkProxyElevationForbidden {
	return &GetAirlinkProxyElevationForbidden{}
}

/* GetAirlinkProxyElevationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAirlinkProxyElevationForbidden struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyElevationForbidden) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/elevation][%d] getAirlinkProxyElevationForbidden  %+v", 403, o.Payload)
}
func (o *GetAirlinkProxyElevationForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyElevationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyElevationInternalServerError creates a GetAirlinkProxyElevationInternalServerError with default headers values
func NewGetAirlinkProxyElevationInternalServerError() *GetAirlinkProxyElevationInternalServerError {
	return &GetAirlinkProxyElevationInternalServerError{}
}

/* GetAirlinkProxyElevationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAirlinkProxyElevationInternalServerError struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyElevationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/elevation][%d] getAirlinkProxyElevationInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAirlinkProxyElevationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyElevationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyElevationBadGateway creates a GetAirlinkProxyElevationBadGateway with default headers values
func NewGetAirlinkProxyElevationBadGateway() *GetAirlinkProxyElevationBadGateway {
	return &GetAirlinkProxyElevationBadGateway{}
}

/* GetAirlinkProxyElevationBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type GetAirlinkProxyElevationBadGateway struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyElevationBadGateway) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/elevation][%d] getAirlinkProxyElevationBadGateway  %+v", 502, o.Payload)
}
func (o *GetAirlinkProxyElevationBadGateway) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyElevationBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
