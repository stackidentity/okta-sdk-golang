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

// OINSaml20ApplicationSettingsSignOnAllOf struct for OINSaml20ApplicationSettingsSignOnAllOf
type OINSaml20ApplicationSettingsSignOnAllOf struct {
	SignOnMode interface{} `json:"signOnMode,omitempty"`
	// Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm)
	DestinationOverride string `json:"destinationOverride"`
	// Set to `true` to prompt users for their credentials when a SAML request has the `ForceAuthn` attribute set to `true`
	HonorForceAuthn               *bool                    `json:"honorForceAuthn,omitempty"`
	ConfiguredAttributeStatements []SamlAttributeStatement `json:"configuredAttributeStatements,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _OINSaml20ApplicationSettingsSignOnAllOf OINSaml20ApplicationSettingsSignOnAllOf

// NewOINSaml20ApplicationSettingsSignOnAllOf instantiates a new OINSaml20ApplicationSettingsSignOnAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOINSaml20ApplicationSettingsSignOnAllOf(destinationOverride string) *OINSaml20ApplicationSettingsSignOnAllOf {
	this := OINSaml20ApplicationSettingsSignOnAllOf{}
	this.DestinationOverride = destinationOverride
	return &this
}

// NewOINSaml20ApplicationSettingsSignOnAllOfWithDefaults instantiates a new OINSaml20ApplicationSettingsSignOnAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOINSaml20ApplicationSettingsSignOnAllOfWithDefaults() *OINSaml20ApplicationSettingsSignOnAllOf {
	this := OINSaml20ApplicationSettingsSignOnAllOf{}
	return &this
}

// GetSignOnMode returns the SignOnMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetSignOnMode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetSignOnModeOk() (*interface{}, bool) {
	if o == nil || o.SignOnMode == nil {
		return nil, false
	}
	return &o.SignOnMode, true
}

// HasSignOnMode returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) HasSignOnMode() bool {
	if o != nil && o.SignOnMode != nil {
		return true
	}

	return false
}

// SetSignOnMode gets a reference to the given interface{} and assigns it to the SignOnMode field.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) SetSignOnMode(v interface{}) {
	o.SignOnMode = v
}

// GetDestinationOverride returns the DestinationOverride field value
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetDestinationOverride() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationOverride
}

// GetDestinationOverrideOk returns a tuple with the DestinationOverride field value
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetDestinationOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationOverride, true
}

// SetDestinationOverride sets field value
func (o *OINSaml20ApplicationSettingsSignOnAllOf) SetDestinationOverride(v string) {
	o.DestinationOverride = v
}

// GetHonorForceAuthn returns the HonorForceAuthn field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetHonorForceAuthn() bool {
	if o == nil || o.HonorForceAuthn == nil {
		var ret bool
		return ret
	}
	return *o.HonorForceAuthn
}

// GetHonorForceAuthnOk returns a tuple with the HonorForceAuthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetHonorForceAuthnOk() (*bool, bool) {
	if o == nil || o.HonorForceAuthn == nil {
		return nil, false
	}
	return o.HonorForceAuthn, true
}

// HasHonorForceAuthn returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) HasHonorForceAuthn() bool {
	if o != nil && o.HonorForceAuthn != nil {
		return true
	}

	return false
}

// SetHonorForceAuthn gets a reference to the given bool and assigns it to the HonorForceAuthn field.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) SetHonorForceAuthn(v bool) {
	o.HonorForceAuthn = &v
}

// GetConfiguredAttributeStatements returns the ConfiguredAttributeStatements field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetConfiguredAttributeStatements() []SamlAttributeStatement {
	if o == nil || o.ConfiguredAttributeStatements == nil {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.ConfiguredAttributeStatements
}

// GetConfiguredAttributeStatementsOk returns a tuple with the ConfiguredAttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) GetConfiguredAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || o.ConfiguredAttributeStatements == nil {
		return nil, false
	}
	return o.ConfiguredAttributeStatements, true
}

// HasConfiguredAttributeStatements returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) HasConfiguredAttributeStatements() bool {
	if o != nil && o.ConfiguredAttributeStatements != nil {
		return true
	}

	return false
}

// SetConfiguredAttributeStatements gets a reference to the given []SamlAttributeStatement and assigns it to the ConfiguredAttributeStatements field.
func (o *OINSaml20ApplicationSettingsSignOnAllOf) SetConfiguredAttributeStatements(v []SamlAttributeStatement) {
	o.ConfiguredAttributeStatements = v
}

func (o OINSaml20ApplicationSettingsSignOnAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SignOnMode != nil {
		toSerialize["signOnMode"] = o.SignOnMode
	}
	if true {
		toSerialize["destinationOverride"] = o.DestinationOverride
	}
	if o.HonorForceAuthn != nil {
		toSerialize["honorForceAuthn"] = o.HonorForceAuthn
	}
	if o.ConfiguredAttributeStatements != nil {
		toSerialize["configuredAttributeStatements"] = o.ConfiguredAttributeStatements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OINSaml20ApplicationSettingsSignOnAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOINSaml20ApplicationSettingsSignOnAllOf := _OINSaml20ApplicationSettingsSignOnAllOf{}

	err = json.Unmarshal(bytes, &varOINSaml20ApplicationSettingsSignOnAllOf)
	if err == nil {
		*o = OINSaml20ApplicationSettingsSignOnAllOf(varOINSaml20ApplicationSettingsSignOnAllOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signOnMode")
		delete(additionalProperties, "destinationOverride")
		delete(additionalProperties, "honorForceAuthn")
		delete(additionalProperties, "configuredAttributeStatements")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOINSaml20ApplicationSettingsSignOnAllOf struct {
	value *OINSaml20ApplicationSettingsSignOnAllOf
	isSet bool
}

func (v NullableOINSaml20ApplicationSettingsSignOnAllOf) Get() *OINSaml20ApplicationSettingsSignOnAllOf {
	return v.value
}

func (v *NullableOINSaml20ApplicationSettingsSignOnAllOf) Set(val *OINSaml20ApplicationSettingsSignOnAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOINSaml20ApplicationSettingsSignOnAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOINSaml20ApplicationSettingsSignOnAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOINSaml20ApplicationSettingsSignOnAllOf(val *OINSaml20ApplicationSettingsSignOnAllOf) *NullableOINSaml20ApplicationSettingsSignOnAllOf {
	return &NullableOINSaml20ApplicationSettingsSignOnAllOf{value: val, isSet: true}
}

func (v NullableOINSaml20ApplicationSettingsSignOnAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOINSaml20ApplicationSettingsSignOnAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}