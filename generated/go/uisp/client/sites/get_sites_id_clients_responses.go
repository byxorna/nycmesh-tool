// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetSitesIDClientsReader is a Reader for the GetSitesIDClients structure.
type GetSitesIDClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesIDClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesIDClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSitesIDClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSitesIDClientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSitesIDClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSitesIDClientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesIDClientsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSitesIDClientsOK creates a GetSitesIDClientsOK with default headers values
func NewGetSitesIDClientsOK() *GetSitesIDClientsOK {
	return &GetSitesIDClientsOK{}
}

/* GetSitesIDClientsOK describes a response with status code 200, with default header values.

Successful
*/
type GetSitesIDClientsOK struct {
	Payload models.SitesList
}

func (o *GetSitesIDClientsOK) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/clients][%d] getSitesIdClientsOK  %+v", 200, o.Payload)
}
func (o *GetSitesIDClientsOK) GetPayload() models.SitesList {
	return o.Payload
}

func (o *GetSitesIDClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDClientsBadRequest creates a GetSitesIDClientsBadRequest with default headers values
func NewGetSitesIDClientsBadRequest() *GetSitesIDClientsBadRequest {
	return &GetSitesIDClientsBadRequest{}
}

/* GetSitesIDClientsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSitesIDClientsBadRequest struct {
	Payload *models.Error
}

func (o *GetSitesIDClientsBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/clients][%d] getSitesIdClientsBadRequest  %+v", 400, o.Payload)
}
func (o *GetSitesIDClientsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDClientsUnauthorized creates a GetSitesIDClientsUnauthorized with default headers values
func NewGetSitesIDClientsUnauthorized() *GetSitesIDClientsUnauthorized {
	return &GetSitesIDClientsUnauthorized{}
}

/* GetSitesIDClientsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSitesIDClientsUnauthorized struct {
	Payload *models.Error
}

func (o *GetSitesIDClientsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/clients][%d] getSitesIdClientsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSitesIDClientsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDClientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDClientsForbidden creates a GetSitesIDClientsForbidden with default headers values
func NewGetSitesIDClientsForbidden() *GetSitesIDClientsForbidden {
	return &GetSitesIDClientsForbidden{}
}

/* GetSitesIDClientsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSitesIDClientsForbidden struct {
	Payload *models.Error
}

func (o *GetSitesIDClientsForbidden) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/clients][%d] getSitesIdClientsForbidden  %+v", 403, o.Payload)
}
func (o *GetSitesIDClientsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDClientsNotFound creates a GetSitesIDClientsNotFound with default headers values
func NewGetSitesIDClientsNotFound() *GetSitesIDClientsNotFound {
	return &GetSitesIDClientsNotFound{}
}

/* GetSitesIDClientsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSitesIDClientsNotFound struct {
	Payload *models.Error
}

func (o *GetSitesIDClientsNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/clients][%d] getSitesIdClientsNotFound  %+v", 404, o.Payload)
}
func (o *GetSitesIDClientsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDClientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesIDClientsInternalServerError creates a GetSitesIDClientsInternalServerError with default headers values
func NewGetSitesIDClientsInternalServerError() *GetSitesIDClientsInternalServerError {
	return &GetSitesIDClientsInternalServerError{}
}

/* GetSitesIDClientsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSitesIDClientsInternalServerError struct {
	Payload *models.Error
}

func (o *GetSitesIDClientsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sites/{id}/clients][%d] getSitesIdClientsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSitesIDClientsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesIDClientsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
