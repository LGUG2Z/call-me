// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostMaybeParams creates a new PostMaybeParams object
// with the default values initialized.
func NewPostMaybeParams() *PostMaybeParams {
	var ()
	return &PostMaybeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMaybeParamsWithTimeout creates a new PostMaybeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMaybeParamsWithTimeout(timeout time.Duration) *PostMaybeParams {
	var ()
	return &PostMaybeParams{

		timeout: timeout,
	}
}

// NewPostMaybeParamsWithContext creates a new PostMaybeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMaybeParamsWithContext(ctx context.Context) *PostMaybeParams {
	var ()
	return &PostMaybeParams{

		Context: ctx,
	}
}

// NewPostMaybeParamsWithHTTPClient creates a new PostMaybeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMaybeParamsWithHTTPClient(client *http.Client) *PostMaybeParams {
	var ()
	return &PostMaybeParams{
		HTTPClient: client,
	}
}

/*PostMaybeParams contains all the parameters to send to the API endpoint
for the post maybe operation typically these are written to a http.Request
*/
type PostMaybeParams struct {

	/*Environment
	  Environment name

	*/
	Environment string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post maybe params
func (o *PostMaybeParams) WithTimeout(timeout time.Duration) *PostMaybeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post maybe params
func (o *PostMaybeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post maybe params
func (o *PostMaybeParams) WithContext(ctx context.Context) *PostMaybeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post maybe params
func (o *PostMaybeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post maybe params
func (o *PostMaybeParams) WithHTTPClient(client *http.Client) *PostMaybeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post maybe params
func (o *PostMaybeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironment adds the environment to the post maybe params
func (o *PostMaybeParams) WithEnvironment(environment string) *PostMaybeParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the post maybe params
func (o *PostMaybeParams) SetEnvironment(environment string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *PostMaybeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param environment
	qrEnvironment := o.Environment
	qEnvironment := qrEnvironment
	if qEnvironment != "" {
		if err := r.SetQueryParam("environment", qEnvironment); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
