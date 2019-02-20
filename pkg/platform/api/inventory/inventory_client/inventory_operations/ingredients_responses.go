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

// IngredientsReader is a Reader for the Ingredients structure.
type IngredientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IngredientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIngredientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIngredientsOK creates a IngredientsOK with default headers values
func NewIngredientsOK() *IngredientsOK {
	return &IngredientsOK{}
}

/*IngredientsOK handles this case with default header values.

Returns a list of ingredients and any of their versions which provide the requested package. If nothing provides the requested package, returns an empty list.
*/
type IngredientsOK struct {
	Payload []*inventory_models.IngredientAndVersions
}

func (o *IngredientsOK) Error() string {
	return fmt.Sprintf("[GET /ingredients][%d] ingredientsOK  %+v", 200, o.Payload)
}

func (o *IngredientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
