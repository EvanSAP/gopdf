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

// OrderHistoryWsDTO order history ws d t o
// swagger:model OrderHistoryWsDTO
type OrderHistoryWsDTO struct {

	// code
	Code string `json:"code,omitempty"`

	// guid
	GUID string `json:"guid,omitempty"`

	// placed
	// Format: date-time
	Placed strfmt.DateTime `json:"placed,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status display
	StatusDisplay string `json:"statusDisplay,omitempty"`

	// total
	Total *PriceWsDTO `json:"total,omitempty"`
}

// Validate validates this order history ws d t o
func (m *OrderHistoryWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlaced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderHistoryWsDTO) validatePlaced(formats strfmt.Registry) error {

	if swag.IsZero(m.Placed) { // not required
		return nil
	}

	if err := validate.FormatOf("placed", "body", "date-time", m.Placed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrderHistoryWsDTO) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	if m.Total != nil {
		if err := m.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderHistoryWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderHistoryWsDTO) UnmarshalBinary(b []byte) error {
	var res OrderHistoryWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
