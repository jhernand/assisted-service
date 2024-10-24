// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIVip The virtual IP used to reach the OpenShift cluster's API.
//
// swagger:model api_vip
type APIVip struct {

	// The cluster that this VIP is associated with.
	// Format: uuid
	ClusterID strfmt.UUID `json:"cluster_id,omitempty" gorm:"primaryKey"`

	// The IP address.
	IP IP `json:"ip,omitempty" gorm:"primaryKey"`
}

// Validate validates this api vip
func (m *APIVip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIVip) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster_id", "body", "uuid", m.ClusterID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIVip) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if err := m.IP.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ip")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ip")
		}
		return err
	}

	return nil
}

// ContextValidate validate this api vip based on the context it is used
func (m *APIVip) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIVip) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IP.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ip")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ip")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIVip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIVip) UnmarshalBinary(b []byte) error {
	var res APIVip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
