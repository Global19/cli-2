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

// NewGetNamespaceIngredientParams creates a new GetNamespaceIngredientParams object
// with the default values initialized.
func NewGetNamespaceIngredientParams() *GetNamespaceIngredientParams {
	var ()
	return &GetNamespaceIngredientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceIngredientParamsWithTimeout creates a new GetNamespaceIngredientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNamespaceIngredientParamsWithTimeout(timeout time.Duration) *GetNamespaceIngredientParams {
	var ()
	return &GetNamespaceIngredientParams{

		timeout: timeout,
	}
}

// NewGetNamespaceIngredientParamsWithContext creates a new GetNamespaceIngredientParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNamespaceIngredientParamsWithContext(ctx context.Context) *GetNamespaceIngredientParams {
	var ()
	return &GetNamespaceIngredientParams{

		Context: ctx,
	}
}

// NewGetNamespaceIngredientParamsWithHTTPClient creates a new GetNamespaceIngredientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNamespaceIngredientParamsWithHTTPClient(client *http.Client) *GetNamespaceIngredientParams {
	var ()
	return &GetNamespaceIngredientParams{
		HTTPClient: client,
	}
}

/*GetNamespaceIngredientParams contains all the parameters to send to the API endpoint
for the get namespace ingredient operation typically these are written to a http.Request
*/
type GetNamespaceIngredientParams struct {

	/*Name*/
	Name string
	/*Namespace*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) WithTimeout(timeout time.Duration) *GetNamespaceIngredientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) WithContext(ctx context.Context) *GetNamespaceIngredientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) WithHTTPClient(client *http.Client) *GetNamespaceIngredientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) WithName(name string) *GetNamespaceIngredientParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) WithNamespace(namespace string) *GetNamespaceIngredientParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get namespace ingredient params
func (o *GetNamespaceIngredientParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespaceIngredientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {
		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
