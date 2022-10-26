// Code generated by go-swagger; DO NOT EDIT.

package vsphere

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

// NewListVSphereFoldersParams creates a new ListVSphereFoldersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVSphereFoldersParams() *ListVSphereFoldersParams {
	return &ListVSphereFoldersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVSphereFoldersParamsWithTimeout creates a new ListVSphereFoldersParams object
// with the ability to set a timeout on a request.
func NewListVSphereFoldersParamsWithTimeout(timeout time.Duration) *ListVSphereFoldersParams {
	return &ListVSphereFoldersParams{
		timeout: timeout,
	}
}

// NewListVSphereFoldersParamsWithContext creates a new ListVSphereFoldersParams object
// with the ability to set a context for a request.
func NewListVSphereFoldersParamsWithContext(ctx context.Context) *ListVSphereFoldersParams {
	return &ListVSphereFoldersParams{
		Context: ctx,
	}
}

// NewListVSphereFoldersParamsWithHTTPClient creates a new ListVSphereFoldersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVSphereFoldersParamsWithHTTPClient(client *http.Client) *ListVSphereFoldersParams {
	return &ListVSphereFoldersParams{
		HTTPClient: client,
	}
}

/*
ListVSphereFoldersParams contains all the parameters to send to the API endpoint

	for the list v sphere folders operation.

	Typically these are written to a http.Request.
*/
type ListVSphereFoldersParams struct {

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Password.
	Password *string

	// Username.
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list v sphere folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVSphereFoldersParams) WithDefaults() *ListVSphereFoldersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list v sphere folders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVSphereFoldersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithTimeout(timeout time.Duration) *ListVSphereFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithContext(ctx context.Context) *ListVSphereFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithHTTPClient(client *http.Client) *ListVSphereFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithCredential(credential *string) *ListVSphereFoldersParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithDatacenterName(datacenterName *string) *ListVSphereFoldersParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithPassword adds the password to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithPassword(password *string) *ListVSphereFoldersParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithUsername(username *string) *ListVSphereFoldersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *ListVSphereFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}
	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}