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

// PasswordProtectionWarningTrigger Indicates whether the Password Protection Warning feature is enabled
type PasswordProtectionWarningTrigger string

// List of PasswordProtectionWarningTrigger
const (
	PASSWORDPROTECTIONWARNINGTRIGGER_PASSWORD_PROTECTION_OFF PasswordProtectionWarningTrigger = "PASSWORD_PROTECTION_OFF"
	PASSWORDPROTECTIONWARNINGTRIGGER_PASSWORD_REUSE          PasswordProtectionWarningTrigger = "PASSWORD_REUSE"
	PASSWORDPROTECTIONWARNINGTRIGGER_PHISHING_REUSE          PasswordProtectionWarningTrigger = "PHISHING_REUSE"
)

// All allowed values of PasswordProtectionWarningTrigger enum
var AllowedPasswordProtectionWarningTriggerEnumValues = []PasswordProtectionWarningTrigger{
	"PASSWORD_PROTECTION_OFF",
	"PASSWORD_REUSE",
	"PHISHING_REUSE",
}

func (v *PasswordProtectionWarningTrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PasswordProtectionWarningTrigger(value)
	for _, existing := range AllowedPasswordProtectionWarningTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PasswordProtectionWarningTrigger", value)
}

// NewPasswordProtectionWarningTriggerFromValue returns a pointer to a valid PasswordProtectionWarningTrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPasswordProtectionWarningTriggerFromValue(v string) (*PasswordProtectionWarningTrigger, error) {
	ev := PasswordProtectionWarningTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PasswordProtectionWarningTrigger: valid values are %v", v, AllowedPasswordProtectionWarningTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PasswordProtectionWarningTrigger) IsValid() bool {
	for _, existing := range AllowedPasswordProtectionWarningTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PasswordProtectionWarningTrigger value
func (v PasswordProtectionWarningTrigger) Ptr() *PasswordProtectionWarningTrigger {
	return &v
}

type NullablePasswordProtectionWarningTrigger struct {
	value *PasswordProtectionWarningTrigger
	isSet bool
}

func (v NullablePasswordProtectionWarningTrigger) Get() *PasswordProtectionWarningTrigger {
	return v.value
}

func (v *NullablePasswordProtectionWarningTrigger) Set(val *PasswordProtectionWarningTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordProtectionWarningTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordProtectionWarningTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordProtectionWarningTrigger(val *PasswordProtectionWarningTrigger) *NullablePasswordProtectionWarningTrigger {
	return &NullablePasswordProtectionWarningTrigger{value: val, isSet: true}
}

func (v NullablePasswordProtectionWarningTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordProtectionWarningTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}