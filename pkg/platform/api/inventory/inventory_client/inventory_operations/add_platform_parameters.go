// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

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

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddPlatformParams creates a new AddPlatformParams object
// with the default values initialized.
func NewAddPlatformParams() *AddPlatformParams {
	var ()
	return &AddPlatformParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPlatformParamsWithTimeout creates a new AddPlatformParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPlatformParamsWithTimeout(timeout time.Duration) *AddPlatformParams {
	var ()
	return &AddPlatformParams{

		timeout: timeout,
	}
}

// NewAddPlatformParamsWithContext creates a new AddPlatformParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPlatformParamsWithContext(ctx context.Context) *AddPlatformParams {
	var ()
	return &AddPlatformParams{

		Context: ctx,
	}
}

// NewAddPlatformParamsWithHTTPClient creates a new AddPlatformParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPlatformParamsWithHTTPClient(client *http.Client) *AddPlatformParams {
	var ()
	return &AddPlatformParams{
		HTTPClient: client,
	}
}

/*AddPlatformParams contains all the parameters to send to the API endpoint
for the add platform operation typically these are written to a http.Request
*/
type AddPlatformParams struct {

	/*Platform*/
	Platform *inventory_models.V1PlatformCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add platform params
func (o *AddPlatformParams) WithTimeout(timeout time.Duration) *AddPlatformParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add platform params
func (o *AddPlatformParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add platform params
func (o *AddPlatformParams) WithContext(ctx context.Context) *AddPlatformParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add platform params
func (o *AddPlatformParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add platform params
func (o *AddPlatformParams) WithHTTPClient(client *http.Client) *AddPlatformParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add platform params
func (o *AddPlatformParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPlatform adds the platform to the add platform params
func (o *AddPlatformParams) WithPlatform(platform *inventory_models.V1PlatformCreate) *AddPlatformParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the add platform params
func (o *AddPlatformParams) SetPlatform(platform *inventory_models.V1PlatformCreate) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *AddPlatformParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Platform != nil {
		if err := r.SetBodyParam(o.Platform); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
