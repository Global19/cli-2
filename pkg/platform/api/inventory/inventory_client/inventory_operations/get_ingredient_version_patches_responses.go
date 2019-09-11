// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetIngredientVersionPatchesReader is a Reader for the GetIngredientVersionPatches structure.
type GetIngredientVersionPatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIngredientVersionPatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIngredientVersionPatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetIngredientVersionPatchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIngredientVersionPatchesOK creates a GetIngredientVersionPatchesOK with default headers values
func NewGetIngredientVersionPatchesOK() *GetIngredientVersionPatchesOK {
	return &GetIngredientVersionPatchesOK{}
}

/*GetIngredientVersionPatchesOK handles this case with default header values.

A paginated list of patches
*/
type GetIngredientVersionPatchesOK struct {
	Payload *inventory_models.V1PatchPagedList
}

func (o *GetIngredientVersionPatchesOK) Error() string {
	return fmt.Sprintf("[GET /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/patches][%d] getIngredientVersionPatchesOK  %+v", 200, o.Payload)
}

func (o *GetIngredientVersionPatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1PatchPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIngredientVersionPatchesDefault creates a GetIngredientVersionPatchesDefault with default headers values
func NewGetIngredientVersionPatchesDefault(code int) *GetIngredientVersionPatchesDefault {
	return &GetIngredientVersionPatchesDefault{
		_statusCode: code,
	}
}

/*GetIngredientVersionPatchesDefault handles this case with default header values.

generic error response
*/
type GetIngredientVersionPatchesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get ingredient version patches default response
func (o *GetIngredientVersionPatchesDefault) Code() int {
	return o._statusCode
}

func (o *GetIngredientVersionPatchesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}/patches][%d] getIngredientVersionPatches default  %+v", o._statusCode, o.Payload)
}

func (o *GetIngredientVersionPatchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
