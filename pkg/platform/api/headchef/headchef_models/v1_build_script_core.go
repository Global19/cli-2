// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildScriptCore Build Script Core
//
// A short piece of scripting code that can be used to build an ingredient. This model only contains mutable fields of a build script.
//
// swagger:model v1BuildScriptCore
type V1BuildScriptCore struct {

	// The features that must already be present in the recipe for this build script to be used. For example, can be used to create build scripts that only work on specific operating systems.
	Conditions []*V1SubSchemaCondition `json:"conditions"`

	// The scripting language that the build script is written in
	// Required: true
	// Enum: [bash perl python]
	Language *string `json:"language"`

	// The build script itself
	// Required: true
	Script *string `json:"script"`
}

// Validate validates this v1 build script core
func (m *V1BuildScriptCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScript(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildScriptCore) validateConditions(formats strfmt.Registry) error {

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

var v1BuildScriptCoreTypeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bash","perl","python"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1BuildScriptCoreTypeLanguagePropEnum = append(v1BuildScriptCoreTypeLanguagePropEnum, v)
	}
}

const (

	// V1BuildScriptCoreLanguageBash captures enum value "bash"
	V1BuildScriptCoreLanguageBash string = "bash"

	// V1BuildScriptCoreLanguagePerl captures enum value "perl"
	V1BuildScriptCoreLanguagePerl string = "perl"

	// V1BuildScriptCoreLanguagePython captures enum value "python"
	V1BuildScriptCoreLanguagePython string = "python"
)

// prop value enum
func (m *V1BuildScriptCore) validateLanguageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1BuildScriptCoreTypeLanguagePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1BuildScriptCore) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	// value enum
	if err := m.validateLanguageEnum("language", "body", *m.Language); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptCore) validateScript(formats strfmt.Registry) error {

	if err := validate.Required("script", "body", m.Script); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScriptCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptCore) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
