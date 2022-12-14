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

// ServerGetServerListResponseBody struct for ServerGetServerListResponseBody
type ServerGetServerListResponseBody struct {
	Paging *PagingResponseBody `json:"paging,omitempty"`
	// Servers list
	Servers []ServerListResponseBody `json:"servers,omitempty"`
}

// NewServerGetServerListResponseBody instantiates a new ServerGetServerListResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerGetServerListResponseBody() *ServerGetServerListResponseBody {
	this := ServerGetServerListResponseBody{}
	return &this
}

// NewServerGetServerListResponseBodyWithDefaults instantiates a new ServerGetServerListResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerGetServerListResponseBodyWithDefaults() *ServerGetServerListResponseBody {
	this := ServerGetServerListResponseBody{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *ServerGetServerListResponseBody) GetPaging() PagingResponseBody {
	if o == nil || o.Paging == nil {
		var ret PagingResponseBody
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerListResponseBody) GetPagingOk() (*PagingResponseBody, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *ServerGetServerListResponseBody) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given PagingResponseBody and assigns it to the Paging field.
func (o *ServerGetServerListResponseBody) SetPaging(v PagingResponseBody) {
	o.Paging = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *ServerGetServerListResponseBody) GetServers() []ServerListResponseBody {
	if o == nil || o.Servers == nil {
		var ret []ServerListResponseBody
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerGetServerListResponseBody) GetServersOk() ([]ServerListResponseBody, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *ServerGetServerListResponseBody) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerListResponseBody and assigns it to the Servers field.
func (o *ServerGetServerListResponseBody) SetServers(v []ServerListResponseBody) {
	o.Servers = v
}

func (o ServerGetServerListResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableServerGetServerListResponseBody struct {
	value *ServerGetServerListResponseBody
	isSet bool
}

func (v NullableServerGetServerListResponseBody) Get() *ServerGetServerListResponseBody {
	return v.value
}

func (v *NullableServerGetServerListResponseBody) Set(val *ServerGetServerListResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableServerGetServerListResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableServerGetServerListResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerGetServerListResponseBody(val *ServerGetServerListResponseBody) *NullableServerGetServerListResponseBody {
	return &NullableServerGetServerListResponseBody{value: val, isSet: true}
}

func (v NullableServerGetServerListResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerGetServerListResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


