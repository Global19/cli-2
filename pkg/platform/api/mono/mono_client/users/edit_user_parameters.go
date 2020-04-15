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

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewEditUserParams creates a new EditUserParams object
// with the default values initialized.
func NewEditUserParams() *EditUserParams {
	var ()
	return &EditUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditUserParamsWithTimeout creates a new EditUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditUserParamsWithTimeout(timeout time.Duration) *EditUserParams {
	var ()
	return &EditUserParams{

		timeout: timeout,
	}
}

// NewEditUserParamsWithContext creates a new EditUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditUserParamsWithContext(ctx context.Context) *EditUserParams {
	var ()
	return &EditUserParams{

		Context: ctx,
	}
}

// NewEditUserParamsWithHTTPClient creates a new EditUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditUserParamsWithHTTPClient(client *http.Client) *EditUserParams {
	var ()
	return &EditUserParams{
		HTTPClient: client,
	}
}

/*EditUserParams contains all the parameters to send to the API endpoint
for the edit user operation typically these are written to a http.Request
*/
type EditUserParams struct {

	/*User*/
	User *mono_models.UserEditable
	/*Username
	  username of desired User

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit user params
func (o *EditUserParams) WithTimeout(timeout time.Duration) *EditUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit user params
func (o *EditUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit user params
func (o *EditUserParams) WithContext(ctx context.Context) *EditUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit user params
func (o *EditUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit user params
func (o *EditUserParams) WithHTTPClient(client *http.Client) *EditUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit user params
func (o *EditUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the edit user params
func (o *EditUserParams) WithUser(user *mono_models.UserEditable) *EditUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the edit user params
func (o *EditUserParams) SetUser(user *mono_models.UserEditable) {
	o.User = user
}

// WithUsername adds the username to the edit user params
func (o *EditUserParams) WithUsername(username string) *EditUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the edit user params
func (o *EditUserParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *EditUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.User != nil {
		if err := r.SetBodyParam(o.User); err != nil {
			return err
		}
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
