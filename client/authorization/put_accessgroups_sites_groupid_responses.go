// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PutAccessgroupsSitesGroupidReader is a Reader for the PutAccessgroupsSitesGroupid structure.
type PutAccessgroupsSitesGroupidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAccessgroupsSitesGroupidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAccessgroupsSitesGroupidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAccessgroupsSitesGroupidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAccessgroupsSitesGroupidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAccessgroupsSitesGroupidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAccessgroupsSitesGroupidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutAccessgroupsSitesGroupidConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAccessgroupsSitesGroupidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAccessgroupsSitesGroupidOK creates a PutAccessgroupsSitesGroupidOK with default headers values
func NewPutAccessgroupsSitesGroupidOK() *PutAccessgroupsSitesGroupidOK {
	return &PutAccessgroupsSitesGroupidOK{}
}

/* PutAccessgroupsSitesGroupidOK describes a response with status code 200, with default header values.

Successful
*/
type PutAccessgroupsSitesGroupidOK struct {
	Payload *models.SiteAccessGroup
}

func (o *PutAccessgroupsSitesGroupidOK) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidOK  %+v", 200, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidOK) GetPayload() *models.SiteAccessGroup {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteAccessGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAccessgroupsSitesGroupidBadRequest creates a PutAccessgroupsSitesGroupidBadRequest with default headers values
func NewPutAccessgroupsSitesGroupidBadRequest() *PutAccessgroupsSitesGroupidBadRequest {
	return &PutAccessgroupsSitesGroupidBadRequest{}
}

/* PutAccessgroupsSitesGroupidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutAccessgroupsSitesGroupidBadRequest struct {
	Payload *models.Error
}

func (o *PutAccessgroupsSitesGroupidBadRequest) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidBadRequest  %+v", 400, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAccessgroupsSitesGroupidUnauthorized creates a PutAccessgroupsSitesGroupidUnauthorized with default headers values
func NewPutAccessgroupsSitesGroupidUnauthorized() *PutAccessgroupsSitesGroupidUnauthorized {
	return &PutAccessgroupsSitesGroupidUnauthorized{}
}

/* PutAccessgroupsSitesGroupidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutAccessgroupsSitesGroupidUnauthorized struct {
	Payload *models.Error
}

func (o *PutAccessgroupsSitesGroupidUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidUnauthorized  %+v", 401, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAccessgroupsSitesGroupidForbidden creates a PutAccessgroupsSitesGroupidForbidden with default headers values
func NewPutAccessgroupsSitesGroupidForbidden() *PutAccessgroupsSitesGroupidForbidden {
	return &PutAccessgroupsSitesGroupidForbidden{}
}

/* PutAccessgroupsSitesGroupidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutAccessgroupsSitesGroupidForbidden struct {
	Payload *models.Error
}

func (o *PutAccessgroupsSitesGroupidForbidden) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidForbidden  %+v", 403, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAccessgroupsSitesGroupidNotFound creates a PutAccessgroupsSitesGroupidNotFound with default headers values
func NewPutAccessgroupsSitesGroupidNotFound() *PutAccessgroupsSitesGroupidNotFound {
	return &PutAccessgroupsSitesGroupidNotFound{}
}

/* PutAccessgroupsSitesGroupidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutAccessgroupsSitesGroupidNotFound struct {
	Payload *models.Error
}

func (o *PutAccessgroupsSitesGroupidNotFound) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidNotFound  %+v", 404, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAccessgroupsSitesGroupidConflict creates a PutAccessgroupsSitesGroupidConflict with default headers values
func NewPutAccessgroupsSitesGroupidConflict() *PutAccessgroupsSitesGroupidConflict {
	return &PutAccessgroupsSitesGroupidConflict{}
}

/* PutAccessgroupsSitesGroupidConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutAccessgroupsSitesGroupidConflict struct {
	Payload *models.Error
}

func (o *PutAccessgroupsSitesGroupidConflict) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidConflict  %+v", 409, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAccessgroupsSitesGroupidInternalServerError creates a PutAccessgroupsSitesGroupidInternalServerError with default headers values
func NewPutAccessgroupsSitesGroupidInternalServerError() *PutAccessgroupsSitesGroupidInternalServerError {
	return &PutAccessgroupsSitesGroupidInternalServerError{}
}

/* PutAccessgroupsSitesGroupidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutAccessgroupsSitesGroupidInternalServerError struct {
	Payload *models.Error
}

func (o *PutAccessgroupsSitesGroupidInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /access-groups/sites/{groupId}][%d] putAccessgroupsSitesGroupidInternalServerError  %+v", 500, o.Payload)
}
func (o *PutAccessgroupsSitesGroupidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutAccessgroupsSitesGroupidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
