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

// NucleotideContext struct for NucleotideContext
type NucleotideContext struct {
	Hgvs *string `json:"hgvs,omitempty"`
	Id *string `json:"id,omitempty"`
	Molecule *string `json:"molecule,omitempty"`
	Query *string `json:"query,omitempty"`
	// Nucleotide context sequence
	Seq string `json:"seq"`
}

// NewNucleotideContext instantiates a new NucleotideContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNucleotideContext(seq string) *NucleotideContext {
	this := NucleotideContext{}
	this.Seq = seq
	return &this
}

// NewNucleotideContextWithDefaults instantiates a new NucleotideContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNucleotideContextWithDefaults() *NucleotideContext {
	this := NucleotideContext{}
	return &this
}

// GetHgvs returns the Hgvs field value if set, zero value otherwise.
func (o *NucleotideContext) GetHgvs() string {
	if o == nil || isNil(o.Hgvs) {
		var ret string
		return ret
	}
	return *o.Hgvs
}

// GetHgvsOk returns a tuple with the Hgvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NucleotideContext) GetHgvsOk() (*string, bool) {
	if o == nil || isNil(o.Hgvs) {
    return nil, false
	}
	return o.Hgvs, true
}

// HasHgvs returns a boolean if a field has been set.
func (o *NucleotideContext) HasHgvs() bool {
	if o != nil && !isNil(o.Hgvs) {
		return true
	}

	return false
}

// SetHgvs gets a reference to the given string and assigns it to the Hgvs field.
func (o *NucleotideContext) SetHgvs(v string) {
	o.Hgvs = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NucleotideContext) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NucleotideContext) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NucleotideContext) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NucleotideContext) SetId(v string) {
	o.Id = &v
}

// GetMolecule returns the Molecule field value if set, zero value otherwise.
func (o *NucleotideContext) GetMolecule() string {
	if o == nil || isNil(o.Molecule) {
		var ret string
		return ret
	}
	return *o.Molecule
}

// GetMoleculeOk returns a tuple with the Molecule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NucleotideContext) GetMoleculeOk() (*string, bool) {
	if o == nil || isNil(o.Molecule) {
    return nil, false
	}
	return o.Molecule, true
}

// HasMolecule returns a boolean if a field has been set.
func (o *NucleotideContext) HasMolecule() bool {
	if o != nil && !isNil(o.Molecule) {
		return true
	}

	return false
}

// SetMolecule gets a reference to the given string and assigns it to the Molecule field.
func (o *NucleotideContext) SetMolecule(v string) {
	o.Molecule = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *NucleotideContext) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NucleotideContext) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
    return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *NucleotideContext) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *NucleotideContext) SetQuery(v string) {
	o.Query = &v
}

// GetSeq returns the Seq field value
func (o *NucleotideContext) GetSeq() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seq
}

// GetSeqOk returns a tuple with the Seq field value
// and a boolean to check if the value has been set.
func (o *NucleotideContext) GetSeqOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Seq, true
}

// SetSeq sets field value
func (o *NucleotideContext) SetSeq(v string) {
	o.Seq = v
}

func (o NucleotideContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Hgvs) {
		toSerialize["hgvs"] = o.Hgvs
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Molecule) {
		toSerialize["molecule"] = o.Molecule
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["seq"] = o.Seq
	}
	return json.Marshal(toSerialize)
}

type NullableNucleotideContext struct {
	value *NucleotideContext
	isSet bool
}

func (v NullableNucleotideContext) Get() *NucleotideContext {
	return v.value
}

func (v *NullableNucleotideContext) Set(val *NucleotideContext) {
	v.value = val
	v.isSet = true
}

func (v NullableNucleotideContext) IsSet() bool {
	return v.isSet
}

func (v *NullableNucleotideContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNucleotideContext(val *NucleotideContext) *NullableNucleotideContext {
	return &NullableNucleotideContext{value: val, isSet: true}
}

func (v NullableNucleotideContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNucleotideContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


