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

package ipam

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

// NewIpamRirsUpdateParams creates a new IpamRirsUpdateParams object
// with the default values initialized.
func NewIpamRirsUpdateParams() *IpamRirsUpdateParams {
	var ()
	return &IpamRirsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsUpdateParamsWithTimeout creates a new IpamRirsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamRirsUpdateParamsWithTimeout(timeout time.Duration) *IpamRirsUpdateParams {
	var ()
	return &IpamRirsUpdateParams{

		timeout: timeout,
	}
}

// NewIpamRirsUpdateParamsWithContext creates a new IpamRirsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamRirsUpdateParamsWithContext(ctx context.Context) *IpamRirsUpdateParams {
	var ()
	return &IpamRirsUpdateParams{

		Context: ctx,
	}
}

// NewIpamRirsUpdateParamsWithHTTPClient creates a new IpamRirsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamRirsUpdateParamsWithHTTPClient(client *http.Client) *IpamRirsUpdateParams {
	var ()
	return &IpamRirsUpdateParams{
		HTTPClient: client,
	}
}

/*IpamRirsUpdateParams contains all the parameters to send to the API endpoint
for the ipam rirs update operation typically these are written to a http.Request
*/
type IpamRirsUpdateParams struct {

	/*Data*/
	Data *models.RIR
	/*ID
	  A unique integer value identifying this RIR.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithTimeout(timeout time.Duration) *IpamRirsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithContext(ctx context.Context) *IpamRirsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithHTTPClient(client *http.Client) *IpamRirsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithData(data *models.RIR) *IpamRirsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetData(data *models.RIR) {
	o.Data = data
}

// WithID adds the id to the ipam rirs update params
func (o *IpamRirsUpdateParams) WithID(id int64) *IpamRirsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam rirs update params
func (o *IpamRirsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
