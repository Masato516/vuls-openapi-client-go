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

// PkgCpeListResponseBody struct for PkgCpeListResponseBody
type PkgCpeListResponseBody struct {
	// ApplicationName of library package
	ApplicationName *string `json:"applicationName,omitempty"`
	// Cpe FS of cpe
	CpeFS *string `json:"cpeFS,omitempty"`
	// CpeID of cpe
	CpeID *int64 `json:"cpeID,omitempty"`
	// Cpe URI of cpe
	CpeURI *string `json:"cpeURI,omitempty"`
	// crated time of package or cpe
	CreatedAt time.Time `json:"createdAt"`
	// githubPKGID of github pkg
	GithubPkgID *int64 `json:"githubPkgID,omitempty"`
	// LibraryPath of library package
	LibraryPath *string `json:"libraryPath,omitempty"`
	// libraryPKGID of library pkg
	LibraryPkgID *int64 `json:"libraryPkgID,omitempty"`
	// Name of package or cpe
	Name string `json:"name"`
	// NeedRestartProcess list of package
	NeedRestartProcs []NeedRestartProcResponseBody `json:"needRestartProcs,omitempty"`
	// New Release of package
	NewRelease *string `json:"newRelease,omitempty"`
	// New Version of package
	NewVersion *string `json:"newVersion,omitempty"`
	// Flag of Not fixed yet of package
	NotFixedYet *bool `json:"notFixedYet,omitempty"`
	// ossLicense of library package
	OssLicense *string `json:"ossLicense,omitempty"`
	// Package ID of package
	PkgID *int64 `json:"pkgID,omitempty"`
	// Release of package
	Release *string `json:"release,omitempty"`
	// Repository of package
	Repository *string `json:"repository,omitempty"`
	// ServerID of package or cpe
	ServerID int64 `json:"serverID"`
	// ServerName of package or cpe
	ServerName string `json:"serverName"`
	// ServerUUID of package or cpe
	ServerUuid string `json:"serverUuid"`
	// updated time of package or cpe
	UpdatedAt time.Time `json:"updatedAt"`
	// Version of package or cpe
	Version string `json:"version"`
	// WordpressPackageStatus of wordpress package
	WordpressPackageStatus *string `json:"wordpressPackageStatus,omitempty"`
	// wordpressPKGID of wordpress pkg
	WordpressPkgID *int64 `json:"wordpressPkgID,omitempty"`
}

// NewPkgCpeListResponseBody instantiates a new PkgCpeListResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgCpeListResponseBody(createdAt time.Time, name string, serverID int64, serverName string, serverUuid string, updatedAt time.Time, version string) *PkgCpeListResponseBody {
	this := PkgCpeListResponseBody{}
	this.CreatedAt = createdAt
	this.Name = name
	this.ServerID = serverID
	this.ServerName = serverName
	this.ServerUuid = serverUuid
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewPkgCpeListResponseBodyWithDefaults instantiates a new PkgCpeListResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgCpeListResponseBodyWithDefaults() *PkgCpeListResponseBody {
	this := PkgCpeListResponseBody{}
	return &this
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *PkgCpeListResponseBody) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetCpeFS returns the CpeFS field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetCpeFS() string {
	if o == nil || o.CpeFS == nil {
		var ret string
		return ret
	}
	return *o.CpeFS
}

// GetCpeFSOk returns a tuple with the CpeFS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetCpeFSOk() (*string, bool) {
	if o == nil || o.CpeFS == nil {
		return nil, false
	}
	return o.CpeFS, true
}

// HasCpeFS returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasCpeFS() bool {
	if o != nil && o.CpeFS != nil {
		return true
	}

	return false
}

// SetCpeFS gets a reference to the given string and assigns it to the CpeFS field.
func (o *PkgCpeListResponseBody) SetCpeFS(v string) {
	o.CpeFS = &v
}

// GetCpeID returns the CpeID field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetCpeID() int64 {
	if o == nil || o.CpeID == nil {
		var ret int64
		return ret
	}
	return *o.CpeID
}

// GetCpeIDOk returns a tuple with the CpeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetCpeIDOk() (*int64, bool) {
	if o == nil || o.CpeID == nil {
		return nil, false
	}
	return o.CpeID, true
}

// HasCpeID returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasCpeID() bool {
	if o != nil && o.CpeID != nil {
		return true
	}

	return false
}

// SetCpeID gets a reference to the given int64 and assigns it to the CpeID field.
func (o *PkgCpeListResponseBody) SetCpeID(v int64) {
	o.CpeID = &v
}

// GetCpeURI returns the CpeURI field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetCpeURI() string {
	if o == nil || o.CpeURI == nil {
		var ret string
		return ret
	}
	return *o.CpeURI
}

// GetCpeURIOk returns a tuple with the CpeURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetCpeURIOk() (*string, bool) {
	if o == nil || o.CpeURI == nil {
		return nil, false
	}
	return o.CpeURI, true
}

// HasCpeURI returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasCpeURI() bool {
	if o != nil && o.CpeURI != nil {
		return true
	}

	return false
}

// SetCpeURI gets a reference to the given string and assigns it to the CpeURI field.
func (o *PkgCpeListResponseBody) SetCpeURI(v string) {
	o.CpeURI = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PkgCpeListResponseBody) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PkgCpeListResponseBody) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetGithubPkgID returns the GithubPkgID field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetGithubPkgID() int64 {
	if o == nil || o.GithubPkgID == nil {
		var ret int64
		return ret
	}
	return *o.GithubPkgID
}

// GetGithubPkgIDOk returns a tuple with the GithubPkgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetGithubPkgIDOk() (*int64, bool) {
	if o == nil || o.GithubPkgID == nil {
		return nil, false
	}
	return o.GithubPkgID, true
}

// HasGithubPkgID returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasGithubPkgID() bool {
	if o != nil && o.GithubPkgID != nil {
		return true
	}

	return false
}

// SetGithubPkgID gets a reference to the given int64 and assigns it to the GithubPkgID field.
func (o *PkgCpeListResponseBody) SetGithubPkgID(v int64) {
	o.GithubPkgID = &v
}

// GetLibraryPath returns the LibraryPath field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetLibraryPath() string {
	if o == nil || o.LibraryPath == nil {
		var ret string
		return ret
	}
	return *o.LibraryPath
}

// GetLibraryPathOk returns a tuple with the LibraryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetLibraryPathOk() (*string, bool) {
	if o == nil || o.LibraryPath == nil {
		return nil, false
	}
	return o.LibraryPath, true
}

// HasLibraryPath returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasLibraryPath() bool {
	if o != nil && o.LibraryPath != nil {
		return true
	}

	return false
}

// SetLibraryPath gets a reference to the given string and assigns it to the LibraryPath field.
func (o *PkgCpeListResponseBody) SetLibraryPath(v string) {
	o.LibraryPath = &v
}

// GetLibraryPkgID returns the LibraryPkgID field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetLibraryPkgID() int64 {
	if o == nil || o.LibraryPkgID == nil {
		var ret int64
		return ret
	}
	return *o.LibraryPkgID
}

// GetLibraryPkgIDOk returns a tuple with the LibraryPkgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetLibraryPkgIDOk() (*int64, bool) {
	if o == nil || o.LibraryPkgID == nil {
		return nil, false
	}
	return o.LibraryPkgID, true
}

// HasLibraryPkgID returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasLibraryPkgID() bool {
	if o != nil && o.LibraryPkgID != nil {
		return true
	}

	return false
}

// SetLibraryPkgID gets a reference to the given int64 and assigns it to the LibraryPkgID field.
func (o *PkgCpeListResponseBody) SetLibraryPkgID(v int64) {
	o.LibraryPkgID = &v
}

// GetName returns the Name field value
func (o *PkgCpeListResponseBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PkgCpeListResponseBody) SetName(v string) {
	o.Name = v
}

// GetNeedRestartProcs returns the NeedRestartProcs field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetNeedRestartProcs() []NeedRestartProcResponseBody {
	if o == nil || o.NeedRestartProcs == nil {
		var ret []NeedRestartProcResponseBody
		return ret
	}
	return o.NeedRestartProcs
}

// GetNeedRestartProcsOk returns a tuple with the NeedRestartProcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetNeedRestartProcsOk() ([]NeedRestartProcResponseBody, bool) {
	if o == nil || o.NeedRestartProcs == nil {
		return nil, false
	}
	return o.NeedRestartProcs, true
}

// HasNeedRestartProcs returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasNeedRestartProcs() bool {
	if o != nil && o.NeedRestartProcs != nil {
		return true
	}

	return false
}

// SetNeedRestartProcs gets a reference to the given []NeedRestartProcResponseBody and assigns it to the NeedRestartProcs field.
func (o *PkgCpeListResponseBody) SetNeedRestartProcs(v []NeedRestartProcResponseBody) {
	o.NeedRestartProcs = v
}

// GetNewRelease returns the NewRelease field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetNewRelease() string {
	if o == nil || o.NewRelease == nil {
		var ret string
		return ret
	}
	return *o.NewRelease
}

// GetNewReleaseOk returns a tuple with the NewRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetNewReleaseOk() (*string, bool) {
	if o == nil || o.NewRelease == nil {
		return nil, false
	}
	return o.NewRelease, true
}

// HasNewRelease returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasNewRelease() bool {
	if o != nil && o.NewRelease != nil {
		return true
	}

	return false
}

// SetNewRelease gets a reference to the given string and assigns it to the NewRelease field.
func (o *PkgCpeListResponseBody) SetNewRelease(v string) {
	o.NewRelease = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetNewVersion() string {
	if o == nil || o.NewVersion == nil {
		var ret string
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetNewVersionOk() (*string, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given string and assigns it to the NewVersion field.
func (o *PkgCpeListResponseBody) SetNewVersion(v string) {
	o.NewVersion = &v
}

// GetNotFixedYet returns the NotFixedYet field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetNotFixedYet() bool {
	if o == nil || o.NotFixedYet == nil {
		var ret bool
		return ret
	}
	return *o.NotFixedYet
}

// GetNotFixedYetOk returns a tuple with the NotFixedYet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetNotFixedYetOk() (*bool, bool) {
	if o == nil || o.NotFixedYet == nil {
		return nil, false
	}
	return o.NotFixedYet, true
}

// HasNotFixedYet returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasNotFixedYet() bool {
	if o != nil && o.NotFixedYet != nil {
		return true
	}

	return false
}

// SetNotFixedYet gets a reference to the given bool and assigns it to the NotFixedYet field.
func (o *PkgCpeListResponseBody) SetNotFixedYet(v bool) {
	o.NotFixedYet = &v
}

// GetOssLicense returns the OssLicense field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetOssLicense() string {
	if o == nil || o.OssLicense == nil {
		var ret string
		return ret
	}
	return *o.OssLicense
}

// GetOssLicenseOk returns a tuple with the OssLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetOssLicenseOk() (*string, bool) {
	if o == nil || o.OssLicense == nil {
		return nil, false
	}
	return o.OssLicense, true
}

// HasOssLicense returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasOssLicense() bool {
	if o != nil && o.OssLicense != nil {
		return true
	}

	return false
}

// SetOssLicense gets a reference to the given string and assigns it to the OssLicense field.
func (o *PkgCpeListResponseBody) SetOssLicense(v string) {
	o.OssLicense = &v
}

// GetPkgID returns the PkgID field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetPkgID() int64 {
	if o == nil || o.PkgID == nil {
		var ret int64
		return ret
	}
	return *o.PkgID
}

// GetPkgIDOk returns a tuple with the PkgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetPkgIDOk() (*int64, bool) {
	if o == nil || o.PkgID == nil {
		return nil, false
	}
	return o.PkgID, true
}

// HasPkgID returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasPkgID() bool {
	if o != nil && o.PkgID != nil {
		return true
	}

	return false
}

// SetPkgID gets a reference to the given int64 and assigns it to the PkgID field.
func (o *PkgCpeListResponseBody) SetPkgID(v int64) {
	o.PkgID = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *PkgCpeListResponseBody) SetRelease(v string) {
	o.Release = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *PkgCpeListResponseBody) SetRepository(v string) {
	o.Repository = &v
}

// GetServerID returns the ServerID field value
func (o *PkgCpeListResponseBody) GetServerID() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerID
}

// GetServerIDOk returns a tuple with the ServerID field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetServerIDOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerID, true
}

// SetServerID sets field value
func (o *PkgCpeListResponseBody) SetServerID(v int64) {
	o.ServerID = v
}

// GetServerName returns the ServerName field value
func (o *PkgCpeListResponseBody) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *PkgCpeListResponseBody) SetServerName(v string) {
	o.ServerName = v
}

// GetServerUuid returns the ServerUuid field value
func (o *PkgCpeListResponseBody) GetServerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerUuid
}

// GetServerUuidOk returns a tuple with the ServerUuid field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetServerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerUuid, true
}

// SetServerUuid sets field value
func (o *PkgCpeListResponseBody) SetServerUuid(v string) {
	o.ServerUuid = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PkgCpeListResponseBody) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PkgCpeListResponseBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value
func (o *PkgCpeListResponseBody) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PkgCpeListResponseBody) SetVersion(v string) {
	o.Version = v
}

// GetWordpressPackageStatus returns the WordpressPackageStatus field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetWordpressPackageStatus() string {
	if o == nil || o.WordpressPackageStatus == nil {
		var ret string
		return ret
	}
	return *o.WordpressPackageStatus
}

// GetWordpressPackageStatusOk returns a tuple with the WordpressPackageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetWordpressPackageStatusOk() (*string, bool) {
	if o == nil || o.WordpressPackageStatus == nil {
		return nil, false
	}
	return o.WordpressPackageStatus, true
}

// HasWordpressPackageStatus returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasWordpressPackageStatus() bool {
	if o != nil && o.WordpressPackageStatus != nil {
		return true
	}

	return false
}

// SetWordpressPackageStatus gets a reference to the given string and assigns it to the WordpressPackageStatus field.
func (o *PkgCpeListResponseBody) SetWordpressPackageStatus(v string) {
	o.WordpressPackageStatus = &v
}

// GetWordpressPkgID returns the WordpressPkgID field value if set, zero value otherwise.
func (o *PkgCpeListResponseBody) GetWordpressPkgID() int64 {
	if o == nil || o.WordpressPkgID == nil {
		var ret int64
		return ret
	}
	return *o.WordpressPkgID
}

// GetWordpressPkgIDOk returns a tuple with the WordpressPkgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgCpeListResponseBody) GetWordpressPkgIDOk() (*int64, bool) {
	if o == nil || o.WordpressPkgID == nil {
		return nil, false
	}
	return o.WordpressPkgID, true
}

// HasWordpressPkgID returns a boolean if a field has been set.
func (o *PkgCpeListResponseBody) HasWordpressPkgID() bool {
	if o != nil && o.WordpressPkgID != nil {
		return true
	}

	return false
}

// SetWordpressPkgID gets a reference to the given int64 and assigns it to the WordpressPkgID field.
func (o *PkgCpeListResponseBody) SetWordpressPkgID(v int64) {
	o.WordpressPkgID = &v
}

func (o PkgCpeListResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationName != nil {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if o.CpeFS != nil {
		toSerialize["cpeFS"] = o.CpeFS
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
	if o.GithubPkgID != nil {
		toSerialize["githubPkgID"] = o.GithubPkgID
	}
	if o.LibraryPath != nil {
		toSerialize["libraryPath"] = o.LibraryPath
	}
	if o.LibraryPkgID != nil {
		toSerialize["libraryPkgID"] = o.LibraryPkgID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NeedRestartProcs != nil {
		toSerialize["needRestartProcs"] = o.NeedRestartProcs
	}
	if o.NewRelease != nil {
		toSerialize["newRelease"] = o.NewRelease
	}
	if o.NewVersion != nil {
		toSerialize["newVersion"] = o.NewVersion
	}
	if o.NotFixedYet != nil {
		toSerialize["notFixedYet"] = o.NotFixedYet
	}
	if o.OssLicense != nil {
		toSerialize["ossLicense"] = o.OssLicense
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
		toSerialize["serverName"] = o.ServerName
	}
	if true {
		toSerialize["serverUuid"] = o.ServerUuid
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.WordpressPackageStatus != nil {
		toSerialize["wordpressPackageStatus"] = o.WordpressPackageStatus
	}
	if o.WordpressPkgID != nil {
		toSerialize["wordpressPkgID"] = o.WordpressPkgID
	}
	return json.Marshal(toSerialize)
}

type NullablePkgCpeListResponseBody struct {
	value *PkgCpeListResponseBody
	isSet bool
}

func (v NullablePkgCpeListResponseBody) Get() *PkgCpeListResponseBody {
	return v.value
}

func (v *NullablePkgCpeListResponseBody) Set(val *PkgCpeListResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgCpeListResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgCpeListResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgCpeListResponseBody(val *PkgCpeListResponseBody) *NullablePkgCpeListResponseBody {
	return &NullablePkgCpeListResponseBody{value: val, isSet: true}
}

func (v NullablePkgCpeListResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgCpeListResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


