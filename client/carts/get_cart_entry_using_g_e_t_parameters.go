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
	"golang.org/x/net/context"
)

// NewGetCartEntryUsingGETParams creates a new GetCartEntryUsingGETParams object
// with the default values initialized.
func NewGetCartEntryUsingGETParams() *GetCartEntryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCartEntryUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCartEntryUsingGETParamsWithTimeout creates a new GetCartEntryUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCartEntryUsingGETParamsWithTimeout(timeout time.Duration) *GetCartEntryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCartEntryUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetCartEntryUsingGETParamsWithContext creates a new GetCartEntryUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCartEntryUsingGETParamsWithContext(ctx context.Context) *GetCartEntryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCartEntryUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetCartEntryUsingGETParamsWithHTTPClient creates a new GetCartEntryUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCartEntryUsingGETParamsWithHTTPClient(client *http.Client) *GetCartEntryUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetCartEntryUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetCartEntryUsingGETParams contains all the parameters to send to the API endpoint
for the get cart entry using g e t operation typically these are written to a http.Request
*/
type GetCartEntryUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CartID
	  Cart identifier: cart code for logged in user, cart guid for anonymous user, 'current' for the last modified cart

	*/
	CartID string
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

// WithTimeout adds the timeout to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithTimeout(timeout time.Duration) *GetCartEntryUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithContext(ctx context.Context) *GetCartEntryUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithHTTPClient(client *http.Client) *GetCartEntryUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithBaseSiteID(baseSiteID string) *GetCartEntryUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithCartID(cartID string) *GetCartEntryUsingGETParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithEntryNumber adds the entryNumber to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithEntryNumber(entryNumber int64) *GetCartEntryUsingGETParams {
	o.SetEntryNumber(entryNumber)
	return o
}

// SetEntryNumber adds the entryNumber to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetEntryNumber(entryNumber int64) {
	o.EntryNumber = entryNumber
}

// WithFields adds the fields to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithFields(fields *string) *GetCartEntryUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) WithUserID(userID string) *GetCartEntryUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get cart entry using g e t params
func (o *GetCartEntryUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCartEntryUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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