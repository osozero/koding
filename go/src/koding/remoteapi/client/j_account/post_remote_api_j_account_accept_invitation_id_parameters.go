package j_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJAccountAcceptInvitationIDParams creates a new PostRemoteAPIJAccountAcceptInvitationIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountAcceptInvitationIDParams() *PostRemoteAPIJAccountAcceptInvitationIDParams {
	var ()
	return &PostRemoteAPIJAccountAcceptInvitationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountAcceptInvitationIDParamsWithTimeout creates a new PostRemoteAPIJAccountAcceptInvitationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountAcceptInvitationIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountAcceptInvitationIDParams {
	var ()
	return &PostRemoteAPIJAccountAcceptInvitationIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountAcceptInvitationIDParamsWithContext creates a new PostRemoteAPIJAccountAcceptInvitationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountAcceptInvitationIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountAcceptInvitationIDParams {
	var ()
	return &PostRemoteAPIJAccountAcceptInvitationIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountAcceptInvitationIDParams contains all the parameters to send to the API endpoint
for the post remote API j account accept invitation ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountAcceptInvitationIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j account accept invitation ID params
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountAcceptInvitationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account accept invitation ID params
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account accept invitation ID params
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountAcceptInvitationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account accept invitation ID params
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account accept invitation ID params
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) WithID(id string) *PostRemoteAPIJAccountAcceptInvitationIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account accept invitation ID params
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountAcceptInvitationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
