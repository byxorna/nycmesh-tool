// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/models"
)

// GetNmsDownloadsTokenReader is a Reader for the GetNmsDownloadsToken structure.
type GetNmsDownloadsTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNmsDownloadsTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewGetNmsDownloadsTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNmsDownloadsTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNmsDownloadsTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNmsDownloadsTokenBadRequest creates a GetNmsDownloadsTokenBadRequest with default headers values
func NewGetNmsDownloadsTokenBadRequest() *GetNmsDownloadsTokenBadRequest {
	return &GetNmsDownloadsTokenBadRequest{}
}

/* GetNmsDownloadsTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNmsDownloadsTokenBadRequest struct {
	Payload *models.Error
}

func (o *GetNmsDownloadsTokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /nms/downloads/{token}][%d] getNmsDownloadsTokenBadRequest  %+v", 400, o.Payload)
}
func (o *GetNmsDownloadsTokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsDownloadsTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsDownloadsTokenNotFound creates a GetNmsDownloadsTokenNotFound with default headers values
func NewGetNmsDownloadsTokenNotFound() *GetNmsDownloadsTokenNotFound {
	return &GetNmsDownloadsTokenNotFound{}
}

/* GetNmsDownloadsTokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNmsDownloadsTokenNotFound struct {
	Payload *models.Error
}

func (o *GetNmsDownloadsTokenNotFound) Error() string {
	return fmt.Sprintf("[GET /nms/downloads/{token}][%d] getNmsDownloadsTokenNotFound  %+v", 404, o.Payload)
}
func (o *GetNmsDownloadsTokenNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsDownloadsTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNmsDownloadsTokenInternalServerError creates a GetNmsDownloadsTokenInternalServerError with default headers values
func NewGetNmsDownloadsTokenInternalServerError() *GetNmsDownloadsTokenInternalServerError {
	return &GetNmsDownloadsTokenInternalServerError{}
}

/* GetNmsDownloadsTokenInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNmsDownloadsTokenInternalServerError struct {
	Payload *models.Error
}

func (o *GetNmsDownloadsTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nms/downloads/{token}][%d] getNmsDownloadsTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNmsDownloadsTokenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNmsDownloadsTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
