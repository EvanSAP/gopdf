// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCartUsingDELETEReader is a Reader for the DeleteCartUsingDELETE structure.
type DeleteCartUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCartUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCartUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewDeleteCartUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteCartUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteCartUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCartUsingDELETEOK creates a DeleteCartUsingDELETEOK with default headers values
func NewDeleteCartUsingDELETEOK() *DeleteCartUsingDELETEOK {
	return &DeleteCartUsingDELETEOK{}
}

/*DeleteCartUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteCartUsingDELETEOK struct {
}

func (o *DeleteCartUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}][%d] deleteCartUsingDELETEOK ", 200)
}

func (o *DeleteCartUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCartUsingDELETENoContent creates a DeleteCartUsingDELETENoContent with default headers values
func NewDeleteCartUsingDELETENoContent() *DeleteCartUsingDELETENoContent {
	return &DeleteCartUsingDELETENoContent{}
}

/*DeleteCartUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteCartUsingDELETENoContent struct {
}

func (o *DeleteCartUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}][%d] deleteCartUsingDELETENoContent ", 204)
}

func (o *DeleteCartUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCartUsingDELETEUnauthorized creates a DeleteCartUsingDELETEUnauthorized with default headers values
func NewDeleteCartUsingDELETEUnauthorized() *DeleteCartUsingDELETEUnauthorized {
	return &DeleteCartUsingDELETEUnauthorized{}
}

/*DeleteCartUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCartUsingDELETEUnauthorized struct {
}

func (o *DeleteCartUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}][%d] deleteCartUsingDELETEUnauthorized ", 401)
}

func (o *DeleteCartUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCartUsingDELETEForbidden creates a DeleteCartUsingDELETEForbidden with default headers values
func NewDeleteCartUsingDELETEForbidden() *DeleteCartUsingDELETEForbidden {
	return &DeleteCartUsingDELETEForbidden{}
}

/*DeleteCartUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteCartUsingDELETEForbidden struct {
}

func (o *DeleteCartUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}][%d] deleteCartUsingDELETEForbidden ", 403)
}

func (o *DeleteCartUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
