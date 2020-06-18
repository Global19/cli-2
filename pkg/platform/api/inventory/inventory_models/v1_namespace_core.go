// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1NamespaceCore Namespace Core
//
// The properties of a namespace needed to create a new one
//
// swagger:model v1NamespaceCore
type V1NamespaceCore struct {

	// is case sensitive
	// Required: true
	IsCaseSensitive *bool `json:"is_case_sensitive"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this v1 namespace core
func (m *V1NamespaceCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsCaseSensitive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NamespaceCore) validateIsCaseSensitive(formats strfmt.Registry) error {

	if err := validate.Required("is_case_sensitive", "body", m.IsCaseSensitive); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespaceCore) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NamespaceCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NamespaceCore) UnmarshalBinary(b []byte) error {
	var res V1NamespaceCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
