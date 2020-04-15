// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetEmailsByUserParams creates a new GetEmailsByUserParams object
// with the default values initialized.
func NewGetEmailsByUserParams() *GetEmailsByUserParams {
	var ()
	return &GetEmailsByUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailsByUserParamsWithTimeout creates a new GetEmailsByUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailsByUserParamsWithTimeout(timeout time.Duration) *GetEmailsByUserParams {
	var ()
	return &GetEmailsByUserParams{

		timeout: timeout,
	}
}

// NewGetEmailsByUserParamsWithContext creates a new GetEmailsByUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailsByUserParamsWithContext(ctx context.Context) *GetEmailsByUserParams {
	var ()
	return &GetEmailsByUserParams{

		Context: ctx,
	}
}

// NewGetEmailsByUserParamsWithHTTPClient creates a new GetEmailsByUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailsByUserParamsWithHTTPClient(client *http.Client) *GetEmailsByUserParams {
	var ()
	return &GetEmailsByUserParams{
		HTTPClient: client,
	}
}

/*GetEmailsByUserParams contains all the parameters to send to the API endpoint
for the get emails by user operation typically these are written to a http.Request
*/
type GetEmailsByUserParams struct {

	/*Username
	  username of desired User

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get emails by user params
func (o *GetEmailsByUserParams) WithTimeout(timeout time.Duration) *GetEmailsByUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get emails by user params
func (o *GetEmailsByUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get emails by user params
func (o *GetEmailsByUserParams) WithContext(ctx context.Context) *GetEmailsByUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get emails by user params
func (o *GetEmailsByUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get emails by user params
func (o *GetEmailsByUserParams) WithHTTPClient(client *http.Client) *GetEmailsByUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get emails by user params
func (o *GetEmailsByUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get emails by user params
func (o *GetEmailsByUserParams) WithUsername(username string) *GetEmailsByUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get emails by user params
func (o *GetEmailsByUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailsByUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
