// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
)

// SetCartEntryUsingPUT1Reader is a Reader for the SetCartEntryUsingPUT1 structure.
type SetCartEntryUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCartEntryUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetCartEntryUsingPUT1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewSetCartEntryUsingPUT1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSetCartEntryUsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSetCartEntryUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSetCartEntryUsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetCartEntryUsingPUT1OK creates a SetCartEntryUsingPUT1OK with default headers values
func NewSetCartEntryUsingPUT1OK() *SetCartEntryUsingPUT1OK {
	return &SetCartEntryUsingPUT1OK{}
}

/*SetCartEntryUsingPUT1OK handles this case with default header values.

OK
*/
type SetCartEntryUsingPUT1OK struct {
	Payload *models.CartModificationWsDTO
}

func (o *SetCartEntryUsingPUT1OK) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] setCartEntryUsingPUT1OK  %+v", 200, o.Payload)
}

func (o *SetCartEntryUsingPUT1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CartModificationWsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetCartEntryUsingPUT1Created creates a SetCartEntryUsingPUT1Created with default headers values
func NewSetCartEntryUsingPUT1Created() *SetCartEntryUsingPUT1Created {
	return &SetCartEntryUsingPUT1Created{}
}

/*SetCartEntryUsingPUT1Created handles this case with default header values.

Created
*/
type SetCartEntryUsingPUT1Created struct {
}

func (o *SetCartEntryUsingPUT1Created) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] setCartEntryUsingPUT1Created ", 201)
}

func (o *SetCartEntryUsingPUT1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCartEntryUsingPUT1Unauthorized creates a SetCartEntryUsingPUT1Unauthorized with default headers values
func NewSetCartEntryUsingPUT1Unauthorized() *SetCartEntryUsingPUT1Unauthorized {
	return &SetCartEntryUsingPUT1Unauthorized{}
}

/*SetCartEntryUsingPUT1Unauthorized handles this case with default header values.

Unauthorized
*/
type SetCartEntryUsingPUT1Unauthorized struct {
}

func (o *SetCartEntryUsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] setCartEntryUsingPUT1Unauthorized ", 401)
}

func (o *SetCartEntryUsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCartEntryUsingPUT1Forbidden creates a SetCartEntryUsingPUT1Forbidden with default headers values
func NewSetCartEntryUsingPUT1Forbidden() *SetCartEntryUsingPUT1Forbidden {
	return &SetCartEntryUsingPUT1Forbidden{}
}

/*SetCartEntryUsingPUT1Forbidden handles this case with default header values.

Forbidden
*/
type SetCartEntryUsingPUT1Forbidden struct {
}

func (o *SetCartEntryUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] setCartEntryUsingPUT1Forbidden ", 403)
}

func (o *SetCartEntryUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetCartEntryUsingPUT1NotFound creates a SetCartEntryUsingPUT1NotFound with default headers values
func NewSetCartEntryUsingPUT1NotFound() *SetCartEntryUsingPUT1NotFound {
	return &SetCartEntryUsingPUT1NotFound{}
}

/*SetCartEntryUsingPUT1NotFound handles this case with default header values.

Not Found
*/
type SetCartEntryUsingPUT1NotFound struct {
}

func (o *SetCartEntryUsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /{baseSiteId}/users/{userId}/carts/{cartId}/entries/{entryNumber}][%d] setCartEntryUsingPUT1NotFound ", 404)
}

func (o *SetCartEntryUsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
