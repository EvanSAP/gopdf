// Code generated by go-swagger; DO NOT EDIT.

package customer_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/gopdf/models"
	"golang.org/x/net/context"
)

// NewAssignUserToCustomerGroupUsingPATCH1Params creates a new AssignUserToCustomerGroupUsingPATCH1Params object
// with the default values initialized.
func NewAssignUserToCustomerGroupUsingPATCH1Params() *AssignUserToCustomerGroupUsingPATCH1Params {
	var ()
	return &AssignUserToCustomerGroupUsingPATCH1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAssignUserToCustomerGroupUsingPATCH1ParamsWithTimeout creates a new AssignUserToCustomerGroupUsingPATCH1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAssignUserToCustomerGroupUsingPATCH1ParamsWithTimeout(timeout time.Duration) *AssignUserToCustomerGroupUsingPATCH1Params {
	var ()
	return &AssignUserToCustomerGroupUsingPATCH1Params{

		timeout: timeout,
	}
}

// NewAssignUserToCustomerGroupUsingPATCH1ParamsWithContext creates a new AssignUserToCustomerGroupUsingPATCH1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAssignUserToCustomerGroupUsingPATCH1ParamsWithContext(ctx context.Context) *AssignUserToCustomerGroupUsingPATCH1Params {
	var ()
	return &AssignUserToCustomerGroupUsingPATCH1Params{

		Context: ctx,
	}
}

// NewAssignUserToCustomerGroupUsingPATCH1ParamsWithHTTPClient creates a new AssignUserToCustomerGroupUsingPATCH1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAssignUserToCustomerGroupUsingPATCH1ParamsWithHTTPClient(client *http.Client) *AssignUserToCustomerGroupUsingPATCH1Params {
	var ()
	return &AssignUserToCustomerGroupUsingPATCH1Params{
		HTTPClient: client,
	}
}

/*AssignUserToCustomerGroupUsingPATCH1Params contains all the parameters to send to the API endpoint
for the assign user to customer group using p a t c h 1 operation typically these are written to a http.Request
*/
type AssignUserToCustomerGroupUsingPATCH1Params struct {

	/*BaseSiteID
	  Base site identifier

	*/
	BaseSiteID string
	/*GroupID
	  Group identifier.

	*/
	GroupID string
	/*Members
	  List of users to assign to customer group.

	*/
	Members *models.MemberListWsDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WithTimeout(timeout time.Duration) *AssignUserToCustomerGroupUsingPATCH1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WithContext(ctx context.Context) *AssignUserToCustomerGroupUsingPATCH1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WithHTTPClient(client *http.Client) *AssignUserToCustomerGroupUsingPATCH1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseSiteID adds the baseSiteID to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WithBaseSiteID(baseSiteID string) *AssignUserToCustomerGroupUsingPATCH1Params {
	o.SetBaseSiteID(baseSiteID)
	return o
}

// SetBaseSiteID adds the baseSiteId to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) SetBaseSiteID(baseSiteID string) {
	o.BaseSiteID = baseSiteID
}

// WithGroupID adds the groupID to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WithGroupID(groupID string) *AssignUserToCustomerGroupUsingPATCH1Params {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithMembers adds the members to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WithMembers(members *models.MemberListWsDTO) *AssignUserToCustomerGroupUsingPATCH1Params {
	o.SetMembers(members)
	return o
}

// SetMembers adds the members to the assign user to customer group using p a t c h 1 params
func (o *AssignUserToCustomerGroupUsingPATCH1Params) SetMembers(members *models.MemberListWsDTO) {
	o.Members = members
}

// WriteToRequest writes these params to a swagger request
func (o *AssignUserToCustomerGroupUsingPATCH1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param baseSiteId
	if err := r.SetPathParam("baseSiteId", o.BaseSiteID); err != nil {
		return err
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if o.Members != nil {
		if err := r.SetBodyParam(o.Members); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
