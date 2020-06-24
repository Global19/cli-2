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

// UpdateBuildScriptReader is a Reader for the UpdateBuildScript structure.
type UpdateBuildScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBuildScriptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateBuildScriptBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateBuildScriptDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBuildScriptOK creates a UpdateBuildScriptOK with default headers values
func NewUpdateBuildScriptOK() *UpdateBuildScriptOK {
	return &UpdateBuildScriptOK{}
}

/*UpdateBuildScriptOK handles this case with default header values.

The updated build script
*/
type UpdateBuildScriptOK struct {
	Payload *inventory_models.V1BuildScript
}

func (o *UpdateBuildScriptOK) Error() string {
	return fmt.Sprintf("[PUT /v1/build-scripts/{build_script_id}][%d] updateBuildScriptOK  %+v", 200, o.Payload)
}

func (o *UpdateBuildScriptOK) GetPayload() *inventory_models.V1BuildScript {
	return o.Payload
}

func (o *UpdateBuildScriptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1BuildScript)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildScriptBadRequest creates a UpdateBuildScriptBadRequest with default headers values
func NewUpdateBuildScriptBadRequest() *UpdateBuildScriptBadRequest {
	return &UpdateBuildScriptBadRequest{}
}

/*UpdateBuildScriptBadRequest handles this case with default header values.

If the build script update is invalid or the build script cannot be updated because it is in use by a stable ingredient version revision
*/
type UpdateBuildScriptBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *UpdateBuildScriptBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/build-scripts/{build_script_id}][%d] updateBuildScriptBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateBuildScriptBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *UpdateBuildScriptBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildScriptDefault creates a UpdateBuildScriptDefault with default headers values
func NewUpdateBuildScriptDefault(code int) *UpdateBuildScriptDefault {
	return &UpdateBuildScriptDefault{
		_statusCode: code,
	}
}

/*UpdateBuildScriptDefault handles this case with default header values.

If there is an error processing the request
*/
type UpdateBuildScriptDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the update build script default response
func (o *UpdateBuildScriptDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBuildScriptDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/build-scripts/{build_script_id}][%d] updateBuildScript default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBuildScriptDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *UpdateBuildScriptDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
