// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkMetadata Metadata associated with an entity
// swagger:model networkMetadata
type NetworkMetadata struct {

	// annotations
	Annotations NetworkStringMapEntry `json:"annotations,omitempty"`

	// The date that a metadata entry was created
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// The date that a network policy was requested for deletion
	// Read Only: true
	// Format: date-time
	DeleteRequestedAt strfmt.DateTime `json:"deleteRequestedAt,omitempty"`

	// labels
	Labels NetworkStringMapEntry `json:"labels,omitempty"`

	// The date that a metadata entry was last updated
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// An entity's version number
	//
	// Versions start at 1 when they are created and increment by 1 every time they are updated.
	Version string `json:"version,omitempty"`
}

// Validate validates this network metadata
func (m *NetworkMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkMetadata) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	if err := m.Annotations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("annotations")
		}
		return err
	}

	return nil
}

func (m *NetworkMetadata) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkMetadata) validateDeleteRequestedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.DeleteRequestedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deleteRequestedAt", "body", "date-time", m.DeleteRequestedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkMetadata) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *NetworkMetadata) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkMetadata) UnmarshalBinary(b []byte) error {
	var res NetworkMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
