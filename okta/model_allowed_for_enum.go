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

// AllowedForEnum the model 'AllowedForEnum'
type AllowedForEnum string

// List of AllowedForEnum
const (
	ALLOWEDFORENUM_ANY      AllowedForEnum = "any"
	ALLOWEDFORENUM_NONE     AllowedForEnum = "none"
	ALLOWEDFORENUM_RECOVERY AllowedForEnum = "recovery"
	ALLOWEDFORENUM_SSO      AllowedForEnum = "sso"
)

// All allowed values of AllowedForEnum enum
var AllowedAllowedForEnumEnumValues = []AllowedForEnum{
	"any",
	"none",
	"recovery",
	"sso",
}

func (v *AllowedForEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllowedForEnum(value)
	for _, existing := range AllowedAllowedForEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllowedForEnum", value)
}

// NewAllowedForEnumFromValue returns a pointer to a valid AllowedForEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllowedForEnumFromValue(v string) (*AllowedForEnum, error) {
	ev := AllowedForEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllowedForEnum: valid values are %v", v, AllowedAllowedForEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllowedForEnum) IsValid() bool {
	for _, existing := range AllowedAllowedForEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AllowedForEnum value
func (v AllowedForEnum) Ptr() *AllowedForEnum {
	return &v
}

type NullableAllowedForEnum struct {
	value *AllowedForEnum
	isSet bool
}

func (v NullableAllowedForEnum) Get() *AllowedForEnum {
	return v.value
}

func (v *NullableAllowedForEnum) Set(val *AllowedForEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedForEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedForEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedForEnum(val *AllowedForEnum) *NullableAllowedForEnum {
	return &NullableAllowedForEnum{value: val, isSet: true}
}

func (v NullableAllowedForEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedForEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
