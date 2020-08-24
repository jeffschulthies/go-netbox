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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerPortTemplate power port template
//
// swagger:model PowerPortTemplate
type PowerPortTemplate struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
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
	// Maximum power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// type
	Type *PowerPortTemplateType `json:"type,omitempty"`
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

	if err := m.validateType(formats); err != nil {
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

func (m *PowerPortTemplate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
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

// PowerPortTemplateType Type
//
// swagger:model PowerPortTemplateType
type PowerPortTemplateType struct {

	// label
	// Required: true
	// Enum: [C6 C8 C14 C16 C20 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 5-15P NEMA 5-20P NEMA 5-30P NEMA 5-50P NEMA 6-15P NEMA 6-20P NEMA 6-30P NEMA 6-50P NEMA L5-15P NEMA L5-20P NEMA L5-30P NEMA L6-15P NEMA L6-20P NEMA L6-30P NEMA L6-50P CS6361C CS6365C CS8165C CS8265C CS8365C CS8465C ITA Type E (CEE 7/5) ITA Type F (CEE 7/4) ITA Type E/F (CEE 7/7) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-20p nema-l6-30p nema-l6-50p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o]
	Value *string `json:"value"`
}

// Validate validates this power port template type
func (m *PowerPortTemplateType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C6","C8","C14","C16","C20","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 5-15P","NEMA 5-20P","NEMA 5-30P","NEMA 5-50P","NEMA 6-15P","NEMA 6-20P","NEMA 6-30P","NEMA 6-50P","NEMA L5-15P","NEMA L5-20P","NEMA L5-30P","NEMA L6-15P","NEMA L6-20P","NEMA L6-30P","NEMA L6-50P","CS6361C","CS6365C","CS8165C","CS8265C","CS8365C","CS8465C","ITA Type E (CEE 7/5)","ITA Type F (CEE 7/4)","ITA Type E/F (CEE 7/7)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTemplateTypeTypeLabelPropEnum = append(powerPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerPortTemplateTypeLabelC6 captures enum value "C6"
	PowerPortTemplateTypeLabelC6 string = "C6"

	// PowerPortTemplateTypeLabelC8 captures enum value "C8"
	PowerPortTemplateTypeLabelC8 string = "C8"

	// PowerPortTemplateTypeLabelC14 captures enum value "C14"
	PowerPortTemplateTypeLabelC14 string = "C14"

	// PowerPortTemplateTypeLabelC16 captures enum value "C16"
	PowerPortTemplateTypeLabelC16 string = "C16"

	// PowerPortTemplateTypeLabelC20 captures enum value "C20"
	PowerPortTemplateTypeLabelC20 string = "C20"

	// PowerPortTemplateTypeLabelPNE4H captures enum value "P+N+E 4H"
	PowerPortTemplateTypeLabelPNE4H string = "P+N+E 4H"

	// PowerPortTemplateTypeLabelPNE6H captures enum value "P+N+E 6H"
	PowerPortTemplateTypeLabelPNE6H string = "P+N+E 6H"

	// PowerPortTemplateTypeLabelPNE9H captures enum value "P+N+E 9H"
	PowerPortTemplateTypeLabelPNE9H string = "P+N+E 9H"

	// PowerPortTemplateTypeLabelNr2PE4H captures enum value "2P+E 4H"
	PowerPortTemplateTypeLabelNr2PE4H string = "2P+E 4H"

	// PowerPortTemplateTypeLabelNr2PE6H captures enum value "2P+E 6H"
	PowerPortTemplateTypeLabelNr2PE6H string = "2P+E 6H"

	// PowerPortTemplateTypeLabelNr2PE9H captures enum value "2P+E 9H"
	PowerPortTemplateTypeLabelNr2PE9H string = "2P+E 9H"

	// PowerPortTemplateTypeLabelNr3PE4H captures enum value "3P+E 4H"
	PowerPortTemplateTypeLabelNr3PE4H string = "3P+E 4H"

	// PowerPortTemplateTypeLabelNr3PE6H captures enum value "3P+E 6H"
	PowerPortTemplateTypeLabelNr3PE6H string = "3P+E 6H"

	// PowerPortTemplateTypeLabelNr3PE9H captures enum value "3P+E 9H"
	PowerPortTemplateTypeLabelNr3PE9H string = "3P+E 9H"

	// PowerPortTemplateTypeLabelNr3PNE4H captures enum value "3P+N+E 4H"
	PowerPortTemplateTypeLabelNr3PNE4H string = "3P+N+E 4H"

	// PowerPortTemplateTypeLabelNr3PNE6H captures enum value "3P+N+E 6H"
	PowerPortTemplateTypeLabelNr3PNE6H string = "3P+N+E 6H"

	// PowerPortTemplateTypeLabelNr3PNE9H captures enum value "3P+N+E 9H"
	PowerPortTemplateTypeLabelNr3PNE9H string = "3P+N+E 9H"

	// PowerPortTemplateTypeLabelNEMA515P captures enum value "NEMA 5-15P"
	PowerPortTemplateTypeLabelNEMA515P string = "NEMA 5-15P"

	// PowerPortTemplateTypeLabelNEMA520P captures enum value "NEMA 5-20P"
	PowerPortTemplateTypeLabelNEMA520P string = "NEMA 5-20P"

	// PowerPortTemplateTypeLabelNEMA530P captures enum value "NEMA 5-30P"
	PowerPortTemplateTypeLabelNEMA530P string = "NEMA 5-30P"

	// PowerPortTemplateTypeLabelNEMA550P captures enum value "NEMA 5-50P"
	PowerPortTemplateTypeLabelNEMA550P string = "NEMA 5-50P"

	// PowerPortTemplateTypeLabelNEMA615P captures enum value "NEMA 6-15P"
	PowerPortTemplateTypeLabelNEMA615P string = "NEMA 6-15P"

	// PowerPortTemplateTypeLabelNEMA620P captures enum value "NEMA 6-20P"
	PowerPortTemplateTypeLabelNEMA620P string = "NEMA 6-20P"

	// PowerPortTemplateTypeLabelNEMA630P captures enum value "NEMA 6-30P"
	PowerPortTemplateTypeLabelNEMA630P string = "NEMA 6-30P"

	// PowerPortTemplateTypeLabelNEMA650P captures enum value "NEMA 6-50P"
	PowerPortTemplateTypeLabelNEMA650P string = "NEMA 6-50P"

	// PowerPortTemplateTypeLabelNEMAL515P captures enum value "NEMA L5-15P"
	PowerPortTemplateTypeLabelNEMAL515P string = "NEMA L5-15P"

	// PowerPortTemplateTypeLabelNEMAL520P captures enum value "NEMA L5-20P"
	PowerPortTemplateTypeLabelNEMAL520P string = "NEMA L5-20P"

	// PowerPortTemplateTypeLabelNEMAL530P captures enum value "NEMA L5-30P"
	PowerPortTemplateTypeLabelNEMAL530P string = "NEMA L5-30P"

	// PowerPortTemplateTypeLabelNEMAL615P captures enum value "NEMA L6-15P"
	PowerPortTemplateTypeLabelNEMAL615P string = "NEMA L6-15P"

	// PowerPortTemplateTypeLabelNEMAL620P captures enum value "NEMA L6-20P"
	PowerPortTemplateTypeLabelNEMAL620P string = "NEMA L6-20P"

	// PowerPortTemplateTypeLabelNEMAL630P captures enum value "NEMA L6-30P"
	PowerPortTemplateTypeLabelNEMAL630P string = "NEMA L6-30P"

	// PowerPortTemplateTypeLabelNEMAL650P captures enum value "NEMA L6-50P"
	PowerPortTemplateTypeLabelNEMAL650P string = "NEMA L6-50P"

	// PowerPortTemplateTypeLabelCS6361C captures enum value "CS6361C"
	PowerPortTemplateTypeLabelCS6361C string = "CS6361C"

	// PowerPortTemplateTypeLabelCS6365C captures enum value "CS6365C"
	PowerPortTemplateTypeLabelCS6365C string = "CS6365C"

	// PowerPortTemplateTypeLabelCS8165C captures enum value "CS8165C"
	PowerPortTemplateTypeLabelCS8165C string = "CS8165C"

	// PowerPortTemplateTypeLabelCS8265C captures enum value "CS8265C"
	PowerPortTemplateTypeLabelCS8265C string = "CS8265C"

	// PowerPortTemplateTypeLabelCS8365C captures enum value "CS8365C"
	PowerPortTemplateTypeLabelCS8365C string = "CS8365C"

	// PowerPortTemplateTypeLabelCS8465C captures enum value "CS8465C"
	PowerPortTemplateTypeLabelCS8465C string = "CS8465C"

	// PowerPortTemplateTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE 7/5)"
	PowerPortTemplateTypeLabelITATypeECEE75 string = "ITA Type E (CEE 7/5)"

	// PowerPortTemplateTypeLabelITATypeFCEE74 captures enum value "ITA Type F (CEE 7/4)"
	PowerPortTemplateTypeLabelITATypeFCEE74 string = "ITA Type F (CEE 7/4)"

	// PowerPortTemplateTypeLabelITATypeEFCEE77 captures enum value "ITA Type E/F (CEE 7/7)"
	PowerPortTemplateTypeLabelITATypeEFCEE77 string = "ITA Type E/F (CEE 7/7)"

	// PowerPortTemplateTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerPortTemplateTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerPortTemplateTypeLabelITATypeH captures enum value "ITA Type H"
	PowerPortTemplateTypeLabelITATypeH string = "ITA Type H"

	// PowerPortTemplateTypeLabelITATypeI captures enum value "ITA Type I"
	PowerPortTemplateTypeLabelITATypeI string = "ITA Type I"

	// PowerPortTemplateTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerPortTemplateTypeLabelITATypeJ string = "ITA Type J"

	// PowerPortTemplateTypeLabelITATypeK captures enum value "ITA Type K"
	PowerPortTemplateTypeLabelITATypeK string = "ITA Type K"

	// PowerPortTemplateTypeLabelITATypeLCEI2350 captures enum value "ITA Type L (CEI 23-50)"
	PowerPortTemplateTypeLabelITATypeLCEI2350 string = "ITA Type L (CEI 23-50)"

	// PowerPortTemplateTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerPortTemplateTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerPortTemplateTypeLabelITATypeN captures enum value "ITA Type N"
	PowerPortTemplateTypeLabelITATypeN string = "ITA Type N"

	// PowerPortTemplateTypeLabelITATypeO captures enum value "ITA Type O"
	PowerPortTemplateTypeLabelITATypeO string = "ITA Type O"
)

// prop value enum
func (m *PowerPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-20p","nema-l6-30p","nema-l6-50p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPortTemplateTypeTypeValuePropEnum = append(powerPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerPortTemplateTypeValueIec60320C6 captures enum value "iec-60320-c6"
	PowerPortTemplateTypeValueIec60320C6 string = "iec-60320-c6"

	// PowerPortTemplateTypeValueIec60320C8 captures enum value "iec-60320-c8"
	PowerPortTemplateTypeValueIec60320C8 string = "iec-60320-c8"

	// PowerPortTemplateTypeValueIec60320C14 captures enum value "iec-60320-c14"
	PowerPortTemplateTypeValueIec60320C14 string = "iec-60320-c14"

	// PowerPortTemplateTypeValueIec60320C16 captures enum value "iec-60320-c16"
	PowerPortTemplateTypeValueIec60320C16 string = "iec-60320-c16"

	// PowerPortTemplateTypeValueIec60320C20 captures enum value "iec-60320-c20"
	PowerPortTemplateTypeValueIec60320C20 string = "iec-60320-c20"

	// PowerPortTemplateTypeValueIec60309pne4h captures enum value "iec-60309-p-n-e-4h"
	PowerPortTemplateTypeValueIec60309pne4h string = "iec-60309-p-n-e-4h"

	// PowerPortTemplateTypeValueIec60309pne6h captures enum value "iec-60309-p-n-e-6h"
	PowerPortTemplateTypeValueIec60309pne6h string = "iec-60309-p-n-e-6h"

	// PowerPortTemplateTypeValueIec60309pne9h captures enum value "iec-60309-p-n-e-9h"
	PowerPortTemplateTypeValueIec60309pne9h string = "iec-60309-p-n-e-9h"

	// PowerPortTemplateTypeValueIec603092pe4h captures enum value "iec-60309-2p-e-4h"
	PowerPortTemplateTypeValueIec603092pe4h string = "iec-60309-2p-e-4h"

	// PowerPortTemplateTypeValueIec603092pe6h captures enum value "iec-60309-2p-e-6h"
	PowerPortTemplateTypeValueIec603092pe6h string = "iec-60309-2p-e-6h"

	// PowerPortTemplateTypeValueIec603092pe9h captures enum value "iec-60309-2p-e-9h"
	PowerPortTemplateTypeValueIec603092pe9h string = "iec-60309-2p-e-9h"

	// PowerPortTemplateTypeValueIec603093pe4h captures enum value "iec-60309-3p-e-4h"
	PowerPortTemplateTypeValueIec603093pe4h string = "iec-60309-3p-e-4h"

	// PowerPortTemplateTypeValueIec603093pe6h captures enum value "iec-60309-3p-e-6h"
	PowerPortTemplateTypeValueIec603093pe6h string = "iec-60309-3p-e-6h"

	// PowerPortTemplateTypeValueIec603093pe9h captures enum value "iec-60309-3p-e-9h"
	PowerPortTemplateTypeValueIec603093pe9h string = "iec-60309-3p-e-9h"

	// PowerPortTemplateTypeValueIec603093pne4h captures enum value "iec-60309-3p-n-e-4h"
	PowerPortTemplateTypeValueIec603093pne4h string = "iec-60309-3p-n-e-4h"

	// PowerPortTemplateTypeValueIec603093pne6h captures enum value "iec-60309-3p-n-e-6h"
	PowerPortTemplateTypeValueIec603093pne6h string = "iec-60309-3p-n-e-6h"

	// PowerPortTemplateTypeValueIec603093pne9h captures enum value "iec-60309-3p-n-e-9h"
	PowerPortTemplateTypeValueIec603093pne9h string = "iec-60309-3p-n-e-9h"

	// PowerPortTemplateTypeValueNema515p captures enum value "nema-5-15p"
	PowerPortTemplateTypeValueNema515p string = "nema-5-15p"

	// PowerPortTemplateTypeValueNema520p captures enum value "nema-5-20p"
	PowerPortTemplateTypeValueNema520p string = "nema-5-20p"

	// PowerPortTemplateTypeValueNema530p captures enum value "nema-5-30p"
	PowerPortTemplateTypeValueNema530p string = "nema-5-30p"

	// PowerPortTemplateTypeValueNema550p captures enum value "nema-5-50p"
	PowerPortTemplateTypeValueNema550p string = "nema-5-50p"

	// PowerPortTemplateTypeValueNema615p captures enum value "nema-6-15p"
	PowerPortTemplateTypeValueNema615p string = "nema-6-15p"

	// PowerPortTemplateTypeValueNema620p captures enum value "nema-6-20p"
	PowerPortTemplateTypeValueNema620p string = "nema-6-20p"

	// PowerPortTemplateTypeValueNema630p captures enum value "nema-6-30p"
	PowerPortTemplateTypeValueNema630p string = "nema-6-30p"

	// PowerPortTemplateTypeValueNema650p captures enum value "nema-6-50p"
	PowerPortTemplateTypeValueNema650p string = "nema-6-50p"

	// PowerPortTemplateTypeValueNemaL515p captures enum value "nema-l5-15p"
	PowerPortTemplateTypeValueNemaL515p string = "nema-l5-15p"

	// PowerPortTemplateTypeValueNemaL520p captures enum value "nema-l5-20p"
	PowerPortTemplateTypeValueNemaL520p string = "nema-l5-20p"

	// PowerPortTemplateTypeValueNemaL530p captures enum value "nema-l5-30p"
	PowerPortTemplateTypeValueNemaL530p string = "nema-l5-30p"

	// PowerPortTemplateTypeValueNemaL550p captures enum value "nema-l5-50p"
	PowerPortTemplateTypeValueNemaL550p string = "nema-l5-50p"

	// PowerPortTemplateTypeValueNemaL620p captures enum value "nema-l6-20p"
	PowerPortTemplateTypeValueNemaL620p string = "nema-l6-20p"

	// PowerPortTemplateTypeValueNemaL630p captures enum value "nema-l6-30p"
	PowerPortTemplateTypeValueNemaL630p string = "nema-l6-30p"

	// PowerPortTemplateTypeValueNemaL650p captures enum value "nema-l6-50p"
	PowerPortTemplateTypeValueNemaL650p string = "nema-l6-50p"

	// PowerPortTemplateTypeValueCs6361c captures enum value "cs6361c"
	PowerPortTemplateTypeValueCs6361c string = "cs6361c"

	// PowerPortTemplateTypeValueCs6365c captures enum value "cs6365c"
	PowerPortTemplateTypeValueCs6365c string = "cs6365c"

	// PowerPortTemplateTypeValueCs8165c captures enum value "cs8165c"
	PowerPortTemplateTypeValueCs8165c string = "cs8165c"

	// PowerPortTemplateTypeValueCs8265c captures enum value "cs8265c"
	PowerPortTemplateTypeValueCs8265c string = "cs8265c"

	// PowerPortTemplateTypeValueCs8365c captures enum value "cs8365c"
	PowerPortTemplateTypeValueCs8365c string = "cs8365c"

	// PowerPortTemplateTypeValueCs8465c captures enum value "cs8465c"
	PowerPortTemplateTypeValueCs8465c string = "cs8465c"

	// PowerPortTemplateTypeValueItae captures enum value "ita-e"
	PowerPortTemplateTypeValueItae string = "ita-e"

	// PowerPortTemplateTypeValueItaf captures enum value "ita-f"
	PowerPortTemplateTypeValueItaf string = "ita-f"

	// PowerPortTemplateTypeValueItaEf captures enum value "ita-ef"
	PowerPortTemplateTypeValueItaEf string = "ita-ef"

	// PowerPortTemplateTypeValueItag captures enum value "ita-g"
	PowerPortTemplateTypeValueItag string = "ita-g"

	// PowerPortTemplateTypeValueItah captures enum value "ita-h"
	PowerPortTemplateTypeValueItah string = "ita-h"

	// PowerPortTemplateTypeValueItai captures enum value "ita-i"
	PowerPortTemplateTypeValueItai string = "ita-i"

	// PowerPortTemplateTypeValueItaj captures enum value "ita-j"
	PowerPortTemplateTypeValueItaj string = "ita-j"

	// PowerPortTemplateTypeValueItak captures enum value "ita-k"
	PowerPortTemplateTypeValueItak string = "ita-k"

	// PowerPortTemplateTypeValueItal captures enum value "ita-l"
	PowerPortTemplateTypeValueItal string = "ita-l"

	// PowerPortTemplateTypeValueItam captures enum value "ita-m"
	PowerPortTemplateTypeValueItam string = "ita-m"

	// PowerPortTemplateTypeValueItan captures enum value "ita-n"
	PowerPortTemplateTypeValueItan string = "ita-n"

	// PowerPortTemplateTypeValueItao captures enum value "ita-o"
	PowerPortTemplateTypeValueItao string = "ita-o"
)

// prop value enum
func (m *PowerPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerPortTemplateType) UnmarshalBinary(b []byte) error {
	var res PowerPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
