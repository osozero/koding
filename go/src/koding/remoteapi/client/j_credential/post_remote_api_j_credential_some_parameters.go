package j_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJCredentialSomeParams creates a new PostRemoteAPIJCredentialSomeParams object
// with the default values initialized.
func NewPostRemoteAPIJCredentialSomeParams() *PostRemoteAPIJCredentialSomeParams {
	var ()
	return &PostRemoteAPIJCredentialSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJCredentialSomeParamsWithTimeout creates a new PostRemoteAPIJCredentialSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJCredentialSomeParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJCredentialSomeParams {
	var ()
	return &PostRemoteAPIJCredentialSomeParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJCredentialSomeParamsWithContext creates a new PostRemoteAPIJCredentialSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJCredentialSomeParamsWithContext(ctx context.Context) *PostRemoteAPIJCredentialSomeParams {
	var ()
	return &PostRemoteAPIJCredentialSomeParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJCredentialSomeParams contains all the parameters to send to the API endpoint
for the post remote API j credential some operation typically these are written to a http.Request
*/
type PostRemoteAPIJCredentialSomeParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j credential some params
func (o *PostRemoteAPIJCredentialSomeParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJCredentialSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j credential some params
func (o *PostRemoteAPIJCredentialSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j credential some params
func (o *PostRemoteAPIJCredentialSomeParams) WithContext(ctx context.Context) *PostRemoteAPIJCredentialSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j credential some params
func (o *PostRemoteAPIJCredentialSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j credential some params
func (o *PostRemoteAPIJCredentialSomeParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJCredentialSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j credential some params
func (o *PostRemoteAPIJCredentialSomeParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJCredentialSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
