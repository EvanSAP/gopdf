// Code generated by go-swagger; DO NOT EDIT.

package save_cart

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
)

// NewGetSavedCartUsingGETParams creates a new GetSavedCartUsingGETParams object
// with the default values initialized.
func NewGetSavedCartUsingGETParams() *GetSavedCartUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSavedCartUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSavedCartUsingGETParamsWithTimeout creates a new GetSavedCartUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSavedCartUsingGETParamsWithTimeout(timeout time.Duration) *GetSavedCartUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSavedCartUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetSavedCartUsingGETParamsWithContext creates a new GetSavedCartUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSavedCartUsingGETParamsWithContext(ctx context.Context) *GetSavedCartUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSavedCartUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetSavedCartUsingGETParamsWithHTTPClient creates a new GetSavedCartUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSavedCartUsingGETParamsWithHTTPClient(client *http.Client) *GetSavedCartUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetSavedCartUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetSavedCartUsingGETParams contains all the parameters to send to the API endpoint
for the get saved cart using g e t operation typically these are written to a http.Request
*/
type GetSavedCartUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*CartID
	  Cart identifier: cart code for logged in user, cart guid for anonymous user, 'current' for the last modified cart

	*/
	CartID string
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

// WithTimeout adds the timeout to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithTimeout(timeout time.Duration) *GetSavedCartUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithContext(ctx context.Context) *GetSavedCartUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithHTTPClient(client *http.Client) *GetSavedCartUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithBaseSiteID(baseSiteID string) *GetSavedCartUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithCartID adds the cartID to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithCartID(cartID string) *GetSavedCartUsingGETParams {
	o.SetCartID(cartID)
	return o
}

// SetCartID adds the cartId to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetCartID(cartID string) {
	o.CartID = cartID
}

// WithFields adds the fields to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithFields(fields *string) *GetSavedCartUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) WithUserID(userID string) *GetSavedCartUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get saved cart using g e t params
func (o *GetSavedCartUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSavedCartUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
