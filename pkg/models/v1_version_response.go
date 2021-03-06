// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VersionResponse v1 version response
//
// swagger:model v1.VersionResponse
type V1VersionResponse struct {

	// build timestamp
	BuildTimestamp string `json:"build_timestamp,omitempty"`

	// commit hash
	CommitHash string `json:"commit_hash,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 version response
func (m *V1VersionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 version response based on context it is used
func (m *V1VersionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1VersionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VersionResponse) UnmarshalBinary(b []byte) error {
	var res V1VersionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
