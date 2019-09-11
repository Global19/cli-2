// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CPUExtension CPU Extension
//
// The full CPU extension data model
// swagger:model v1CpuExtension
type V1CPUExtension struct {

	// cpu extension id
	// Required: true
	// Format: uuid
	CPUExtensionID *strfmt.UUID `json:"cpu_extension_id"`

	// links
	// Required: true
	Links *V1CPUExtensionAO0Links `json:"links"`

	V1CPUExtensionAllOf1

	// The revision number of this revision of the resource. This number increases monotonically with each new revision.
	// Required: true
	// Minimum: 1
	Revision *int64 `json:"revision"`

	// The date and time at which this revision of the resource was created
	// Required: true
	// Format: date-time
	RevisionTimestamp *strfmt.DateTime `json:"revision_timestamp"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1CPUExtension) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		CPUExtensionID *strfmt.UUID `json:"cpu_extension_id"`

		Links *V1CPUExtensionAO0Links `json:"links"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.CPUExtensionID = dataAO0.CPUExtensionID

	m.Links = dataAO0.Links

	// AO1
	var aO1 V1CPUExtensionAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1CPUExtensionAllOf1 = aO1

	// AO2
	var dataAO2 struct {
		Revision *int64 `json:"revision"`

		RevisionTimestamp *strfmt.DateTime `json:"revision_timestamp"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Revision = dataAO2.Revision

	m.RevisionTimestamp = dataAO2.RevisionTimestamp

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1CPUExtension) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	var dataAO0 struct {
		CPUExtensionID *strfmt.UUID `json:"cpu_extension_id"`

		Links *V1CPUExtensionAO0Links `json:"links"`
	}

	dataAO0.CPUExtensionID = m.CPUExtensionID

	dataAO0.Links = m.Links

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.V1CPUExtensionAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	var dataAO2 struct {
		Revision *int64 `json:"revision"`

		RevisionTimestamp *strfmt.DateTime `json:"revision_timestamp"`
	}

	dataAO2.Revision = m.Revision

	dataAO2.RevisionTimestamp = m.RevisionTimestamp

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 Cpu extension
func (m *V1CPUExtension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUExtensionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with V1CPUExtensionAllOf1
	if err := m.V1CPUExtensionAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevisionTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CPUExtension) validateCPUExtensionID(formats strfmt.Registry) error {

	if err := validate.Required("cpu_extension_id", "body", m.CPUExtensionID); err != nil {
		return err
	}

	if err := validate.FormatOf("cpu_extension_id", "body", "uuid", m.CPUExtensionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUExtension) validateLinks(formats strfmt.Registry) error {

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

func (m *V1CPUExtension) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	if err := validate.MinimumInt("revision", "body", int64(*m.Revision), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUExtension) validateRevisionTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("revision_timestamp", "body", m.RevisionTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("revision_timestamp", "body", "date-time", m.RevisionTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CPUExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CPUExtension) UnmarshalBinary(b []byte) error {
	var res V1CPUExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1CPUExtensionAO0Links Self Link
//
// A self link
// swagger:model V1CPUExtensionAO0Links
type V1CPUExtensionAO0Links struct {

	// The URI of this resource
	// Required: true
	// Format: uri
	Self *strfmt.URI `json:"self"`
}

// Validate validates this v1 CPU extension a o0 links
func (m *V1CPUExtensionAO0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CPUExtensionAO0Links) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("links"+"."+"self", "body", m.Self); err != nil {
		return err
	}

	if err := validate.FormatOf("links"+"."+"self", "body", "uri", m.Self.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CPUExtensionAO0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CPUExtensionAO0Links) UnmarshalBinary(b []byte) error {
	var res V1CPUExtensionAO0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1CPUExtensionAllOf1 CPU Extension Core
//
// The properties of a CPU extension needed to create a new one
// swagger:model V1CPUExtensionAllOf1
type V1CPUExtensionAllOf1 struct {

	// The name of the CPU extension
	// Required: true
	Name *string `json:"name"`

	// Whether this revision should be considered 'stable'. When a new stable revision is created, it supercedes any existing stable revision and becomes the default revision of the revisioned resource going forward.
	IsStableRevision *bool `json:"is_stable_revision,omitempty"`

	// provided features
	// Required: true
	ProvidedFeatures []*V1CPUExtensionAllOf1ProvidedFeaturesItems0 `json:"provided_features"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1CPUExtensionAllOf1) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	// AO1
	var dataAO1 struct {
		IsStableRevision *bool `json:"is_stable_revision,omitempty"`

		ProvidedFeatures []*V1CPUExtensionAllOf1ProvidedFeaturesItems0 `json:"provided_features"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.IsStableRevision = dataAO1.IsStableRevision

	m.ProvidedFeatures = dataAO1.ProvidedFeatures

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1CPUExtensionAllOf1) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Name *string `json:"name"`
	}

	dataAO0.Name = m.Name

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		IsStableRevision *bool `json:"is_stable_revision,omitempty"`

		ProvidedFeatures []*V1CPUExtensionAllOf1ProvidedFeaturesItems0 `json:"provided_features"`
	}

	dataAO1.IsStableRevision = m.IsStableRevision

	dataAO1.ProvidedFeatures = m.ProvidedFeatures

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 CPU extension all of1
func (m *V1CPUExtensionAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvidedFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1CPUExtensionAllOf1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUExtensionAllOf1) validateProvidedFeatures(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V1CPUExtensionAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CPUExtensionAllOf1) UnmarshalBinary(b []byte) error {
	var res V1CPUExtensionAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// V1CPUExtensionAllOf1ProvidedFeaturesItems0 Provided Feature
//
// A feature that is provided by a revisioned resource
// swagger:model V1CPUExtensionAllOf1ProvidedFeaturesItems0
type V1CPUExtensionAllOf1ProvidedFeaturesItems0 struct {

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
func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) UnmarshalJSON(raw []byte) error {
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
func (m V1CPUExtensionAllOf1ProvidedFeaturesItems0) MarshalJSON() ([]byte, error) {
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

// Validate validates this v1 CPU extension all of1 provided features items0
func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) Validate(formats strfmt.Registry) error {
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

func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) validateIsDefaultProvider(formats strfmt.Registry) error {

	if err := validate.Required("is_default_provider", "body", m.IsDefaultProvider); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) validateSortableVersion(formats strfmt.Registry) error {

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

func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CPUExtensionAllOf1ProvidedFeaturesItems0) UnmarshalBinary(b []byte) error {
	var res V1CPUExtensionAllOf1ProvidedFeaturesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
