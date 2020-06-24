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

// AddCPUArchitectureCPUExtensionReader is a Reader for the AddCPUArchitectureCPUExtension structure.
type AddCPUArchitectureCPUExtensionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCPUArchitectureCPUExtensionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddCPUArchitectureCPUExtensionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddCPUArchitectureCPUExtensionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddCPUArchitectureCPUExtensionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddCPUArchitectureCPUExtensionOK creates a AddCPUArchitectureCPUExtensionOK with default headers values
func NewAddCPUArchitectureCPUExtensionOK() *AddCPUArchitectureCPUExtensionOK {
	return &AddCPUArchitectureCPUExtensionOK{}
}

/*AddCPUArchitectureCPUExtensionOK handles this case with default header values.

The CPU extension added to the kernel
*/
type AddCPUArchitectureCPUExtensionOK struct {
	Payload *inventory_models.V1CPUExtension
}

func (o *AddCPUArchitectureCPUExtensionOK) Error() string {
	return fmt.Sprintf("[POST /v1/cpu-architectures/{cpu_architecture_id}/extensions][%d] addCpuArchitectureCpuExtensionOK  %+v", 200, o.Payload)
}

func (o *AddCPUArchitectureCPUExtensionOK) GetPayload() *inventory_models.V1CPUExtension {
	return o.Payload
}

func (o *AddCPUArchitectureCPUExtensionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1CPUExtension)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCPUArchitectureCPUExtensionBadRequest creates a AddCPUArchitectureCPUExtensionBadRequest with default headers values
func NewAddCPUArchitectureCPUExtensionBadRequest() *AddCPUArchitectureCPUExtensionBadRequest {
	return &AddCPUArchitectureCPUExtensionBadRequest{}
}

/*AddCPUArchitectureCPUExtensionBadRequest handles this case with default header values.

If the CPU extension ID doesn't exist
*/
type AddCPUArchitectureCPUExtensionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddCPUArchitectureCPUExtensionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/cpu-architectures/{cpu_architecture_id}/extensions][%d] addCpuArchitectureCpuExtensionBadRequest  %+v", 400, o.Payload)
}

func (o *AddCPUArchitectureCPUExtensionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddCPUArchitectureCPUExtensionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCPUArchitectureCPUExtensionDefault creates a AddCPUArchitectureCPUExtensionDefault with default headers values
func NewAddCPUArchitectureCPUExtensionDefault(code int) *AddCPUArchitectureCPUExtensionDefault {
	return &AddCPUArchitectureCPUExtensionDefault{
		_statusCode: code,
	}
}

/*AddCPUArchitectureCPUExtensionDefault handles this case with default header values.

generic error response
*/
type AddCPUArchitectureCPUExtensionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add Cpu architecture Cpu extension default response
func (o *AddCPUArchitectureCPUExtensionDefault) Code() int {
	return o._statusCode
}

func (o *AddCPUArchitectureCPUExtensionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cpu-architectures/{cpu_architecture_id}/extensions][%d] addCpuArchitectureCpuExtension default  %+v", o._statusCode, o.Payload)
}

func (o *AddCPUArchitectureCPUExtensionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddCPUArchitectureCPUExtensionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
