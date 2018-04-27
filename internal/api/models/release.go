// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Release release
// swagger:model Release
type Release struct {

	// added
	Added strfmt.DateTime `json:"added,omitempty"`

	// build number
	BuildNumber int64 `json:"buildNumber,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// manifest
	Manifest []*ReleaseManifestItems `json:"manifest"`

	// name
	Name string `json:"name,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// project ID
	ProjectID strfmt.UUID `json:"projectID,omitempty"`

	// release ID
	ReleaseID strfmt.UUID `json:"releaseID,omitempty"`
}

// Validate validates this release
func (m *Release) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManifest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReleaseID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateManifest(formats strfmt.Registry) error {

	if swag.IsZero(m.Manifest) { // not required
		return nil
	}

	for i := 0; i < len(m.Manifest); i++ {

		if swag.IsZero(m.Manifest[i]) { // not required
			continue
		}

		if m.Manifest[i] != nil {

			if err := m.Manifest[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("manifest" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Release) validateProjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("projectID", "body", "uuid", m.ProjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateReleaseID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReleaseID) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseID", "body", "uuid", m.ReleaseID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Release) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Release) UnmarshalBinary(b []byte) error {
	var res Release
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}