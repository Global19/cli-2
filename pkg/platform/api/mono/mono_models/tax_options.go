// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaxOptions tax options
// swagger:model TaxOptions
type TaxOptions struct {

	// org URL name
	// Required: true
	OrgURLName string `json:"orgURLName"`

	// shipping address
	// Required: true
	ShippingAddress *AddressInfo `json:"shippingAddress"`

	// tier name
	// Required: true
	TierName string `json:"tierName"`

	// users
	// Required: true
	Users int64 `json:"users"`
}

// Validate validates this tax options
func (m *TaxOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrgURLName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxOptions) validateOrgURLName(formats strfmt.Registry) error {

	if err := validate.RequiredString("orgURLName", "body", string(m.OrgURLName)); err != nil {
		return err
	}

	return nil
}

func (m *TaxOptions) validateShippingAddress(formats strfmt.Registry) error {

	if err := validate.Required("shippingAddress", "body", m.ShippingAddress); err != nil {
		return err
	}

	if m.ShippingAddress != nil {
		if err := m.ShippingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shippingAddress")
			}
			return err
		}
	}

	return nil
}

func (m *TaxOptions) validateTierName(formats strfmt.Registry) error {

	if err := validate.RequiredString("tierName", "body", string(m.TierName)); err != nil {
		return err
	}

	return nil
}

func (m *TaxOptions) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", int64(m.Users)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxOptions) UnmarshalBinary(b []byte) error {
	var res TaxOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}