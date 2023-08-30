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
	"reflect"
	"strings"
)

// EmailUserFactor struct for EmailUserFactor
type EmailUserFactor struct {
	UserFactor
	Profile              *EmailUserFactorProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailUserFactor EmailUserFactor

// NewEmailUserFactor instantiates a new EmailUserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailUserFactor() *EmailUserFactor {
	this := EmailUserFactor{}
	return &this
}

// NewEmailUserFactorWithDefaults instantiates a new EmailUserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailUserFactorWithDefaults() *EmailUserFactor {
	this := EmailUserFactor{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *EmailUserFactor) GetProfile() EmailUserFactorProfile {
	if o == nil || o.Profile == nil {
		var ret EmailUserFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailUserFactor) GetProfileOk() (*EmailUserFactorProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *EmailUserFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given EmailUserFactorProfile and assigns it to the Profile field.
func (o *EmailUserFactor) SetProfile(v EmailUserFactorProfile) {
	o.Profile = &v
}

func (o EmailUserFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUserFactor, errUserFactor := json.Marshal(o.UserFactor)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	errUserFactor = json.Unmarshal([]byte(serializedUserFactor), &toSerialize)
	if errUserFactor != nil {
		return []byte{}, errUserFactor
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailUserFactor) UnmarshalJSON(bytes []byte) (err error) {
	type EmailUserFactorWithoutEmbeddedStruct struct {
		Profile *EmailUserFactorProfile `json:"profile,omitempty"`
	}

	varEmailUserFactorWithoutEmbeddedStruct := EmailUserFactorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEmailUserFactorWithoutEmbeddedStruct)
	if err == nil {
		varEmailUserFactor := _EmailUserFactor{}
		varEmailUserFactor.Profile = varEmailUserFactorWithoutEmbeddedStruct.Profile
		*o = EmailUserFactor(varEmailUserFactor)
	} else {
		return err
	}

	varEmailUserFactor := _EmailUserFactor{}

	err = json.Unmarshal(bytes, &varEmailUserFactor)
	if err == nil {
		o.UserFactor = varEmailUserFactor.UserFactor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "profile")

		// remove fields from embedded structs
		reflectUserFactor := reflect.ValueOf(o.UserFactor)
		for i := 0; i < reflectUserFactor.Type().NumField(); i++ {
			t := reflectUserFactor.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailUserFactor struct {
	value *EmailUserFactor
	isSet bool
}

func (v NullableEmailUserFactor) Get() *EmailUserFactor {
	return v.value
}

func (v *NullableEmailUserFactor) Set(val *EmailUserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailUserFactor(val *EmailUserFactor) *NullableEmailUserFactor {
	return &NullableEmailUserFactor{value: val, isSet: true}
}

func (v NullableEmailUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}