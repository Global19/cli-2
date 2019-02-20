// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestAPIValidationError REST API validation error body
//
// A error for when invalid input is supplied to an endpoint.
// swagger:model restApiValidationError
type RestAPIValidationError struct {

	// message
	// Required: true
	Message *string `json:"message"`

	// The list of validations errors found in the request
	// Required: true
	ValidationErrors []*RestAPIValidationErrorValidationErrorsItems0 `json:"validation_errors"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RestAPIValidationError) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Message = dataAO0.Message

	// AO1
	var dataAO1 struct {
		ValidationErrors []*RestAPIValidationErrorValidationErrorsItems0 `json:"validation_errors"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ValidationErrors = dataAO1.ValidationErrors

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RestAPIValidationError) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Message *string `json:"message"`
	}

	dataAO0.Message = m.Message

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		ValidationErrors []*RestAPIValidationErrorValidationErrorsItems0 `json:"validation_errors"`
	}

	dataAO1.ValidationErrors = m.ValidationErrors

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this rest Api validation error
func (m *RestAPIValidationError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAPIValidationError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *RestAPIValidationError) validateValidationErrors(formats strfmt.Registry) error {

	if err := validate.Required("validation_errors", "body", m.ValidationErrors); err != nil {
		return err
	}

	for i := 0; i < len(m.ValidationErrors); i++ {
		if swag.IsZero(m.ValidationErrors[i]) { // not required
			continue
		}

		if m.ValidationErrors[i] != nil {
			if err := m.ValidationErrors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validation_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestAPIValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAPIValidationError) UnmarshalBinary(b []byte) error {
	var res RestAPIValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RestAPIValidationErrorValidationErrorsItems0 A single validation error
// swagger:model RestAPIValidationErrorValidationErrorsItems0
type RestAPIValidationErrorValidationErrorsItems0 struct {

	// Human-readable message about this specific validation error
	// Required: true
	Error *string `json:"error"`

	// A JSON path to the specific field that was invalid, if applicable
	JSONPath string `json:"jsonPath,omitempty"`
}

// Validate validates this rest API validation error validation errors items0
func (m *RestAPIValidationErrorValidationErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAPIValidationErrorValidationErrorsItems0) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestAPIValidationErrorValidationErrorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAPIValidationErrorValidationErrorsItems0) UnmarshalBinary(b []byte) error {
	var res RestAPIValidationErrorValidationErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
