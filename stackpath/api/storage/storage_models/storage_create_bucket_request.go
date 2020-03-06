// Code generated by go-swagger; DO NOT EDIT.

package storage_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageCreateBucketRequest storage create bucket request
// swagger:model storageCreateBucketRequest
type StorageCreateBucketRequest struct {

	// The name of the bucket to be created
	Label string `json:"label,omitempty"`

	// The region where to create the bucket, defaults to us-east-1
	Region string `json:"region,omitempty"`
}

// Validate validates this storage create bucket request
func (m *StorageCreateBucketRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageCreateBucketRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageCreateBucketRequest) UnmarshalBinary(b []byte) error {
	var res StorageCreateBucketRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
