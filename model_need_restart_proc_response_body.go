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

// NeedRestartProcResponseBody struct for NeedRestartProcResponseBody
type NeedRestartProcResponseBody struct {
	// InitSystem of NeedRestartProc
	InitSystem string `json:"initSystem"`
	// Path of NeedRestartProc
	Path string `json:"path"`
	// PID
	Pid string `json:"pid"`
	// ServiceName of NeedRestartProc
	ServiceName string `json:"serviceName"`
}

// NewNeedRestartProcResponseBody instantiates a new NeedRestartProcResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeedRestartProcResponseBody(initSystem string, path string, pid string, serviceName string) *NeedRestartProcResponseBody {
	this := NeedRestartProcResponseBody{}
	this.InitSystem = initSystem
	this.Path = path
	this.Pid = pid
	this.ServiceName = serviceName
	return &this
}

// NewNeedRestartProcResponseBodyWithDefaults instantiates a new NeedRestartProcResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeedRestartProcResponseBodyWithDefaults() *NeedRestartProcResponseBody {
	this := NeedRestartProcResponseBody{}
	return &this
}

// GetInitSystem returns the InitSystem field value
func (o *NeedRestartProcResponseBody) GetInitSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitSystem
}

// GetInitSystemOk returns a tuple with the InitSystem field value
// and a boolean to check if the value has been set.
func (o *NeedRestartProcResponseBody) GetInitSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitSystem, true
}

// SetInitSystem sets field value
func (o *NeedRestartProcResponseBody) SetInitSystem(v string) {
	o.InitSystem = v
}

// GetPath returns the Path field value
func (o *NeedRestartProcResponseBody) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *NeedRestartProcResponseBody) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *NeedRestartProcResponseBody) SetPath(v string) {
	o.Path = v
}

// GetPid returns the Pid field value
func (o *NeedRestartProcResponseBody) GetPid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *NeedRestartProcResponseBody) GetPidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value
func (o *NeedRestartProcResponseBody) SetPid(v string) {
	o.Pid = v
}

// GetServiceName returns the ServiceName field value
func (o *NeedRestartProcResponseBody) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *NeedRestartProcResponseBody) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *NeedRestartProcResponseBody) SetServiceName(v string) {
	o.ServiceName = v
}

func (o NeedRestartProcResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["initSystem"] = o.InitSystem
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["pid"] = o.Pid
	}
	if true {
		toSerialize["serviceName"] = o.ServiceName
	}
	return json.Marshal(toSerialize)
}

type NullableNeedRestartProcResponseBody struct {
	value *NeedRestartProcResponseBody
	isSet bool
}

func (v NullableNeedRestartProcResponseBody) Get() *NeedRestartProcResponseBody {
	return v.value
}

func (v *NullableNeedRestartProcResponseBody) Set(val *NeedRestartProcResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableNeedRestartProcResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableNeedRestartProcResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeedRestartProcResponseBody(val *NeedRestartProcResponseBody) *NullableNeedRestartProcResponseBody {
	return &NullableNeedRestartProcResponseBody{value: val, isSet: true}
}

func (v NullableNeedRestartProcResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeedRestartProcResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


