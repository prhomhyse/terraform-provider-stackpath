// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1VolumeClaim A claim for a volume
//
// Volumes may be attached to workload instance containers or virtual machines
// swagger:model v1VolumeClaim
type V1VolumeClaim struct {

	// A volume claim's unique identifier
	// Read Only: true
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *V1Metadata `json:"metadata,omitempty"`

	// A volume claim's name as it appears in the StackPath portal
	Name string `json:"name,omitempty"`

	// phase
	// Read Only: true
	Phase VolumeClaimVolumeClaimPhase `json:"phase,omitempty"`

	// A volume claim's programmatic name
	//
	// Volume claim slugs are used to programatically label a claim
	Slug string `json:"slug,omitempty"`

	// spec
	Spec *V1VolumeClaimSpec `json:"spec,omitempty"`

	// The ID of the stack that a volume claim belongs to
	// Read Only: true
	StackID string `json:"stackId,omitempty"`
}

// Validate validates this v1 volume claim
func (m *V1VolumeClaim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
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

func (m *V1VolumeClaim) validateMetadata(formats strfmt.Registry) error {

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

func (m *V1VolumeClaim) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if err := m.Phase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phase")
		}
		return err
	}

	return nil
}

func (m *V1VolumeClaim) validateSpec(formats strfmt.Registry) error {

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
func (m *V1VolumeClaim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VolumeClaim) UnmarshalBinary(b []byte) error {
	var res V1VolumeClaim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
