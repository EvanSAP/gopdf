// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserSignUpWsDTO user sign up ws d t o
// swagger:model UserSignUpWsDTO
type UserSignUpWsDTO struct {

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// title code
	TitleCode string `json:"titleCode,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this user sign up ws d t o
func (m *UserSignUpWsDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserSignUpWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSignUpWsDTO) UnmarshalBinary(b []byte) error {
	var res UserSignUpWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
