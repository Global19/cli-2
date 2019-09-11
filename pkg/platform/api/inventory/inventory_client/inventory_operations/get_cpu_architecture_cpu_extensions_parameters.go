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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCPUArchitectureCPUExtensionsParams creates a new GetCPUArchitectureCPUExtensionsParams object
// with the default values initialized.
func NewGetCPUArchitectureCPUExtensionsParams() *GetCPUArchitectureCPUExtensionsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitectureCPUExtensionsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCPUArchitectureCPUExtensionsParamsWithTimeout creates a new GetCPUArchitectureCPUExtensionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCPUArchitectureCPUExtensionsParamsWithTimeout(timeout time.Duration) *GetCPUArchitectureCPUExtensionsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitectureCPUExtensionsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: timeout,
	}
}

// NewGetCPUArchitectureCPUExtensionsParamsWithContext creates a new GetCPUArchitectureCPUExtensionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCPUArchitectureCPUExtensionsParamsWithContext(ctx context.Context) *GetCPUArchitectureCPUExtensionsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitectureCPUExtensionsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		Context: ctx,
	}
}

// NewGetCPUArchitectureCPUExtensionsParamsWithHTTPClient creates a new GetCPUArchitectureCPUExtensionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCPUArchitectureCPUExtensionsParamsWithHTTPClient(client *http.Client) *GetCPUArchitectureCPUExtensionsParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitectureCPUExtensionsParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
		HTTPClient:    client,
	}
}

/*GetCPUArchitectureCPUExtensionsParams contains all the parameters to send to the API endpoint
for the get Cpu architecture Cpu extensions operation typically these are written to a http.Request
*/
type GetCPUArchitectureCPUExtensionsParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*CPUArchitectureID*/
	CPUArchitectureID strfmt.UUID
	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
	/*Page
	  The page number returned

	*/
	Page *int64
	/*StateAt
	  Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithTimeout(timeout time.Duration) *GetCPUArchitectureCPUExtensionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithContext(ctx context.Context) *GetCPUArchitectureCPUExtensionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithHTTPClient(client *http.Client) *GetCPUArchitectureCPUExtensionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithAllowUnstable(allowUnstable *bool) *GetCPUArchitectureCPUExtensionsParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithCPUArchitectureID adds the cPUArchitectureID to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithCPUArchitectureID(cPUArchitectureID strfmt.UUID) *GetCPUArchitectureCPUExtensionsParams {
	o.SetCPUArchitectureID(cPUArchitectureID)
	return o
}

// SetCPUArchitectureID adds the cpuArchitectureId to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetCPUArchitectureID(cPUArchitectureID strfmt.UUID) {
	o.CPUArchitectureID = cPUArchitectureID
}

// WithLimit adds the limit to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithLimit(limit *int64) *GetCPUArchitectureCPUExtensionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithPage(page *int64) *GetCPUArchitectureCPUExtensionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) WithStateAt(stateAt *strfmt.DateTime) *GetCPUArchitectureCPUExtensionsParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get Cpu architecture Cpu extensions params
func (o *GetCPUArchitectureCPUExtensionsParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetCPUArchitectureCPUExtensionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool
		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {
			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}

	}

	// path param cpu_architecture_id
	if err := r.SetPathParam("cpu_architecture_id", o.CPUArchitectureID.String()); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime
		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {
			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
