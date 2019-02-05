// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserSecret user secret
// swagger:model UserSecret
type UserSecret struct {

	// is user
	// Required: true
	IsUser *bool `json:"is_user"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization id
	// Required: true
	// Format: uuid
	OrganizationID *strfmt.UUID `json:"organization_id"`

	// project id
	// Format: uuid
	ProjectID strfmt.UUID `json:"project_id,omitempty"`

	// secret id
	// Required: true
	// Format: uuid
	SecretID *strfmt.UUID `json:"secret_id"`

	// user id
	// Required: true
	// Format: uuid
	UserID *strfmt.UUID `json:"user_id"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this user secret
func (m *UserSecret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSecret) validateIsUser(formats strfmt.Registry) error {

	if err := validate.Required("is_user", "body", m.IsUser); err != nil {
		return err
	}

	return nil
}

func (m *UserSecret) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UserSecret) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organization_id", "body", m.OrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organization_id", "body", "uuid", m.OrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserSecret) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("project_id", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserSecret) validateSecretID(formats strfmt.Registry) error {

	if err := validate.Required("secret_id", "body", m.SecretID); err != nil {
		return err
	}

	if err := validate.FormatOf("secret_id", "body", "uuid", m.SecretID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserSecret) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserSecret) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSecret) UnmarshalBinary(b []byte) error {
	var res UserSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}