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

// GetSitesSiteidTrafficSummaryReader is a Reader for the GetSitesSiteidTrafficSummary structure.
type GetSitesSiteidTrafficSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesSiteidTrafficSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesSiteidTrafficSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSitesSiteidTrafficSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSitesSiteidTrafficSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSitesSiteidTrafficSummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesSiteidTrafficSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSitesSiteidTrafficSummaryOK creates a GetSitesSiteidTrafficSummaryOK with default headers values
func NewGetSitesSiteidTrafficSummaryOK() *GetSitesSiteidTrafficSummaryOK {
	return &GetSitesSiteidTrafficSummaryOK{}
}

/* GetSitesSiteidTrafficSummaryOK describes a response with status code 200, with default header values.

Successful
*/
type GetSitesSiteidTrafficSummaryOK struct {
	Payload *models.TrafficSummary
}

func (o *GetSitesSiteidTrafficSummaryOK) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/summary][%d] getSitesSiteidTrafficSummaryOK  %+v", 200, o.Payload)
}
func (o *GetSitesSiteidTrafficSummaryOK) GetPayload() *models.TrafficSummary {
	return o.Payload
}

func (o *GetSitesSiteidTrafficSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrafficSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficSummaryBadRequest creates a GetSitesSiteidTrafficSummaryBadRequest with default headers values
func NewGetSitesSiteidTrafficSummaryBadRequest() *GetSitesSiteidTrafficSummaryBadRequest {
	return &GetSitesSiteidTrafficSummaryBadRequest{}
}

/* GetSitesSiteidTrafficSummaryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSitesSiteidTrafficSummaryBadRequest struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficSummaryBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/summary][%d] getSitesSiteidTrafficSummaryBadRequest  %+v", 400, o.Payload)
}
func (o *GetSitesSiteidTrafficSummaryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficSummaryUnauthorized creates a GetSitesSiteidTrafficSummaryUnauthorized with default headers values
func NewGetSitesSiteidTrafficSummaryUnauthorized() *GetSitesSiteidTrafficSummaryUnauthorized {
	return &GetSitesSiteidTrafficSummaryUnauthorized{}
}

/* GetSitesSiteidTrafficSummaryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSitesSiteidTrafficSummaryUnauthorized struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/summary][%d] getSitesSiteidTrafficSummaryUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSitesSiteidTrafficSummaryUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficSummaryForbidden creates a GetSitesSiteidTrafficSummaryForbidden with default headers values
func NewGetSitesSiteidTrafficSummaryForbidden() *GetSitesSiteidTrafficSummaryForbidden {
	return &GetSitesSiteidTrafficSummaryForbidden{}
}

/* GetSitesSiteidTrafficSummaryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSitesSiteidTrafficSummaryForbidden struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficSummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/summary][%d] getSitesSiteidTrafficSummaryForbidden  %+v", 403, o.Payload)
}
func (o *GetSitesSiteidTrafficSummaryForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficSummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficSummaryInternalServerError creates a GetSitesSiteidTrafficSummaryInternalServerError with default headers values
func NewGetSitesSiteidTrafficSummaryInternalServerError() *GetSitesSiteidTrafficSummaryInternalServerError {
	return &GetSitesSiteidTrafficSummaryInternalServerError{}
}

/* GetSitesSiteidTrafficSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSitesSiteidTrafficSummaryInternalServerError struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficSummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/summary][%d] getSitesSiteidTrafficSummaryInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSitesSiteidTrafficSummaryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
