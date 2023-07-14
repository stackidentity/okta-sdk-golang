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

API version: 4.0.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"fmt"
)

// PolicyAccess the model 'PolicyAccess'
type PolicyAccess string

// List of PolicyAccess
const (
	POLICYACCESS_ALLOW PolicyAccess = "ALLOW"
	POLICYACCESS_DENY  PolicyAccess = "DENY"
)

// All allowed values of PolicyAccess enum
var AllowedPolicyAccessEnumValues = []PolicyAccess{
	"ALLOW",
	"DENY",
}

func (v *PolicyAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyAccess(value)
	for _, existing := range AllowedPolicyAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyAccess", value)
}

// NewPolicyAccessFromValue returns a pointer to a valid PolicyAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyAccessFromValue(v string) (*PolicyAccess, error) {
	ev := PolicyAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyAccess: valid values are %v", v, AllowedPolicyAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyAccess) IsValid() bool {
	for _, existing := range AllowedPolicyAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyAccess value
func (v PolicyAccess) Ptr() *PolicyAccess {
	return &v
}

type NullablePolicyAccess struct {
	value *PolicyAccess
	isSet bool
}

func (v NullablePolicyAccess) Get() *PolicyAccess {
	return v.value
}

func (v *NullablePolicyAccess) Set(val *PolicyAccess) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAccess) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAccess(val *PolicyAccess) *NullablePolicyAccess {
	return &NullablePolicyAccess{value: val, isSet: true}
}

func (v NullablePolicyAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}