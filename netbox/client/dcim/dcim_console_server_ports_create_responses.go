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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jeffschulthies/go-netbox/netbox/models"
)

// DcimConsoleServerPortsCreateReader is a Reader for the DcimConsoleServerPortsCreate structure.
type DcimConsoleServerPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimConsoleServerPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsCreateCreated creates a DcimConsoleServerPortsCreateCreated with default headers values
func NewDcimConsoleServerPortsCreateCreated() *DcimConsoleServerPortsCreateCreated {
	return &DcimConsoleServerPortsCreateCreated{}
}

/*DcimConsoleServerPortsCreateCreated handles this case with default header values.

DcimConsoleServerPortsCreateCreated dcim console server ports create created
*/
type DcimConsoleServerPortsCreateCreated struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/console-server-ports/][%d] dcimConsoleServerPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimConsoleServerPortsCreateCreated) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
