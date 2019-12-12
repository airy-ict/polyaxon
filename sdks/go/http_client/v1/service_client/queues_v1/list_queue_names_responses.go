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

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListQueueNamesReader is a Reader for the ListQueueNames structure.
type ListQueueNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListQueueNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListQueueNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListQueueNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListQueueNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListQueueNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListQueueNamesOK creates a ListQueueNamesOK with default headers values
func NewListQueueNamesOK() *ListQueueNamesOK {
	return &ListQueueNamesOK{}
}

/*ListQueueNamesOK handles this case with default header values.

A successful response.
*/
type ListQueueNamesOK struct {
	Payload *service_model.V1ListQueuesResponse
}

func (o *ListQueueNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{agent}/queues/names][%d] listQueueNamesOK  %+v", 200, o.Payload)
}

func (o *ListQueueNamesOK) GetPayload() *service_model.V1ListQueuesResponse {
	return o.Payload
}

func (o *ListQueueNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListQueuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueueNamesNoContent creates a ListQueueNamesNoContent with default headers values
func NewListQueueNamesNoContent() *ListQueueNamesNoContent {
	return &ListQueueNamesNoContent{}
}

/*ListQueueNamesNoContent handles this case with default header values.

No content.
*/
type ListQueueNamesNoContent struct {
	Payload interface{}
}

func (o *ListQueueNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{agent}/queues/names][%d] listQueueNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListQueueNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListQueueNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueueNamesForbidden creates a ListQueueNamesForbidden with default headers values
func NewListQueueNamesForbidden() *ListQueueNamesForbidden {
	return &ListQueueNamesForbidden{}
}

/*ListQueueNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListQueueNamesForbidden struct {
	Payload interface{}
}

func (o *ListQueueNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{agent}/queues/names][%d] listQueueNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListQueueNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListQueueNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueueNamesNotFound creates a ListQueueNamesNotFound with default headers values
func NewListQueueNamesNotFound() *ListQueueNamesNotFound {
	return &ListQueueNamesNotFound{}
}

/*ListQueueNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListQueueNamesNotFound struct {
	Payload interface{}
}

func (o *ListQueueNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{agent}/queues/names][%d] listQueueNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListQueueNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListQueueNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
