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

// SlackApplicationAllOf struct for SlackApplicationAllOf
type SlackApplicationAllOf struct {
	Name                 string                   `json:"name"`
	SignOnMode           interface{}              `json:"signOnMode,omitempty"`
	Settings             SlackApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _SlackApplicationAllOf SlackApplicationAllOf

// NewSlackApplicationAllOf instantiates a new SlackApplicationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackApplicationAllOf(name string, settings SlackApplicationSettings) *SlackApplicationAllOf {
	this := SlackApplicationAllOf{}
	this.Name = name
	this.Settings = settings
	return &this
}

// NewSlackApplicationAllOfWithDefaults instantiates a new SlackApplicationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackApplicationAllOfWithDefaults() *SlackApplicationAllOf {
	this := SlackApplicationAllOf{}
	var name string = "slack"
	this.Name = name
	return &this
}

// GetName returns the Name field value
func (o *SlackApplicationAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlackApplicationAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlackApplicationAllOf) SetName(v string) {
	o.Name = v
}

// GetSignOnMode returns the SignOnMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SlackApplicationAllOf) GetSignOnMode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SlackApplicationAllOf) GetSignOnModeOk() (*interface{}, bool) {
	if o == nil || o.SignOnMode == nil {
		return nil, false
	}
	return &o.SignOnMode, true
}

// HasSignOnMode returns a boolean if a field has been set.
func (o *SlackApplicationAllOf) HasSignOnMode() bool {
	if o != nil && o.SignOnMode != nil {
		return true
	}

	return false
}

// SetSignOnMode gets a reference to the given interface{} and assigns it to the SignOnMode field.
func (o *SlackApplicationAllOf) SetSignOnMode(v interface{}) {
	o.SignOnMode = v
}

// GetSettings returns the Settings field value
func (o *SlackApplicationAllOf) GetSettings() SlackApplicationSettings {
	if o == nil {
		var ret SlackApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *SlackApplicationAllOf) GetSettingsOk() (*SlackApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *SlackApplicationAllOf) SetSettings(v SlackApplicationSettings) {
	o.Settings = v
}

func (o SlackApplicationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SignOnMode != nil {
		toSerialize["signOnMode"] = o.SignOnMode
	}
	if true {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SlackApplicationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSlackApplicationAllOf := _SlackApplicationAllOf{}

	err = json.Unmarshal(bytes, &varSlackApplicationAllOf)
	if err == nil {
		*o = SlackApplicationAllOf(varSlackApplicationAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "signOnMode")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSlackApplicationAllOf struct {
	value *SlackApplicationAllOf
	isSet bool
}

func (v NullableSlackApplicationAllOf) Get() *SlackApplicationAllOf {
	return v.value
}

func (v *NullableSlackApplicationAllOf) Set(val *SlackApplicationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackApplicationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackApplicationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackApplicationAllOf(val *SlackApplicationAllOf) *NullableSlackApplicationAllOf {
	return &NullableSlackApplicationAllOf{value: val, isSet: true}
}

func (v NullableSlackApplicationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackApplicationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}