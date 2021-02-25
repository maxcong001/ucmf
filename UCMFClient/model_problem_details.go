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
)

// ProblemDetails struct for ProblemDetails
type ProblemDetails struct {
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Status *int32 `json:"status,omitempty"`
	Detail *string `json:"detail,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Cause *string `json:"cause,omitempty"`
	InvalidParams *[]InvalidParam `json:"invalidParams,omitempty"`
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewProblemDetails instantiates a new ProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetails() *ProblemDetails {
	this := ProblemDetails{}
	return &this
}

// NewProblemDetailsWithDefaults instantiates a new ProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailsWithDefaults() *ProblemDetails {
	this := ProblemDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProblemDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProblemDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProblemDetails) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProblemDetails) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProblemDetails) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProblemDetails) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProblemDetails) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProblemDetails) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ProblemDetails) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ProblemDetails) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemDetails) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ProblemDetails) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ProblemDetails) GetInstance() string {
	if o == nil || o.Instance == nil {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetInstanceOk() (*string, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ProblemDetails) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ProblemDetails) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ProblemDetails) GetCause() string {
	if o == nil || o.Cause == nil {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetCauseOk() (*string, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ProblemDetails) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ProblemDetails) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *ProblemDetails) GetInvalidParams() []InvalidParam {
	if o == nil || o.InvalidParams == nil {
		var ret []InvalidParam
		return ret
	}
	return *o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetInvalidParamsOk() (*[]InvalidParam, bool) {
	if o == nil || o.InvalidParams == nil {
		return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *ProblemDetails) HasInvalidParams() bool {
	if o != nil && o.InvalidParams != nil {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam and assigns it to the InvalidParams field.
func (o *ProblemDetails) SetInvalidParams(v []InvalidParam) {
	o.InvalidParams = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ProblemDetails) GetSupportedFeatures() string {
	if o == nil || o.SupportedFeatures == nil {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetails) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || o.SupportedFeatures == nil {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ProblemDetails) HasSupportedFeatures() bool {
	if o != nil && o.SupportedFeatures != nil {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ProblemDetails) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o ProblemDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.Cause != nil {
		toSerialize["cause"] = o.Cause
	}
	if o.InvalidParams != nil {
		toSerialize["invalidParams"] = o.InvalidParams
	}
	if o.SupportedFeatures != nil {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return json.Marshal(toSerialize)
}

type NullableProblemDetails struct {
	value *ProblemDetails
	isSet bool
}

func (v NullableProblemDetails) Get() *ProblemDetails {
	return v.value
}

func (v *NullableProblemDetails) Set(val *ProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetails(val *ProblemDetails) *NullableProblemDetails {
	return &NullableProblemDetails{value: val, isSet: true}
}

func (v NullableProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


