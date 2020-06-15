// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1IngredientVersionCore Ingredient Version Core
//
// The fields of an ingredient version that can be set when an ingredient version is created
//
// swagger:model v1IngredientVersionCore
type V1IngredientVersionCore struct {
	V1IngredientVersionUpdate

	V1SubSchemaVersionInfo
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1IngredientVersionCore) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1IngredientVersionUpdate
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1IngredientVersionUpdate = aO0

	// AO1
	var aO1 V1SubSchemaVersionInfo
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1SubSchemaVersionInfo = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1IngredientVersionCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1IngredientVersionUpdate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1SubSchemaVersionInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 ingredient version core
func (m *V1IngredientVersionCore) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1IngredientVersionUpdate
	if err := m.V1IngredientVersionUpdate.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1SubSchemaVersionInfo
	if err := m.V1SubSchemaVersionInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionCore) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
