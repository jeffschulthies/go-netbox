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

package tenancy

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

	models "github.com/netbox-community/go-netbox/netbox/models"
)

// NewTenancyTenantGroupsCreateParams creates a new TenancyTenantGroupsCreateParams object
// with the default values initialized.
func NewTenancyTenantGroupsCreateParams() *TenancyTenantGroupsCreateParams {
	var ()
	return &TenancyTenantGroupsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsCreateParamsWithTimeout creates a new TenancyTenantGroupsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTenancyTenantGroupsCreateParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsCreateParams {
	var ()
	return &TenancyTenantGroupsCreateParams{

		timeout: timeout,
	}
}

// NewTenancyTenantGroupsCreateParamsWithContext creates a new TenancyTenantGroupsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTenancyTenantGroupsCreateParamsWithContext(ctx context.Context) *TenancyTenantGroupsCreateParams {
	var ()
	return &TenancyTenantGroupsCreateParams{

		Context: ctx,
	}
}

// NewTenancyTenantGroupsCreateParamsWithHTTPClient creates a new TenancyTenantGroupsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTenancyTenantGroupsCreateParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsCreateParams {
	var ()
	return &TenancyTenantGroupsCreateParams{
		HTTPClient: client,
	}
}

/*TenancyTenantGroupsCreateParams contains all the parameters to send to the API endpoint
for the tenancy tenant groups create operation typically these are written to a http.Request
*/
type TenancyTenantGroupsCreateParams struct {

	/*Data*/
	Data *models.TenantGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithContext(ctx context.Context) *TenancyTenantGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithData(data *models.TenantGroup) *TenancyTenantGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetData(data *models.TenantGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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