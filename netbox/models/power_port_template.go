// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerPortTemplate Power port
// swagger:model PowerPortTemplate
type PowerPortTemplate struct {

	// Allocated draw
	//
	// Allocated current draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Maximum draw
	//
	// Maximum current draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this power port template
func (m *PowerPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximumDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerPortTemplate) validateAllocatedDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocatedDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *PowerPortTemplate) validateMaximumDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.MaximumDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("maximum_draw", "body", int64(*m.MaximumDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maximum_draw", "body", int64(*m.MaximumDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *PowerPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortTemplate) UnmarshalBinary(b []byte) error {
	var res PowerPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}