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

// V1ImageCore Image Core
//
// The properties of an image needed to create a new one
// swagger:model v1ImageCore
type V1ImageCore struct {

	// The name of the image (excluding any version information). This will be something like "activestate/centos-6.9-build" (for a Docker image) or "ami-foo-12345" (WindowsInstance).
	// Required: true
	Name *string `json:"name"`

	// The platform to which this image belongs.
	// Required: true
	PlatformID *string `json:"platform_id"`

	// The type of the image.
	// Required: true
	// Enum: [Docker WindowsInstance]
	Type *string `json:"type"`

	V1ImageCoreAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1ImageCore) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name *string `json:"name"`

		PlatformID *string `json:"platform_id"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	m.PlatformID = dataAO0.PlatformID

	m.Type = dataAO0.Type

	// AO1
	var aO1 V1ImageCoreAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1ImageCoreAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1ImageCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Name *string `json:"name"`

		PlatformID *string `json:"platform_id"`

		Type *string `json:"type"`
	}

	dataAO0.Name = m.Name

	dataAO0.PlatformID = m.PlatformID

	dataAO0.Type = m.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.V1ImageCoreAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 image core
func (m *V1ImageCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with V1ImageCoreAllOf1
	if err := m.V1ImageCoreAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageCore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCore) validatePlatformID(formats strfmt.Registry) error {

	if err := validate.Required("platform_id", "body", m.PlatformID); err != nil {
		return err
	}

	return nil
}

var v1ImageCoreTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Docker","WindowsInstance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ImageCoreTypeTypePropEnum = append(v1ImageCoreTypeTypePropEnum, v)
	}
}

// property enum
func (m *V1ImageCore) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1ImageCoreTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1ImageCore) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCore) UnmarshalBinary(b []byte) error {
	var res V1ImageCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1ImageCoreAllOf1 Image Revision Core
//
// The properties of an image revision needed to create a new one
// swagger:model V1ImageCoreAllOf1
type V1ImageCoreAllOf1 struct {

	// An array of decimal values representing all segments of a version, ordered from most to least significant. How a version string is rendered into a list of decimals will vary depending on the format of the source string and is therefore left up to the caller, but it must be done consistently across all versions of the same resource for sorting to work properly. This is represented as a string to avoid losing precision when converting to a floating point number.
	// Required: true
	// Min Length: 1
	SortableVersion []string `json:"sortable_version"`

	// The canonical version string for the resource. Should be as specific as possible (e.g. '10.9.6' of macOS instead of just '10.9'). May contain non-numeric version segments and other formatting characters if necessary.
	// Required: true
	Version *string `json:"version"`

	// Whether this revision should be considered 'stable'. When a new stable revision is created, it supercedes any existing stable revision and becomes the default revision of the revisioned resource going forward.
	IsStableRevision *bool `json:"is_stable_revision,omitempty"`

	// provided features
	// Required: true
	ProvidedFeatures []*V1ImageCoreAllOf1ProvidedFeaturesItems0 `json:"provided_features"`

	// conditions
	Conditions []*V1ImageCoreAllOf1ConditionsItems0 `json:"conditions"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1ImageCoreAllOf1) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		SortableVersion []string `json:"sortable_version"`

		Version *string `json:"version"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.SortableVersion = dataAO0.SortableVersion

	m.Version = dataAO0.Version

	// AO1
	var dataAO1 struct {
		IsStableRevision *bool `json:"is_stable_revision,omitempty"`

		ProvidedFeatures []*V1ImageCoreAllOf1ProvidedFeaturesItems0 `json:"provided_features"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.IsStableRevision = dataAO1.IsStableRevision

	m.ProvidedFeatures = dataAO1.ProvidedFeatures

	// AO2
	var dataAO2 struct {
		Conditions []*V1ImageCoreAllOf1ConditionsItems0 `json:"conditions,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Conditions = dataAO2.Conditions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1ImageCoreAllOf1) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	var dataAO0 struct {
		SortableVersion []string `json:"sortable_version"`

		Version *string `json:"version"`
	}

	dataAO0.SortableVersion = m.SortableVersion

	dataAO0.Version = m.Version

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		IsStableRevision *bool `json:"is_stable_revision,omitempty"`

		ProvidedFeatures []*V1ImageCoreAllOf1ProvidedFeaturesItems0 `json:"provided_features"`
	}

	dataAO1.IsStableRevision = m.IsStableRevision

	dataAO1.ProvidedFeatures = m.ProvidedFeatures

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	var dataAO2 struct {
		Conditions []*V1ImageCoreAllOf1ConditionsItems0 `json:"conditions,omitempty"`
	}

	dataAO2.Conditions = m.Conditions

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 image core all of1
func (m *V1ImageCoreAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSortableVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvidedFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageCoreAllOf1) validateSortableVersion(formats strfmt.Registry) error {

	if err := validate.Required("sortable_version", "body", m.SortableVersion); err != nil {
		return err
	}

	for i := 0; i < len(m.SortableVersion); i++ {

		if err := validate.MinLength("sortable_version"+"."+strconv.Itoa(i), "body", string(m.SortableVersion[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1ImageCoreAllOf1) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1) validateProvidedFeatures(formats strfmt.Registry) error {

	if err := validate.Required("provided_features", "body", m.ProvidedFeatures); err != nil {
		return err
	}

	for i := 0; i < len(m.ProvidedFeatures); i++ {
		if swag.IsZero(m.ProvidedFeatures[i]) { // not required
			continue
		}

		if m.ProvidedFeatures[i] != nil {
			if err := m.ProvidedFeatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("provided_features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1ImageCoreAllOf1) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageCoreAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCoreAllOf1) UnmarshalBinary(b []byte) error {
	var res V1ImageCoreAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1ImageCoreAllOf1ConditionsItems0 Condition Sub Schema
//
// A feature that must be present in a recipe for the containing entity to apply. If nothing in the recipe matches this condition, the containing entity is disable/cannot be used.
// swagger:model V1ImageCoreAllOf1ConditionsItems0
type V1ImageCoreAllOf1ConditionsItems0 struct {

	// What feature must be present for the containing entity to apply
	// Required: true
	Feature *string `json:"feature"`

	// The namespace the conditional feature is contained in
	// Required: true
	Namespace *string `json:"namespace"`

	// Requirements Sub Schema
	//
	// The version constraints that an ingredient version's requirement or condition puts on a feature
	// Required: true
	// Min Length: 1
	Requirements []*V1ImageCoreAllOf1ConditionsItems0RequirementsItems0 `json:"requirements"`
}

// Validate validates this v1 image core all of1 conditions items0
func (m *V1ImageCoreAllOf1ConditionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageCoreAllOf1ConditionsItems0) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1ConditionsItems0) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1ConditionsItems0) validateRequirements(formats strfmt.Registry) error {

	if err := validate.Required("requirements", "body", m.Requirements); err != nil {
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

// MarshalBinary interface implementation
func (m *V1ImageCoreAllOf1ConditionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCoreAllOf1ConditionsItems0) UnmarshalBinary(b []byte) error {
	var res V1ImageCoreAllOf1ConditionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0 v1 image core all of1 conditions items0 requirements items0
// swagger:model V1ImageCoreAllOf1ConditionsItems0RequirementsItems0
type V1ImageCoreAllOf1ConditionsItems0RequirementsItems0 struct {

	// The operator used to compare the sortable_version against a given provided feature to determine if it meets the requirement
	// Required: true
	// Enum: [eq gt gte lt lte ne]
	Comparator *string `json:"comparator"`

	// An array of decimal values representing all segments of a version, ordered from most to least significant. How a version string is rendered into a list of decimals will vary depending on the format of the source string and is therefore left up to the caller, but it must be done consistently across all versions of the same resource for sorting to work properly. This is represented as a string to avoid losing precision when converting to a floating point number.
	// Min Length: 1
	SortableVersion []string `json:"sortable_version"`

	// The required version in its original form.
	// Min Length: 1
	Version *string `json:"version,omitempty"`
}

// Validate validates this v1 image core all of1 conditions items0 requirements items0
func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortableVersion(formats); err != nil {
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

var v1ImageCoreAllOf1ConditionsItems0RequirementsItems0TypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ImageCoreAllOf1ConditionsItems0RequirementsItems0TypeComparatorPropEnum = append(v1ImageCoreAllOf1ConditionsItems0RequirementsItems0TypeComparatorPropEnum, v)
	}
}

const (

	// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorEq captures enum value "eq"
	V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorEq string = "eq"

	// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorGt captures enum value "gt"
	V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorGt string = "gt"

	// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorGte captures enum value "gte"
	V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorGte string = "gte"

	// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorLt captures enum value "lt"
	V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorLt string = "lt"

	// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorLte captures enum value "lte"
	V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorLte string = "lte"

	// V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorNe captures enum value "ne"
	V1ImageCoreAllOf1ConditionsItems0RequirementsItems0ComparatorNe string = "ne"
)

// prop value enum
func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1ImageCoreAllOf1ConditionsItems0RequirementsItems0TypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) validateSortableVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.SortableVersion) { // not required
		return nil
	}

	for i := 0; i < len(m.SortableVersion); i++ {

		if err := validate.MinLength("sortable_version"+"."+strconv.Itoa(i), "body", string(m.SortableVersion[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCoreAllOf1ConditionsItems0RequirementsItems0) UnmarshalBinary(b []byte) error {
	var res V1ImageCoreAllOf1ConditionsItems0RequirementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1ImageCoreAllOf1ProvidedFeaturesItems0 Provided Feature
//
// A feature that is provided by a revisioned resource
// swagger:model V1ImageCoreAllOf1ProvidedFeaturesItems0
type V1ImageCoreAllOf1ProvidedFeaturesItems0 struct {

	// feature
	// Required: true
	Feature *string `json:"feature"`

	// If this is true then it means that we assigned a version to this feature ourselves rather than getting it directly from metadata in the source ingredient.
	IsActivestateVersion *bool `json:"is_activestate_version,omitempty"`

	// Whether the provider of this feature is the default provider. There can only be one default provider per feature namespace, name, and version.
	// Required: true
	IsDefaultProvider *bool `json:"is_default_provider"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// An array of decimal values representing all segments of a version, ordered from most to least significant. How a version string is rendered into a list of decimals will vary depending on the format of the source string and is therefore left up to the caller, but it must be done consistently across all versions of the same resource for sorting to work properly. This is represented as a string to avoid losing precision when converting to a floating point number.
	// Required: true
	// Min Length: 1
	SortableVersion []string `json:"sortable_version"`

	// The canonical version string for the resource. Should be as specific as possible (e.g. '10.9.6' of macOS instead of just '10.9'). May contain non-numeric version segments and other formatting characters if necessary.
	// Required: true
	Version *string `json:"version"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Feature *string `json:"feature"`

		IsActivestateVersion *bool `json:"is_activestate_version,omitempty"`

		IsDefaultProvider *bool `json:"is_default_provider"`

		Namespace *string `json:"namespace"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Feature = dataAO0.Feature

	m.IsActivestateVersion = dataAO0.IsActivestateVersion

	m.IsDefaultProvider = dataAO0.IsDefaultProvider

	m.Namespace = dataAO0.Namespace

	// AO1
	var dataAO1 struct {
		SortableVersion []string `json:"sortable_version"`

		Version *string `json:"version"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.SortableVersion = dataAO1.SortableVersion

	m.Version = dataAO1.Version

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1ImageCoreAllOf1ProvidedFeaturesItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Feature *string `json:"feature"`

		IsActivestateVersion *bool `json:"is_activestate_version,omitempty"`

		IsDefaultProvider *bool `json:"is_default_provider"`

		Namespace *string `json:"namespace"`
	}

	dataAO0.Feature = m.Feature

	dataAO0.IsActivestateVersion = m.IsActivestateVersion

	dataAO0.IsDefaultProvider = m.IsDefaultProvider

	dataAO0.Namespace = m.Namespace

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		SortableVersion []string `json:"sortable_version"`

		Version *string `json:"version"`
	}

	dataAO1.SortableVersion = m.SortableVersion

	dataAO1.Version = m.Version

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 image core all of1 provided features items0
func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDefaultProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortableVersion(formats); err != nil {
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

func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) validateIsDefaultProvider(formats strfmt.Registry) error {

	if err := validate.Required("is_default_provider", "body", m.IsDefaultProvider); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) validateSortableVersion(formats strfmt.Registry) error {

	if err := validate.Required("sortable_version", "body", m.SortableVersion); err != nil {
		return err
	}

	for i := 0; i < len(m.SortableVersion); i++ {

		if err := validate.MinLength("sortable_version"+"."+strconv.Itoa(i), "body", string(m.SortableVersion[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCoreAllOf1ProvidedFeaturesItems0) UnmarshalBinary(b []byte) error {
	var res V1ImageCoreAllOf1ProvidedFeaturesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
