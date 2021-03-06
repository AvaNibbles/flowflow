// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApimodelsHTTPError apimodels Http error
//
// swagger:model apimodels.HttpError
type ApimodelsHTTPError struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this apimodels Http error
func (m *ApimodelsHTTPError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apimodels Http error based on context it is used
func (m *ApimodelsHTTPError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApimodelsHTTPError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApimodelsHTTPError) UnmarshalBinary(b []byte) error {
	var res ApimodelsHTTPError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
