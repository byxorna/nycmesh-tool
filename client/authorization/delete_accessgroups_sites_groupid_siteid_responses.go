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

// DeleteAccessgroupsSitesGroupidSiteidReader is a Reader for the DeleteAccessgroupsSitesGroupidSiteid structure.
type DeleteAccessgroupsSitesGroupidSiteidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccessgroupsSitesGroupidSiteidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAccessgroupsSitesGroupidSiteidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAccessgroupsSitesGroupidSiteidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAccessgroupsSitesGroupidSiteidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAccessgroupsSitesGroupidSiteidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAccessgroupsSitesGroupidSiteidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAccessgroupsSitesGroupidSiteidConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAccessgroupsSitesGroupidSiteidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAccessgroupsSitesGroupidSiteidOK creates a DeleteAccessgroupsSitesGroupidSiteidOK with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidOK() *DeleteAccessgroupsSitesGroupidSiteidOK {
	return &DeleteAccessgroupsSitesGroupidSiteidOK{}
}

/* DeleteAccessgroupsSitesGroupidSiteidOK describes a response with status code 200, with default header values.

Successful
*/
type DeleteAccessgroupsSitesGroupidSiteidOK struct {
	Payload *models.Status
}

func (o *DeleteAccessgroupsSitesGroupidSiteidOK) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidOK  %+v", 200, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessgroupsSitesGroupidSiteidBadRequest creates a DeleteAccessgroupsSitesGroupidSiteidBadRequest with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidBadRequest() *DeleteAccessgroupsSitesGroupidSiteidBadRequest {
	return &DeleteAccessgroupsSitesGroupidSiteidBadRequest{}
}

/* DeleteAccessgroupsSitesGroupidSiteidBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteAccessgroupsSitesGroupidSiteidBadRequest struct {
	Payload *models.Error
}

func (o *DeleteAccessgroupsSitesGroupidSiteidBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessgroupsSitesGroupidSiteidUnauthorized creates a DeleteAccessgroupsSitesGroupidSiteidUnauthorized with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidUnauthorized() *DeleteAccessgroupsSitesGroupidSiteidUnauthorized {
	return &DeleteAccessgroupsSitesGroupidSiteidUnauthorized{}
}

/* DeleteAccessgroupsSitesGroupidSiteidUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteAccessgroupsSitesGroupidSiteidUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteAccessgroupsSitesGroupidSiteidUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessgroupsSitesGroupidSiteidForbidden creates a DeleteAccessgroupsSitesGroupidSiteidForbidden with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidForbidden() *DeleteAccessgroupsSitesGroupidSiteidForbidden {
	return &DeleteAccessgroupsSitesGroupidSiteidForbidden{}
}

/* DeleteAccessgroupsSitesGroupidSiteidForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAccessgroupsSitesGroupidSiteidForbidden struct {
	Payload *models.Error
}

func (o *DeleteAccessgroupsSitesGroupidSiteidForbidden) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessgroupsSitesGroupidSiteidNotFound creates a DeleteAccessgroupsSitesGroupidSiteidNotFound with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidNotFound() *DeleteAccessgroupsSitesGroupidSiteidNotFound {
	return &DeleteAccessgroupsSitesGroupidSiteidNotFound{}
}

/* DeleteAccessgroupsSitesGroupidSiteidNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAccessgroupsSitesGroupidSiteidNotFound struct {
	Payload *models.Error
}

func (o *DeleteAccessgroupsSitesGroupidSiteidNotFound) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessgroupsSitesGroupidSiteidConflict creates a DeleteAccessgroupsSitesGroupidSiteidConflict with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidConflict() *DeleteAccessgroupsSitesGroupidSiteidConflict {
	return &DeleteAccessgroupsSitesGroupidSiteidConflict{}
}

/* DeleteAccessgroupsSitesGroupidSiteidConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteAccessgroupsSitesGroupidSiteidConflict struct {
	Payload *models.Error
}

func (o *DeleteAccessgroupsSitesGroupidSiteidConflict) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidConflict  %+v", 409, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccessgroupsSitesGroupidSiteidInternalServerError creates a DeleteAccessgroupsSitesGroupidSiteidInternalServerError with default headers values
func NewDeleteAccessgroupsSitesGroupidSiteidInternalServerError() *DeleteAccessgroupsSitesGroupidSiteidInternalServerError {
	return &DeleteAccessgroupsSitesGroupidSiteidInternalServerError{}
}

/* DeleteAccessgroupsSitesGroupidSiteidInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAccessgroupsSitesGroupidSiteidInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteAccessgroupsSitesGroupidSiteidInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /access-groups/sites/{groupId}/{siteId}][%d] deleteAccessgroupsSitesGroupidSiteidInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAccessgroupsSitesGroupidSiteidInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAccessgroupsSitesGroupidSiteidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
