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
	"fmt"
)

// EventTypeAnyOf the model 'EventTypeAnyOf'
type EventTypeAnyOf string

// List of EventType_anyOf
const (
	CREATION_OF_DICTIONARY_ENTRY EventTypeAnyOf = "CREATION_OF_DICTIONARY_ENTRY"
	DELETION_OF_PLMN_ASSIGNED_IDS EventTypeAnyOf = "DELETION_OF_PLMN_ASSIGNED_IDS"
	NEW_VERSION_ID_OF_PLMN_ASSIGNED_IDS EventTypeAnyOf = "NEW_VERSION_ID_OF_PLMN_ASSIGNED_IDs"
)

func (v *EventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeAnyOf(value)
	for _, existing := range []EventTypeAnyOf{ "CREATION_OF_DICTIONARY_ENTRY", "DELETION_OF_PLMN_ASSIGNED_IDS", "NEW_VERSION_ID_OF_PLMN_ASSIGNED_IDs",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeAnyOf", value)
}

// Ptr returns reference to EventType_anyOf value
func (v EventTypeAnyOf) Ptr() *EventTypeAnyOf {
	return &v
}

type NullableEventTypeAnyOf struct {
	value *EventTypeAnyOf
	isSet bool
}

func (v NullableEventTypeAnyOf) Get() *EventTypeAnyOf {
	return v.value
}

func (v *NullableEventTypeAnyOf) Set(val *EventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeAnyOf(val *EventTypeAnyOf) *NullableEventTypeAnyOf {
	return &NullableEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

