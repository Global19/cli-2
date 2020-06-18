// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1ImageCoreAllOf0 v1 image core all of0
//
// swagger:model v1ImageCoreAllOf0
type V1ImageCoreAllOf0 struct {

	// The name of the image (excluding any version information). This will be something like "activestate/centos-6.9-build" (for a Docker image) or "ami-foo-12345" (WindowsInstance).
	// Required: true
	Name *string `json:"name"`

	// The platform to which this image belongs.
	// Required: true
	// Format: uuid
	PlatformID *strfmt.UUID `json:"platform_id"`

	// The type of the image.
	// Required: true
	// Enum: [Docker Mac WindowsDocker WindowsInstance]
	Type *string `json:"type"`
}

// Validate validates this v1 image core all of0
func (m *V1ImageCoreAllOf0) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1ImageCoreAllOf0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1ImageCoreAllOf0) validatePlatformID(formats strfmt.Registry) error {

	if err := validate.Required("platform_id", "body", m.PlatformID); err != nil {
		return err
	}

	if err := validate.FormatOf("platform_id", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

var v1ImageCoreAllOf0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Docker","Mac","WindowsDocker","WindowsInstance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ImageCoreAllOf0TypeTypePropEnum = append(v1ImageCoreAllOf0TypeTypePropEnum, v)
	}
}

const (

	// V1ImageCoreAllOf0TypeDocker captures enum value "Docker"
	V1ImageCoreAllOf0TypeDocker string = "Docker"

	// V1ImageCoreAllOf0TypeMac captures enum value "Mac"
	V1ImageCoreAllOf0TypeMac string = "Mac"

	// V1ImageCoreAllOf0TypeWindowsDocker captures enum value "WindowsDocker"
	V1ImageCoreAllOf0TypeWindowsDocker string = "WindowsDocker"

	// V1ImageCoreAllOf0TypeWindowsInstance captures enum value "WindowsInstance"
	V1ImageCoreAllOf0TypeWindowsInstance string = "WindowsInstance"
)

// prop value enum
func (m *V1ImageCoreAllOf0) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1ImageCoreAllOf0TypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1ImageCoreAllOf0) validateType(formats strfmt.Registry) error {

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
func (m *V1ImageCoreAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ImageCoreAllOf0) UnmarshalBinary(b []byte) error {
	var res V1ImageCoreAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
