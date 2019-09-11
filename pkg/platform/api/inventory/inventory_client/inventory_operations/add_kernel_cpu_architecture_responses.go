// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddKernelCPUArchitectureReader is a Reader for the AddKernelCPUArchitecture structure.
type AddKernelCPUArchitectureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddKernelCPUArchitectureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddKernelCPUArchitectureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddKernelCPUArchitectureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddKernelCPUArchitectureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddKernelCPUArchitectureOK creates a AddKernelCPUArchitectureOK with default headers values
func NewAddKernelCPUArchitectureOK() *AddKernelCPUArchitectureOK {
	return &AddKernelCPUArchitectureOK{}
}

/*AddKernelCPUArchitectureOK handles this case with default header values.

The CPU architecture added to the kernel
*/
type AddKernelCPUArchitectureOK struct {
	Payload *inventory_models.V1CPUArchitecture
}

func (o *AddKernelCPUArchitectureOK) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/cpu-architectures][%d] addKernelCpuArchitectureOK  %+v", 200, o.Payload)
}

func (o *AddKernelCPUArchitectureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1CPUArchitecture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKernelCPUArchitectureBadRequest creates a AddKernelCPUArchitectureBadRequest with default headers values
func NewAddKernelCPUArchitectureBadRequest() *AddKernelCPUArchitectureBadRequest {
	return &AddKernelCPUArchitectureBadRequest{}
}

/*AddKernelCPUArchitectureBadRequest handles this case with default header values.

If the CPU architecture ID doesn't exist
*/
type AddKernelCPUArchitectureBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddKernelCPUArchitectureBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/cpu-architectures][%d] addKernelCpuArchitectureBadRequest  %+v", 400, o.Payload)
}

func (o *AddKernelCPUArchitectureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKernelCPUArchitectureDefault creates a AddKernelCPUArchitectureDefault with default headers values
func NewAddKernelCPUArchitectureDefault(code int) *AddKernelCPUArchitectureDefault {
	return &AddKernelCPUArchitectureDefault{
		_statusCode: code,
	}
}

/*AddKernelCPUArchitectureDefault handles this case with default header values.

generic error response
*/
type AddKernelCPUArchitectureDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add kernel Cpu architecture default response
func (o *AddKernelCPUArchitectureDefault) Code() int {
	return o._statusCode
}

func (o *AddKernelCPUArchitectureDefault) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/cpu-architectures][%d] addKernelCpuArchitecture default  %+v", o._statusCode, o.Payload)
}

func (o *AddKernelCPUArchitectureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddKernelCPUArchitectureBody add kernel CPU architecture body
swagger:model AddKernelCPUArchitectureBody
*/
type AddKernelCPUArchitectureBody struct {

	// The ID of the CPU architecture that can be used with this kernel
	// Required: true
	// Format: uuid
	CPUArchitectureID *strfmt.UUID `json:"cpu_architecture_id"`
}

// Validate validates this add kernel CPU architecture body
func (o *AddKernelCPUArchitectureBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCPUArchitectureID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddKernelCPUArchitectureBody) validateCPUArchitectureID(formats strfmt.Registry) error {

	if err := validate.Required("cpu_architecture_id"+"."+"cpu_architecture_id", "body", o.CPUArchitectureID); err != nil {
		return err
	}

	if err := validate.FormatOf("cpu_architecture_id"+"."+"cpu_architecture_id", "body", "uuid", o.CPUArchitectureID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddKernelCPUArchitectureBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddKernelCPUArchitectureBody) UnmarshalBinary(b []byte) error {
	var res AddKernelCPUArchitectureBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
