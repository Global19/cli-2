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

// AddLibcVersionReader is a Reader for the AddLibcVersion structure.
type AddLibcVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLibcVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddLibcVersionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddLibcVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddLibcVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddLibcVersionCreated creates a AddLibcVersionCreated with default headers values
func NewAddLibcVersionCreated() *AddLibcVersionCreated {
	return &AddLibcVersionCreated{}
}

/*AddLibcVersionCreated handles this case with default header values.

The added libc version
*/
type AddLibcVersionCreated struct {
	Payload *inventory_models.V1LibcVersion
}

func (o *AddLibcVersionCreated) Error() string {
	return fmt.Sprintf("[POST /v1/libcs/{libc_id}/versions][%d] addLibcVersionCreated  %+v", 201, o.Payload)
}

func (o *AddLibcVersionCreated) GetPayload() *inventory_models.V1LibcVersion {
	return o.Payload
}

func (o *AddLibcVersionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1LibcVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLibcVersionBadRequest creates a AddLibcVersionBadRequest with default headers values
func NewAddLibcVersionBadRequest() *AddLibcVersionBadRequest {
	return &AddLibcVersionBadRequest{}
}

/*AddLibcVersionBadRequest handles this case with default header values.

If the libc version is invalid
*/
type AddLibcVersionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddLibcVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/libcs/{libc_id}/versions][%d] addLibcVersionBadRequest  %+v", 400, o.Payload)
}

func (o *AddLibcVersionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddLibcVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLibcVersionDefault creates a AddLibcVersionDefault with default headers values
func NewAddLibcVersionDefault(code int) *AddLibcVersionDefault {
	return &AddLibcVersionDefault{
		_statusCode: code,
	}
}

/*AddLibcVersionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddLibcVersionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add libc version default response
func (o *AddLibcVersionDefault) Code() int {
	return o._statusCode
}

func (o *AddLibcVersionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/libcs/{libc_id}/versions][%d] addLibcVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AddLibcVersionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddLibcVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
