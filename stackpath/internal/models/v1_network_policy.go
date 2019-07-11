// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1NetworkPolicy A network policy
//
// Network policies define an ACL that applies to one or many workload instances
// swagger:model v1NetworkPolicy
type V1NetworkPolicy struct {

	// Detailed summary of what the policy does
	Description string `json:"description,omitempty"`

	// A network policy's unique identifier
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *V1Metadata `json:"metadata,omitempty"`

	// A network policy's name as it appears in the StackPath portal
	Name string `json:"name,omitempty"`

	// A network policy's programmatic name
	//
	// Network policy slugs are used to build its instances names
	Slug string `json:"slug,omitempty"`

	// spec
	Spec *V1NetworkPolicySpec `json:"spec,omitempty"`

	// The ID of the stack that a network policy belongs to
	StackID string `json:"stackId,omitempty"`
}

// Validate validates this v1 network policy
func (m *V1NetworkPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkPolicy) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *V1NetworkPolicy) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkPolicy) UnmarshalBinary(b []byte) error {
	var res V1NetworkPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}