// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetAccessgroupsSitesGroupidReader is a Reader for the GetAccessgroupsSitesGroupid structure.
type GetAccessgroupsSitesGroupidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessgroupsSitesGroupidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessgroupsSitesGroupidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccessgroupsSitesGroupidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccessgroupsSitesGroupidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessgroupsSitesGroupidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccessgroupsSitesGroupidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccessgroupsSitesGroupidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccessgroupsSitesGroupidOK creates a GetAccessgroupsSitesGroupidOK with default headers values
func NewGetAccessgroupsSitesGroupidOK() *GetAccessgroupsSitesGroupidOK {
	return &GetAccessgroupsSitesGroupidOK{}
}

/* GetAccessgroupsSitesGroupidOK describes a response with status code 200, with default header values.

Successful
*/
type GetAccessgroupsSitesGroupidOK struct {
	Payload *models.SiteAccessGroup
}

func (o *GetAccessgroupsSitesGroupidOK) Error() string {
	return fmt.Sprintf("[GET /access-groups/sites/{groupId}][%d] getAccessgroupsSitesGroupidOK  %+v", 200, o.Payload)
}
func (o *GetAccessgroupsSitesGroupidOK) GetPayload() *models.SiteAccessGroup {
	return o.Payload
}

func (o *GetAccessgroupsSitesGroupidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteAccessGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessgroupsSitesGroupidBadRequest creates a GetAccessgroupsSitesGroupidBadRequest with default headers values
func NewGetAccessgroupsSitesGroupidBadRequest() *GetAccessgroupsSitesGroupidBadRequest {
	return &GetAccessgroupsSitesGroupidBadRequest{}
}

/* GetAccessgroupsSitesGroupidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAccessgroupsSitesGroupidBadRequest struct {
	Payload *models.Error
}

func (o *GetAccessgroupsSitesGroupidBadRequest) Error() string {
	return fmt.Sprintf("[GET /access-groups/sites/{groupId}][%d] getAccessgroupsSitesGroupidBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccessgroupsSitesGroupidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccessgroupsSitesGroupidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessgroupsSitesGroupidUnauthorized creates a GetAccessgroupsSitesGroupidUnauthorized with default headers values
func NewGetAccessgroupsSitesGroupidUnauthorized() *GetAccessgroupsSitesGroupidUnauthorized {
	return &GetAccessgroupsSitesGroupidUnauthorized{}
}

/* GetAccessgroupsSitesGroupidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccessgroupsSitesGroupidUnauthorized struct {
	Payload *models.Error
}

func (o *GetAccessgroupsSitesGroupidUnauthorized) Error() string {
	return fmt.Sprintf("[GET /access-groups/sites/{groupId}][%d] getAccessgroupsSitesGroupidUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAccessgroupsSitesGroupidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccessgroupsSitesGroupidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessgroupsSitesGroupidForbidden creates a GetAccessgroupsSitesGroupidForbidden with default headers values
func NewGetAccessgroupsSitesGroupidForbidden() *GetAccessgroupsSitesGroupidForbidden {
	return &GetAccessgroupsSitesGroupidForbidden{}
}

/* GetAccessgroupsSitesGroupidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccessgroupsSitesGroupidForbidden struct {
	Payload *models.Error
}

func (o *GetAccessgroupsSitesGroupidForbidden) Error() string {
	return fmt.Sprintf("[GET /access-groups/sites/{groupId}][%d] getAccessgroupsSitesGroupidForbidden  %+v", 403, o.Payload)
}
func (o *GetAccessgroupsSitesGroupidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccessgroupsSitesGroupidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessgroupsSitesGroupidNotFound creates a GetAccessgroupsSitesGroupidNotFound with default headers values
func NewGetAccessgroupsSitesGroupidNotFound() *GetAccessgroupsSitesGroupidNotFound {
	return &GetAccessgroupsSitesGroupidNotFound{}
}

/* GetAccessgroupsSitesGroupidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAccessgroupsSitesGroupidNotFound struct {
	Payload *models.Error
}

func (o *GetAccessgroupsSitesGroupidNotFound) Error() string {
	return fmt.Sprintf("[GET /access-groups/sites/{groupId}][%d] getAccessgroupsSitesGroupidNotFound  %+v", 404, o.Payload)
}
func (o *GetAccessgroupsSitesGroupidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccessgroupsSitesGroupidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessgroupsSitesGroupidInternalServerError creates a GetAccessgroupsSitesGroupidInternalServerError with default headers values
func NewGetAccessgroupsSitesGroupidInternalServerError() *GetAccessgroupsSitesGroupidInternalServerError {
	return &GetAccessgroupsSitesGroupidInternalServerError{}
}

/* GetAccessgroupsSitesGroupidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccessgroupsSitesGroupidInternalServerError struct {
	Payload *models.Error
}

func (o *GetAccessgroupsSitesGroupidInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access-groups/sites/{groupId}][%d] getAccessgroupsSitesGroupidInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccessgroupsSitesGroupidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccessgroupsSitesGroupidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
