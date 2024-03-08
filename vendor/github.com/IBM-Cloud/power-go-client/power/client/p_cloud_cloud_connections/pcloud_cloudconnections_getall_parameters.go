// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

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

// NewPcloudCloudconnectionsGetallParams creates a new PcloudCloudconnectionsGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudconnectionsGetallParams() *PcloudCloudconnectionsGetallParams {
	return &PcloudCloudconnectionsGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudconnectionsGetallParamsWithTimeout creates a new PcloudCloudconnectionsGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudconnectionsGetallParamsWithTimeout(timeout time.Duration) *PcloudCloudconnectionsGetallParams {
	return &PcloudCloudconnectionsGetallParams{
		timeout: timeout,
	}
}

// NewPcloudCloudconnectionsGetallParamsWithContext creates a new PcloudCloudconnectionsGetallParams object
// with the ability to set a context for a request.
func NewPcloudCloudconnectionsGetallParamsWithContext(ctx context.Context) *PcloudCloudconnectionsGetallParams {
	return &PcloudCloudconnectionsGetallParams{
		Context: ctx,
	}
}

// NewPcloudCloudconnectionsGetallParamsWithHTTPClient creates a new PcloudCloudconnectionsGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudconnectionsGetallParamsWithHTTPClient(client *http.Client) *PcloudCloudconnectionsGetallParams {
	return &PcloudCloudconnectionsGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudCloudconnectionsGetallParams contains all the parameters to send to the API endpoint

	for the pcloud cloudconnections getall operation.

	Typically these are written to a http.Request.
*/
type PcloudCloudconnectionsGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudconnections getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudconnectionsGetallParams) WithDefaults() *PcloudCloudconnectionsGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudconnections getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudconnectionsGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) WithTimeout(timeout time.Duration) *PcloudCloudconnectionsGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) WithContext(ctx context.Context) *PcloudCloudconnectionsGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) WithHTTPClient(client *http.Client) *PcloudCloudconnectionsGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudconnectionsGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudconnections getall params
func (o *PcloudCloudconnectionsGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudconnectionsGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}