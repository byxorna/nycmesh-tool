// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDevicesEpowersIDConfigurationPowerReader is a Reader for the GetDevicesEpowersIDConfigurationPower structure.
type GetDevicesEpowersIDConfigurationPowerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDevicesEpowersIDConfigurationPowerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetDevicesEpowersIDConfigurationPowerDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetDevicesEpowersIDConfigurationPowerDefault creates a GetDevicesEpowersIDConfigurationPowerDefault with default headers values
func NewGetDevicesEpowersIDConfigurationPowerDefault(code int) *GetDevicesEpowersIDConfigurationPowerDefault {
	return &GetDevicesEpowersIDConfigurationPowerDefault{
		_statusCode: code,
	}
}

/* GetDevicesEpowersIDConfigurationPowerDefault describes a response with status code -1, with default header values.

Successful
*/
type GetDevicesEpowersIDConfigurationPowerDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the get devices epowers Id configuration power default response
func (o *GetDevicesEpowersIDConfigurationPowerDefault) Code() int {
	return o._statusCode
}

func (o *GetDevicesEpowersIDConfigurationPowerDefault) Error() string {
	return fmt.Sprintf("[GET /devices/epowers/{id}/configuration/power][%d] getDevicesEpowersIdConfigurationPower default  %+v", o._statusCode, o.Payload)
}
func (o *GetDevicesEpowersIDConfigurationPowerDefault) GetPayload() string {
	return o.Payload
}

func (o *GetDevicesEpowersIDConfigurationPowerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}