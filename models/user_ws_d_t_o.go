// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserWsDTO user ws d t o
// swagger:model UserWsDTO
type UserWsDTO struct {

	// currency
	Currency *CurrencyWsDTO `json:"currency,omitempty"`

	// customer Id
	CustomerID string `json:"customerId,omitempty"`

	// deactivation date
	// Format: date-time
	DeactivationDate CustomDateTime `json:"deactivationDate,omitempty"`

	// default address
	DefaultAddress *AddressWsDTO `json:"defaultAddress,omitempty"`

	// display Uid
	DisplayUID string `json:"displayUid,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// language
	Language *LanguageWsDTO `json:"language,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// title code
	TitleCode string `json:"titleCode,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this user ws d t o
func (m *UserWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeactivationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserWsDTO) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *UserWsDTO) validateDeactivationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeactivationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("deactivationDate", "body", "date-time", m.DeactivationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserWsDTO) validateDefaultAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultAddress) { // not required
		return nil
	}

	if m.DefaultAddress != nil {
		if err := m.DefaultAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defaultAddress")
			}
			return err
		}
	}

	return nil
}

func (m *UserWsDTO) validateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.Language) { // not required
		return nil
	}

	if m.Language != nil {
		if err := m.Language.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserWsDTO) UnmarshalBinary(b []byte) error {
	var res UserWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
