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

package registry_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListRegistryAccessNamesReader is a Reader for the ListRegistryAccessNames structure.
type ListRegistryAccessNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistryAccessNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistryAccessNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListRegistryAccessNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListRegistryAccessNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRegistryAccessNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRegistryAccessNamesOK creates a ListRegistryAccessNamesOK with default headers values
func NewListRegistryAccessNamesOK() *ListRegistryAccessNamesOK {
	return &ListRegistryAccessNamesOK{}
}

/*ListRegistryAccessNamesOK handles this case with default header values.

A successful response.
*/
type ListRegistryAccessNamesOK struct {
	Payload *service_model.V1ListHostAccessesResponse
}

func (o *ListRegistryAccessNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/registry_accesses/names][%d] listRegistryAccessNamesOK  %+v", 200, o.Payload)
}

func (o *ListRegistryAccessNamesOK) GetPayload() *service_model.V1ListHostAccessesResponse {
	return o.Payload
}

func (o *ListRegistryAccessNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListHostAccessesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryAccessNamesNoContent creates a ListRegistryAccessNamesNoContent with default headers values
func NewListRegistryAccessNamesNoContent() *ListRegistryAccessNamesNoContent {
	return &ListRegistryAccessNamesNoContent{}
}

/*ListRegistryAccessNamesNoContent handles this case with default header values.

No content.
*/
type ListRegistryAccessNamesNoContent struct {
	Payload interface{}
}

func (o *ListRegistryAccessNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/registry_accesses/names][%d] listRegistryAccessNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListRegistryAccessNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRegistryAccessNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryAccessNamesForbidden creates a ListRegistryAccessNamesForbidden with default headers values
func NewListRegistryAccessNamesForbidden() *ListRegistryAccessNamesForbidden {
	return &ListRegistryAccessNamesForbidden{}
}

/*ListRegistryAccessNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListRegistryAccessNamesForbidden struct {
	Payload interface{}
}

func (o *ListRegistryAccessNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/registry_accesses/names][%d] listRegistryAccessNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryAccessNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRegistryAccessNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryAccessNamesNotFound creates a ListRegistryAccessNamesNotFound with default headers values
func NewListRegistryAccessNamesNotFound() *ListRegistryAccessNamesNotFound {
	return &ListRegistryAccessNamesNotFound{}
}

/*ListRegistryAccessNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListRegistryAccessNamesNotFound struct {
	Payload interface{}
}

func (o *ListRegistryAccessNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/registry_accesses/names][%d] listRegistryAccessNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListRegistryAccessNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRegistryAccessNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
