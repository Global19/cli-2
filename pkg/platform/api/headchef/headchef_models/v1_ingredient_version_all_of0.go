// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IngredientVersionAllOf0 v1 ingredient version all of0
//
// swagger:model v1IngredientVersionAllOf0
type V1IngredientVersionAllOf0 struct {

	// creation timestamp
	// Required: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

	// ingredient id
	// Required: true
	// Format: uuid
	IngredientID *strfmt.UUID `json:"ingredient_id"`

	// ingredient version id
	// Required: true
	// Format: uuid
	IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`

	// links
	// Required: true
	Links *V1IngredientVersionAllOf0Links `json:"links"`
}

// Validate validates this v1 ingredient version all of0
func (m *V1IngredientVersionAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientVersionAllOf0) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("creation_timestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionAllOf0) validateIngredientID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_id", "body", m.IngredientID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_id", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionAllOf0) validateIngredientVersionID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version_id", "body", m.IngredientVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionAllOf0) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionAllOf0) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
