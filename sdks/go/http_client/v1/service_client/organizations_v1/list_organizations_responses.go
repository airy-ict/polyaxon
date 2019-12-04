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

// ListOrganizationsReader is a Reader for the ListOrganizations structure.
type ListOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListOrganizationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListOrganizationsOK creates a ListOrganizationsOK with default headers values
func NewListOrganizationsOK() *ListOrganizationsOK {
	return &ListOrganizationsOK{}
}

/*ListOrganizationsOK handles this case with default header values.

A successful response.
*/
type ListOrganizationsOK struct {
	Payload *service_model.V1ListOrganizationsResponse
}

func (o *ListOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/list][%d] listOrganizationsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationsOK) GetPayload() *service_model.V1ListOrganizationsResponse {
	return o.Payload
}

func (o *ListOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListOrganizationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsNoContent creates a ListOrganizationsNoContent with default headers values
func NewListOrganizationsNoContent() *ListOrganizationsNoContent {
	return &ListOrganizationsNoContent{}
}

/*ListOrganizationsNoContent handles this case with default header values.

No content.
*/
type ListOrganizationsNoContent struct {
	Payload interface{}
}

func (o *ListOrganizationsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/list][%d] listOrganizationsNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsForbidden creates a ListOrganizationsForbidden with default headers values
func NewListOrganizationsForbidden() *ListOrganizationsForbidden {
	return &ListOrganizationsForbidden{}
}

/*ListOrganizationsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListOrganizationsForbidden struct {
	Payload interface{}
}

func (o *ListOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/list][%d] listOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsNotFound creates a ListOrganizationsNotFound with default headers values
func NewListOrganizationsNotFound() *ListOrganizationsNotFound {
	return &ListOrganizationsNotFound{}
}

/*ListOrganizationsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListOrganizationsNotFound struct {
	Payload interface{}
}

func (o *ListOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/list][%d] listOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
