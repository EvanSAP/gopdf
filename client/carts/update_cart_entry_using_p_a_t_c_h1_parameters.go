// Code generated by go-swagger; DO NOT EDIT.

package carts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	models "github.com/gopdf/models"
	"golang.org/x/net/context"
)

// NewUpdateCartEntryUsingPATCH1Params creates a new UpdateCartEntryUsingPATCH1Params object
// with the default values initialized.
func NewUpdateCartEntryUsingPATCH1Params() *UpdateCartEntryUsingPATCH1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &UpdateCartEntryUsingPATCH1Params{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCartEntryUsingPATCH1ParamsWithTimeout creates a new UpdateCartEntryUsingPATCH1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCartEntryUsingPATCH1ParamsWithTimeout(timeout time.Duration) *UpdateCartEntryUsingPATCH1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &UpdateCartEntryUsingPATCH1Params{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewUpdateCartEntryUsingPATCH1ParamsWithContext creates a new UpdateCartEntryUsingPATCH1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCartEntryUsingPATCH1ParamsWithContext(ctx context.Context) *UpdateCartEntryUsingPATCH1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &UpdateCartEntryUsingPATCH1Params{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewUpdateCartEntryUsingPATCH1ParamsWithHTTPClient creates a new UpdateCartEntryUsingPATCH1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCartEntryUsingPATCH1ParamsWithHTTPClient(client *http.Client) *UpdateCartEntryUsingPATCH1Params {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &UpdateCartEntryUsingPATCH1Params{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*UpdateCartEntryUsingPATCH1Params contains all the parameters to send to the API endpoint
for the update cart entry using p a t c h 1 operation typically these are written to a http.Request
*/
type UpdateCartEntryUsingPATCH1Params struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CartID
	  Cart identifier: cart code for logged in user, cart guid for anonymous user, 'current' for the last modified cart

	*/
	CartID string
	/*Entry
	  Request body parameter that contains details such as the quantity of product (quantity), and the pickup store name (deliveryPointOfService.name)

	The DTO is in XML or .json format.

	*/
	Entry *models.OrderEntryWsDTO
	/*EntryNumber
	  The entry number. Each entry in a cart has an entry number. Cart entries are numbered in ascending order, starting with zero (0).

	*/
	EntryNumber int64
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithTimeout(timeout time.Duration) *UpdateCartEntryUsingPATCH1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithContext(ctx context.Context) *UpdateCartEntryUsingPATCH1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithHTTPClient(client *http.Client) *UpdateCartEntryUsingPATCH1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithBaseSiteID(baseSiteID string) *UpdateCartEntryUsingPATCH1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithCartID(cartID string) *UpdateCartEntryUsingPATCH1Params {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithEntry adds the entry to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithEntry(entry *models.OrderEntryWsDTO) *UpdateCartEntryUsingPATCH1Params {
	o.SetEntry(entry)
	return o
}

// SetEntry adds the entry to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetEntry(entry *models.OrderEntryWsDTO) {
	o.Entry = entry
}

// WithEntryNumber adds the entryNumber to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithEntryNumber(entryNumber int64) *UpdateCartEntryUsingPATCH1Params {
	o.SetEntryNumber(entryNumber)
	return o
}

// SetEntryNumber adds the entryNumber to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetEntryNumber(entryNumber int64) {
	o.EntryNumber = entryNumber
}

// WithFields adds the fields to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithFields(fields *string) *UpdateCartEntryUsingPATCH1Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) WithUserID(userID string) *UpdateCartEntryUsingPATCH1Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update cart entry using p a t c h 1 params
func (o *UpdateCartEntryUsingPATCH1Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCartEntryUsingPATCH1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param cartId
	if err := r.SetPathParam("cartId", o.CartID); err != nil {
		return err
	}

	if o.Entry != nil {
		if err := r.SetBodyParam(o.Entry); err != nil {
			return err
		}
	}

	// path param entryNumber
	if err := r.SetPathParam("entryNumber", swag.FormatInt64(o.EntryNumber)); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
