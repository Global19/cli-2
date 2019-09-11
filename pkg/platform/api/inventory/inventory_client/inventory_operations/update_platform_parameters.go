// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdatePlatformParams creates a new UpdatePlatformParams object
// with the default values initialized.
func NewUpdatePlatformParams() *UpdatePlatformParams {
	var ()
	return &UpdatePlatformParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePlatformParamsWithTimeout creates a new UpdatePlatformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePlatformParamsWithTimeout(timeout time.Duration) *UpdatePlatformParams {
	var ()
	return &UpdatePlatformParams{

		timeout: timeout,
	}
}

// NewUpdatePlatformParamsWithContext creates a new UpdatePlatformParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePlatformParamsWithContext(ctx context.Context) *UpdatePlatformParams {
	var ()
	return &UpdatePlatformParams{

		Context: ctx,
	}
}

// NewUpdatePlatformParamsWithHTTPClient creates a new UpdatePlatformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePlatformParamsWithHTTPClient(client *http.Client) *UpdatePlatformParams {
	var ()
	return &UpdatePlatformParams{
		HTTPClient: client,
	}
}

/*UpdatePlatformParams contains all the parameters to send to the API endpoint
for the update platform operation typically these are written to a http.Request
*/
type UpdatePlatformParams struct {

	/*PlatformID*/
	PlatformID strfmt.UUID
	/*PlatformUpdate*/
	PlatformUpdate UpdatePlatformBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update platform params
func (o *UpdatePlatformParams) WithTimeout(timeout time.Duration) *UpdatePlatformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update platform params
func (o *UpdatePlatformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update platform params
func (o *UpdatePlatformParams) WithContext(ctx context.Context) *UpdatePlatformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update platform params
func (o *UpdatePlatformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update platform params
func (o *UpdatePlatformParams) WithHTTPClient(client *http.Client) *UpdatePlatformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update platform params
func (o *UpdatePlatformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlatformID adds the platformID to the update platform params
func (o *UpdatePlatformParams) WithPlatformID(platformID strfmt.UUID) *UpdatePlatformParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the update platform params
func (o *UpdatePlatformParams) SetPlatformID(platformID strfmt.UUID) {
	o.PlatformID = platformID
}

// WithPlatformUpdate adds the platformUpdate to the update platform params
func (o *UpdatePlatformParams) WithPlatformUpdate(platformUpdate UpdatePlatformBody) *UpdatePlatformParams {
	o.SetPlatformUpdate(platformUpdate)
	return o
}

// SetPlatformUpdate adds the platformUpdate to the update platform params
func (o *UpdatePlatformParams) SetPlatformUpdate(platformUpdate UpdatePlatformBody) {
	o.PlatformUpdate = platformUpdate
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePlatformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param platform_id
	if err := r.SetPathParam("platform_id", o.PlatformID.String()); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.PlatformUpdate); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
