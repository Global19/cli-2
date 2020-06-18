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

// GetKernelVersionsReader is a Reader for the GetKernelVersions structure.
type GetKernelVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKernelVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKernelVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKernelVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKernelVersionsOK creates a GetKernelVersionsOK with default headers values
func NewGetKernelVersionsOK() *GetKernelVersionsOK {
	return &GetKernelVersionsOK{}
}

/*GetKernelVersionsOK handles this case with default header values.

A paginated list of kernel versions
*/
type GetKernelVersionsOK struct {
	Payload *inventory_models.V1KernelVersionPagedList
}

func (o *GetKernelVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/kernels/{kernel_id}/versions][%d] getKernelVersionsOK  %+v", 200, o.Payload)
}

func (o *GetKernelVersionsOK) GetPayload() *inventory_models.V1KernelVersionPagedList {
	return o.Payload
}

func (o *GetKernelVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1KernelVersionPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKernelVersionsDefault creates a GetKernelVersionsDefault with default headers values
func NewGetKernelVersionsDefault(code int) *GetKernelVersionsDefault {
	return &GetKernelVersionsDefault{
		_statusCode: code,
	}
}

/*GetKernelVersionsDefault handles this case with default header values.

generic error response
*/
type GetKernelVersionsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get kernel versions default response
func (o *GetKernelVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetKernelVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/kernels/{kernel_id}/versions][%d] getKernelVersions default  %+v", o._statusCode, o.Payload)
}

func (o *GetKernelVersionsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetKernelVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
