// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SubSchemaVersionRequirement Version Requirement Sub Schema
//
// The version constraint for a feature
//
// swagger:model v1SubSchemaVersionRequirement
type V1SubSchemaVersionRequirement struct {

	// The operator used to compare the version against a given provided feature to determine if it meets this requirement
	// Required: true
	// Enum: [eq gt gte lt lte ne]
	Comparator *string `json:"comparator"`

	// The required version in its original form
	// Required: true
	// Min Length: 1
	Version *string `json:"version"`
}

// Validate validates this v1 sub schema version requirement
func (m *V1SubSchemaVersionRequirement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1SubSchemaVersionRequirementTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SubSchemaVersionRequirementTypeComparatorPropEnum = append(v1SubSchemaVersionRequirementTypeComparatorPropEnum, v)
	}
}

const (

	// V1SubSchemaVersionRequirementComparatorEq captures enum value "eq"
	V1SubSchemaVersionRequirementComparatorEq string = "eq"

	// V1SubSchemaVersionRequirementComparatorGt captures enum value "gt"
	V1SubSchemaVersionRequirementComparatorGt string = "gt"

	// V1SubSchemaVersionRequirementComparatorGte captures enum value "gte"
	V1SubSchemaVersionRequirementComparatorGte string = "gte"

	// V1SubSchemaVersionRequirementComparatorLt captures enum value "lt"
	V1SubSchemaVersionRequirementComparatorLt string = "lt"

	// V1SubSchemaVersionRequirementComparatorLte captures enum value "lte"
	V1SubSchemaVersionRequirementComparatorLte string = "lte"

	// V1SubSchemaVersionRequirementComparatorNe captures enum value "ne"
	V1SubSchemaVersionRequirementComparatorNe string = "ne"
)

// prop value enum
func (m *V1SubSchemaVersionRequirement) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1SubSchemaVersionRequirementTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1SubSchemaVersionRequirement) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1SubSchemaVersionRequirement) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SubSchemaVersionRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SubSchemaVersionRequirement) UnmarshalBinary(b []byte) error {
	var res V1SubSchemaVersionRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}