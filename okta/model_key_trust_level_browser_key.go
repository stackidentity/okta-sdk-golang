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

// KeyTrustLevelBrowserKey Represents the attestation strength used by the Chrome Verified Access API
type KeyTrustLevelBrowserKey string

// List of KeyTrustLevelBrowserKey
const (
	KEYTRUSTLEVELBROWSERKEY_HW_KEY KeyTrustLevelBrowserKey = "CHROME_BROWSER_HW_KEY"
	KEYTRUSTLEVELBROWSERKEY_OS_KEY KeyTrustLevelBrowserKey = "CHROME_BROWSER_OS_KEY"
)

// All allowed values of KeyTrustLevelBrowserKey enum
var AllowedKeyTrustLevelBrowserKeyEnumValues = []KeyTrustLevelBrowserKey{
	"CHROME_BROWSER_HW_KEY",
	"CHROME_BROWSER_OS_KEY",
}

func (v *KeyTrustLevelBrowserKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyTrustLevelBrowserKey(value)
	for _, existing := range AllowedKeyTrustLevelBrowserKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyTrustLevelBrowserKey", value)
}

// NewKeyTrustLevelBrowserKeyFromValue returns a pointer to a valid KeyTrustLevelBrowserKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyTrustLevelBrowserKeyFromValue(v string) (*KeyTrustLevelBrowserKey, error) {
	ev := KeyTrustLevelBrowserKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyTrustLevelBrowserKey: valid values are %v", v, AllowedKeyTrustLevelBrowserKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyTrustLevelBrowserKey) IsValid() bool {
	for _, existing := range AllowedKeyTrustLevelBrowserKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyTrustLevelBrowserKey value
func (v KeyTrustLevelBrowserKey) Ptr() *KeyTrustLevelBrowserKey {
	return &v
}

type NullableKeyTrustLevelBrowserKey struct {
	value *KeyTrustLevelBrowserKey
	isSet bool
}

func (v NullableKeyTrustLevelBrowserKey) Get() *KeyTrustLevelBrowserKey {
	return v.value
}

func (v *NullableKeyTrustLevelBrowserKey) Set(val *KeyTrustLevelBrowserKey) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyTrustLevelBrowserKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyTrustLevelBrowserKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyTrustLevelBrowserKey(val *KeyTrustLevelBrowserKey) *NullableKeyTrustLevelBrowserKey {
	return &NullableKeyTrustLevelBrowserKey{value: val, isSet: true}
}

func (v NullableKeyTrustLevelBrowserKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyTrustLevelBrowserKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
