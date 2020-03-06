// Code generated by go-swagger; DO NOT EDIT.

package user_credentials

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

// NewDeleteCredentialParams creates a new DeleteCredentialParams object
// with the default values initialized.
func NewDeleteCredentialParams() *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCredentialParamsWithTimeout creates a new DeleteCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCredentialParamsWithTimeout(timeout time.Duration) *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{

		timeout: timeout,
	}
}

// NewDeleteCredentialParamsWithContext creates a new DeleteCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCredentialParamsWithContext(ctx context.Context) *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{

		Context: ctx,
	}
}

// NewDeleteCredentialParamsWithHTTPClient creates a new DeleteCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCredentialParamsWithHTTPClient(client *http.Client) *DeleteCredentialParams {
	var ()
	return &DeleteCredentialParams{
		HTTPClient: client,
	}
}

/*DeleteCredentialParams contains all the parameters to send to the API endpoint
for the delete credential operation typically these are written to a http.Request
*/
type DeleteCredentialParams struct {

	/*AccessKey
	  The credentials access key to be removed

	*/
	AccessKey string
	/*StackID
	  The stack's ID for which the user belongs to

	*/
	StackID string
	/*UserID
	  The user's ID for which the credentials will be generated

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete credential params
func (o *DeleteCredentialParams) WithTimeout(timeout time.Duration) *DeleteCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete credential params
func (o *DeleteCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete credential params
func (o *DeleteCredentialParams) WithContext(ctx context.Context) *DeleteCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete credential params
func (o *DeleteCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete credential params
func (o *DeleteCredentialParams) WithHTTPClient(client *http.Client) *DeleteCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete credential params
func (o *DeleteCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKey adds the accessKey to the delete credential params
func (o *DeleteCredentialParams) WithAccessKey(accessKey string) *DeleteCredentialParams {
	o.SetAccessKey(accessKey)
	return o
}

// SetAccessKey adds the accessKey to the delete credential params
func (o *DeleteCredentialParams) SetAccessKey(accessKey string) {
	o.AccessKey = accessKey
}

// WithStackID adds the stackID to the delete credential params
func (o *DeleteCredentialParams) WithStackID(stackID string) *DeleteCredentialParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete credential params
func (o *DeleteCredentialParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithUserID adds the userID to the delete credential params
func (o *DeleteCredentialParams) WithUserID(userID string) *DeleteCredentialParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete credential params
func (o *DeleteCredentialParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param access_key
	if err := r.SetPathParam("access_key", o.AccessKey); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
