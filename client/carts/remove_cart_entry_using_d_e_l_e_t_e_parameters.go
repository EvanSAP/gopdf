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

// NewRemoveCartEntryUsingDELETEParams creates a new RemoveCartEntryUsingDELETEParams object
// with the default values initialized.
func NewRemoveCartEntryUsingDELETEParams() *RemoveCartEntryUsingDELETEParams {
	var ()
	return &RemoveCartEntryUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveCartEntryUsingDELETEParamsWithTimeout creates a new RemoveCartEntryUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveCartEntryUsingDELETEParamsWithTimeout(timeout time.Duration) *RemoveCartEntryUsingDELETEParams {
	var ()
	return &RemoveCartEntryUsingDELETEParams{

		timeout: timeout,
	}
}

// NewRemoveCartEntryUsingDELETEParamsWithContext creates a new RemoveCartEntryUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveCartEntryUsingDELETEParamsWithContext(ctx context.Context) *RemoveCartEntryUsingDELETEParams {
	var ()
	return &RemoveCartEntryUsingDELETEParams{

		Context: ctx,
	}
}

// NewRemoveCartEntryUsingDELETEParamsWithHTTPClient creates a new RemoveCartEntryUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveCartEntryUsingDELETEParamsWithHTTPClient(client *http.Client) *RemoveCartEntryUsingDELETEParams {
	var ()
	return &RemoveCartEntryUsingDELETEParams{
		HTTPClient: client,
	}
}

/*RemoveCartEntryUsingDELETEParams contains all the parameters to send to the API endpoint
for the remove cart entry using d e l e t e operation typically these are written to a http.Request
*/
type RemoveCartEntryUsingDELETEParams struct {

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
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithTimeout(timeout time.Duration) *RemoveCartEntryUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithContext(ctx context.Context) *RemoveCartEntryUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithHTTPClient(client *http.Client) *RemoveCartEntryUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithBaseSiteID(baseSiteID string) *RemoveCartEntryUsingDELETEParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithCartID(cartID string) *RemoveCartEntryUsingDELETEParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithEntryNumber adds the entryNumber to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithEntryNumber(entryNumber int64) *RemoveCartEntryUsingDELETEParams {
	o.SetEntryNumber(entryNumber)
	return o
}

// SetEntryNumber adds the entryNumber to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetEntryNumber(entryNumber int64) {
	o.EntryNumber = entryNumber
}

// WithUserID adds the userID to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) WithUserID(userID string) *RemoveCartEntryUsingDELETEParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the remove cart entry using d e l e t e params
func (o *RemoveCartEntryUsingDELETEParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveCartEntryUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
