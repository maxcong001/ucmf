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

// ManAssOpRequestlist - struct for ManAssOpRequestlist
type ManAssOpRequestlist struct {
	interface{} *interface{}
}

// interface{}AsManAssOpRequestlist is a convenience function that returns interface{} wrapped in ManAssOpRequestlist
func interface{}AsManAssOpRequestlist(v *interface{}) ManAssOpRequestlist {
	return ManAssOpRequestlist{ interface{}: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ManAssOpRequestlist) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into interface{}
	err = json.Unmarshal(data, &dst.interface{})
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			match++
		}
	} else {
		dst.interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.interface{} = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ManAssOpRequestlist)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ManAssOpRequestlist)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ManAssOpRequestlist) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ManAssOpRequestlist) GetActualInstance() (interface{}) {
	if obj.interface{} != nil {
		return obj.interface{}
	}

	// all schemas are nil
	return nil
}

type NullableManAssOpRequestlist struct {
	value *ManAssOpRequestlist
	isSet bool
}

func (v NullableManAssOpRequestlist) Get() *ManAssOpRequestlist {
	return v.value
}

func (v *NullableManAssOpRequestlist) Set(val *ManAssOpRequestlist) {
	v.value = val
	v.isSet = true
}

func (v NullableManAssOpRequestlist) IsSet() bool {
	return v.isSet
}

func (v *NullableManAssOpRequestlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManAssOpRequestlist(val *ManAssOpRequestlist) *NullableManAssOpRequestlist {
	return &NullableManAssOpRequestlist{value: val, isSet: true}
}

func (v NullableManAssOpRequestlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManAssOpRequestlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


