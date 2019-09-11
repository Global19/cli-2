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

// NewAddLibcVersionRevisionParams creates a new AddLibcVersionRevisionParams object
// with the default values initialized.
func NewAddLibcVersionRevisionParams() *AddLibcVersionRevisionParams {
	var ()
	return &AddLibcVersionRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddLibcVersionRevisionParamsWithTimeout creates a new AddLibcVersionRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddLibcVersionRevisionParamsWithTimeout(timeout time.Duration) *AddLibcVersionRevisionParams {
	var ()
	return &AddLibcVersionRevisionParams{

		timeout: timeout,
	}
}

// NewAddLibcVersionRevisionParamsWithContext creates a new AddLibcVersionRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddLibcVersionRevisionParamsWithContext(ctx context.Context) *AddLibcVersionRevisionParams {
	var ()
	return &AddLibcVersionRevisionParams{

		Context: ctx,
	}
}

// NewAddLibcVersionRevisionParamsWithHTTPClient creates a new AddLibcVersionRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddLibcVersionRevisionParamsWithHTTPClient(client *http.Client) *AddLibcVersionRevisionParams {
	var ()
	return &AddLibcVersionRevisionParams{
		HTTPClient: client,
	}
}

/*AddLibcVersionRevisionParams contains all the parameters to send to the API endpoint
for the add libc version revision operation typically these are written to a http.Request
*/
type AddLibcVersionRevisionParams struct {

	/*LibcID*/
	LibcID strfmt.UUID
	/*LibcVersionID*/
	LibcVersionID strfmt.UUID
	/*LibcVersionRevision*/
	LibcVersionRevision *inventory_models.V1Revision

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add libc version revision params
func (o *AddLibcVersionRevisionParams) WithTimeout(timeout time.Duration) *AddLibcVersionRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add libc version revision params
func (o *AddLibcVersionRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add libc version revision params
func (o *AddLibcVersionRevisionParams) WithContext(ctx context.Context) *AddLibcVersionRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add libc version revision params
func (o *AddLibcVersionRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add libc version revision params
func (o *AddLibcVersionRevisionParams) WithHTTPClient(client *http.Client) *AddLibcVersionRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add libc version revision params
func (o *AddLibcVersionRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibcID adds the libcID to the add libc version revision params
func (o *AddLibcVersionRevisionParams) WithLibcID(libcID strfmt.UUID) *AddLibcVersionRevisionParams {
	o.SetLibcID(libcID)
	return o
}

// SetLibcID adds the libcId to the add libc version revision params
func (o *AddLibcVersionRevisionParams) SetLibcID(libcID strfmt.UUID) {
	o.LibcID = libcID
}

// WithLibcVersionID adds the libcVersionID to the add libc version revision params
func (o *AddLibcVersionRevisionParams) WithLibcVersionID(libcVersionID strfmt.UUID) *AddLibcVersionRevisionParams {
	o.SetLibcVersionID(libcVersionID)
	return o
}

// SetLibcVersionID adds the libcVersionId to the add libc version revision params
func (o *AddLibcVersionRevisionParams) SetLibcVersionID(libcVersionID strfmt.UUID) {
	o.LibcVersionID = libcVersionID
}

// WithLibcVersionRevision adds the libcVersionRevision to the add libc version revision params
func (o *AddLibcVersionRevisionParams) WithLibcVersionRevision(libcVersionRevision *inventory_models.V1Revision) *AddLibcVersionRevisionParams {
	o.SetLibcVersionRevision(libcVersionRevision)
	return o
}

// SetLibcVersionRevision adds the libcVersionRevision to the add libc version revision params
func (o *AddLibcVersionRevisionParams) SetLibcVersionRevision(libcVersionRevision *inventory_models.V1Revision) {
	o.LibcVersionRevision = libcVersionRevision
}

// WriteToRequest writes these params to a swagger request
func (o *AddLibcVersionRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param libc_id
	if err := r.SetPathParam("libc_id", o.LibcID.String()); err != nil {
		return err
	}

	// path param libc_version_id
	if err := r.SetPathParam("libc_version_id", o.LibcVersionID.String()); err != nil {
		return err
	}

	if o.LibcVersionRevision != nil {
		if err := r.SetBodyParam(o.LibcVersionRevision); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
