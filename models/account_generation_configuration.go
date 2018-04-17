// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountGenerationConfiguration account generation configuration
// swagger:model AccountGenerationConfiguration
type AccountGenerationConfiguration struct {

	// country
	Country string `json:"country,omitempty"`

	// valid account ranges
	ValidAccountRanges AccountGenerationConfigurationValidAccountRanges `json:"valid_account_ranges"`
}

// Validate validates this account generation configuration
func (m *AccountGenerationConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AccountGenerationConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountGenerationConfiguration) UnmarshalBinary(b []byte) error {
	var res AccountGenerationConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}