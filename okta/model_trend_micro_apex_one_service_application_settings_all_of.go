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
)

// TrendMicroApexOneServiceApplicationSettingsAllOf struct for TrendMicroApexOneServiceApplicationSettingsAllOf
type TrendMicroApexOneServiceApplicationSettingsAllOf struct {
	App                  TrendMicroApexOneServiceApplicationSettingsApplication `json:"app"`
	AdditionalProperties map[string]interface{}
}

type _TrendMicroApexOneServiceApplicationSettingsAllOf TrendMicroApexOneServiceApplicationSettingsAllOf

// NewTrendMicroApexOneServiceApplicationSettingsAllOf instantiates a new TrendMicroApexOneServiceApplicationSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrendMicroApexOneServiceApplicationSettingsAllOf(app TrendMicroApexOneServiceApplicationSettingsApplication) *TrendMicroApexOneServiceApplicationSettingsAllOf {
	this := TrendMicroApexOneServiceApplicationSettingsAllOf{}
	this.App = app
	return &this
}

// NewTrendMicroApexOneServiceApplicationSettingsAllOfWithDefaults instantiates a new TrendMicroApexOneServiceApplicationSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrendMicroApexOneServiceApplicationSettingsAllOfWithDefaults() *TrendMicroApexOneServiceApplicationSettingsAllOf {
	this := TrendMicroApexOneServiceApplicationSettingsAllOf{}
	return &this
}

// GetApp returns the App field value
func (o *TrendMicroApexOneServiceApplicationSettingsAllOf) GetApp() TrendMicroApexOneServiceApplicationSettingsApplication {
	if o == nil {
		var ret TrendMicroApexOneServiceApplicationSettingsApplication
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *TrendMicroApexOneServiceApplicationSettingsAllOf) GetAppOk() (*TrendMicroApexOneServiceApplicationSettingsApplication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *TrendMicroApexOneServiceApplicationSettingsAllOf) SetApp(v TrendMicroApexOneServiceApplicationSettingsApplication) {
	o.App = v
}

func (o TrendMicroApexOneServiceApplicationSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TrendMicroApexOneServiceApplicationSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTrendMicroApexOneServiceApplicationSettingsAllOf := _TrendMicroApexOneServiceApplicationSettingsAllOf{}

	err = json.Unmarshal(bytes, &varTrendMicroApexOneServiceApplicationSettingsAllOf)
	if err == nil {
		*o = TrendMicroApexOneServiceApplicationSettingsAllOf(varTrendMicroApexOneServiceApplicationSettingsAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTrendMicroApexOneServiceApplicationSettingsAllOf struct {
	value *TrendMicroApexOneServiceApplicationSettingsAllOf
	isSet bool
}

func (v NullableTrendMicroApexOneServiceApplicationSettingsAllOf) Get() *TrendMicroApexOneServiceApplicationSettingsAllOf {
	return v.value
}

func (v *NullableTrendMicroApexOneServiceApplicationSettingsAllOf) Set(val *TrendMicroApexOneServiceApplicationSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendMicroApexOneServiceApplicationSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendMicroApexOneServiceApplicationSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendMicroApexOneServiceApplicationSettingsAllOf(val *TrendMicroApexOneServiceApplicationSettingsAllOf) *NullableTrendMicroApexOneServiceApplicationSettingsAllOf {
	return &NullableTrendMicroApexOneServiceApplicationSettingsAllOf{value: val, isSet: true}
}

func (v NullableTrendMicroApexOneServiceApplicationSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrendMicroApexOneServiceApplicationSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
