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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1RandomSearch Parallelism based on randomly generated search space
// swagger:model v1RandomSearch
type V1RandomSearch struct {

	// Number of concurrent runs
	Concurrency string `json:"concurrency,omitempty"`

	// A list of Early stoppings, accpets both metric and failure early stopping mechanisms
	EarlyStopping []interface{} `json:"early_stopping"`

	// Kind of parallel, should be equal to "random_search"
	Kind string `json:"kind,omitempty"`

	// Matrix/Space definition of params to traverse
	Matrix string `json:"matrix,omitempty"`

	// Number of runs to generate and search
	NRuns string `json:"n_runs,omitempty"`

	// Seed for the random generator
	Seed string `json:"seed,omitempty"`
}

// Validate validates this v1 random search
func (m *V1RandomSearch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RandomSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RandomSearch) UnmarshalBinary(b []byte) error {
	var res V1RandomSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
