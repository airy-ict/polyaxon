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

package dashboard_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchDashboardReader is a Reader for the PatchDashboard structure.
type PatchDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchDashboardOK creates a PatchDashboardOK with default headers values
func NewPatchDashboardOK() *PatchDashboardOK {
	return &PatchDashboardOK{}
}

/*PatchDashboardOK handles this case with default header values.

A successful response.
*/
type PatchDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *PatchDashboardOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchDashboardOK  %+v", 200, o.Payload)
}

func (o *PatchDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *PatchDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDashboardNoContent creates a PatchDashboardNoContent with default headers values
func NewPatchDashboardNoContent() *PatchDashboardNoContent {
	return &PatchDashboardNoContent{}
}

/*PatchDashboardNoContent handles this case with default header values.

No content.
*/
type PatchDashboardNoContent struct {
	Payload interface{}
}

func (o *PatchDashboardNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchDashboardNoContent  %+v", 204, o.Payload)
}

func (o *PatchDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDashboardForbidden creates a PatchDashboardForbidden with default headers values
func NewPatchDashboardForbidden() *PatchDashboardForbidden {
	return &PatchDashboardForbidden{}
}

/*PatchDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchDashboardForbidden struct {
	Payload interface{}
}

func (o *PatchDashboardForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PatchDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDashboardNotFound creates a PatchDashboardNotFound with default headers values
func NewPatchDashboardNotFound() *PatchDashboardNotFound {
	return &PatchDashboardNotFound{}
}

/*PatchDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchDashboardNotFound struct {
	Payload interface{}
}

func (o *PatchDashboardNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/{project}/dashboards/{dashboard.uuid}][%d] patchDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PatchDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
