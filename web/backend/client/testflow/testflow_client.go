// Code generated by go-swagger; DO NOT EDIT.

package testflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new testflow API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for testflow API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Changetestflow(params *ChangetestflowParams, opts ...ClientOption) (*ChangetestflowOK, error)

	CountTestFlowsInGroup(params *CountTestFlowsInGroupParams, opts ...ClientOption) (*CountTestFlowsInGroupOK, error)

	DeleteTestFlow(params *DeleteTestFlowParams, opts ...ClientOption) (*DeleteTestFlowOK, error)

	GetTestFlow(params *GetTestFlowParams, opts ...ClientOption) (*GetTestFlowOK, error)

	ListTestFlows(params *ListTestFlowsParams, opts ...ClientOption) (*ListTestFlowsOK, error)

	SaveTestFlow(params *SaveTestFlowParams, opts ...ClientOption) (*SaveTestFlowOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Changetestflow change testflow group id
*/
func (a *Client) Changetestflow(params *ChangetestflowParams, opts ...ClientOption) (*ChangetestflowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangetestflowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changetestflow",
		Method:             "POST",
		PathPattern:        "/changegroup",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangetestflowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangetestflowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changetestflow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CountTestFlowsInGroup Count testflow numbers in group
*/
func (a *Client) CountTestFlowsInGroup(params *CountTestFlowsInGroupParams, opts ...ClientOption) (*CountTestFlowsInGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCountTestFlowsInGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "countTestFlowsInGroup",
		Method:             "GET",
		PathPattern:        "/testflow/count",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CountTestFlowsInGroupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CountTestFlowsInGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for countTestFlowsInGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTestFlow Delete test flow by id
*/
func (a *Client) DeleteTestFlow(params *DeleteTestFlowParams, opts ...ClientOption) (*DeleteTestFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTestFlowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTestFlow",
		Method:             "DELETE",
		PathPattern:        "/testflow/{id}",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTestFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTestFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTestFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTestFlow gets specific test case by condition
*/
func (a *Client) GetTestFlow(params *GetTestFlowParams, opts ...ClientOption) (*GetTestFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestFlowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTestFlow",
		Method:             "GET",
		PathPattern:        "/testflow",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTestFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTestFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTestFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTestFlows lists test flows
*/
func (a *Client) ListTestFlows(params *ListTestFlowsParams, opts ...ClientOption) (*ListTestFlowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTestFlowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTestFlows",
		Method:             "GET",
		PathPattern:        "/testflow/list",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTestFlowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListTestFlowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTestFlows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SaveTestFlow save test case, create if not exist
*/
func (a *Client) SaveTestFlow(params *SaveTestFlowParams, opts ...ClientOption) (*SaveTestFlowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveTestFlowParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "saveTestFlow",
		Method:             "POST",
		PathPattern:        "/testflow",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SaveTestFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SaveTestFlowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for saveTestFlow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}