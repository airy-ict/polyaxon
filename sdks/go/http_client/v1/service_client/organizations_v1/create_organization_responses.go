// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/*CreateOrganizationOK handles this case with default header values.

A successful response.
*/
type CreateOrganizationOK struct {
	Payload *service_model.V1Organization
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) GetPayload() *service_model.V1Organization {
	return o.Payload
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationNoContent creates a CreateOrganizationNoContent with default headers values
func NewCreateOrganizationNoContent() *CreateOrganizationNoContent {
	return &CreateOrganizationNoContent{}
}

/*CreateOrganizationNoContent handles this case with default header values.

No content.
*/
type CreateOrganizationNoContent struct {
	Payload interface{}
}

func (o *CreateOrganizationNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationNoContent  %+v", 204, o.Payload)
}

func (o *CreateOrganizationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationForbidden creates a CreateOrganizationForbidden with default headers values
func NewCreateOrganizationForbidden() *CreateOrganizationForbidden {
	return &CreateOrganizationForbidden{}
}

/*CreateOrganizationForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateOrganizationForbidden struct {
	Payload interface{}
}

func (o *CreateOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrganizationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationNotFound creates a CreateOrganizationNotFound with default headers values
func NewCreateOrganizationNotFound() *CreateOrganizationNotFound {
	return &CreateOrganizationNotFound{}
}

/*CreateOrganizationNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateOrganizationNotFound struct {
	Payload interface{}
}

func (o *CreateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/create][%d] createOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrganizationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
