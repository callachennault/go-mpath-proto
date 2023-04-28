/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_public_api

import (
	"encoding/json"
)

// GenomeNexusInfo struct for GenomeNexusInfo
type GenomeNexusInfo struct {
	Database *Version `json:"database,omitempty"`
	Server *Version `json:"server,omitempty"`
}

// NewGenomeNexusInfo instantiates a new GenomeNexusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenomeNexusInfo() *GenomeNexusInfo {
	this := GenomeNexusInfo{}
	return &this
}

// NewGenomeNexusInfoWithDefaults instantiates a new GenomeNexusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenomeNexusInfoWithDefaults() *GenomeNexusInfo {
	this := GenomeNexusInfo{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *GenomeNexusInfo) GetDatabase() Version {
	if o == nil || isNil(o.Database) {
		var ret Version
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenomeNexusInfo) GetDatabaseOk() (*Version, bool) {
	if o == nil || isNil(o.Database) {
    return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *GenomeNexusInfo) HasDatabase() bool {
	if o != nil && !isNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given Version and assigns it to the Database field.
func (o *GenomeNexusInfo) SetDatabase(v Version) {
	o.Database = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *GenomeNexusInfo) GetServer() Version {
	if o == nil || isNil(o.Server) {
		var ret Version
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenomeNexusInfo) GetServerOk() (*Version, bool) {
	if o == nil || isNil(o.Server) {
    return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *GenomeNexusInfo) HasServer() bool {
	if o != nil && !isNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given Version and assigns it to the Server field.
func (o *GenomeNexusInfo) SetServer(v Version) {
	o.Server = &v
}

func (o GenomeNexusInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !isNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableGenomeNexusInfo struct {
	value *GenomeNexusInfo
	isSet bool
}

func (v NullableGenomeNexusInfo) Get() *GenomeNexusInfo {
	return v.value
}

func (v *NullableGenomeNexusInfo) Set(val *GenomeNexusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGenomeNexusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGenomeNexusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenomeNexusInfo(val *GenomeNexusInfo) *NullableGenomeNexusInfo {
	return &NullableGenomeNexusInfo{value: val, isSet: true}
}

func (v NullableGenomeNexusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenomeNexusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


