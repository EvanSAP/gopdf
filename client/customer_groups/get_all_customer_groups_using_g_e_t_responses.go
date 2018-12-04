// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// GetAllCustomerGroupsUsingGETReader is a Reader for the GetAllCustomerGroupsUsingGET structure.
type GetAllCustomerGroupsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllCustomerGroupsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllCustomerGroupsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAllCustomerGroupsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllCustomerGroupsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAllCustomerGroupsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllCustomerGroupsUsingGETOK creates a GetAllCustomerGroupsUsingGETOK with default headers values
func NewGetAllCustomerGroupsUsingGETOK() *GetAllCustomerGroupsUsingGETOK {
	return &GetAllCustomerGroupsUsingGETOK{}
}

/*GetAllCustomerGroupsUsingGETOK handles this case with default header values.

OK
*/
type GetAllCustomerGroupsUsingGETOK struct {
	Payload *models.UserGroupListWsDTO
}

func (o *GetAllCustomerGroupsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/customergroups][%d] getAllCustomerGroupsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllCustomerGroupsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupListWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllCustomerGroupsUsingGETUnauthorized creates a GetAllCustomerGroupsUsingGETUnauthorized with default headers values
func NewGetAllCustomerGroupsUsingGETUnauthorized() *GetAllCustomerGroupsUsingGETUnauthorized {
	return &GetAllCustomerGroupsUsingGETUnauthorized{}
}

/*GetAllCustomerGroupsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllCustomerGroupsUsingGETUnauthorized struct {
}

func (o *GetAllCustomerGroupsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/customergroups][%d] getAllCustomerGroupsUsingGETUnauthorized ", 401)
}

func (o *GetAllCustomerGroupsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCustomerGroupsUsingGETForbidden creates a GetAllCustomerGroupsUsingGETForbidden with default headers values
func NewGetAllCustomerGroupsUsingGETForbidden() *GetAllCustomerGroupsUsingGETForbidden {
	return &GetAllCustomerGroupsUsingGETForbidden{}
}

/*GetAllCustomerGroupsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllCustomerGroupsUsingGETForbidden struct {
}

func (o *GetAllCustomerGroupsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/customergroups][%d] getAllCustomerGroupsUsingGETForbidden ", 403)
}

func (o *GetAllCustomerGroupsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCustomerGroupsUsingGETNotFound creates a GetAllCustomerGroupsUsingGETNotFound with default headers values
func NewGetAllCustomerGroupsUsingGETNotFound() *GetAllCustomerGroupsUsingGETNotFound {
	return &GetAllCustomerGroupsUsingGETNotFound{}
}

/*GetAllCustomerGroupsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllCustomerGroupsUsingGETNotFound struct {
}

func (o *GetAllCustomerGroupsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /{baseSiteId}/customergroups][%d] getAllCustomerGroupsUsingGETNotFound ", 404)
}

func (o *GetAllCustomerGroupsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
