// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAddMemberSetup account add member setup
//
// swagger:model accountAddMemberSetup
type AccountAddMemberSetup struct {

	// email
	Email string `json:"email,omitempty"`

	// role
	// Enum: [Owner Collaborator Controller]
	Role string `json:"role,omitempty"`
}

// Validate validates this account add member setup
func (m *AccountAddMemberSetup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountAddMemberSetupTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Owner","Collaborator","Controller"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountAddMemberSetupTypeRolePropEnum = append(accountAddMemberSetupTypeRolePropEnum, v)
	}
}

const (

	// AccountAddMemberSetupRoleOwner captures enum value "Owner"
	AccountAddMemberSetupRoleOwner string = "Owner"

	// AccountAddMemberSetupRoleCollaborator captures enum value "Collaborator"
	AccountAddMemberSetupRoleCollaborator string = "Collaborator"

	// AccountAddMemberSetupRoleController captures enum value "Controller"
	AccountAddMemberSetupRoleController string = "Controller"
)

// prop value enum
func (m *AccountAddMemberSetup) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountAddMemberSetupTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountAddMemberSetup) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAddMemberSetup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAddMemberSetup) UnmarshalBinary(b []byte) error {
	var res AccountAddMemberSetup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
