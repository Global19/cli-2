// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetOperatingSystemKernelsReader is a Reader for the GetOperatingSystemKernels structure.
type GetOperatingSystemKernelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperatingSystemKernelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOperatingSystemKernelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOperatingSystemKernelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOperatingSystemKernelsOK creates a GetOperatingSystemKernelsOK with default headers values
func NewGetOperatingSystemKernelsOK() *GetOperatingSystemKernelsOK {
	return &GetOperatingSystemKernelsOK{}
}

/*GetOperatingSystemKernelsOK handles this case with default header values.

A paginated list of kernels
*/
type GetOperatingSystemKernelsOK struct {
	Payload *inventory_models.V1KernelPagedList
}

func (o *GetOperatingSystemKernelsOK) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/kernels][%d] getOperatingSystemKernelsOK  %+v", 200, o.Payload)
}

func (o *GetOperatingSystemKernelsOK) GetPayload() *inventory_models.V1KernelPagedList {
	return o.Payload
}

func (o *GetOperatingSystemKernelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1KernelPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperatingSystemKernelsDefault creates a GetOperatingSystemKernelsDefault with default headers values
func NewGetOperatingSystemKernelsDefault(code int) *GetOperatingSystemKernelsDefault {
	return &GetOperatingSystemKernelsDefault{
		_statusCode: code,
	}
}

/*GetOperatingSystemKernelsDefault handles this case with default header values.

generic error response
*/
type GetOperatingSystemKernelsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get operating system kernels default response
func (o *GetOperatingSystemKernelsDefault) Code() int {
	return o._statusCode
}

func (o *GetOperatingSystemKernelsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/kernels][%d] getOperatingSystemKernels default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperatingSystemKernelsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetOperatingSystemKernelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
