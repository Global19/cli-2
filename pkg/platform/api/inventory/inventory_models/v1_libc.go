// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Libc Libc
//
// The full libc data model
//
// swagger:model v1Libc
type V1Libc struct {
	V1LibcAllOf0

	V1LibcCore
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1Libc) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1LibcAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1LibcAllOf0 = aO0

	// AO1
	var aO1 V1LibcCore
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1LibcCore = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1Libc) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1LibcAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1LibcCore)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 libc
func (m *V1Libc) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1LibcAllOf0
	if err := m.V1LibcAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1LibcCore
	if err := m.V1LibcCore.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1Libc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Libc) UnmarshalBinary(b []byte) error {
	var res V1Libc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
