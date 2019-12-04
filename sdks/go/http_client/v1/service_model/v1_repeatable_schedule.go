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

// V1RepeatableSchedule Repeatable schedule specification
// swagger:model v1RepeatableSchedule
type V1RepeatableSchedule struct {

	// A flag to set a dependency on past executions
	DependsOnPast bool `json:"depends_on_past,omitempty"`

	// Kind of parallel, should be equal to "repeatable"
	Kind string `json:"kind,omitempty"`

	// Limit to stop executing this schedule
	Limit string `json:"limit,omitempty"`
}

// Validate validates this v1 repeatable schedule
func (m *V1RepeatableSchedule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RepeatableSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RepeatableSchedule) UnmarshalBinary(b []byte) error {
	var res V1RepeatableSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
