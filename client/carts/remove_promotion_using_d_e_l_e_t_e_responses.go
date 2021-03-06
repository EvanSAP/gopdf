// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// RemovePromotionUsingDELETEReader is a Reader for the RemovePromotionUsingDELETE structure.
type RemovePromotionUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemovePromotionUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemovePromotionUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewRemovePromotionUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRemovePromotionUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemovePromotionUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemovePromotionUsingDELETEOK creates a RemovePromotionUsingDELETEOK with default headers values
func NewRemovePromotionUsingDELETEOK() *RemovePromotionUsingDELETEOK {
	return &RemovePromotionUsingDELETEOK{}
}

/*RemovePromotionUsingDELETEOK handles this case with default header values.

OK
*/
type RemovePromotionUsingDELETEOK struct {
}

func (o *RemovePromotionUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] removePromotionUsingDELETEOK ", 200)
}

func (o *RemovePromotionUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemovePromotionUsingDELETENoContent creates a RemovePromotionUsingDELETENoContent with default headers values
func NewRemovePromotionUsingDELETENoContent() *RemovePromotionUsingDELETENoContent {
	return &RemovePromotionUsingDELETENoContent{}
}

/*RemovePromotionUsingDELETENoContent handles this case with default header values.

No Content
*/
type RemovePromotionUsingDELETENoContent struct {
}

func (o *RemovePromotionUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] removePromotionUsingDELETENoContent ", 204)
}

func (o *RemovePromotionUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemovePromotionUsingDELETEUnauthorized creates a RemovePromotionUsingDELETEUnauthorized with default headers values
func NewRemovePromotionUsingDELETEUnauthorized() *RemovePromotionUsingDELETEUnauthorized {
	return &RemovePromotionUsingDELETEUnauthorized{}
}

/*RemovePromotionUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type RemovePromotionUsingDELETEUnauthorized struct {
}

func (o *RemovePromotionUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] removePromotionUsingDELETEUnauthorized ", 401)
}

func (o *RemovePromotionUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemovePromotionUsingDELETEForbidden creates a RemovePromotionUsingDELETEForbidden with default headers values
func NewRemovePromotionUsingDELETEForbidden() *RemovePromotionUsingDELETEForbidden {
	return &RemovePromotionUsingDELETEForbidden{}
}

/*RemovePromotionUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type RemovePromotionUsingDELETEForbidden struct {
}

func (o *RemovePromotionUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /{baseSiteId}/users/{userId}/carts/{cartId}/promotions/{promotionId}][%d] removePromotionUsingDELETEForbidden ", 403)
}

func (o *RemovePromotionUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
