// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserAttributes user attributes
// swagger:model userAttributes
type UserAttributes struct {

	// email
	// Pattern: ^[A-Za-z0-9-+@.]*$
	Email string `json:"email,omitempty"`

	// role ids
	RoleIds []strfmt.UUID `json:"role_ids"`

	// username
	// Pattern: ^[A-Za-z0-9-+@.]*$
	Username string `json:"username,omitempty"`
}

// Validate validates this user attributes
func (m *UserAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoleIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAttributes) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(m.Email), `^[A-Za-z0-9-+@.]*$`); err != nil {
		return err
	}

	return nil
}

func (m *UserAttributes) validateRoleIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleIds) { // not required
		return nil
	}

	return nil
}

func (m *UserAttributes) validateUsername(formats strfmt.Registry) error {

	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.Pattern("username", "body", string(m.Username), `^[A-Za-z0-9-+@.]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAttributes) UnmarshalBinary(b []byte) error {
	var res UserAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
