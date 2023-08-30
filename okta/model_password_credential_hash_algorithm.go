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

// PasswordCredentialHashAlgorithm the model 'PasswordCredentialHashAlgorithm'
type PasswordCredentialHashAlgorithm string

// List of PasswordCredentialHashAlgorithm
const (
	PASSWORDCREDENTIALHASHALGORITHM_BCRYPT  PasswordCredentialHashAlgorithm = "BCRYPT"
	PASSWORDCREDENTIALHASHALGORITHM_MD5     PasswordCredentialHashAlgorithm = "MD5"
	PASSWORDCREDENTIALHASHALGORITHM_PBKDF2  PasswordCredentialHashAlgorithm = "PBKDF2"
	PASSWORDCREDENTIALHASHALGORITHM_SHA_1   PasswordCredentialHashAlgorithm = "SHA-1"
	PASSWORDCREDENTIALHASHALGORITHM_SHA_256 PasswordCredentialHashAlgorithm = "SHA-256"
	PASSWORDCREDENTIALHASHALGORITHM_SHA_512 PasswordCredentialHashAlgorithm = "SHA-512"
)

// All allowed values of PasswordCredentialHashAlgorithm enum
var AllowedPasswordCredentialHashAlgorithmEnumValues = []PasswordCredentialHashAlgorithm{
	"BCRYPT",
	"MD5",
	"PBKDF2",
	"SHA-1",
	"SHA-256",
	"SHA-512",
}

func (v *PasswordCredentialHashAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PasswordCredentialHashAlgorithm(value)
	for _, existing := range AllowedPasswordCredentialHashAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PasswordCredentialHashAlgorithm", value)
}

// NewPasswordCredentialHashAlgorithmFromValue returns a pointer to a valid PasswordCredentialHashAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPasswordCredentialHashAlgorithmFromValue(v string) (*PasswordCredentialHashAlgorithm, error) {
	ev := PasswordCredentialHashAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PasswordCredentialHashAlgorithm: valid values are %v", v, AllowedPasswordCredentialHashAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PasswordCredentialHashAlgorithm) IsValid() bool {
	for _, existing := range AllowedPasswordCredentialHashAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PasswordCredentialHashAlgorithm value
func (v PasswordCredentialHashAlgorithm) Ptr() *PasswordCredentialHashAlgorithm {
	return &v
}

type NullablePasswordCredentialHashAlgorithm struct {
	value *PasswordCredentialHashAlgorithm
	isSet bool
}

func (v NullablePasswordCredentialHashAlgorithm) Get() *PasswordCredentialHashAlgorithm {
	return v.value
}

func (v *NullablePasswordCredentialHashAlgorithm) Set(val *PasswordCredentialHashAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordCredentialHashAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordCredentialHashAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordCredentialHashAlgorithm(val *PasswordCredentialHashAlgorithm) *NullablePasswordCredentialHashAlgorithm {
	return &NullablePasswordCredentialHashAlgorithm{value: val, isSet: true}
}

func (v NullablePasswordCredentialHashAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordCredentialHashAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
