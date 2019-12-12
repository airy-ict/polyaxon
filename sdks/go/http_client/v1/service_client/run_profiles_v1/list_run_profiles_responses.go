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

package run_profiles_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListRunProfilesReader is a Reader for the ListRunProfiles structure.
type ListRunProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRunProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRunProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListRunProfilesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListRunProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRunProfilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRunProfilesOK creates a ListRunProfilesOK with default headers values
func NewListRunProfilesOK() *ListRunProfilesOK {
	return &ListRunProfilesOK{}
}

/*ListRunProfilesOK handles this case with default header values.

A successful response.
*/
type ListRunProfilesOK struct {
	Payload *service_model.V1ListRunProfilesResponse
}

func (o *ListRunProfilesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/run_profiles][%d] listRunProfilesOK  %+v", 200, o.Payload)
}

func (o *ListRunProfilesOK) GetPayload() *service_model.V1ListRunProfilesResponse {
	return o.Payload
}

func (o *ListRunProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunProfilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunProfilesNoContent creates a ListRunProfilesNoContent with default headers values
func NewListRunProfilesNoContent() *ListRunProfilesNoContent {
	return &ListRunProfilesNoContent{}
}

/*ListRunProfilesNoContent handles this case with default header values.

No content.
*/
type ListRunProfilesNoContent struct {
	Payload interface{}
}

func (o *ListRunProfilesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/run_profiles][%d] listRunProfilesNoContent  %+v", 204, o.Payload)
}

func (o *ListRunProfilesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunProfilesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunProfilesForbidden creates a ListRunProfilesForbidden with default headers values
func NewListRunProfilesForbidden() *ListRunProfilesForbidden {
	return &ListRunProfilesForbidden{}
}

/*ListRunProfilesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListRunProfilesForbidden struct {
	Payload interface{}
}

func (o *ListRunProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/run_profiles][%d] listRunProfilesForbidden  %+v", 403, o.Payload)
}

func (o *ListRunProfilesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunProfilesNotFound creates a ListRunProfilesNotFound with default headers values
func NewListRunProfilesNotFound() *ListRunProfilesNotFound {
	return &ListRunProfilesNotFound{}
}

/*ListRunProfilesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListRunProfilesNotFound struct {
	Payload interface{}
}

func (o *ListRunProfilesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/run_profiles][%d] listRunProfilesNotFound  %+v", 404, o.Payload)
}

func (o *ListRunProfilesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListRunProfilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
