// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetNmsTrafficSubnetsReader is a Reader for the GetNmsTrafficSubnets structure.
type GetNmsTrafficSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsTrafficSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNmsTrafficSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNmsTrafficSubnetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNmsTrafficSubnetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNmsTrafficSubnetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNmsTrafficSubnetsOK creates a GetNmsTrafficSubnetsOK with default headers values
func NewGetNmsTrafficSubnetsOK() *GetNmsTrafficSubnetsOK {
	return &GetNmsTrafficSubnetsOK{}
}

/* GetNmsTrafficSubnetsOK describes a response with status code 200, with default header values.

Successful
*/
type GetNmsTrafficSubnetsOK struct {
	Payload *models.SubnetList
}

func (o *GetNmsTrafficSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /nms/traffic/subnets][%d] getNmsTrafficSubnetsOK  %+v", 200, o.Payload)
}
func (o *GetNmsTrafficSubnetsOK) GetPayload() *models.SubnetList {
	return o.Payload
}

func (o *GetNmsTrafficSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubnetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsTrafficSubnetsUnauthorized creates a GetNmsTrafficSubnetsUnauthorized with default headers values
func NewGetNmsTrafficSubnetsUnauthorized() *GetNmsTrafficSubnetsUnauthorized {
	return &GetNmsTrafficSubnetsUnauthorized{}
}

/* GetNmsTrafficSubnetsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNmsTrafficSubnetsUnauthorized struct {
	Payload *models.Error
}

func (o *GetNmsTrafficSubnetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /nms/traffic/subnets][%d] getNmsTrafficSubnetsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetNmsTrafficSubnetsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsTrafficSubnetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsTrafficSubnetsForbidden creates a GetNmsTrafficSubnetsForbidden with default headers values
func NewGetNmsTrafficSubnetsForbidden() *GetNmsTrafficSubnetsForbidden {
	return &GetNmsTrafficSubnetsForbidden{}
}

/* GetNmsTrafficSubnetsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNmsTrafficSubnetsForbidden struct {
	Payload *models.Error
}

func (o *GetNmsTrafficSubnetsForbidden) Error() string {
	return fmt.Sprintf("[GET /nms/traffic/subnets][%d] getNmsTrafficSubnetsForbidden  %+v", 403, o.Payload)
}
func (o *GetNmsTrafficSubnetsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsTrafficSubnetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsTrafficSubnetsInternalServerError creates a GetNmsTrafficSubnetsInternalServerError with default headers values
func NewGetNmsTrafficSubnetsInternalServerError() *GetNmsTrafficSubnetsInternalServerError {
	return &GetNmsTrafficSubnetsInternalServerError{}
}

/* GetNmsTrafficSubnetsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNmsTrafficSubnetsInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsTrafficSubnetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/traffic/subnets][%d] getNmsTrafficSubnetsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNmsTrafficSubnetsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsTrafficSubnetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}