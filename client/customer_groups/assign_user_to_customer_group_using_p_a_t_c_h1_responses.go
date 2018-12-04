// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// AssignUserToCustomerGroupUsingPATCH1Reader is a Reader for the AssignUserToCustomerGroupUsingPATCH1 structure.
type AssignUserToCustomerGroupUsingPATCH1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignUserToCustomerGroupUsingPATCH1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAssignUserToCustomerGroupUsingPATCH1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewAssignUserToCustomerGroupUsingPATCH1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAssignUserToCustomerGroupUsingPATCH1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAssignUserToCustomerGroupUsingPATCH1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAssignUserToCustomerGroupUsingPATCH1OK creates a AssignUserToCustomerGroupUsingPATCH1OK with default headers values
func NewAssignUserToCustomerGroupUsingPATCH1OK() *AssignUserToCustomerGroupUsingPATCH1OK {
	return &AssignUserToCustomerGroupUsingPATCH1OK{}
}

/*AssignUserToCustomerGroupUsingPATCH1OK handles this case with default header values.

OK
*/
type AssignUserToCustomerGroupUsingPATCH1OK struct {
}

func (o *AssignUserToCustomerGroupUsingPATCH1OK) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/customergroups/{groupId}/members][%d] assignUserToCustomerGroupUsingPATCH1OK ", 200)
}

func (o *AssignUserToCustomerGroupUsingPATCH1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignUserToCustomerGroupUsingPATCH1NoContent creates a AssignUserToCustomerGroupUsingPATCH1NoContent with default headers values
func NewAssignUserToCustomerGroupUsingPATCH1NoContent() *AssignUserToCustomerGroupUsingPATCH1NoContent {
	return &AssignUserToCustomerGroupUsingPATCH1NoContent{}
}

/*AssignUserToCustomerGroupUsingPATCH1NoContent handles this case with default header values.

No Content
*/
type AssignUserToCustomerGroupUsingPATCH1NoContent struct {
}

func (o *AssignUserToCustomerGroupUsingPATCH1NoContent) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/customergroups/{groupId}/members][%d] assignUserToCustomerGroupUsingPATCH1NoContent ", 204)
}

func (o *AssignUserToCustomerGroupUsingPATCH1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignUserToCustomerGroupUsingPATCH1Unauthorized creates a AssignUserToCustomerGroupUsingPATCH1Unauthorized with default headers values
func NewAssignUserToCustomerGroupUsingPATCH1Unauthorized() *AssignUserToCustomerGroupUsingPATCH1Unauthorized {
	return &AssignUserToCustomerGroupUsingPATCH1Unauthorized{}
}

/*AssignUserToCustomerGroupUsingPATCH1Unauthorized handles this case with default header values.

Unauthorized
*/
type AssignUserToCustomerGroupUsingPATCH1Unauthorized struct {
}

func (o *AssignUserToCustomerGroupUsingPATCH1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/customergroups/{groupId}/members][%d] assignUserToCustomerGroupUsingPATCH1Unauthorized ", 401)
}

func (o *AssignUserToCustomerGroupUsingPATCH1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignUserToCustomerGroupUsingPATCH1Forbidden creates a AssignUserToCustomerGroupUsingPATCH1Forbidden with default headers values
func NewAssignUserToCustomerGroupUsingPATCH1Forbidden() *AssignUserToCustomerGroupUsingPATCH1Forbidden {
	return &AssignUserToCustomerGroupUsingPATCH1Forbidden{}
}

/*AssignUserToCustomerGroupUsingPATCH1Forbidden handles this case with default header values.

Forbidden
*/
type AssignUserToCustomerGroupUsingPATCH1Forbidden struct {
}

func (o *AssignUserToCustomerGroupUsingPATCH1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /{baseSiteId}/customergroups/{groupId}/members][%d] assignUserToCustomerGroupUsingPATCH1Forbidden ", 403)
}

func (o *AssignUserToCustomerGroupUsingPATCH1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}