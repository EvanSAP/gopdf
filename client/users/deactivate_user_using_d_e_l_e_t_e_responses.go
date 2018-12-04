// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// DeactivateUserUsingDELETEReader is a Reader for the DeactivateUserUsingDELETE structure.
type DeactivateUserUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeactivateUserUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeactivateUserUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeactivateUserUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeactivateUserUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeactivateUserUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeactivateUserUsingDELETEOK creates a DeactivateUserUsingDELETEOK with default headers values
func NewDeactivateUserUsingDELETEOK() *DeactivateUserUsingDELETEOK {
	return &DeactivateUserUsingDELETEOK{}
}

/*DeactivateUserUsingDELETEOK handles this case with default header values.

OK
*/
type DeactivateUserUsingDELETEOK struct {
}

func (o *DeactivateUserUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}][%d] deactivateUserUsingDELETEOK ", 200)
}

func (o *DeactivateUserUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeactivateUserUsingDELETENoContent creates a DeactivateUserUsingDELETENoContent with default headers values
func NewDeactivateUserUsingDELETENoContent() *DeactivateUserUsingDELETENoContent {
	return &DeactivateUserUsingDELETENoContent{}
}

/*DeactivateUserUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeactivateUserUsingDELETENoContent struct {
}

func (o *DeactivateUserUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}][%d] deactivateUserUsingDELETENoContent ", 204)
}

func (o *DeactivateUserUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeactivateUserUsingDELETEUnauthorized creates a DeactivateUserUsingDELETEUnauthorized with default headers values
func NewDeactivateUserUsingDELETEUnauthorized() *DeactivateUserUsingDELETEUnauthorized {
	return &DeactivateUserUsingDELETEUnauthorized{}
}

/*DeactivateUserUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeactivateUserUsingDELETEUnauthorized struct {
}

func (o *DeactivateUserUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}][%d] deactivateUserUsingDELETEUnauthorized ", 401)
}

func (o *DeactivateUserUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeactivateUserUsingDELETEForbidden creates a DeactivateUserUsingDELETEForbidden with default headers values
func NewDeactivateUserUsingDELETEForbidden() *DeactivateUserUsingDELETEForbidden {
	return &DeactivateUserUsingDELETEForbidden{}
}

/*DeactivateUserUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeactivateUserUsingDELETEForbidden struct {
}

func (o *DeactivateUserUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}][%d] deactivateUserUsingDELETEForbidden ", 403)
}

func (o *DeactivateUserUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
