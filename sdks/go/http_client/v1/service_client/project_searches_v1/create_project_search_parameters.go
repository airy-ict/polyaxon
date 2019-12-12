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

// NewCreateProjectSearchParams creates a new CreateProjectSearchParams object
// with the default values initialized.
func NewCreateProjectSearchParams() *CreateProjectSearchParams {
	var ()
	return &CreateProjectSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectSearchParamsWithTimeout creates a new CreateProjectSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProjectSearchParamsWithTimeout(timeout time.Duration) *CreateProjectSearchParams {
	var ()
	return &CreateProjectSearchParams{

		timeout: timeout,
	}
}

// NewCreateProjectSearchParamsWithContext creates a new CreateProjectSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProjectSearchParamsWithContext(ctx context.Context) *CreateProjectSearchParams {
	var ()
	return &CreateProjectSearchParams{

		Context: ctx,
	}
}

// NewCreateProjectSearchParamsWithHTTPClient creates a new CreateProjectSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProjectSearchParamsWithHTTPClient(client *http.Client) *CreateProjectSearchParams {
	var ()
	return &CreateProjectSearchParams{
		HTTPClient: client,
	}
}

/*CreateProjectSearchParams contains all the parameters to send to the API endpoint
for the create project search operation typically these are written to a http.Request
*/
type CreateProjectSearchParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create project search params
func (o *CreateProjectSearchParams) WithTimeout(timeout time.Duration) *CreateProjectSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project search params
func (o *CreateProjectSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project search params
func (o *CreateProjectSearchParams) WithContext(ctx context.Context) *CreateProjectSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project search params
func (o *CreateProjectSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project search params
func (o *CreateProjectSearchParams) WithHTTPClient(client *http.Client) *CreateProjectSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project search params
func (o *CreateProjectSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create project search params
func (o *CreateProjectSearchParams) WithBody(body *service_model.V1Search) *CreateProjectSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create project search params
func (o *CreateProjectSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the create project search params
func (o *CreateProjectSearchParams) WithOwner(owner string) *CreateProjectSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create project search params
func (o *CreateProjectSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the create project search params
func (o *CreateProjectSearchParams) WithProject(project string) *CreateProjectSearchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create project search params
func (o *CreateProjectSearchParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
