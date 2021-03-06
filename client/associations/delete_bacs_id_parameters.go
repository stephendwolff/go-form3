// Code generated by go-swagger; DO NOT EDIT.

package associations

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

// NewDeleteBacsIDParams creates a new DeleteBacsIDParams object
// with the default values initialized.
func NewDeleteBacsIDParams() *DeleteBacsIDParams {
	var ()
	return &DeleteBacsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBacsIDParamsWithTimeout creates a new DeleteBacsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBacsIDParamsWithTimeout(timeout time.Duration) *DeleteBacsIDParams {
	var ()
	return &DeleteBacsIDParams{

		timeout: timeout,
	}
}

// NewDeleteBacsIDParamsWithContext creates a new DeleteBacsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBacsIDParamsWithContext(ctx context.Context) *DeleteBacsIDParams {
	var ()
	return &DeleteBacsIDParams{

		Context: ctx,
	}
}

// NewDeleteBacsIDParamsWithHTTPClient creates a new DeleteBacsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBacsIDParamsWithHTTPClient(client *http.Client) *DeleteBacsIDParams {
	var ()
	return &DeleteBacsIDParams{
		HTTPClient: client,
	}
}

/*DeleteBacsIDParams contains all the parameters to send to the API endpoint
for the delete bacs ID operation typically these are written to a http.Request
*/
type DeleteBacsIDParams struct {

	/*ID
	  Association Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete bacs ID params
func (o *DeleteBacsIDParams) WithTimeout(timeout time.Duration) *DeleteBacsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bacs ID params
func (o *DeleteBacsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bacs ID params
func (o *DeleteBacsIDParams) WithContext(ctx context.Context) *DeleteBacsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bacs ID params
func (o *DeleteBacsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bacs ID params
func (o *DeleteBacsIDParams) WithHTTPClient(client *http.Client) *DeleteBacsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bacs ID params
func (o *DeleteBacsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete bacs ID params
func (o *DeleteBacsIDParams) WithID(id strfmt.UUID) *DeleteBacsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete bacs ID params
func (o *DeleteBacsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete bacs ID params
func (o *DeleteBacsIDParams) WithVersion(version int64) *DeleteBacsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete bacs ID params
func (o *DeleteBacsIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBacsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
