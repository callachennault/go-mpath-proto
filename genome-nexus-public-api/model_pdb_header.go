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

// PdbHeader struct for PdbHeader
type PdbHeader struct {
	Compound map[string]interface{} `json:"compound,omitempty"`
	// PDB id
	PdbId string `json:"pdbId"`
	Source map[string]interface{} `json:"source,omitempty"`
	// PDB description
	Title string `json:"title"`
}

// NewPdbHeader instantiates a new PdbHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdbHeader(pdbId string, title string) *PdbHeader {
	this := PdbHeader{}
	this.PdbId = pdbId
	this.Title = title
	return &this
}

// NewPdbHeaderWithDefaults instantiates a new PdbHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPdbHeaderWithDefaults() *PdbHeader {
	this := PdbHeader{}
	return &this
}

// GetCompound returns the Compound field value if set, zero value otherwise.
func (o *PdbHeader) GetCompound() map[string]interface{} {
	if o == nil || isNil(o.Compound) {
		var ret map[string]interface{}
		return ret
	}
	return o.Compound
}

// GetCompoundOk returns a tuple with the Compound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdbHeader) GetCompoundOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Compound) {
    return map[string]interface{}{}, false
	}
	return o.Compound, true
}

// HasCompound returns a boolean if a field has been set.
func (o *PdbHeader) HasCompound() bool {
	if o != nil && !isNil(o.Compound) {
		return true
	}

	return false
}

// SetCompound gets a reference to the given map[string]interface{} and assigns it to the Compound field.
func (o *PdbHeader) SetCompound(v map[string]interface{}) {
	o.Compound = v
}

// GetPdbId returns the PdbId field value
func (o *PdbHeader) GetPdbId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PdbId
}

// GetPdbIdOk returns a tuple with the PdbId field value
// and a boolean to check if the value has been set.
func (o *PdbHeader) GetPdbIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PdbId, true
}

// SetPdbId sets field value
func (o *PdbHeader) SetPdbId(v string) {
	o.PdbId = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PdbHeader) GetSource() map[string]interface{} {
	if o == nil || isNil(o.Source) {
		var ret map[string]interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdbHeader) GetSourceOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Source) {
    return map[string]interface{}{}, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PdbHeader) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given map[string]interface{} and assigns it to the Source field.
func (o *PdbHeader) SetSource(v map[string]interface{}) {
	o.Source = v
}

// GetTitle returns the Title field value
func (o *PdbHeader) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PdbHeader) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PdbHeader) SetTitle(v string) {
	o.Title = v
}

func (o PdbHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Compound) {
		toSerialize["compound"] = o.Compound
	}
	if true {
		toSerialize["pdbId"] = o.PdbId
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullablePdbHeader struct {
	value *PdbHeader
	isSet bool
}

func (v NullablePdbHeader) Get() *PdbHeader {
	return v.value
}

func (v *NullablePdbHeader) Set(val *PdbHeader) {
	v.value = val
	v.isSet = true
}

func (v NullablePdbHeader) IsSet() bool {
	return v.isSet
}

func (v *NullablePdbHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdbHeader(val *PdbHeader) *NullablePdbHeader {
	return &NullablePdbHeader{value: val, isSet: true}
}

func (v NullablePdbHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdbHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


