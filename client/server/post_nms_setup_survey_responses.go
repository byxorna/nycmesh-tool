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

// PostNmsSetupSurveyReader is a Reader for the PostNmsSetupSurvey structure.
type PostNmsSetupSurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNmsSetupSurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNmsSetupSurveyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostNmsSetupSurveyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNmsSetupSurveyOK creates a PostNmsSetupSurveyOK with default headers values
func NewPostNmsSetupSurveyOK() *PostNmsSetupSurveyOK {
	return &PostNmsSetupSurveyOK{}
}

/* PostNmsSetupSurveyOK describes a response with status code 200, with default header values.

Successful
*/
type PostNmsSetupSurveyOK struct {
	Payload *models.Status
}

func (o *PostNmsSetupSurveyOK) Error() string {
	return fmt.Sprintf("[POST /nms/setup/survey][%d] postNmsSetupSurveyOK  %+v", 200, o.Payload)
}
func (o *PostNmsSetupSurveyOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *PostNmsSetupSurveyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNmsSetupSurveyInternalServerError creates a PostNmsSetupSurveyInternalServerError with default headers values
func NewPostNmsSetupSurveyInternalServerError() *PostNmsSetupSurveyInternalServerError {
	return &PostNmsSetupSurveyInternalServerError{}
}

/* PostNmsSetupSurveyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostNmsSetupSurveyInternalServerError struct {
	Payload *models.Error
}

func (o *PostNmsSetupSurveyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nms/setup/survey][%d] postNmsSetupSurveyInternalServerError  %+v", 500, o.Payload)
}
func (o *PostNmsSetupSurveyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostNmsSetupSurveyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
