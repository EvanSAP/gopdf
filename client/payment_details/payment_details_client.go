// Code generated by go-swagger; DO NOT EDIT.

package payment_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payment details API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment details API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePaymentInfoUsingDELETE deletes customer s credit card payment details

Removes a customer's credit card payment details based on a specified paymentDetailsId.
*/
func (a *Client) DeletePaymentInfoUsingDELETE(params *DeletePaymentInfoUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePaymentInfoUsingDELETEOK, *DeletePaymentInfoUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentInfoUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePaymentInfoUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/{baseSiteId}/users/{userId}/paymentdetails/{paymentDetailsId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePaymentInfoUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeletePaymentInfoUsingDELETEOK:
		return value, nil, nil
	case *DeletePaymentInfoUsingDELETENoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetPaymentDetailsUsingGET gets customer s credit card payment details

Returns a customer's credit card payment details for the specified paymentDetailsId.
*/
func (a *Client) GetPaymentDetailsUsingGET(params *GetPaymentDetailsUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentDetailsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentDetailsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentDetailsUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/users/{userId}/paymentdetails/{paymentDetailsId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentDetailsUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentDetailsUsingGETOK), nil

}

/*
GetPaymentInfosUsingGET gets customer s credit card payment details list

Return customer's credit card payment details list.
*/
func (a *Client) GetPaymentInfosUsingGET(params *GetPaymentInfosUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentInfosUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentInfosUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentInfosUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/users/{userId}/paymentdetails",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentInfosUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentInfosUsingGETOK), nil

}

/*
PutPaymentInfoUsingPUT1 updates existing customer s credit card payment info

Updates existing customer's credit card payment info based on the payment info ID. Attributes not given in request will be defined again (set to null or default).
*/
func (a *Client) PutPaymentInfoUsingPUT1(params *PutPaymentInfoUsingPUT1Params, authInfo runtime.ClientAuthInfoWriter) (*PutPaymentInfoUsingPUT1OK, *PutPaymentInfoUsingPUT1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutPaymentInfoUsingPUT1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putPaymentInfoUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/{baseSiteId}/users/{userId}/paymentdetails/{paymentDetailsId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutPaymentInfoUsingPUT1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutPaymentInfoUsingPUT1OK:
		return value, nil, nil
	case *PutPaymentInfoUsingPUT1Created:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdatePaymentInfoUsingPATCH1 updates existing customer s credit card payment details

Updates an existing customer's credit card payment details based on the specified paymentDetailsId. Only those attributes provided in the request will be updated.
*/
func (a *Client) UpdatePaymentInfoUsingPATCH1(params *UpdatePaymentInfoUsingPATCH1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdatePaymentInfoUsingPATCH1OK, *UpdatePaymentInfoUsingPATCH1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePaymentInfoUsingPATCH1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePaymentInfoUsingPATCH_1",
		Method:             "PATCH",
		PathPattern:        "/{baseSiteId}/users/{userId}/paymentdetails/{paymentDetailsId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePaymentInfoUsingPATCH1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdatePaymentInfoUsingPATCH1OK:
		return value, nil, nil
	case *UpdatePaymentInfoUsingPATCH1NoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
