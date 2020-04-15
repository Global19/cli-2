// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListFormatsReader is a Reader for the ListFormats structure.
type ListFormatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFormatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFormatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListFormatsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListFormatsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListFormatsOK creates a ListFormatsOK with default headers values
func NewListFormatsOK() *ListFormatsOK {
	return &ListFormatsOK{}
}

/*ListFormatsOK handles this case with default header values.

Success
*/
type ListFormatsOK struct {
	Payload []*mono_models.Format
}

func (o *ListFormatsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] listFormatsOK  %+v", 200, o.Payload)
}

func (o *ListFormatsOK) GetPayload() []*mono_models.Format {
	return o.Payload
}

func (o *ListFormatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFormatsNotFound creates a ListFormatsNotFound with default headers values
func NewListFormatsNotFound() *ListFormatsNotFound {
	return &ListFormatsNotFound{}
}

/*ListFormatsNotFound handles this case with default header values.

Not Found
*/
type ListFormatsNotFound struct {
	Payload *mono_models.Message
}

func (o *ListFormatsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] listFormatsNotFound  %+v", 404, o.Payload)
}

func (o *ListFormatsNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *ListFormatsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFormatsInternalServerError creates a ListFormatsInternalServerError with default headers values
func NewListFormatsInternalServerError() *ListFormatsInternalServerError {
	return &ListFormatsInternalServerError{}
}

/*ListFormatsInternalServerError handles this case with default header values.

Server Error
*/
type ListFormatsInternalServerError struct {
	Payload *mono_models.Message
}

func (o *ListFormatsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] listFormatsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFormatsInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *ListFormatsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
