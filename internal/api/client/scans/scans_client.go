// Code generated by go-swagger; DO NOT EDIT.

package scans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new scans API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scans API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetScan gets a single scan record

Retrieve the scan by ID
*/
func (a *Client) GetScan(params *GetScanParams, authInfo runtime.ClientAuthInfoWriter) (*GetScanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScan",
		Method:             "GET",
		PathPattern:        "/scans/{scanID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScanOK), nil

}

/*
ListScans recents scan requests

List of recent scan requests
*/
func (a *Client) ListScans(params *ListScansParams, authInfo runtime.ClientAuthInfoWriter) (*ListScansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScansParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listScans",
		Method:             "GET",
		PathPattern:        "/scans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListScansReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListScansOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}