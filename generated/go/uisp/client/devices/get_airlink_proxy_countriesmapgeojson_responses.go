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

// GetAirlinkProxyCountriesmapgeojsonReader is a Reader for the GetAirlinkProxyCountriesmapgeojson structure.
type GetAirlinkProxyCountriesmapgeojsonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAirlinkProxyCountriesmapgeojsonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetAirlinkProxyCountriesmapgeojsonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAirlinkProxyCountriesmapgeojsonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAirlinkProxyCountriesmapgeojsonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAirlinkProxyCountriesmapgeojsonInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetAirlinkProxyCountriesmapgeojsonBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAirlinkProxyCountriesmapgeojsonBadRequest creates a GetAirlinkProxyCountriesmapgeojsonBadRequest with default headers values
func NewGetAirlinkProxyCountriesmapgeojsonBadRequest() *GetAirlinkProxyCountriesmapgeojsonBadRequest {
	return &GetAirlinkProxyCountriesmapgeojsonBadRequest{}
}

/* GetAirlinkProxyCountriesmapgeojsonBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAirlinkProxyCountriesmapgeojsonBadRequest struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyCountriesmapgeojsonBadRequest) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/countries-map.geo.json][%d] getAirlinkProxyCountriesmapgeojsonBadRequest  %+v", 400, o.Payload)
}
func (o *GetAirlinkProxyCountriesmapgeojsonBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyCountriesmapgeojsonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyCountriesmapgeojsonUnauthorized creates a GetAirlinkProxyCountriesmapgeojsonUnauthorized with default headers values
func NewGetAirlinkProxyCountriesmapgeojsonUnauthorized() *GetAirlinkProxyCountriesmapgeojsonUnauthorized {
	return &GetAirlinkProxyCountriesmapgeojsonUnauthorized{}
}

/* GetAirlinkProxyCountriesmapgeojsonUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAirlinkProxyCountriesmapgeojsonUnauthorized struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyCountriesmapgeojsonUnauthorized) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/countries-map.geo.json][%d] getAirlinkProxyCountriesmapgeojsonUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAirlinkProxyCountriesmapgeojsonUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyCountriesmapgeojsonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyCountriesmapgeojsonForbidden creates a GetAirlinkProxyCountriesmapgeojsonForbidden with default headers values
func NewGetAirlinkProxyCountriesmapgeojsonForbidden() *GetAirlinkProxyCountriesmapgeojsonForbidden {
	return &GetAirlinkProxyCountriesmapgeojsonForbidden{}
}

/* GetAirlinkProxyCountriesmapgeojsonForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAirlinkProxyCountriesmapgeojsonForbidden struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyCountriesmapgeojsonForbidden) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/countries-map.geo.json][%d] getAirlinkProxyCountriesmapgeojsonForbidden  %+v", 403, o.Payload)
}
func (o *GetAirlinkProxyCountriesmapgeojsonForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyCountriesmapgeojsonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyCountriesmapgeojsonInternalServerError creates a GetAirlinkProxyCountriesmapgeojsonInternalServerError with default headers values
func NewGetAirlinkProxyCountriesmapgeojsonInternalServerError() *GetAirlinkProxyCountriesmapgeojsonInternalServerError {
	return &GetAirlinkProxyCountriesmapgeojsonInternalServerError{}
}

/* GetAirlinkProxyCountriesmapgeojsonInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAirlinkProxyCountriesmapgeojsonInternalServerError struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyCountriesmapgeojsonInternalServerError) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/countries-map.geo.json][%d] getAirlinkProxyCountriesmapgeojsonInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAirlinkProxyCountriesmapgeojsonInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyCountriesmapgeojsonInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAirlinkProxyCountriesmapgeojsonBadGateway creates a GetAirlinkProxyCountriesmapgeojsonBadGateway with default headers values
func NewGetAirlinkProxyCountriesmapgeojsonBadGateway() *GetAirlinkProxyCountriesmapgeojsonBadGateway {
	return &GetAirlinkProxyCountriesmapgeojsonBadGateway{}
}

/* GetAirlinkProxyCountriesmapgeojsonBadGateway describes a response with status code 502, with default header values.

Bad Gateway
*/
type GetAirlinkProxyCountriesmapgeojsonBadGateway struct {
	Payload *models.Error
}

func (o *GetAirlinkProxyCountriesmapgeojsonBadGateway) Error() string {
	return fmt.Sprintf("[GET /airlink/proxy/countries-map.geo.json][%d] getAirlinkProxyCountriesmapgeojsonBadGateway  %+v", 502, o.Payload)
}
func (o *GetAirlinkProxyCountriesmapgeojsonBadGateway) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAirlinkProxyCountriesmapgeojsonBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
