// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new task API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for task API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteTask(params *DeleteTaskParams, opts ...ClientOption) (*DeleteTaskOK, error)

	GetTask(params *GetTaskParams, opts ...ClientOption) (*GetTaskOK, error)

	ListTasks(params *ListTasksParams, opts ...ClientOption) (*ListTasksOK, error)

	StopTask(params *StopTaskParams, opts ...ClientOption) (*StopTaskOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteTask Delete task by id
*/
func (a *Client) DeleteTask(params *DeleteTaskParams, opts ...ClientOption) (*DeleteTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTask",
		Method:             "DELETE",
		PathPattern:        "/task/{id}",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTaskReader{formats: a.formats},
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
	success, ok := result.(*DeleteTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTask Get task by id
*/
func (a *Client) GetTask(params *GetTaskParams, opts ...ClientOption) (*GetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTask",
		Method:             "GET",
		PathPattern:        "/task/{id}",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTaskReader{formats: a.formats},
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
	success, ok := result.(*GetTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTasks lists all tasks
*/
func (a *Client) ListTasks(params *ListTasksParams, opts ...ClientOption) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTasks",
		Method:             "GET",
		PathPattern:        "/task",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
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
	success, ok := result.(*ListTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopTask stop task
*/
func (a *Client) StopTask(params *StopTaskParams, opts ...ClientOption) (*StopTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopTask",
		Method:             "DELETE",
		PathPattern:        "/task/stop/{id}",
		ProducesMediaTypes: []string{"application/json", "application/text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopTaskReader{formats: a.formats},
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
	success, ok := result.(*StopTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}