// Code generated by go-swagger; DO NOT EDIT.

package log

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

// NewPodLogReqParams creates a new PodLogReqParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodLogReqParams() *PodLogReqParams {
	return &PodLogReqParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodLogReqParamsWithTimeout creates a new PodLogReqParams object
// with the ability to set a timeout on a request.
func NewPodLogReqParamsWithTimeout(timeout time.Duration) *PodLogReqParams {
	return &PodLogReqParams{
		timeout: timeout,
	}
}

// NewPodLogReqParamsWithContext creates a new PodLogReqParams object
// with the ability to set a context for a request.
func NewPodLogReqParamsWithContext(ctx context.Context) *PodLogReqParams {
	return &PodLogReqParams{
		Context: ctx,
	}
}

// NewPodLogReqParamsWithHTTPClient creates a new PodLogReqParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodLogReqParamsWithHTTPClient(client *http.Client) *PodLogReqParams {
	return &PodLogReqParams{
		HTTPClient: client,
	}
}

/*
PodLogReqParams contains all the parameters to send to the API endpoint

	for the pod log req operation.

	Typically these are written to a http.Request.
*/
type PodLogReqParams struct {

	/* PodLog.

	   pod name
	*/
	PodName string

	/* TestID.

	   testid of task
	*/
	TestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod log req params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodLogReqParams) WithDefaults() *PodLogReqParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod log req params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodLogReqParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod log req params
func (o *PodLogReqParams) WithTimeout(timeout time.Duration) *PodLogReqParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod log req params
func (o *PodLogReqParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod log req params
func (o *PodLogReqParams) WithContext(ctx context.Context) *PodLogReqParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod log req params
func (o *PodLogReqParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod log req params
func (o *PodLogReqParams) WithHTTPClient(client *http.Client) *PodLogReqParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod log req params
func (o *PodLogReqParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPodName adds the podLog to the pod log req params
func (o *PodLogReqParams) WithPodName(podLog string) *PodLogReqParams {
	o.SetPodName(podLog)
	return o
}

// SetPodName adds the podLog to the pod log req params
func (o *PodLogReqParams) SetPodName(podLog string) {
	o.PodName = podLog
}

// WithTestID adds the testID to the pod log req params
func (o *PodLogReqParams) WithTestID(testID string) *PodLogReqParams {
	o.SetTestID(testID)
	return o
}

// SetTestID adds the testId to the pod log req params
func (o *PodLogReqParams) SetTestID(testID string) {
	o.TestID = testID
}

// WriteToRequest writes these params to a swagger request
func (o *PodLogReqParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param podLog
	qrPodLog := o.PodName
	qPodLog := qrPodLog
	if qPodLog != "" {

		if err := r.SetQueryParam("podLog", qPodLog); err != nil {
			return err
		}
	}

	// query param testID
	qrTestID := o.TestID
	qTestID := qrTestID
	if qTestID != "" {

		if err := r.SetQueryParam("testID", qTestID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}