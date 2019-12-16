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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewPatchProjectSearchParams creates a new PatchProjectSearchParams object
// with the default values initialized.
func NewPatchProjectSearchParams() *PatchProjectSearchParams {
	var ()
	return &PatchProjectSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchProjectSearchParamsWithTimeout creates a new PatchProjectSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchProjectSearchParamsWithTimeout(timeout time.Duration) *PatchProjectSearchParams {
	var ()
	return &PatchProjectSearchParams{

		timeout: timeout,
	}
}

// NewPatchProjectSearchParamsWithContext creates a new PatchProjectSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchProjectSearchParamsWithContext(ctx context.Context) *PatchProjectSearchParams {
	var ()
	return &PatchProjectSearchParams{

		Context: ctx,
	}
}

// NewPatchProjectSearchParamsWithHTTPClient creates a new PatchProjectSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchProjectSearchParamsWithHTTPClient(client *http.Client) *PatchProjectSearchParams {
	var ()
	return &PatchProjectSearchParams{
		HTTPClient: client,
	}
}

/*PatchProjectSearchParams contains all the parameters to send to the API endpoint
for the patch project search operation typically these are written to a http.Request
*/
type PatchProjectSearchParams struct {

	/*Body
	  Search body

	*/
	Body *service_model.V1Search
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project under namesapce

	*/
	Project string
	/*SearchUUID
	  UUID

	*/
	SearchUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch project search params
func (o *PatchProjectSearchParams) WithTimeout(timeout time.Duration) *PatchProjectSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch project search params
func (o *PatchProjectSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch project search params
func (o *PatchProjectSearchParams) WithContext(ctx context.Context) *PatchProjectSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch project search params
func (o *PatchProjectSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch project search params
func (o *PatchProjectSearchParams) WithHTTPClient(client *http.Client) *PatchProjectSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch project search params
func (o *PatchProjectSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch project search params
func (o *PatchProjectSearchParams) WithBody(body *service_model.V1Search) *PatchProjectSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch project search params
func (o *PatchProjectSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the patch project search params
func (o *PatchProjectSearchParams) WithOwner(owner string) *PatchProjectSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch project search params
func (o *PatchProjectSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the patch project search params
func (o *PatchProjectSearchParams) WithProject(project string) *PatchProjectSearchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch project search params
func (o *PatchProjectSearchParams) SetProject(project string) {
	o.Project = project
}

// WithSearchUUID adds the searchUUID to the patch project search params
func (o *PatchProjectSearchParams) WithSearchUUID(searchUUID string) *PatchProjectSearchParams {
	o.SetSearchUUID(searchUUID)
	return o
}

// SetSearchUUID adds the searchUuid to the patch project search params
func (o *PatchProjectSearchParams) SetSearchUUID(searchUUID string) {
	o.SearchUUID = searchUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchProjectSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param search.uuid
	if err := r.SetPathParam("search.uuid", o.SearchUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}