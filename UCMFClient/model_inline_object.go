/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"os"
)

// InlineObject struct for InlineObject
type InlineObject struct {
	JsonData *DicEntryCreateData `json:"jsonData,omitempty"`
	BinaryDataUeRadioCapability5GS **os.File `json:"binaryDataUeRadioCapability5GS,omitempty"`
	BinaryDataUeRadioCapabilityEPS **os.File `json:"binaryDataUeRadioCapabilityEPS,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *InlineObject) GetJsonData() DicEntryCreateData {
	if o == nil || o.JsonData == nil {
		var ret DicEntryCreateData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetJsonDataOk() (*DicEntryCreateData, bool) {
	if o == nil || o.JsonData == nil {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *InlineObject) HasJsonData() bool {
	if o != nil && o.JsonData != nil {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given DicEntryCreateData and assigns it to the JsonData field.
func (o *InlineObject) SetJsonData(v DicEntryCreateData) {
	o.JsonData = &v
}

// GetBinaryDataUeRadioCapability5GS returns the BinaryDataUeRadioCapability5GS field value if set, zero value otherwise.
func (o *InlineObject) GetBinaryDataUeRadioCapability5GS() *os.File {
	if o == nil || o.BinaryDataUeRadioCapability5GS == nil {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapability5GS
}

// GetBinaryDataUeRadioCapability5GSOk returns a tuple with the BinaryDataUeRadioCapability5GS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetBinaryDataUeRadioCapability5GSOk() (**os.File, bool) {
	if o == nil || o.BinaryDataUeRadioCapability5GS == nil {
		return nil, false
	}
	return o.BinaryDataUeRadioCapability5GS, true
}

// HasBinaryDataUeRadioCapability5GS returns a boolean if a field has been set.
func (o *InlineObject) HasBinaryDataUeRadioCapability5GS() bool {
	if o != nil && o.BinaryDataUeRadioCapability5GS != nil {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapability5GS gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapability5GS field.
func (o *InlineObject) SetBinaryDataUeRadioCapability5GS(v *os.File) {
	o.BinaryDataUeRadioCapability5GS = &v
}

// GetBinaryDataUeRadioCapabilityEPS returns the BinaryDataUeRadioCapabilityEPS field value if set, zero value otherwise.
func (o *InlineObject) GetBinaryDataUeRadioCapabilityEPS() *os.File {
	if o == nil || o.BinaryDataUeRadioCapabilityEPS == nil {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataUeRadioCapabilityEPS
}

// GetBinaryDataUeRadioCapabilityEPSOk returns a tuple with the BinaryDataUeRadioCapabilityEPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetBinaryDataUeRadioCapabilityEPSOk() (**os.File, bool) {
	if o == nil || o.BinaryDataUeRadioCapabilityEPS == nil {
		return nil, false
	}
	return o.BinaryDataUeRadioCapabilityEPS, true
}

// HasBinaryDataUeRadioCapabilityEPS returns a boolean if a field has been set.
func (o *InlineObject) HasBinaryDataUeRadioCapabilityEPS() bool {
	if o != nil && o.BinaryDataUeRadioCapabilityEPS != nil {
		return true
	}

	return false
}

// SetBinaryDataUeRadioCapabilityEPS gets a reference to the given *os.File and assigns it to the BinaryDataUeRadioCapabilityEPS field.
func (o *InlineObject) SetBinaryDataUeRadioCapabilityEPS(v *os.File) {
	o.BinaryDataUeRadioCapabilityEPS = &v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonData != nil {
		toSerialize["jsonData"] = o.JsonData
	}
	if o.BinaryDataUeRadioCapability5GS != nil {
		toSerialize["binaryDataUeRadioCapability5GS"] = o.BinaryDataUeRadioCapability5GS
	}
	if o.BinaryDataUeRadioCapabilityEPS != nil {
		toSerialize["binaryDataUeRadioCapabilityEPS"] = o.BinaryDataUeRadioCapabilityEPS
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


