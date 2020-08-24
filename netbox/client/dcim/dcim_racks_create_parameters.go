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

	"github.com/jeffschulthies/go-netbox/netbox/models"
)

// NewDcimRacksCreateParams creates a new DcimRacksCreateParams object
// with the default values initialized.
func NewDcimRacksCreateParams() *DcimRacksCreateParams {
	var ()
	return &DcimRacksCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksCreateParamsWithTimeout creates a new DcimRacksCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksCreateParamsWithTimeout(timeout time.Duration) *DcimRacksCreateParams {
	var ()
	return &DcimRacksCreateParams{

		timeout: timeout,
	}
}

// NewDcimRacksCreateParamsWithContext creates a new DcimRacksCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksCreateParamsWithContext(ctx context.Context) *DcimRacksCreateParams {
	var ()
	return &DcimRacksCreateParams{

		Context: ctx,
	}
}

// NewDcimRacksCreateParamsWithHTTPClient creates a new DcimRacksCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksCreateParamsWithHTTPClient(client *http.Client) *DcimRacksCreateParams {
	var ()
	return &DcimRacksCreateParams{
		HTTPClient: client,
	}
}

/*DcimRacksCreateParams contains all the parameters to send to the API endpoint
for the dcim racks create operation typically these are written to a http.Request
*/
type DcimRacksCreateParams struct {

	/*Data*/
	Data *models.WritableRack

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks create params
func (o *DcimRacksCreateParams) WithTimeout(timeout time.Duration) *DcimRacksCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks create params
func (o *DcimRacksCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks create params
func (o *DcimRacksCreateParams) WithContext(ctx context.Context) *DcimRacksCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks create params
func (o *DcimRacksCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks create params
func (o *DcimRacksCreateParams) WithHTTPClient(client *http.Client) *DcimRacksCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks create params
func (o *DcimRacksCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks create params
func (o *DcimRacksCreateParams) WithData(data *models.WritableRack) *DcimRacksCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks create params
func (o *DcimRacksCreateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
