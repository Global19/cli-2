// Code generated by go-swagger; DO NOT EDIT.

package secrets_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UID UID
// swagger:model UID
type UID struct {

	// uid
	// Required: true
	// Format: uuid
	UID *strfmt.UUID `json:"uid"`
}

// Validate validates this UID
func (m *UID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UID) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UID) UnmarshalBinary(b []byte) error {
	var res UID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}