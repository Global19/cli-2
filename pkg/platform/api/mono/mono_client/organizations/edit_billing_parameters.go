// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewEditBillingParams creates a new EditBillingParams object
// with the default values initialized.
func NewEditBillingParams() *EditBillingParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditBillingParams{
		IdentifierType: &identifierTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEditBillingParamsWithTimeout creates a new EditBillingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditBillingParamsWithTimeout(timeout time.Duration) *EditBillingParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditBillingParams{
		IdentifierType: &identifierTypeDefault,

		timeout: timeout,
	}
}

// NewEditBillingParamsWithContext creates a new EditBillingParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditBillingParamsWithContext(ctx context.Context) *EditBillingParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditBillingParams{
		IdentifierType: &identifierTypeDefault,

		Context: ctx,
	}
}

// NewEditBillingParamsWithHTTPClient creates a new EditBillingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditBillingParamsWithHTTPClient(client *http.Client) *EditBillingParams {
	var (
		identifierTypeDefault = string("URLname")
	)
	return &EditBillingParams{
		IdentifierType: &identifierTypeDefault,
		HTTPClient:     client,
	}
}

/*EditBillingParams contains all the parameters to send to the API endpoint
for the edit billing operation typically these are written to a http.Request
*/
type EditBillingParams struct {

	/*BillingInformation
	  updated billing information

	*/
	BillingInformation *mono_models.BillingInformationEditable
	/*IdentifierType
	  what kind of thing the provided organizationIdentifier is

	*/
	IdentifierType *string
	/*OrganizationIdentifier
	  identifier (URLname, by default) of the desired organization

	*/
	OrganizationIdentifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit billing params
func (o *EditBillingParams) WithTimeout(timeout time.Duration) *EditBillingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit billing params
func (o *EditBillingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit billing params
func (o *EditBillingParams) WithContext(ctx context.Context) *EditBillingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit billing params
func (o *EditBillingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit billing params
func (o *EditBillingParams) WithHTTPClient(client *http.Client) *EditBillingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit billing params
func (o *EditBillingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingInformation adds the billingInformation to the edit billing params
func (o *EditBillingParams) WithBillingInformation(billingInformation *mono_models.BillingInformationEditable) *EditBillingParams {
	o.SetBillingInformation(billingInformation)
	return o
}

// SetBillingInformation adds the billingInformation to the edit billing params
func (o *EditBillingParams) SetBillingInformation(billingInformation *mono_models.BillingInformationEditable) {
	o.BillingInformation = billingInformation
}

// WithIdentifierType adds the identifierType to the edit billing params
func (o *EditBillingParams) WithIdentifierType(identifierType *string) *EditBillingParams {
	o.SetIdentifierType(identifierType)
	return o
}

// SetIdentifierType adds the identifierType to the edit billing params
func (o *EditBillingParams) SetIdentifierType(identifierType *string) {
	o.IdentifierType = identifierType
}

// WithOrganizationIdentifier adds the organizationIdentifier to the edit billing params
func (o *EditBillingParams) WithOrganizationIdentifier(organizationIdentifier string) *EditBillingParams {
	o.SetOrganizationIdentifier(organizationIdentifier)
	return o
}

// SetOrganizationIdentifier adds the organizationIdentifier to the edit billing params
func (o *EditBillingParams) SetOrganizationIdentifier(organizationIdentifier string) {
	o.OrganizationIdentifier = organizationIdentifier
}

// WriteToRequest writes these params to a swagger request
func (o *EditBillingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BillingInformation != nil {
		if err := r.SetBodyParam(o.BillingInformation); err != nil {
			return err
		}
	}

	if o.IdentifierType != nil {

		// query param identifierType
		var qrIdentifierType string
		if o.IdentifierType != nil {
			qrIdentifierType = *o.IdentifierType
		}
		qIdentifierType := qrIdentifierType
		if qIdentifierType != "" {
			if err := r.SetQueryParam("identifierType", qIdentifierType); err != nil {
				return err
			}
		}

	}

	// path param organizationIdentifier
	if err := r.SetPathParam("organizationIdentifier", o.OrganizationIdentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
