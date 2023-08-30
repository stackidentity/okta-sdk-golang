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

// ProvisioningConnectionProfileTokenAllOf struct for ProvisioningConnectionProfileTokenAllOf
type ProvisioningConnectionProfileTokenAllOf struct {
	// Token used to authenticate with the app
	Token                string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionProfileTokenAllOf ProvisioningConnectionProfileTokenAllOf

// NewProvisioningConnectionProfileTokenAllOf instantiates a new ProvisioningConnectionProfileTokenAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionProfileTokenAllOf(token string) *ProvisioningConnectionProfileTokenAllOf {
	this := ProvisioningConnectionProfileTokenAllOf{}
	this.Token = token
	return &this
}

// NewProvisioningConnectionProfileTokenAllOfWithDefaults instantiates a new ProvisioningConnectionProfileTokenAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionProfileTokenAllOfWithDefaults() *ProvisioningConnectionProfileTokenAllOf {
	this := ProvisioningConnectionProfileTokenAllOf{}
	return &this
}

// GetToken returns the Token field value
func (o *ProvisioningConnectionProfileTokenAllOf) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionProfileTokenAllOf) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ProvisioningConnectionProfileTokenAllOf) SetToken(v string) {
	o.Token = v
}

func (o ProvisioningConnectionProfileTokenAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionProfileTokenAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionProfileTokenAllOf := _ProvisioningConnectionProfileTokenAllOf{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionProfileTokenAllOf)
	if err == nil {
		*o = ProvisioningConnectionProfileTokenAllOf(varProvisioningConnectionProfileTokenAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnectionProfileTokenAllOf struct {
	value *ProvisioningConnectionProfileTokenAllOf
	isSet bool
}

func (v NullableProvisioningConnectionProfileTokenAllOf) Get() *ProvisioningConnectionProfileTokenAllOf {
	return v.value
}

func (v *NullableProvisioningConnectionProfileTokenAllOf) Set(val *ProvisioningConnectionProfileTokenAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionProfileTokenAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionProfileTokenAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionProfileTokenAllOf(val *ProvisioningConnectionProfileTokenAllOf) *NullableProvisioningConnectionProfileTokenAllOf {
	return &NullableProvisioningConnectionProfileTokenAllOf{value: val, isSet: true}
}

func (v NullableProvisioningConnectionProfileTokenAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionProfileTokenAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
