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

// CatalogWsDTO catalog ws d t o
// swagger:model CatalogWsDTO
type CatalogWsDTO struct {

	// catalog versions
	CatalogVersions []*CatalogVersionWsDTO `json:"catalogVersions"`

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

// Validate validates this catalog ws d t o
func (m *CatalogWsDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogVersions(formats); err != nil {
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

func (m *CatalogWsDTO) validateCatalogVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.CatalogVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.CatalogVersions); i++ {
		if swag.IsZero(m.CatalogVersions[i]) { // not required
			continue
		}

		if m.CatalogVersions[i] != nil {
			if err := m.CatalogVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalogVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CatalogWsDTO) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogWsDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogWsDTO) UnmarshalBinary(b []byte) error {
	var res CatalogWsDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
