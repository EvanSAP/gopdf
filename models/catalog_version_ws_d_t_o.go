// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogVersionWsDTO catalog version ws d t o
// swagger:model CatalogVersionWsDTO
type CatalogVersionWsDTO struct {

	// categories
	Categories []*CategoryHierarchyWsDTO `json:"categories"`

	// id
	ID string `json:"id,omitempty"`

	// last modified
	// Format: date-time
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this catalog version ws d t o
func (m *CatalogVersionWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogVersionWsDTO) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {
		if swag.IsZero(m.Categories[i]) { // not required
			continue
		}

		if m.Categories[i] != nil {
			if err := m.Categories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogVersionWsDTO) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogVersionWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogVersionWsDTO) UnmarshalBinary(b []byte) error {
	var res CatalogVersionWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
