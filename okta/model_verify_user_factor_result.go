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

// VerifyUserFactorResult the model 'VerifyUserFactorResult'
type VerifyUserFactorResult string

// List of VerifyUserFactorResult
const (
	VERIFYUSERFACTORRESULT_CHALLENGE            VerifyUserFactorResult = "CHALLENGE"
	VERIFYUSERFACTORRESULT_ERROR                VerifyUserFactorResult = "ERROR"
	VERIFYUSERFACTORRESULT_EXPIRED              VerifyUserFactorResult = "EXPIRED"
	VERIFYUSERFACTORRESULT_FAILED               VerifyUserFactorResult = "FAILED"
	VERIFYUSERFACTORRESULT_PASSCODE_REPLAYED    VerifyUserFactorResult = "PASSCODE_REPLAYED"
	VERIFYUSERFACTORRESULT_REJECTED             VerifyUserFactorResult = "REJECTED"
	VERIFYUSERFACTORRESULT_SUCCESS              VerifyUserFactorResult = "SUCCESS"
	VERIFYUSERFACTORRESULT_TIMEOUT              VerifyUserFactorResult = "TIMEOUT"
	VERIFYUSERFACTORRESULT_TIME_WINDOW_EXCEEDED VerifyUserFactorResult = "TIME_WINDOW_EXCEEDED"
	VERIFYUSERFACTORRESULT_WAITING              VerifyUserFactorResult = "WAITING"
)

// All allowed values of VerifyUserFactorResult enum
var AllowedVerifyUserFactorResultEnumValues = []VerifyUserFactorResult{
	"CHALLENGE",
	"ERROR",
	"EXPIRED",
	"FAILED",
	"PASSCODE_REPLAYED",
	"REJECTED",
	"SUCCESS",
	"TIMEOUT",
	"TIME_WINDOW_EXCEEDED",
	"WAITING",
}

func (v *VerifyUserFactorResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerifyUserFactorResult(value)
	for _, existing := range AllowedVerifyUserFactorResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerifyUserFactorResult", value)
}

// NewVerifyUserFactorResultFromValue returns a pointer to a valid VerifyUserFactorResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerifyUserFactorResultFromValue(v string) (*VerifyUserFactorResult, error) {
	ev := VerifyUserFactorResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerifyUserFactorResult: valid values are %v", v, AllowedVerifyUserFactorResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerifyUserFactorResult) IsValid() bool {
	for _, existing := range AllowedVerifyUserFactorResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerifyUserFactorResult value
func (v VerifyUserFactorResult) Ptr() *VerifyUserFactorResult {
	return &v
}

type NullableVerifyUserFactorResult struct {
	value *VerifyUserFactorResult
	isSet bool
}

func (v NullableVerifyUserFactorResult) Get() *VerifyUserFactorResult {
	return v.value
}

func (v *NullableVerifyUserFactorResult) Set(val *VerifyUserFactorResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyUserFactorResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyUserFactorResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyUserFactorResult(val *VerifyUserFactorResult) *NullableVerifyUserFactorResult {
	return &NullableVerifyUserFactorResult{value: val, isSet: true}
}

func (v NullableVerifyUserFactorResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyUserFactorResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}