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

package k8s_config_maps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchK8sConfigMapReader is a Reader for the PatchK8sConfigMap structure.
type PatchK8sConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchK8sConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchK8sConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchK8sConfigMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchK8sConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchK8sConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchK8sConfigMapOK creates a PatchK8sConfigMapOK with default headers values
func NewPatchK8sConfigMapOK() *PatchK8sConfigMapOK {
	return &PatchK8sConfigMapOK{}
}

/*PatchK8sConfigMapOK handles this case with default header values.

A successful response.
*/
type PatchK8sConfigMapOK struct {
	Payload *service_model.V1K8sResource
}

func (o *PatchK8sConfigMapOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8sConfigMapOK  %+v", 200, o.Payload)
}

func (o *PatchK8sConfigMapOK) GetPayload() *service_model.V1K8sResource {
	return o.Payload
}

func (o *PatchK8sConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8sResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchK8sConfigMapNoContent creates a PatchK8sConfigMapNoContent with default headers values
func NewPatchK8sConfigMapNoContent() *PatchK8sConfigMapNoContent {
	return &PatchK8sConfigMapNoContent{}
}

/*PatchK8sConfigMapNoContent handles this case with default header values.

No content.
*/
type PatchK8sConfigMapNoContent struct {
	Payload interface{}
}

func (o *PatchK8sConfigMapNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8sConfigMapNoContent  %+v", 204, o.Payload)
}

func (o *PatchK8sConfigMapNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchK8sConfigMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchK8sConfigMapForbidden creates a PatchK8sConfigMapForbidden with default headers values
func NewPatchK8sConfigMapForbidden() *PatchK8sConfigMapForbidden {
	return &PatchK8sConfigMapForbidden{}
}

/*PatchK8sConfigMapForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchK8sConfigMapForbidden struct {
	Payload interface{}
}

func (o *PatchK8sConfigMapForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8sConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *PatchK8sConfigMapForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchK8sConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchK8sConfigMapNotFound creates a PatchK8sConfigMapNotFound with default headers values
func NewPatchK8sConfigMapNotFound() *PatchK8sConfigMapNotFound {
	return &PatchK8sConfigMapNotFound{}
}

/*PatchK8sConfigMapNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchK8sConfigMapNotFound struct {
	Payload interface{}
}

func (o *PatchK8sConfigMapNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8sConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *PatchK8sConfigMapNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchK8sConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
