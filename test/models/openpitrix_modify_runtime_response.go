// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRuntimeResponse openpitrix modify runtime response
// swagger:model openpitrixModifyRuntimeResponse
type OpenpitrixModifyRuntimeResponse struct {

	// runtime
	Runtime *OpenpitrixRuntime `json:"runtime,omitempty"`
}

// Validate validates this openpitrix modify runtime response
func (m *OpenpitrixModifyRuntimeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuntime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixModifyRuntimeResponse) validateRuntime(formats strfmt.Registry) error {

	if swag.IsZero(m.Runtime) { // not required
		return nil
	}

	if m.Runtime != nil {

		if err := m.Runtime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRuntimeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRuntimeResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRuntimeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
