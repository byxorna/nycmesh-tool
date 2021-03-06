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

// GetSitesSiteidTrafficIntervalReader is a Reader for the GetSitesSiteidTrafficInterval structure.
type GetSitesSiteidTrafficIntervalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesSiteidTrafficIntervalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesSiteidTrafficIntervalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSitesSiteidTrafficIntervalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSitesSiteidTrafficIntervalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSitesSiteidTrafficIntervalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesSiteidTrafficIntervalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSitesSiteidTrafficIntervalOK creates a GetSitesSiteidTrafficIntervalOK with default headers values
func NewGetSitesSiteidTrafficIntervalOK() *GetSitesSiteidTrafficIntervalOK {
	return &GetSitesSiteidTrafficIntervalOK{}
}

/* GetSitesSiteidTrafficIntervalOK describes a response with status code 200, with default header values.

Successful
*/
type GetSitesSiteidTrafficIntervalOK struct {
	Payload models.TrafficList
}

func (o *GetSitesSiteidTrafficIntervalOK) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/interval][%d] getSitesSiteidTrafficIntervalOK  %+v", 200, o.Payload)
}
func (o *GetSitesSiteidTrafficIntervalOK) GetPayload() models.TrafficList {
	return o.Payload
}

func (o *GetSitesSiteidTrafficIntervalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficIntervalBadRequest creates a GetSitesSiteidTrafficIntervalBadRequest with default headers values
func NewGetSitesSiteidTrafficIntervalBadRequest() *GetSitesSiteidTrafficIntervalBadRequest {
	return &GetSitesSiteidTrafficIntervalBadRequest{}
}

/* GetSitesSiteidTrafficIntervalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSitesSiteidTrafficIntervalBadRequest struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficIntervalBadRequest) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/interval][%d] getSitesSiteidTrafficIntervalBadRequest  %+v", 400, o.Payload)
}
func (o *GetSitesSiteidTrafficIntervalBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficIntervalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficIntervalUnauthorized creates a GetSitesSiteidTrafficIntervalUnauthorized with default headers values
func NewGetSitesSiteidTrafficIntervalUnauthorized() *GetSitesSiteidTrafficIntervalUnauthorized {
	return &GetSitesSiteidTrafficIntervalUnauthorized{}
}

/* GetSitesSiteidTrafficIntervalUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSitesSiteidTrafficIntervalUnauthorized struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficIntervalUnauthorized) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/interval][%d] getSitesSiteidTrafficIntervalUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSitesSiteidTrafficIntervalUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficIntervalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficIntervalForbidden creates a GetSitesSiteidTrafficIntervalForbidden with default headers values
func NewGetSitesSiteidTrafficIntervalForbidden() *GetSitesSiteidTrafficIntervalForbidden {
	return &GetSitesSiteidTrafficIntervalForbidden{}
}

/* GetSitesSiteidTrafficIntervalForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSitesSiteidTrafficIntervalForbidden struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficIntervalForbidden) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/interval][%d] getSitesSiteidTrafficIntervalForbidden  %+v", 403, o.Payload)
}
func (o *GetSitesSiteidTrafficIntervalForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficIntervalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesSiteidTrafficIntervalInternalServerError creates a GetSitesSiteidTrafficIntervalInternalServerError with default headers values
func NewGetSitesSiteidTrafficIntervalInternalServerError() *GetSitesSiteidTrafficIntervalInternalServerError {
	return &GetSitesSiteidTrafficIntervalInternalServerError{}
}

/* GetSitesSiteidTrafficIntervalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSitesSiteidTrafficIntervalInternalServerError struct {
	Payload *models.Error
}

func (o *GetSitesSiteidTrafficIntervalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sites/{siteId}/traffic/interval][%d] getSitesSiteidTrafficIntervalInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSitesSiteidTrafficIntervalInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSitesSiteidTrafficIntervalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
