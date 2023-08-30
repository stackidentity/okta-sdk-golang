/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// MDMEnrollmentPolicyEnrollment the model 'MDMEnrollmentPolicyEnrollment'
type MDMEnrollmentPolicyEnrollment string

// List of MDMEnrollmentPolicyEnrollment
const (
	MDMENROLLMENTPOLICYENROLLMENT_ANY_OR_NONE MDMEnrollmentPolicyEnrollment = "ANY_OR_NONE"
	MDMENROLLMENTPOLICYENROLLMENT_OMM         MDMEnrollmentPolicyEnrollment = "OMM"
)

// All allowed values of MDMEnrollmentPolicyEnrollment enum
var AllowedMDMEnrollmentPolicyEnrollmentEnumValues = []MDMEnrollmentPolicyEnrollment{
	"ANY_OR_NONE",
	"OMM",
}

func (v *MDMEnrollmentPolicyEnrollment) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MDMEnrollmentPolicyEnrollment(value)
	for _, existing := range AllowedMDMEnrollmentPolicyEnrollmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MDMEnrollmentPolicyEnrollment", value)
}

// NewMDMEnrollmentPolicyEnrollmentFromValue returns a pointer to a valid MDMEnrollmentPolicyEnrollment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMDMEnrollmentPolicyEnrollmentFromValue(v string) (*MDMEnrollmentPolicyEnrollment, error) {
	ev := MDMEnrollmentPolicyEnrollment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MDMEnrollmentPolicyEnrollment: valid values are %v", v, AllowedMDMEnrollmentPolicyEnrollmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MDMEnrollmentPolicyEnrollment) IsValid() bool {
	for _, existing := range AllowedMDMEnrollmentPolicyEnrollmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MDMEnrollmentPolicyEnrollment value
func (v MDMEnrollmentPolicyEnrollment) Ptr() *MDMEnrollmentPolicyEnrollment {
	return &v
}

type NullableMDMEnrollmentPolicyEnrollment struct {
	value *MDMEnrollmentPolicyEnrollment
	isSet bool
}

func (v NullableMDMEnrollmentPolicyEnrollment) Get() *MDMEnrollmentPolicyEnrollment {
	return v.value
}

func (v *NullableMDMEnrollmentPolicyEnrollment) Set(val *MDMEnrollmentPolicyEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableMDMEnrollmentPolicyEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableMDMEnrollmentPolicyEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDMEnrollmentPolicyEnrollment(val *MDMEnrollmentPolicyEnrollment) *NullableMDMEnrollmentPolicyEnrollment {
	return &NullableMDMEnrollmentPolicyEnrollment{value: val, isSet: true}
}

func (v NullableMDMEnrollmentPolicyEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDMEnrollmentPolicyEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
