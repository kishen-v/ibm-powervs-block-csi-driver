// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_shared_processor_pools

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

// NewPcloudSharedprocessorpoolsGetParams creates a new PcloudSharedprocessorpoolsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudSharedprocessorpoolsGetParams() *PcloudSharedprocessorpoolsGetParams {
	return &PcloudSharedprocessorpoolsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudSharedprocessorpoolsGetParamsWithTimeout creates a new PcloudSharedprocessorpoolsGetParams object
// with the ability to set a timeout on a request.
func NewPcloudSharedprocessorpoolsGetParamsWithTimeout(timeout time.Duration) *PcloudSharedprocessorpoolsGetParams {
	return &PcloudSharedprocessorpoolsGetParams{
		timeout: timeout,
	}
}

// NewPcloudSharedprocessorpoolsGetParamsWithContext creates a new PcloudSharedprocessorpoolsGetParams object
// with the ability to set a context for a request.
func NewPcloudSharedprocessorpoolsGetParamsWithContext(ctx context.Context) *PcloudSharedprocessorpoolsGetParams {
	return &PcloudSharedprocessorpoolsGetParams{
		Context: ctx,
	}
}

// NewPcloudSharedprocessorpoolsGetParamsWithHTTPClient creates a new PcloudSharedprocessorpoolsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudSharedprocessorpoolsGetParamsWithHTTPClient(client *http.Client) *PcloudSharedprocessorpoolsGetParams {
	return &PcloudSharedprocessorpoolsGetParams{
		HTTPClient: client,
	}
}

/*
PcloudSharedprocessorpoolsGetParams contains all the parameters to send to the API endpoint

	for the pcloud sharedprocessorpools get operation.

	Typically these are written to a http.Request.
*/
type PcloudSharedprocessorpoolsGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* SharedProcessorPoolID.

	   Shared Processor Pool ID or Name
	*/
	SharedProcessorPoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud sharedprocessorpools get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudSharedprocessorpoolsGetParams) WithDefaults() *PcloudSharedprocessorpoolsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud sharedprocessorpools get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudSharedprocessorpoolsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) WithTimeout(timeout time.Duration) *PcloudSharedprocessorpoolsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) WithContext(ctx context.Context) *PcloudSharedprocessorpoolsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) WithHTTPClient(client *http.Client) *PcloudSharedprocessorpoolsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudSharedprocessorpoolsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithSharedProcessorPoolID adds the sharedProcessorPoolID to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) WithSharedProcessorPoolID(sharedProcessorPoolID string) *PcloudSharedprocessorpoolsGetParams {
	o.SetSharedProcessorPoolID(sharedProcessorPoolID)
	return o
}

// SetSharedProcessorPoolID adds the sharedProcessorPoolId to the pcloud sharedprocessorpools get params
func (o *PcloudSharedprocessorpoolsGetParams) SetSharedProcessorPoolID(sharedProcessorPoolID string) {
	o.SharedProcessorPoolID = sharedProcessorPoolID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudSharedprocessorpoolsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param shared_processor_pool_id
	if err := r.SetPathParam("shared_processor_pool_id", o.SharedProcessorPoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}