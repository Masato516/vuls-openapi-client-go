/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
)

// AffectedProcResponseBody struct for AffectedProcResponseBody
type AffectedProcResponseBody struct {
	// AffectedProc Name
	Name string `json:"name"`
	// PID
	Pid string `json:"pid"`
}

// NewAffectedProcResponseBody instantiates a new AffectedProcResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffectedProcResponseBody(name string, pid string) *AffectedProcResponseBody {
	this := AffectedProcResponseBody{}
	this.Name = name
	this.Pid = pid
	return &this
}

// NewAffectedProcResponseBodyWithDefaults instantiates a new AffectedProcResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffectedProcResponseBodyWithDefaults() *AffectedProcResponseBody {
	this := AffectedProcResponseBody{}
	return &this
}

// GetName returns the Name field value
func (o *AffectedProcResponseBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AffectedProcResponseBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AffectedProcResponseBody) SetName(v string) {
	o.Name = v
}

// GetPid returns the Pid field value
func (o *AffectedProcResponseBody) GetPid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *AffectedProcResponseBody) GetPidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value
func (o *AffectedProcResponseBody) SetPid(v string) {
	o.Pid = v
}

func (o AffectedProcResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["pid"] = o.Pid
	}
	return json.Marshal(toSerialize)
}

type NullableAffectedProcResponseBody struct {
	value *AffectedProcResponseBody
	isSet bool
}

func (v NullableAffectedProcResponseBody) Get() *AffectedProcResponseBody {
	return v.value
}

func (v *NullableAffectedProcResponseBody) Set(val *AffectedProcResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAffectedProcResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAffectedProcResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffectedProcResponseBody(val *AffectedProcResponseBody) *NullableAffectedProcResponseBody {
	return &NullableAffectedProcResponseBody{value: val, isSet: true}
}

func (v NullableAffectedProcResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffectedProcResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


