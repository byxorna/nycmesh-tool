// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetDiscoveryScanstatusReader is a Reader for the GetDiscoveryScanstatus structure.
type GetDiscoveryScanstatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoveryScanstatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoveryScanstatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiscoveryScanstatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDiscoveryScanstatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDiscoveryScanstatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoveryScanstatusOK creates a GetDiscoveryScanstatusOK with default headers values
func NewGetDiscoveryScanstatusOK() *GetDiscoveryScanstatusOK {
	return &GetDiscoveryScanstatusOK{}
}

/* GetDiscoveryScanstatusOK describes a response with status code 200, with default header values.

Successful
*/
type GetDiscoveryScanstatusOK struct {
	Payload string
}

func (o *GetDiscoveryScanstatusOK) Error() string {
	return fmt.Sprintf("[GET /discovery/scan-status][%d] getDiscoveryScanstatusOK  %+v", 200, o.Payload)
}
func (o *GetDiscoveryScanstatusOK) GetPayload() string {
	return o.Payload
}

func (o *GetDiscoveryScanstatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryScanstatusBadRequest creates a GetDiscoveryScanstatusBadRequest with default headers values
func NewGetDiscoveryScanstatusBadRequest() *GetDiscoveryScanstatusBadRequest {
	return &GetDiscoveryScanstatusBadRequest{}
}

/* GetDiscoveryScanstatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDiscoveryScanstatusBadRequest struct {
	Payload *models.Error
}

func (o *GetDiscoveryScanstatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /discovery/scan-status][%d] getDiscoveryScanstatusBadRequest  %+v", 400, o.Payload)
}
func (o *GetDiscoveryScanstatusBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryScanstatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryScanstatusUnauthorized creates a GetDiscoveryScanstatusUnauthorized with default headers values
func NewGetDiscoveryScanstatusUnauthorized() *GetDiscoveryScanstatusUnauthorized {
	return &GetDiscoveryScanstatusUnauthorized{}
}

/* GetDiscoveryScanstatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDiscoveryScanstatusUnauthorized struct {
	Payload *models.Error
}

func (o *GetDiscoveryScanstatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /discovery/scan-status][%d] getDiscoveryScanstatusUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDiscoveryScanstatusUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryScanstatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryScanstatusInternalServerError creates a GetDiscoveryScanstatusInternalServerError with default headers values
func NewGetDiscoveryScanstatusInternalServerError() *GetDiscoveryScanstatusInternalServerError {
	return &GetDiscoveryScanstatusInternalServerError{}
}

/* GetDiscoveryScanstatusInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDiscoveryScanstatusInternalServerError struct {
	Payload *models.Error
}

func (o *GetDiscoveryScanstatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /discovery/scan-status][%d] getDiscoveryScanstatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDiscoveryScanstatusInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiscoveryScanstatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
