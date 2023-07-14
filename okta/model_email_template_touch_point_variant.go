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

// EmailTemplateTouchPointVariant the model 'EmailTemplateTouchPointVariant'
type EmailTemplateTouchPointVariant string

// List of EmailTemplateTouchPointVariant
const (
	EMAILTEMPLATETOUCHPOINTVARIANT_FULL_THEME   EmailTemplateTouchPointVariant = "FULL_THEME"
	EMAILTEMPLATETOUCHPOINTVARIANT_OKTA_DEFAULT EmailTemplateTouchPointVariant = "OKTA_DEFAULT"
)

// All allowed values of EmailTemplateTouchPointVariant enum
var AllowedEmailTemplateTouchPointVariantEnumValues = []EmailTemplateTouchPointVariant{
	"FULL_THEME",
	"OKTA_DEFAULT",
}

func (v *EmailTemplateTouchPointVariant) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailTemplateTouchPointVariant(value)
	for _, existing := range AllowedEmailTemplateTouchPointVariantEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailTemplateTouchPointVariant", value)
}

// NewEmailTemplateTouchPointVariantFromValue returns a pointer to a valid EmailTemplateTouchPointVariant
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailTemplateTouchPointVariantFromValue(v string) (*EmailTemplateTouchPointVariant, error) {
	ev := EmailTemplateTouchPointVariant(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailTemplateTouchPointVariant: valid values are %v", v, AllowedEmailTemplateTouchPointVariantEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailTemplateTouchPointVariant) IsValid() bool {
	for _, existing := range AllowedEmailTemplateTouchPointVariantEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailTemplateTouchPointVariant value
func (v EmailTemplateTouchPointVariant) Ptr() *EmailTemplateTouchPointVariant {
	return &v
}

type NullableEmailTemplateTouchPointVariant struct {
	value *EmailTemplateTouchPointVariant
	isSet bool
}

func (v NullableEmailTemplateTouchPointVariant) Get() *EmailTemplateTouchPointVariant {
	return v.value
}

func (v *NullableEmailTemplateTouchPointVariant) Set(val *EmailTemplateTouchPointVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateTouchPointVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateTouchPointVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateTouchPointVariant(val *EmailTemplateTouchPointVariant) *NullableEmailTemplateTouchPointVariant {
	return &NullableEmailTemplateTouchPointVariant{value: val, isSet: true}
}

func (v NullableEmailTemplateTouchPointVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateTouchPointVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}