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

// Vcf struct for Vcf
type Vcf struct {
	// alt
	Alt *string `json:"alt,omitempty"`
	// position
	Position *string `json:"position,omitempty"`
	// ref
	Ref *string `json:"ref,omitempty"`
}

// NewVcf instantiates a new Vcf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcf() *Vcf {
	this := Vcf{}
	return &this
}

// NewVcfWithDefaults instantiates a new Vcf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcfWithDefaults() *Vcf {
	this := Vcf{}
	return &this
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *Vcf) GetAlt() string {
	if o == nil || isNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vcf) GetAltOk() (*string, bool) {
	if o == nil || isNil(o.Alt) {
    return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *Vcf) HasAlt() bool {
	if o != nil && !isNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *Vcf) SetAlt(v string) {
	o.Alt = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *Vcf) GetPosition() string {
	if o == nil || isNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vcf) GetPositionOk() (*string, bool) {
	if o == nil || isNil(o.Position) {
    return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Vcf) HasPosition() bool {
	if o != nil && !isNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *Vcf) SetPosition(v string) {
	o.Position = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Vcf) GetRef() string {
	if o == nil || isNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vcf) GetRefOk() (*string, bool) {
	if o == nil || isNil(o.Ref) {
    return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Vcf) HasRef() bool {
	if o != nil && !isNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Vcf) SetRef(v string) {
	o.Ref = &v
}

func (o Vcf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	if !isNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !isNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	return json.Marshal(toSerialize)
}

type NullableVcf struct {
	value *Vcf
	isSet bool
}

func (v NullableVcf) Get() *Vcf {
	return v.value
}

func (v *NullableVcf) Set(val *Vcf) {
	v.value = val
	v.isSet = true
}

func (v NullableVcf) IsSet() bool {
	return v.isSet
}

func (v *NullableVcf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcf(val *Vcf) *NullableVcf {
	return &NullableVcf{value: val, isSet: true}
}

func (v NullableVcf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


