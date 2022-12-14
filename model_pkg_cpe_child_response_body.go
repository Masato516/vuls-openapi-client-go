/*
Future Vuls Public API

Future Vuls Public API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"encoding/json"
	"time"
)

// PkgCpeChildResponseBody struct for PkgCpeChildResponseBody
type PkgCpeChildResponseBody struct {
	// AffectedProcess list of package
	AffectedProcs []AffectedProcResponseBody `json:"affectedProcs,omitempty"`
	// CpeID of cpe
	CpeID *int64 `json:"cpeID,omitempty"`
	// Cpe URI of cpe
	CpeURI *string `json:"cpeURI,omitempty"`
	// crated time of package or cpe
	CreatedAt time.Time `json:"createdAt"`
	// Name of package or cpe
	Name string `json:"name"`
	// New Release of package
	NewRelease *string `json:"newRelease,omitempty"`
	// New Version of package
	NewVersion *string `json:"newVersion,omitempty"`
	// Package ID of package
	PkgID *int64 `json:"pkgID,omitempty"`
	// Release of package
	Release *string `json:"release,omitempty"`
	// Repository of package
	Repository *string `json:"repository,omitempty"`
	// ServerID of package or cpe
	ServerID int64 `json:"serverID"`
	// updated time of package or cpe
	UpdatedAt time.Time `json:"updatedAt"`
	// Version of package or cpe
	Version string `json:"version"`
}

// NewPkgCpeChildResponseBody instantiates a new PkgCpeChildResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgCpeChildResponseBody(createdAt time.Time, name string, serverID int64, updatedAt time.Time, version string) *PkgCpeChildResponseBody {
	this := PkgCpeChildResponseBody{}
	this.CreatedAt = createdAt
	this.Name = name
	this.ServerID = serverID
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewPkgCpeChildResponseBodyWithDefaults instantiates a new PkgCpeChildResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgCpeChildResponseBodyWithDefaults() *PkgCpeChildResponseBody {
	this := PkgCpeChildResponseBody{}
	return &this
}

// GetAffectedProcs returns the AffectedProcs field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetAffectedProcs() []AffectedProcResponseBody {
	if o == nil || o.AffectedProcs == nil {
		var ret []AffectedProcResponseBody
		return ret
	}
	return o.AffectedProcs
}

// GetAffectedProcsOk returns a tuple with the AffectedProcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetAffectedProcsOk() ([]AffectedProcResponseBody, bool) {
	if o == nil || o.AffectedProcs == nil {
		return nil, false
	}
	return o.AffectedProcs, true
}

// HasAffectedProcs returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasAffectedProcs() bool {
	if o != nil && o.AffectedProcs != nil {
		return true
	}

	return false
}

// SetAffectedProcs gets a reference to the given []AffectedProcResponseBody and assigns it to the AffectedProcs field.
func (o *PkgCpeChildResponseBody) SetAffectedProcs(v []AffectedProcResponseBody) {
	o.AffectedProcs = v
}

// GetCpeID returns the CpeID field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetCpeID() int64 {
	if o == nil || o.CpeID == nil {
		var ret int64
		return ret
	}
	return *o.CpeID
}

// GetCpeIDOk returns a tuple with the CpeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetCpeIDOk() (*int64, bool) {
	if o == nil || o.CpeID == nil {
		return nil, false
	}
	return o.CpeID, true
}

// HasCpeID returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasCpeID() bool {
	if o != nil && o.CpeID != nil {
		return true
	}

	return false
}

// SetCpeID gets a reference to the given int64 and assigns it to the CpeID field.
func (o *PkgCpeChildResponseBody) SetCpeID(v int64) {
	o.CpeID = &v
}

// GetCpeURI returns the CpeURI field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetCpeURI() string {
	if o == nil || o.CpeURI == nil {
		var ret string
		return ret
	}
	return *o.CpeURI
}

// GetCpeURIOk returns a tuple with the CpeURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetCpeURIOk() (*string, bool) {
	if o == nil || o.CpeURI == nil {
		return nil, false
	}
	return o.CpeURI, true
}

// HasCpeURI returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasCpeURI() bool {
	if o != nil && o.CpeURI != nil {
		return true
	}

	return false
}

// SetCpeURI gets a reference to the given string and assigns it to the CpeURI field.
func (o *PkgCpeChildResponseBody) SetCpeURI(v string) {
	o.CpeURI = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PkgCpeChildResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PkgCpeChildResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetName returns the Name field value
func (o *PkgCpeChildResponseBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PkgCpeChildResponseBody) SetName(v string) {
	o.Name = v
}

// GetNewRelease returns the NewRelease field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetNewRelease() string {
	if o == nil || o.NewRelease == nil {
		var ret string
		return ret
	}
	return *o.NewRelease
}

// GetNewReleaseOk returns a tuple with the NewRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetNewReleaseOk() (*string, bool) {
	if o == nil || o.NewRelease == nil {
		return nil, false
	}
	return o.NewRelease, true
}

// HasNewRelease returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasNewRelease() bool {
	if o != nil && o.NewRelease != nil {
		return true
	}

	return false
}

// SetNewRelease gets a reference to the given string and assigns it to the NewRelease field.
func (o *PkgCpeChildResponseBody) SetNewRelease(v string) {
	o.NewRelease = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetNewVersion() string {
	if o == nil || o.NewVersion == nil {
		var ret string
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetNewVersionOk() (*string, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given string and assigns it to the NewVersion field.
func (o *PkgCpeChildResponseBody) SetNewVersion(v string) {
	o.NewVersion = &v
}

// GetPkgID returns the PkgID field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetPkgID() int64 {
	if o == nil || o.PkgID == nil {
		var ret int64
		return ret
	}
	return *o.PkgID
}

// GetPkgIDOk returns a tuple with the PkgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetPkgIDOk() (*int64, bool) {
	if o == nil || o.PkgID == nil {
		return nil, false
	}
	return o.PkgID, true
}

// HasPkgID returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasPkgID() bool {
	if o != nil && o.PkgID != nil {
		return true
	}

	return false
}

// SetPkgID gets a reference to the given int64 and assigns it to the PkgID field.
func (o *PkgCpeChildResponseBody) SetPkgID(v int64) {
	o.PkgID = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *PkgCpeChildResponseBody) SetRelease(v string) {
	o.Release = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PkgCpeChildResponseBody) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PkgCpeChildResponseBody) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *PkgCpeChildResponseBody) SetRepository(v string) {
	o.Repository = &v
}

// GetServerID returns the ServerID field value
func (o *PkgCpeChildResponseBody) GetServerID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerID
}

// GetServerIDOk returns a tuple with the ServerID field value
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetServerIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerID, true
}

// SetServerID sets field value
func (o *PkgCpeChildResponseBody) SetServerID(v int64) {
	o.ServerID = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PkgCpeChildResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PkgCpeChildResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value
func (o *PkgCpeChildResponseBody) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PkgCpeChildResponseBody) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PkgCpeChildResponseBody) SetVersion(v string) {
	o.Version = v
}

func (o PkgCpeChildResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AffectedProcs != nil {
		toSerialize["affectedProcs"] = o.AffectedProcs
	}
	if o.CpeID != nil {
		toSerialize["cpeID"] = o.CpeID
	}
	if o.CpeURI != nil {
		toSerialize["cpeURI"] = o.CpeURI
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NewRelease != nil {
		toSerialize["newRelease"] = o.NewRelease
	}
	if o.NewVersion != nil {
		toSerialize["newVersion"] = o.NewVersion
	}
	if o.PkgID != nil {
		toSerialize["pkgID"] = o.PkgID
	}
	if o.Release != nil {
		toSerialize["release"] = o.Release
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["serverID"] = o.ServerID
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullablePkgCpeChildResponseBody struct {
	value *PkgCpeChildResponseBody
	isSet bool
}

func (v NullablePkgCpeChildResponseBody) Get() *PkgCpeChildResponseBody {
	return v.value
}

func (v *NullablePkgCpeChildResponseBody) Set(val *PkgCpeChildResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgCpeChildResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgCpeChildResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgCpeChildResponseBody(val *PkgCpeChildResponseBody) *NullablePkgCpeChildResponseBody {
	return &NullablePkgCpeChildResponseBody{value: val, isSet: true}
}

func (v NullablePkgCpeChildResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgCpeChildResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


