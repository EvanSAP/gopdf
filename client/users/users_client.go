// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeLoginUsingPUT changes customer s login name

Changes a customer's login name. Requires the customer's current password.
*/
func (a *Client) ChangeLoginUsingPUT(params *ChangeLoginUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeLoginUsingPUTOK, *ChangeLoginUsingPUTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeLoginUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeLoginUsingPUT",
		Method:             "PUT",
		PathPattern:        "/{baseSiteId}/users/{userId}/login",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangeLoginUsingPUTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ChangeLoginUsingPUTOK:
		return value, nil, nil
	case *ChangeLoginUsingPUTCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ChangePasswordUsingPUT changes customer s password

Changes customer's password.
*/
func (a *Client) ChangePasswordUsingPUT(params *ChangePasswordUsingPUTParams, authInfo runtime.ClientAuthInfoWriter) (*ChangePasswordUsingPUTCreated, *ChangePasswordUsingPUTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changePasswordUsingPUT",
		Method:             "PUT",
		PathPattern:        "/{baseSiteId}/users/{userId}/password",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangePasswordUsingPUTReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ChangePasswordUsingPUTCreated:
		return value, nil, nil
	case *ChangePasswordUsingPUTAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeactivateUserUsingDELETE deletes customer profile

Removes customer profile.
*/
func (a *Client) DeactivateUserUsingDELETE(params *DeactivateUserUsingDELETEParams, authInfo runtime.ClientAuthInfoWriter) (*DeactivateUserUsingDELETEOK, *DeactivateUserUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeactivateUserUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deactivateUserUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/{baseSiteId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeactivateUserUsingDELETEReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeactivateUserUsingDELETEOK:
		return value, nil, nil
	case *DeactivateUserUsingDELETENoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetAllCustomerGroupsForCustomerUsingGET gets all customer groups of a customer

Returns all customer groups of a customer.
*/
func (a *Client) GetAllCustomerGroupsForCustomerUsingGET(params *GetAllCustomerGroupsForCustomerUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllCustomerGroupsForCustomerUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllCustomerGroupsForCustomerUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllCustomerGroupsForCustomerUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/users/{userId}/customergroups",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllCustomerGroupsForCustomerUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllCustomerGroupsForCustomerUsingGETOK), nil

}

/*
GetUserUsingGET gets customer profile

Returns customer profile.
*/
func (a *Client) GetUserUsingGET(params *GetUserUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserUsingGETOK), nil

}

/*
PutUserUsingPUT1 updates customer profile

Updates customer profile. Attributes not provided in the request body will be defined again (set to null or default).
*/
func (a *Client) PutUserUsingPUT1(params *PutUserUsingPUT1Params, authInfo runtime.ClientAuthInfoWriter) (*PutUserUsingPUT1OK, *PutUserUsingPUT1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUserUsingPUT1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/{baseSiteId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutUserUsingPUT1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PutUserUsingPUT1OK:
		return value, nil, nil
	case *PutUserUsingPUT1Created:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
RegisterUserUsingPOST1 registers a customer

Registers a customer. There are two options for registering a customer. The first option requires the following parameters: login, password, firstName, lastName, titleCode. The second option converts a guest to a customer. In this case, the required parameters are: guid, password.
*/
func (a *Client) RegisterUserUsingPOST1(params *RegisterUserUsingPOST1Params, authInfo runtime.ClientAuthInfoWriter) (*RegisterUserUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterUserUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerUserUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/{baseSiteId}/users",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterUserUsingPOST1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterUserUsingPOST1Created), nil

}

/*
UpdateUserUsingPATCH1 updates customer profile

Updates customer profile. Only attributes provided in the request body will be changed.
*/
func (a *Client) UpdateUserUsingPATCH1(params *UpdateUserUsingPATCH1Params, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserUsingPATCH1OK, *UpdateUserUsingPATCH1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserUsingPATCH1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserUsingPATCH_1",
		Method:             "PATCH",
		PathPattern:        "/{baseSiteId}/users/{userId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserUsingPATCH1Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateUserUsingPATCH1OK:
		return value, nil, nil
	case *UpdateUserUsingPATCH1NoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
