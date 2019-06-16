// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteMaybe frees a token slot
*/
func (a *Client) DeleteMaybe(params *DeleteMaybeParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteMaybeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMaybeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMaybe",
		Method:             "DELETE",
		PathPattern:        "/maybe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteMaybeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMaybeNoContent), nil

}

/*
GetMaybe checks if there is an open slot
*/
func (a *Client) GetMaybe(params *GetMaybeParams, authInfo runtime.ClientAuthInfoWriter) (*GetMaybeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMaybeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMaybe",
		Method:             "GET",
		PathPattern:        "/maybe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMaybeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMaybeOK), nil

}

/*
PostMaybe takes an open slot
*/
func (a *Client) PostMaybe(params *PostMaybeParams, authInfo runtime.ClientAuthInfoWriter) (*PostMaybeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMaybeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostMaybe",
		Method:             "POST",
		PathPattern:        "/maybe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostMaybeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostMaybeCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
