// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/byxorna/nycmesh-tool/models"
)

// PostNmsTrafficBlacklistReader is a Reader for the PostNmsTrafficBlacklist structure.
type PostNmsTrafficBlacklistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsTrafficBlacklistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsTrafficBlacklistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostNmsTrafficBlacklistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNmsTrafficBlacklistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNmsTrafficBlacklistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNmsTrafficBlacklistOK creates a PostNmsTrafficBlacklistOK with default headers values
func NewPostNmsTrafficBlacklistOK() *PostNmsTrafficBlacklistOK {
	return &PostNmsTrafficBlacklistOK{}
}

/* PostNmsTrafficBlacklistOK describes a response with status code 200, with default header values.

Successful
*/
type PostNmsTrafficBlacklistOK struct {
	Payload models.BlacklistSchema
}

func (o *PostNmsTrafficBlacklistOK) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/blacklist][%d] postNmsTrafficBlacklistOK  %+v", 200, o.Payload)
}
func (o *PostNmsTrafficBlacklistOK) GetPayload() models.BlacklistSchema {
	return o.Payload
}

func (o *PostNmsTrafficBlacklistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsTrafficBlacklistUnauthorized creates a PostNmsTrafficBlacklistUnauthorized with default headers values
func NewPostNmsTrafficBlacklistUnauthorized() *PostNmsTrafficBlacklistUnauthorized {
	return &PostNmsTrafficBlacklistUnauthorized{}
}

/* PostNmsTrafficBlacklistUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostNmsTrafficBlacklistUnauthorized struct {
	Payload *models.Error
}

func (o *PostNmsTrafficBlacklistUnauthorized) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/blacklist][%d] postNmsTrafficBlacklistUnauthorized  %+v", 401, o.Payload)
}
func (o *PostNmsTrafficBlacklistUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsTrafficBlacklistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsTrafficBlacklistForbidden creates a PostNmsTrafficBlacklistForbidden with default headers values
func NewPostNmsTrafficBlacklistForbidden() *PostNmsTrafficBlacklistForbidden {
	return &PostNmsTrafficBlacklistForbidden{}
}

/* PostNmsTrafficBlacklistForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostNmsTrafficBlacklistForbidden struct {
	Payload *models.Error
}

func (o *PostNmsTrafficBlacklistForbidden) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/blacklist][%d] postNmsTrafficBlacklistForbidden  %+v", 403, o.Payload)
}
func (o *PostNmsTrafficBlacklistForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsTrafficBlacklistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsTrafficBlacklistInternalServerError creates a PostNmsTrafficBlacklistInternalServerError with default headers values
func NewPostNmsTrafficBlacklistInternalServerError() *PostNmsTrafficBlacklistInternalServerError {
	return &PostNmsTrafficBlacklistInternalServerError{}
}

/* PostNmsTrafficBlacklistInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostNmsTrafficBlacklistInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsTrafficBlacklistInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/traffic/blacklist][%d] postNmsTrafficBlacklistInternalServerError  %+v", 500, o.Payload)
}
func (o *PostNmsTrafficBlacklistInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsTrafficBlacklistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
