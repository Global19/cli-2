// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1Order Order
//
// An order to create recipes for a set of requirements, for one or more platforms.
// swagger:model v1Order
type V1Order struct {

	// Camel-specific flags for controlling the build.
	CamelFlags []string `json:"camel_flags"`

	// Order UUID, supplied by client to be copied to all resulting recipes
	// Required: true
	// Format: uuid
	OrderID *strfmt.UUID `json:"order_id"`

	// List of platform IDs for the order
	// Required: true
	// Unique: true
	Platforms []strfmt.UUID `json:"platforms"`

	// The list of required features needed to satisfy this order
	// Required: true
	// Min Items: 1
	Requirements []*V1OrderRequirementsItems0 `json:"requirements"`

	// The date and time that the order was originally submitted
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this v1 order
func (m *V1Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Order) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	if err := validate.FormatOf("order_id", "body", "uuid", m.OrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Order) validatePlatforms(formats strfmt.Registry) error {

	if err := validate.Required("platforms", "body", m.Platforms); err != nil {
		return err
	}

	if err := validate.UniqueItems("platforms", "body", m.Platforms); err != nil {
		return err
	}

	for i := 0; i < len(m.Platforms); i++ {

		if err := validate.FormatOf("platforms"+"."+strconv.Itoa(i), "body", "uuid", m.Platforms[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1Order) validateRequirements(formats strfmt.Registry) error {

	if err := validate.Required("requirements", "body", m.Requirements); err != nil {
		return err
	}

	iRequirementsSize := int64(len(m.Requirements))

	if err := validate.MinItems("requirements", "body", iRequirementsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Requirements); i++ {
		if swag.IsZero(m.Requirements[i]) { // not required
			continue
		}

		if m.Requirements[i] != nil {
			if err := m.Requirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Order) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Order) UnmarshalBinary(b []byte) error {
	var res V1Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1OrderRequirementsItems0 Requirement Sub Schema
//
// An order requirement is a single package name and version specifier requested in an order.
// swagger:model V1OrderRequirementsItems0
type V1OrderRequirementsItems0 struct {

	// The name of the required feature, If no ingredient ID is specified, the default provider of this feature will be chosen.
	// Required: true
	Feature *string `json:"feature"`

	// The ID of the ingredient version that should be used to fulfill this requirement. Can be used to override the default choice of provider for the specified feature. Must be an ingredient version that actually provides the specified feature.
	// Format: uuid
	IngredientVersionID strfmt.UUID `json:"ingredient_version_id,omitempty"`

	// The namespace for the required feature. For now, this can be empty as it is only used to request pre-platform installer ingredients.
	// Required: true
	Namespace *string `json:"namespace"`

	// The requirements for the acceptable versions of this feature. This can be omitted, in which case any version is acceptable.
	// Min Items: 1
	VersionRequirements []*V1OrderRequirementsItems0VersionRequirementsItems0 `json:"version_requirements"`
}

// Validate validates this v1 order requirements items0
func (m *V1OrderRequirementsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OrderRequirementsItems0) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1OrderRequirementsItems0) validateIngredientVersionID(formats strfmt.Registry) error {

	if swag.IsZero(m.IngredientVersionID) { // not required
		return nil
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1OrderRequirementsItems0) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1OrderRequirementsItems0) validateVersionRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionRequirements) { // not required
		return nil
	}

	iVersionRequirementsSize := int64(len(m.VersionRequirements))

	if err := validate.MinItems("version_requirements", "body", iVersionRequirementsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.VersionRequirements); i++ {
		if swag.IsZero(m.VersionRequirements[i]) { // not required
			continue
		}

		if m.VersionRequirements[i] != nil {
			if err := m.VersionRequirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("version_requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OrderRequirementsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OrderRequirementsItems0) UnmarshalBinary(b []byte) error {
	var res V1OrderRequirementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1OrderRequirementsItems0VersionRequirementsItems0 Version Requirement Sub Schema
//
// The version constraint for a feature
// swagger:model V1OrderRequirementsItems0VersionRequirementsItems0
type V1OrderRequirementsItems0VersionRequirementsItems0 struct {

	// The operator used to compare the version against a given provided feature to determine if it meets this requirement
	// Required: true
	// Enum: [eq gt gte lt lte ne]
	Comparator *string `json:"comparator"`

	// The required version in its original form
	// Required: true
	// Min Length: 1
	Version *string `json:"version"`
}

// Validate validates this v1 order requirements items0 version requirements items0
func (m *V1OrderRequirementsItems0VersionRequirementsItems0) Validate(formats strfmt.Registry) error {
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

var v1OrderRequirementsItems0VersionRequirementsItems0TypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1OrderRequirementsItems0VersionRequirementsItems0TypeComparatorPropEnum = append(v1OrderRequirementsItems0VersionRequirementsItems0TypeComparatorPropEnum, v)
	}
}

const (

	// V1OrderRequirementsItems0VersionRequirementsItems0ComparatorEq captures enum value "eq"
	V1OrderRequirementsItems0VersionRequirementsItems0ComparatorEq string = "eq"

	// V1OrderRequirementsItems0VersionRequirementsItems0ComparatorGt captures enum value "gt"
	V1OrderRequirementsItems0VersionRequirementsItems0ComparatorGt string = "gt"

	// V1OrderRequirementsItems0VersionRequirementsItems0ComparatorGte captures enum value "gte"
	V1OrderRequirementsItems0VersionRequirementsItems0ComparatorGte string = "gte"

	// V1OrderRequirementsItems0VersionRequirementsItems0ComparatorLt captures enum value "lt"
	V1OrderRequirementsItems0VersionRequirementsItems0ComparatorLt string = "lt"

	// V1OrderRequirementsItems0VersionRequirementsItems0ComparatorLte captures enum value "lte"
	V1OrderRequirementsItems0VersionRequirementsItems0ComparatorLte string = "lte"

	// V1OrderRequirementsItems0VersionRequirementsItems0ComparatorNe captures enum value "ne"
	V1OrderRequirementsItems0VersionRequirementsItems0ComparatorNe string = "ne"
)

// prop value enum
func (m *V1OrderRequirementsItems0VersionRequirementsItems0) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1OrderRequirementsItems0VersionRequirementsItems0TypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1OrderRequirementsItems0VersionRequirementsItems0) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1OrderRequirementsItems0VersionRequirementsItems0) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OrderRequirementsItems0VersionRequirementsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OrderRequirementsItems0VersionRequirementsItems0) UnmarshalBinary(b []byte) error {
	var res V1OrderRequirementsItems0VersionRequirementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
