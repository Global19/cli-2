// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// LoginAsReader is a Reader for the LoginAs structure.
type LoginAsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginAsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoginAsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLoginAsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewLoginAsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLoginAsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoginAsOK creates a LoginAsOK with default headers values
func NewLoginAsOK() *LoginAsOK {
	return &LoginAsOK{}
}

/*LoginAsOK handles this case with default header values.

Success
*/
type LoginAsOK struct {
	Payload *mono_models.JWT
}

func (o *LoginAsOK) Error() string {
	return fmt.Sprintf("[POST /login/{username}][%d] loginAsOK  %+v", 200, o.Payload)
}

func (o *LoginAsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.JWT)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginAsBadRequest creates a LoginAsBadRequest with default headers values
func NewLoginAsBadRequest() *LoginAsBadRequest {
	return &LoginAsBadRequest{}
}

/*LoginAsBadRequest handles this case with default header values.

Bad Request
*/
type LoginAsBadRequest struct {
	Payload *mono_models.Message
}

func (o *LoginAsBadRequest) Error() string {
	return fmt.Sprintf("[POST /login/{username}][%d] loginAsBadRequest  %+v", 400, o.Payload)
}

func (o *LoginAsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginAsUnauthorized creates a LoginAsUnauthorized with default headers values
func NewLoginAsUnauthorized() *LoginAsUnauthorized {
	return &LoginAsUnauthorized{}
}

/*LoginAsUnauthorized handles this case with default header values.

Invalid credentials
*/
type LoginAsUnauthorized struct {
	Payload *mono_models.Message
}

func (o *LoginAsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /login/{username}][%d] loginAsUnauthorized  %+v", 401, o.Payload)
}

func (o *LoginAsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginAsInternalServerError creates a LoginAsInternalServerError with default headers values
func NewLoginAsInternalServerError() *LoginAsInternalServerError {
	return &LoginAsInternalServerError{}
}

/*LoginAsInternalServerError handles this case with default header values.

Server Error
*/
type LoginAsInternalServerError struct {
	Payload *mono_models.Message
}

func (o *LoginAsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /login/{username}][%d] loginAsInternalServerError  %+v", 500, o.Payload)
}

func (o *LoginAsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
