// Code generated by go-swagger; DO NOT EDIT.

package payment_details

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

// NewGetPaymentDetailsUsingGETParams creates a new GetPaymentDetailsUsingGETParams object
// with the default values initialized.
func NewGetPaymentDetailsUsingGETParams() *GetPaymentDetailsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetPaymentDetailsUsingGETParams{
		Fields: &fieldsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentDetailsUsingGETParamsWithTimeout creates a new GetPaymentDetailsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentDetailsUsingGETParamsWithTimeout(timeout time.Duration) *GetPaymentDetailsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetPaymentDetailsUsingGETParams{
		Fields: &fieldsDefault,

		timeout: timeout,
	}
}

// NewGetPaymentDetailsUsingGETParamsWithContext creates a new GetPaymentDetailsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentDetailsUsingGETParamsWithContext(ctx context.Context) *GetPaymentDetailsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetPaymentDetailsUsingGETParams{
		Fields: &fieldsDefault,

		Context: ctx,
	}
}

// NewGetPaymentDetailsUsingGETParamsWithHTTPClient creates a new GetPaymentDetailsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentDetailsUsingGETParamsWithHTTPClient(client *http.Client) *GetPaymentDetailsUsingGETParams {
	var (
		fieldsDefault = string("DEFAULT")
	)
	return &GetPaymentDetailsUsingGETParams{
		Fields:     &fieldsDefault,
		HTTPClient: client,
	}
}

/*GetPaymentDetailsUsingGETParams contains all the parameters to send to the API endpoint
for the get payment details using g e t operation typically these are written to a http.Request
*/
type GetPaymentDetailsUsingGETParams struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*Fields
	  Response configuration. This is the list of fields that should be returned in the response body.

	*/
	Fields *string
	/*PaymentDetailsID
	  Payment details identifier.

	*/
	PaymentDetailsID string
	/*UserID
	  User identifier or one of the literals : 'current' for currently authenticated user, 'anonymous' for anonymous user

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithTimeout(timeout time.Duration) *GetPaymentDetailsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithContext(ctx context.Context) *GetPaymentDetailsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithHTTPClient(client *http.Client) *GetPaymentDetailsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithBaseSiteID(baseSiteID string) *GetPaymentDetailsUsingGETParams {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithFields adds the fields to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithFields(fields *string) *GetPaymentDetailsUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPaymentDetailsID adds the paymentDetailsID to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithPaymentDetailsID(paymentDetailsID string) *GetPaymentDetailsUsingGETParams {
	o.SetPaymentDetailsID(paymentDetailsID)
	return o
}

// SetPaymentDetailsID adds the paymentDetailsId to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetPaymentDetailsID(paymentDetailsID string) {
	o.PaymentDetailsID = paymentDetailsID
}

// WithUserID adds the userID to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) WithUserID(userID string) *GetPaymentDetailsUsingGETParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get payment details using g e t params
func (o *GetPaymentDetailsUsingGETParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentDetailsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
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

	// path param paymentDetailsId
	if err := r.SetPathParam("paymentDetailsId", o.PaymentDetailsID); err != nil {
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
