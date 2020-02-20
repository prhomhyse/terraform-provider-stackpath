// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1NetworkPolicySpec The specification for the desired state of a network policy
// swagger:model v1NetworkPolicySpec
type V1NetworkPolicySpec struct {

	// A list of rules for outbound traffic from instances
	//
	// If the egress policy type is given but not defined here the default is to allow all
	Egress []*V1Egress `json:"egress"`

	// A list of rules for inbound traffic to instances
	//
	// If the ingress policy type is given but not defined here the default is to block all
	Ingress []*V1Ingress `json:"ingress"`

	// A selector to match workload instances
	InstanceSelectors []*V1MatchExpression `json:"instanceSelectors"`

	// A selector to match networks
	NetworkSelectors []*V1MatchExpression `json:"networkSelectors"`

	// A list of policy types
	//
	// Policy types are used to specify what rules will be defined. If a policy type is given but not defined it will default. If it is not provided then no action will be used.
	PolicyTypes []NetworkPolicySpecPolicyType `json:"policyTypes"`

	// A priority for all rules in the network policy. 1-65000
	Priority int32 `json:"priority,omitempty"`
}

// Validate validates this v1 network policy spec
func (m *V1NetworkPolicySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1NetworkPolicySpec) validateEgress(formats strfmt.Registry) error {

	if swag.IsZero(m.Egress) { // not required
		return nil
	}

	for i := 0; i < len(m.Egress); i++ {
		if swag.IsZero(m.Egress[i]) { // not required
			continue
		}

		if m.Egress[i] != nil {
			if err := m.Egress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("egress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NetworkPolicySpec) validateIngress(formats strfmt.Registry) error {

	if swag.IsZero(m.Ingress) { // not required
		return nil
	}

	for i := 0; i < len(m.Ingress); i++ {
		if swag.IsZero(m.Ingress[i]) { // not required
			continue
		}

		if m.Ingress[i] != nil {
			if err := m.Ingress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NetworkPolicySpec) validateInstanceSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceSelectors); i++ {
		if swag.IsZero(m.InstanceSelectors[i]) { // not required
			continue
		}

		if m.InstanceSelectors[i] != nil {
			if err := m.InstanceSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NetworkPolicySpec) validateNetworkSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkSelectors); i++ {
		if swag.IsZero(m.NetworkSelectors[i]) { // not required
			continue
		}

		if m.NetworkSelectors[i] != nil {
			if err := m.NetworkSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkSelectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1NetworkPolicySpec) validatePolicyTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.PolicyTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyTypes); i++ {

		if err := m.PolicyTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NetworkPolicySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NetworkPolicySpec) UnmarshalBinary(b []byte) error {
	var res V1NetworkPolicySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}