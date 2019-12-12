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

// PatchRunProfileReader is a Reader for the PatchRunProfile structure.
type PatchRunProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRunProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRunProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchRunProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchRunProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRunProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchRunProfileOK creates a PatchRunProfileOK with default headers values
func NewPatchRunProfileOK() *PatchRunProfileOK {
	return &PatchRunProfileOK{}
}

/*PatchRunProfileOK handles this case with default header values.

A successful response.
*/
type PatchRunProfileOK struct {
	Payload *service_model.V1RunProfile
}

func (o *PatchRunProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] patchRunProfileOK  %+v", 200, o.Payload)
}

func (o *PatchRunProfileOK) GetPayload() *service_model.V1RunProfile {
	return o.Payload
}

func (o *PatchRunProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1RunProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRunProfileNoContent creates a PatchRunProfileNoContent with default headers values
func NewPatchRunProfileNoContent() *PatchRunProfileNoContent {
	return &PatchRunProfileNoContent{}
}

/*PatchRunProfileNoContent handles this case with default header values.

No content.
*/
type PatchRunProfileNoContent struct {
	Payload interface{}
}

func (o *PatchRunProfileNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] patchRunProfileNoContent  %+v", 204, o.Payload)
}

func (o *PatchRunProfileNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchRunProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRunProfileForbidden creates a PatchRunProfileForbidden with default headers values
func NewPatchRunProfileForbidden() *PatchRunProfileForbidden {
	return &PatchRunProfileForbidden{}
}

/*PatchRunProfileForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchRunProfileForbidden struct {
	Payload interface{}
}

func (o *PatchRunProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] patchRunProfileForbidden  %+v", 403, o.Payload)
}

func (o *PatchRunProfileForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchRunProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRunProfileNotFound creates a PatchRunProfileNotFound with default headers values
func NewPatchRunProfileNotFound() *PatchRunProfileNotFound {
	return &PatchRunProfileNotFound{}
}

/*PatchRunProfileNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchRunProfileNotFound struct {
	Payload interface{}
}

func (o *PatchRunProfileNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}][%d] patchRunProfileNotFound  %+v", 404, o.Payload)
}

func (o *PatchRunProfileNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchRunProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
