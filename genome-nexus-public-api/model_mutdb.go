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

// checks if the Mutdb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mutdb{}

// Mutdb struct for Mutdb
type Mutdb struct {
	// alt
	Alt *string `json:"alt,omitempty"`
	// chrom
	Chrom *string `json:"chrom,omitempty"`
	// cosmic_id
	CosmicId *string `json:"cosmicId,omitempty"`
	Hg19 *Hg19 `json:"hg19,omitempty"`
	// mutpred_score
	MutpredScore *float64 `json:"mutpredScore,omitempty"`
	// ref
	Ref *string `json:"ref,omitempty"`
	// rsid
	Rsid *string `json:"rsid,omitempty"`
	// uniprot_id
	UniprotId *string `json:"uniprotId,omitempty"`
}

// NewMutdb instantiates a new Mutdb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutdb() *Mutdb {
	this := Mutdb{}
	return &this
}

// NewMutdbWithDefaults instantiates a new Mutdb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutdbWithDefaults() *Mutdb {
	this := Mutdb{}
	return &this
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *Mutdb) GetAlt() string {
	if o == nil || isNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetAltOk() (*string, bool) {
	if o == nil || isNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *Mutdb) HasAlt() bool {
	if o != nil && !isNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *Mutdb) SetAlt(v string) {
	o.Alt = &v
}

// GetChrom returns the Chrom field value if set, zero value otherwise.
func (o *Mutdb) GetChrom() string {
	if o == nil || isNil(o.Chrom) {
		var ret string
		return ret
	}
	return *o.Chrom
}

// GetChromOk returns a tuple with the Chrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetChromOk() (*string, bool) {
	if o == nil || isNil(o.Chrom) {
		return nil, false
	}
	return o.Chrom, true
}

// HasChrom returns a boolean if a field has been set.
func (o *Mutdb) HasChrom() bool {
	if o != nil && !isNil(o.Chrom) {
		return true
	}

	return false
}

// SetChrom gets a reference to the given string and assigns it to the Chrom field.
func (o *Mutdb) SetChrom(v string) {
	o.Chrom = &v
}

// GetCosmicId returns the CosmicId field value if set, zero value otherwise.
func (o *Mutdb) GetCosmicId() string {
	if o == nil || isNil(o.CosmicId) {
		var ret string
		return ret
	}
	return *o.CosmicId
}

// GetCosmicIdOk returns a tuple with the CosmicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetCosmicIdOk() (*string, bool) {
	if o == nil || isNil(o.CosmicId) {
		return nil, false
	}
	return o.CosmicId, true
}

// HasCosmicId returns a boolean if a field has been set.
func (o *Mutdb) HasCosmicId() bool {
	if o != nil && !isNil(o.CosmicId) {
		return true
	}

	return false
}

// SetCosmicId gets a reference to the given string and assigns it to the CosmicId field.
func (o *Mutdb) SetCosmicId(v string) {
	o.CosmicId = &v
}

// GetHg19 returns the Hg19 field value if set, zero value otherwise.
func (o *Mutdb) GetHg19() Hg19 {
	if o == nil || isNil(o.Hg19) {
		var ret Hg19
		return ret
	}
	return *o.Hg19
}

// GetHg19Ok returns a tuple with the Hg19 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetHg19Ok() (*Hg19, bool) {
	if o == nil || isNil(o.Hg19) {
		return nil, false
	}
	return o.Hg19, true
}

// HasHg19 returns a boolean if a field has been set.
func (o *Mutdb) HasHg19() bool {
	if o != nil && !isNil(o.Hg19) {
		return true
	}

	return false
}

// SetHg19 gets a reference to the given Hg19 and assigns it to the Hg19 field.
func (o *Mutdb) SetHg19(v Hg19) {
	o.Hg19 = &v
}

// GetMutpredScore returns the MutpredScore field value if set, zero value otherwise.
func (o *Mutdb) GetMutpredScore() float64 {
	if o == nil || isNil(o.MutpredScore) {
		var ret float64
		return ret
	}
	return *o.MutpredScore
}

// GetMutpredScoreOk returns a tuple with the MutpredScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetMutpredScoreOk() (*float64, bool) {
	if o == nil || isNil(o.MutpredScore) {
		return nil, false
	}
	return o.MutpredScore, true
}

// HasMutpredScore returns a boolean if a field has been set.
func (o *Mutdb) HasMutpredScore() bool {
	if o != nil && !isNil(o.MutpredScore) {
		return true
	}

	return false
}

// SetMutpredScore gets a reference to the given float64 and assigns it to the MutpredScore field.
func (o *Mutdb) SetMutpredScore(v float64) {
	o.MutpredScore = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Mutdb) GetRef() string {
	if o == nil || isNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetRefOk() (*string, bool) {
	if o == nil || isNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Mutdb) HasRef() bool {
	if o != nil && !isNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Mutdb) SetRef(v string) {
	o.Ref = &v
}

// GetRsid returns the Rsid field value if set, zero value otherwise.
func (o *Mutdb) GetRsid() string {
	if o == nil || isNil(o.Rsid) {
		var ret string
		return ret
	}
	return *o.Rsid
}

// GetRsidOk returns a tuple with the Rsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetRsidOk() (*string, bool) {
	if o == nil || isNil(o.Rsid) {
		return nil, false
	}
	return o.Rsid, true
}

// HasRsid returns a boolean if a field has been set.
func (o *Mutdb) HasRsid() bool {
	if o != nil && !isNil(o.Rsid) {
		return true
	}

	return false
}

// SetRsid gets a reference to the given string and assigns it to the Rsid field.
func (o *Mutdb) SetRsid(v string) {
	o.Rsid = &v
}

// GetUniprotId returns the UniprotId field value if set, zero value otherwise.
func (o *Mutdb) GetUniprotId() string {
	if o == nil || isNil(o.UniprotId) {
		var ret string
		return ret
	}
	return *o.UniprotId
}

// GetUniprotIdOk returns a tuple with the UniprotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mutdb) GetUniprotIdOk() (*string, bool) {
	if o == nil || isNil(o.UniprotId) {
		return nil, false
	}
	return o.UniprotId, true
}

// HasUniprotId returns a boolean if a field has been set.
func (o *Mutdb) HasUniprotId() bool {
	if o != nil && !isNil(o.UniprotId) {
		return true
	}

	return false
}

// SetUniprotId gets a reference to the given string and assigns it to the UniprotId field.
func (o *Mutdb) SetUniprotId(v string) {
	o.UniprotId = &v
}

func (o Mutdb) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mutdb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	if !isNil(o.Chrom) {
		toSerialize["chrom"] = o.Chrom
	}
	if !isNil(o.CosmicId) {
		toSerialize["cosmicId"] = o.CosmicId
	}
	if !isNil(o.Hg19) {
		toSerialize["hg19"] = o.Hg19
	}
	if !isNil(o.MutpredScore) {
		toSerialize["mutpredScore"] = o.MutpredScore
	}
	if !isNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !isNil(o.Rsid) {
		toSerialize["rsid"] = o.Rsid
	}
	if !isNil(o.UniprotId) {
		toSerialize["uniprotId"] = o.UniprotId
	}
	return toSerialize, nil
}

type NullableMutdb struct {
	value *Mutdb
	isSet bool
}

func (v NullableMutdb) Get() *Mutdb {
	return v.value
}

func (v *NullableMutdb) Set(val *Mutdb) {
	v.value = val
	v.isSet = true
}

func (v NullableMutdb) IsSet() bool {
	return v.isSet
}

func (v *NullableMutdb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutdb(val *Mutdb) *NullableMutdb {
	return &NullableMutdb{value: val, isSet: true}
}

func (v NullableMutdb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutdb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

