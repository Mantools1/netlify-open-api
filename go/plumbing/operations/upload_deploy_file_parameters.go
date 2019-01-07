// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadDeployFileParams creates a new UploadDeployFileParams object
// with the default values initialized.
func NewUploadDeployFileParams() *UploadDeployFileParams {
	var ()
	return &UploadDeployFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDeployFileParamsWithTimeout creates a new UploadDeployFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadDeployFileParamsWithTimeout(timeout time.Duration) *UploadDeployFileParams {
	var ()
	return &UploadDeployFileParams{

		timeout: timeout,
	}
}

// NewUploadDeployFileParamsWithContext creates a new UploadDeployFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadDeployFileParamsWithContext(ctx context.Context) *UploadDeployFileParams {
	var ()
	return &UploadDeployFileParams{

		Context: ctx,
	}
}

// NewUploadDeployFileParamsWithHTTPClient creates a new UploadDeployFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadDeployFileParamsWithHTTPClient(client *http.Client) *UploadDeployFileParams {
	var ()
	return &UploadDeployFileParams{
		HTTPClient: client,
	}
}

/*UploadDeployFileParams contains all the parameters to send to the API endpoint
for the upload deploy file operation typically these are written to a http.Request
*/
type UploadDeployFileParams struct {

	/*DeployID*/
	DeployID string
	/*FileBody*/
	FileBody io.ReadCloser
	/*Path*/
	Path string
	/*Size*/
	Size *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload deploy file params
func (o *UploadDeployFileParams) WithTimeout(timeout time.Duration) *UploadDeployFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload deploy file params
func (o *UploadDeployFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload deploy file params
func (o *UploadDeployFileParams) WithContext(ctx context.Context) *UploadDeployFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload deploy file params
func (o *UploadDeployFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload deploy file params
func (o *UploadDeployFileParams) WithHTTPClient(client *http.Client) *UploadDeployFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload deploy file params
func (o *UploadDeployFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the upload deploy file params
func (o *UploadDeployFileParams) WithDeployID(deployID string) *UploadDeployFileParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the upload deploy file params
func (o *UploadDeployFileParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WithFileBody adds the fileBody to the upload deploy file params
func (o *UploadDeployFileParams) WithFileBody(fileBody io.ReadCloser) *UploadDeployFileParams {
	o.SetFileBody(fileBody)
	return o
}

// SetFileBody adds the fileBody to the upload deploy file params
func (o *UploadDeployFileParams) SetFileBody(fileBody io.ReadCloser) {
	o.FileBody = fileBody
}

// WithPath adds the path to the upload deploy file params
func (o *UploadDeployFileParams) WithPath(path string) *UploadDeployFileParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the upload deploy file params
func (o *UploadDeployFileParams) SetPath(path string) {
	o.Path = path
}

// WithSize adds the size to the upload deploy file params
func (o *UploadDeployFileParams) WithSize(size *int64) *UploadDeployFileParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the upload deploy file params
func (o *UploadDeployFileParams) SetSize(size *int64) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDeployFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if o.FileBody != nil {
		if err := r.SetBodyParam(o.FileBody); err != nil {
			return err
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if o.Size != nil {

		// query param size
		var qrSize int64
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
