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

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new queues v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for queues v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateQueue lists runs
*/
func (a *Client) CreateQueue(params *CreateQueueParams, authInfo runtime.ClientAuthInfoWriter) (*CreateQueueOK, *CreateQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateQueue",
		Method:             "POST",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{agent}/queues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateQueueOK:
		return value, nil, nil
	case *CreateQueueNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteQueue patches run
*/
func (a *Client) DeleteQueue(params *DeleteQueueParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteQueueOK, *DeleteQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteQueue",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{agent}/queues/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteQueueOK:
		return value, nil, nil
	case *DeleteQueueNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetQueue creates new run
*/
func (a *Client) GetQueue(params *GetQueueParams, authInfo runtime.ClientAuthInfoWriter) (*GetQueueOK, *GetQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetQueue",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{agent}/queues/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetQueueOK:
		return value, nil, nil
	case *GetQueueNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListQueueNames lists bookmarked runs for user
*/
func (a *Client) ListQueueNames(params *ListQueueNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListQueueNamesOK, *ListQueueNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQueueNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListQueueNames",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{agent}/queues/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListQueueNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListQueueNamesOK:
		return value, nil, nil
	case *ListQueueNamesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListQueues lists archived runs for user
*/
func (a *Client) ListQueues(params *ListQueuesParams, authInfo runtime.ClientAuthInfoWriter) (*ListQueuesOK, *ListQueuesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQueuesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListQueues",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{agent}/queues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListQueuesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListQueuesOK:
		return value, nil, nil
	case *ListQueuesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchQueue updates run
*/
func (a *Client) PatchQueue(params *PatchQueueParams, authInfo runtime.ClientAuthInfoWriter) (*PatchQueueOK, *PatchQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchQueue",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{queue.agent}/queues/{queue.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchQueueOK:
		return value, nil, nil
	case *PatchQueueNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateQueue gets run
*/
func (a *Client) UpdateQueue(params *UpdateQueueParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateQueueOK, *UpdateQueueNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateQueueParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateQueue",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/agents/{queue.agent}/queues/{queue.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateQueueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateQueueOK:
		return value, nil, nil
	case *UpdateQueueNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queues_v1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
