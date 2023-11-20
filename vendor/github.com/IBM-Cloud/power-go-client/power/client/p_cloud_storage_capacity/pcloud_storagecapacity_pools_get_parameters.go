// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_storage_capacity

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

// NewPcloudStoragecapacityPoolsGetParams creates a new PcloudStoragecapacityPoolsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudStoragecapacityPoolsGetParams() *PcloudStoragecapacityPoolsGetParams {
	return &PcloudStoragecapacityPoolsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudStoragecapacityPoolsGetParamsWithTimeout creates a new PcloudStoragecapacityPoolsGetParams object
// with the ability to set a timeout on a request.
func NewPcloudStoragecapacityPoolsGetParamsWithTimeout(timeout time.Duration) *PcloudStoragecapacityPoolsGetParams {
	return &PcloudStoragecapacityPoolsGetParams{
		timeout: timeout,
	}
}

// NewPcloudStoragecapacityPoolsGetParamsWithContext creates a new PcloudStoragecapacityPoolsGetParams object
// with the ability to set a context for a request.
func NewPcloudStoragecapacityPoolsGetParamsWithContext(ctx context.Context) *PcloudStoragecapacityPoolsGetParams {
	return &PcloudStoragecapacityPoolsGetParams{
		Context: ctx,
	}
}

// NewPcloudStoragecapacityPoolsGetParamsWithHTTPClient creates a new PcloudStoragecapacityPoolsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudStoragecapacityPoolsGetParamsWithHTTPClient(client *http.Client) *PcloudStoragecapacityPoolsGetParams {
	return &PcloudStoragecapacityPoolsGetParams{
		HTTPClient: client,
	}
}

/*
PcloudStoragecapacityPoolsGetParams contains all the parameters to send to the API endpoint

	for the pcloud storagecapacity pools get operation.

	Typically these are written to a http.Request.
*/
type PcloudStoragecapacityPoolsGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* StoragePoolName.

	   Storage pool name
	*/
	StoragePoolName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud storagecapacity pools get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudStoragecapacityPoolsGetParams) WithDefaults() *PcloudStoragecapacityPoolsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud storagecapacity pools get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudStoragecapacityPoolsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) WithTimeout(timeout time.Duration) *PcloudStoragecapacityPoolsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) WithContext(ctx context.Context) *PcloudStoragecapacityPoolsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) WithHTTPClient(client *http.Client) *PcloudStoragecapacityPoolsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudStoragecapacityPoolsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithStoragePoolName adds the storagePoolName to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) WithStoragePoolName(storagePoolName string) *PcloudStoragecapacityPoolsGetParams {
	o.SetStoragePoolName(storagePoolName)
	return o
}

// SetStoragePoolName adds the storagePoolName to the pcloud storagecapacity pools get params
func (o *PcloudStoragecapacityPoolsGetParams) SetStoragePoolName(storagePoolName string) {
	o.StoragePoolName = storagePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudStoragecapacityPoolsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param storage_pool_name
	if err := r.SetPathParam("storage_pool_name", o.StoragePoolName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}