// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateRuntimeCredentialResponse openpitrix create runtime credential response
// swagger:model openpitrixCreateRuntimeCredentialResponse
type OpenpitrixCreateRuntimeCredentialResponse struct {

	// runtime credential
	RuntimeCredential *OpenpitrixRuntimeCredential `json:"runtime_credential,omitempty"`
}

// Validate validates this openpitrix create runtime credential response
func (m *OpenpitrixCreateRuntimeCredentialResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuntimeCredential(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixCreateRuntimeCredentialResponse) validateRuntimeCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeCredential) { // not required
		return nil
	}

	if m.RuntimeCredential != nil {

		if err := m.RuntimeCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime_credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateRuntimeCredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateRuntimeCredentialResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateRuntimeCredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
