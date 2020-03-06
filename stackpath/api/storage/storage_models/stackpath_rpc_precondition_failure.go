// Code generated by go-swagger; DO NOT EDIT.

package storage_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCPreconditionFailure stackpath rpc precondition failure
// swagger:model stackpath.rpc.PreconditionFailure
type StackpathRPCPreconditionFailure struct {

	// violations
	Violations []*StackpathRPCPreconditionFailureViolation `json:"violations"`
}

// AtType gets the at type of this subtype
func (m *StackpathRPCPreconditionFailure) AtType() string {
	return "stackpath.rpc.PreconditionFailure"
}

// SetAtType sets the at type of this subtype
func (m *StackpathRPCPreconditionFailure) SetAtType(val string) {

}

// Violations gets the violations of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StackpathRPCPreconditionFailure) UnmarshalJSON(raw []byte) error {
	var data struct {

		// violations
		Violations []*StackpathRPCPreconditionFailureViolation `json:"violations"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AtType string `json:"@type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StackpathRPCPreconditionFailure

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}

	result.Violations = data.Violations

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StackpathRPCPreconditionFailure) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// violations
		Violations []*StackpathRPCPreconditionFailureViolation `json:"violations"`
	}{

		Violations: m.Violations,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AtType string `json:"@type"`
	}{

		AtType: m.AtType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this stackpath rpc precondition failure
func (m *StackpathRPCPreconditionFailure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateViolations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackpathRPCPreconditionFailure) validateViolations(formats strfmt.Registry) error {

	if swag.IsZero(m.Violations) { // not required
		return nil
	}

	for i := 0; i < len(m.Violations); i++ {
		if swag.IsZero(m.Violations[i]) { // not required
			continue
		}

		if m.Violations[i] != nil {
			if err := m.Violations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("violations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCPreconditionFailure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCPreconditionFailure) UnmarshalBinary(b []byte) error {
	var res StackpathRPCPreconditionFailure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
