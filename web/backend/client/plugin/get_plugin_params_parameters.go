// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPluginParamsParams creates a new GetPluginParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPluginParamsParams() *GetPluginParamsParams {
	return &GetPluginParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginParamsParamsWithTimeout creates a new GetPluginParamsParams object
// with the ability to set a timeout on a request.
func NewGetPluginParamsParamsWithTimeout(timeout time.Duration) *GetPluginParamsParams {
	return &GetPluginParamsParams{
		timeout: timeout,
	}
}

// NewGetPluginParamsParamsWithContext creates a new GetPluginParamsParams object
// with the ability to set a context for a request.
func NewGetPluginParamsParamsWithContext(ctx context.Context) *GetPluginParamsParams {
	return &GetPluginParamsParams{
		Context: ctx,
	}
}

// NewGetPluginParamsParamsWithHTTPClient creates a new GetPluginParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPluginParamsParamsWithHTTPClient(client *http.Client) *GetPluginParamsParams {
	return &GetPluginParamsParams{
		HTTPClient: client,
	}
}

/*
GetPluginParamsParams contains all the parameters to send to the API endpoint

	for the get plugin params operation.

	Typically these are written to a http.Request.
*/
type GetPluginParamsParams struct {

	/* Name.

	   name of plugin
	*/
	Name *string

	/* PluginType.

	   plugin type
	*/
	PluginType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get plugin params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPluginParamsParams) WithDefaults() *GetPluginParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plugin params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPluginParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get plugin params params
func (o *GetPluginParamsParams) WithTimeout(timeout time.Duration) *GetPluginParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin params params
func (o *GetPluginParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin params params
func (o *GetPluginParamsParams) WithContext(ctx context.Context) *GetPluginParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin params params
func (o *GetPluginParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin params params
func (o *GetPluginParamsParams) WithHTTPClient(client *http.Client) *GetPluginParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin params params
func (o *GetPluginParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get plugin params params
func (o *GetPluginParamsParams) WithName(name *string) *GetPluginParamsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get plugin params params
func (o *GetPluginParamsParams) SetName(name *string) {
	o.Name = name
}

// WithPluginType adds the pluginType to the get plugin params params
func (o *GetPluginParamsParams) WithPluginType(pluginType *string) *GetPluginParamsParams {
	o.SetPluginType(pluginType)
	return o
}

// SetPluginType adds the pluginType to the get plugin params params
func (o *GetPluginParamsParams) SetPluginType(pluginType *string) {
	o.PluginType = pluginType
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.PluginType != nil {

		// query param pluginType
		var qrPluginType string

		if o.PluginType != nil {
			qrPluginType = *o.PluginType
		}
		qPluginType := qrPluginType
		if qPluginType != "" {

			if err := r.SetQueryParam("pluginType", qPluginType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}