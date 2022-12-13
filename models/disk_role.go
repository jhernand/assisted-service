// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DiskRole disk role
//
// swagger:model disk-role
type DiskRole string

func NewDiskRole(value DiskRole) *DiskRole {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DiskRole.
func (m DiskRole) Pointer() *DiskRole {
	return &m
}

const (

	// DiskRoleNone captures enum value "none"
	DiskRoleNone DiskRole = "none"

	// DiskRoleInstall captures enum value "install"
	DiskRoleInstall DiskRole = "install"
)

// for schema
var diskRoleEnum []interface{}

func init() {
	var res []DiskRole
	if err := json.Unmarshal([]byte(`["none","install"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diskRoleEnum = append(diskRoleEnum, v)
	}
}

func (m DiskRole) validateDiskRoleEnum(path, location string, value DiskRole) error {
	if err := validate.EnumCase(path, location, value, diskRoleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this disk role
func (m DiskRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDiskRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this disk role based on context it is used
func (m DiskRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
