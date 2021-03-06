// Code generated by go-swagger; DO NOT EDIT.

package page

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new page API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for page API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPageDataUsingGET gets page data with list of cms content slots

Given a page identifier, return the page data with a list of cms content slots, each of which contains a list of cms component data.
*/
func (a *Client) GetPageDataUsingGET(params *GetPageDataUsingGETParams, authInfo runtime.ClientAuthInfoWriter) (*GetPageDataUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageDataUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPageDataUsingGET",
		Method:             "GET",
		PathPattern:        "/{baseSiteId}/cms/pages",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPageDataUsingGETReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPageDataUsingGETOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
