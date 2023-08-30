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

// OpenIdConnectRefreshTokenRotationType the model 'OpenIdConnectRefreshTokenRotationType'
type OpenIdConnectRefreshTokenRotationType string

// List of OpenIdConnectRefreshTokenRotationType
const (
	OPENIDCONNECTREFRESHTOKENROTATIONTYPE_ROTATE OpenIdConnectRefreshTokenRotationType = "ROTATE"
	OPENIDCONNECTREFRESHTOKENROTATIONTYPE_STATIC OpenIdConnectRefreshTokenRotationType = "STATIC"
)

// All allowed values of OpenIdConnectRefreshTokenRotationType enum
var AllowedOpenIdConnectRefreshTokenRotationTypeEnumValues = []OpenIdConnectRefreshTokenRotationType{
	"ROTATE",
	"STATIC",
}

func (v *OpenIdConnectRefreshTokenRotationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OpenIdConnectRefreshTokenRotationType(value)
	for _, existing := range AllowedOpenIdConnectRefreshTokenRotationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OpenIdConnectRefreshTokenRotationType", value)
}

// NewOpenIdConnectRefreshTokenRotationTypeFromValue returns a pointer to a valid OpenIdConnectRefreshTokenRotationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOpenIdConnectRefreshTokenRotationTypeFromValue(v string) (*OpenIdConnectRefreshTokenRotationType, error) {
	ev := OpenIdConnectRefreshTokenRotationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OpenIdConnectRefreshTokenRotationType: valid values are %v", v, AllowedOpenIdConnectRefreshTokenRotationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OpenIdConnectRefreshTokenRotationType) IsValid() bool {
	for _, existing := range AllowedOpenIdConnectRefreshTokenRotationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpenIdConnectRefreshTokenRotationType value
func (v OpenIdConnectRefreshTokenRotationType) Ptr() *OpenIdConnectRefreshTokenRotationType {
	return &v
}

type NullableOpenIdConnectRefreshTokenRotationType struct {
	value *OpenIdConnectRefreshTokenRotationType
	isSet bool
}

func (v NullableOpenIdConnectRefreshTokenRotationType) Get() *OpenIdConnectRefreshTokenRotationType {
	return v.value
}

func (v *NullableOpenIdConnectRefreshTokenRotationType) Set(val *OpenIdConnectRefreshTokenRotationType) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectRefreshTokenRotationType) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectRefreshTokenRotationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectRefreshTokenRotationType(val *OpenIdConnectRefreshTokenRotationType) *NullableOpenIdConnectRefreshTokenRotationType {
	return &NullableOpenIdConnectRefreshTokenRotationType{value: val, isSet: true}
}

func (v NullableOpenIdConnectRefreshTokenRotationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectRefreshTokenRotationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}