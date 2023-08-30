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

// ProvisioningDeprovisionedAction the model 'ProvisioningDeprovisionedAction'
type ProvisioningDeprovisionedAction string

// List of ProvisioningDeprovisionedAction
const (
	PROVISIONINGDEPROVISIONEDACTION_NONE       ProvisioningDeprovisionedAction = "NONE"
	PROVISIONINGDEPROVISIONEDACTION_REACTIVATE ProvisioningDeprovisionedAction = "REACTIVATE"
)

// All allowed values of ProvisioningDeprovisionedAction enum
var AllowedProvisioningDeprovisionedActionEnumValues = []ProvisioningDeprovisionedAction{
	"NONE",
	"REACTIVATE",
}

func (v *ProvisioningDeprovisionedAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProvisioningDeprovisionedAction(value)
	for _, existing := range AllowedProvisioningDeprovisionedActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProvisioningDeprovisionedAction", value)
}

// NewProvisioningDeprovisionedActionFromValue returns a pointer to a valid ProvisioningDeprovisionedAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProvisioningDeprovisionedActionFromValue(v string) (*ProvisioningDeprovisionedAction, error) {
	ev := ProvisioningDeprovisionedAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProvisioningDeprovisionedAction: valid values are %v", v, AllowedProvisioningDeprovisionedActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProvisioningDeprovisionedAction) IsValid() bool {
	for _, existing := range AllowedProvisioningDeprovisionedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProvisioningDeprovisionedAction value
func (v ProvisioningDeprovisionedAction) Ptr() *ProvisioningDeprovisionedAction {
	return &v
}

type NullableProvisioningDeprovisionedAction struct {
	value *ProvisioningDeprovisionedAction
	isSet bool
}

func (v NullableProvisioningDeprovisionedAction) Get() *ProvisioningDeprovisionedAction {
	return v.value
}

func (v *NullableProvisioningDeprovisionedAction) Set(val *ProvisioningDeprovisionedAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningDeprovisionedAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningDeprovisionedAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningDeprovisionedAction(val *ProvisioningDeprovisionedAction) *NullableProvisioningDeprovisionedAction {
	return &NullableProvisioningDeprovisionedAction{value: val, isSet: true}
}

func (v NullableProvisioningDeprovisionedAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningDeprovisionedAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}