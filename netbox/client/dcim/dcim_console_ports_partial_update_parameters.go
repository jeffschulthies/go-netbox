// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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
	"github.com/go-openapi/swag"

	"github.com/jeffschulthies/go-netbox/netbox/models"
)

// NewDcimConsolePortsPartialUpdateParams creates a new DcimConsolePortsPartialUpdateParams object
// with the default values initialized.
func NewDcimConsolePortsPartialUpdateParams() *DcimConsolePortsPartialUpdateParams {
	var ()
	return &DcimConsolePortsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsPartialUpdateParamsWithTimeout creates a new DcimConsolePortsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsPartialUpdateParams {
	var ()
	return &DcimConsolePortsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortsPartialUpdateParamsWithContext creates a new DcimConsolePortsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortsPartialUpdateParamsWithContext(ctx context.Context) *DcimConsolePortsPartialUpdateParams {
	var ()
	return &DcimConsolePortsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimConsolePortsPartialUpdateParamsWithHTTPClient creates a new DcimConsolePortsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsPartialUpdateParams {
	var ()
	return &DcimConsolePortsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim console ports partial update operation typically these are written to a http.Request
*/
type DcimConsolePortsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableConsolePort
	/*ID
	  A unique integer value identifying this console port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithContext(ctx context.Context) *DcimConsolePortsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WithID adds the id to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) WithID(id int64) *DcimConsolePortsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports partial update params
func (o *DcimConsolePortsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
