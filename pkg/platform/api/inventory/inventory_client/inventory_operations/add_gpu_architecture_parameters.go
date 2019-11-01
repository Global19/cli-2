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

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddGpuArchitectureParams creates a new AddGpuArchitectureParams object
// with the default values initialized.
func NewAddGpuArchitectureParams() *AddGpuArchitectureParams {
	var ()
	return &AddGpuArchitectureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddGpuArchitectureParamsWithTimeout creates a new AddGpuArchitectureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddGpuArchitectureParamsWithTimeout(timeout time.Duration) *AddGpuArchitectureParams {
	var ()
	return &AddGpuArchitectureParams{

		timeout: timeout,
	}
}

// NewAddGpuArchitectureParamsWithContext creates a new AddGpuArchitectureParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddGpuArchitectureParamsWithContext(ctx context.Context) *AddGpuArchitectureParams {
	var ()
	return &AddGpuArchitectureParams{

		Context: ctx,
	}
}

// NewAddGpuArchitectureParamsWithHTTPClient creates a new AddGpuArchitectureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddGpuArchitectureParamsWithHTTPClient(client *http.Client) *AddGpuArchitectureParams {
	var ()
	return &AddGpuArchitectureParams{
		HTTPClient: client,
	}
}

/*AddGpuArchitectureParams contains all the parameters to send to the API endpoint
for the add gpu architecture operation typically these are written to a http.Request
*/
type AddGpuArchitectureParams struct {

	/*GpuArchitecture*/
	GpuArchitecture *inventory_models.V1GpuArchitectureCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add gpu architecture params
func (o *AddGpuArchitectureParams) WithTimeout(timeout time.Duration) *AddGpuArchitectureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add gpu architecture params
func (o *AddGpuArchitectureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add gpu architecture params
func (o *AddGpuArchitectureParams) WithContext(ctx context.Context) *AddGpuArchitectureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add gpu architecture params
func (o *AddGpuArchitectureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add gpu architecture params
func (o *AddGpuArchitectureParams) WithHTTPClient(client *http.Client) *AddGpuArchitectureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add gpu architecture params
func (o *AddGpuArchitectureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGpuArchitecture adds the gpuArchitecture to the add gpu architecture params
func (o *AddGpuArchitectureParams) WithGpuArchitecture(gpuArchitecture *inventory_models.V1GpuArchitectureCore) *AddGpuArchitectureParams {
	o.SetGpuArchitecture(gpuArchitecture)
	return o
}

// SetGpuArchitecture adds the gpuArchitecture to the add gpu architecture params
func (o *AddGpuArchitectureParams) SetGpuArchitecture(gpuArchitecture *inventory_models.V1GpuArchitectureCore) {
	o.GpuArchitecture = gpuArchitecture
}

// WriteToRequest writes these params to a swagger request
func (o *AddGpuArchitectureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GpuArchitecture != nil {
		if err := r.SetBodyParam(o.GpuArchitecture); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}