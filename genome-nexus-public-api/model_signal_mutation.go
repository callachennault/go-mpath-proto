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

// SignalMutation struct for SignalMutation
type SignalMutation struct {
	// Biallelic Counts by Tumor Type
	BiallelicCountsByTumorType []CountByTumorType `json:"biallelicCountsByTumorType,omitempty"`
	// Chromosome
	Chromosome *string `json:"chromosome,omitempty"`
	// Counts by Tumor Type
	CountsByTumorType []CountByTumorType `json:"countsByTumorType,omitempty"`
	// End Position
	EndPosition *int64 `json:"endPosition,omitempty"`
	GeneralPopulationStats *GeneralPopulationStats `json:"generalPopulationStats,omitempty"`
	// Hugo Gene Symbol
	HugoGeneSymbol *string `json:"hugoGeneSymbol,omitempty"`
	// Msk Expert Review
	MskExperReview *bool `json:"mskExperReview,omitempty"`
	// Mutation Status
	MutationStatus *string `json:"mutationStatus,omitempty"`
	OverallNumberOfGermlineHomozygous *int32 `json:"overallNumberOfGermlineHomozygous,omitempty"`
	// Pathogenic
	Pathogenic *string `json:"pathogenic,omitempty"`
	// Penetrance
	Penetrance *string `json:"penetrance,omitempty"`
	// QC Pass Counts by Tumor Type
	QcPassCountsByTumorType []CountByTumorType `json:"qcPassCountsByTumorType,omitempty"`
	// Reference Allele
	ReferenceAllele *string `json:"referenceAllele,omitempty"`
	// Start Position
	StartPosition *int64 `json:"startPosition,omitempty"`
	// Stats By Tumor Type
	StatsByTumorType []StatsByTumorType `json:"statsByTumorType,omitempty"`
	// Variant Allele
	VariantAllele *string `json:"variantAllele,omitempty"`
}

// NewSignalMutation instantiates a new SignalMutation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalMutation() *SignalMutation {
	this := SignalMutation{}
	return &this
}

// NewSignalMutationWithDefaults instantiates a new SignalMutation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalMutationWithDefaults() *SignalMutation {
	this := SignalMutation{}
	return &this
}

// GetBiallelicCountsByTumorType returns the BiallelicCountsByTumorType field value if set, zero value otherwise.
func (o *SignalMutation) GetBiallelicCountsByTumorType() []CountByTumorType {
	if o == nil || isNil(o.BiallelicCountsByTumorType) {
		var ret []CountByTumorType
		return ret
	}
	return o.BiallelicCountsByTumorType
}

// GetBiallelicCountsByTumorTypeOk returns a tuple with the BiallelicCountsByTumorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetBiallelicCountsByTumorTypeOk() ([]CountByTumorType, bool) {
	if o == nil || isNil(o.BiallelicCountsByTumorType) {
    return nil, false
	}
	return o.BiallelicCountsByTumorType, true
}

// HasBiallelicCountsByTumorType returns a boolean if a field has been set.
func (o *SignalMutation) HasBiallelicCountsByTumorType() bool {
	if o != nil && !isNil(o.BiallelicCountsByTumorType) {
		return true
	}

	return false
}

// SetBiallelicCountsByTumorType gets a reference to the given []CountByTumorType and assigns it to the BiallelicCountsByTumorType field.
func (o *SignalMutation) SetBiallelicCountsByTumorType(v []CountByTumorType) {
	o.BiallelicCountsByTumorType = v
}

// GetChromosome returns the Chromosome field value if set, zero value otherwise.
func (o *SignalMutation) GetChromosome() string {
	if o == nil || isNil(o.Chromosome) {
		var ret string
		return ret
	}
	return *o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetChromosomeOk() (*string, bool) {
	if o == nil || isNil(o.Chromosome) {
    return nil, false
	}
	return o.Chromosome, true
}

// HasChromosome returns a boolean if a field has been set.
func (o *SignalMutation) HasChromosome() bool {
	if o != nil && !isNil(o.Chromosome) {
		return true
	}

	return false
}

// SetChromosome gets a reference to the given string and assigns it to the Chromosome field.
func (o *SignalMutation) SetChromosome(v string) {
	o.Chromosome = &v
}

// GetCountsByTumorType returns the CountsByTumorType field value if set, zero value otherwise.
func (o *SignalMutation) GetCountsByTumorType() []CountByTumorType {
	if o == nil || isNil(o.CountsByTumorType) {
		var ret []CountByTumorType
		return ret
	}
	return o.CountsByTumorType
}

// GetCountsByTumorTypeOk returns a tuple with the CountsByTumorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetCountsByTumorTypeOk() ([]CountByTumorType, bool) {
	if o == nil || isNil(o.CountsByTumorType) {
    return nil, false
	}
	return o.CountsByTumorType, true
}

// HasCountsByTumorType returns a boolean if a field has been set.
func (o *SignalMutation) HasCountsByTumorType() bool {
	if o != nil && !isNil(o.CountsByTumorType) {
		return true
	}

	return false
}

// SetCountsByTumorType gets a reference to the given []CountByTumorType and assigns it to the CountsByTumorType field.
func (o *SignalMutation) SetCountsByTumorType(v []CountByTumorType) {
	o.CountsByTumorType = v
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *SignalMutation) GetEndPosition() int64 {
	if o == nil || isNil(o.EndPosition) {
		var ret int64
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetEndPositionOk() (*int64, bool) {
	if o == nil || isNil(o.EndPosition) {
    return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *SignalMutation) HasEndPosition() bool {
	if o != nil && !isNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given int64 and assigns it to the EndPosition field.
func (o *SignalMutation) SetEndPosition(v int64) {
	o.EndPosition = &v
}

// GetGeneralPopulationStats returns the GeneralPopulationStats field value if set, zero value otherwise.
func (o *SignalMutation) GetGeneralPopulationStats() GeneralPopulationStats {
	if o == nil || isNil(o.GeneralPopulationStats) {
		var ret GeneralPopulationStats
		return ret
	}
	return *o.GeneralPopulationStats
}

// GetGeneralPopulationStatsOk returns a tuple with the GeneralPopulationStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetGeneralPopulationStatsOk() (*GeneralPopulationStats, bool) {
	if o == nil || isNil(o.GeneralPopulationStats) {
    return nil, false
	}
	return o.GeneralPopulationStats, true
}

// HasGeneralPopulationStats returns a boolean if a field has been set.
func (o *SignalMutation) HasGeneralPopulationStats() bool {
	if o != nil && !isNil(o.GeneralPopulationStats) {
		return true
	}

	return false
}

// SetGeneralPopulationStats gets a reference to the given GeneralPopulationStats and assigns it to the GeneralPopulationStats field.
func (o *SignalMutation) SetGeneralPopulationStats(v GeneralPopulationStats) {
	o.GeneralPopulationStats = &v
}

// GetHugoGeneSymbol returns the HugoGeneSymbol field value if set, zero value otherwise.
func (o *SignalMutation) GetHugoGeneSymbol() string {
	if o == nil || isNil(o.HugoGeneSymbol) {
		var ret string
		return ret
	}
	return *o.HugoGeneSymbol
}

// GetHugoGeneSymbolOk returns a tuple with the HugoGeneSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetHugoGeneSymbolOk() (*string, bool) {
	if o == nil || isNil(o.HugoGeneSymbol) {
    return nil, false
	}
	return o.HugoGeneSymbol, true
}

// HasHugoGeneSymbol returns a boolean if a field has been set.
func (o *SignalMutation) HasHugoGeneSymbol() bool {
	if o != nil && !isNil(o.HugoGeneSymbol) {
		return true
	}

	return false
}

// SetHugoGeneSymbol gets a reference to the given string and assigns it to the HugoGeneSymbol field.
func (o *SignalMutation) SetHugoGeneSymbol(v string) {
	o.HugoGeneSymbol = &v
}

// GetMskExperReview returns the MskExperReview field value if set, zero value otherwise.
func (o *SignalMutation) GetMskExperReview() bool {
	if o == nil || isNil(o.MskExperReview) {
		var ret bool
		return ret
	}
	return *o.MskExperReview
}

// GetMskExperReviewOk returns a tuple with the MskExperReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetMskExperReviewOk() (*bool, bool) {
	if o == nil || isNil(o.MskExperReview) {
    return nil, false
	}
	return o.MskExperReview, true
}

// HasMskExperReview returns a boolean if a field has been set.
func (o *SignalMutation) HasMskExperReview() bool {
	if o != nil && !isNil(o.MskExperReview) {
		return true
	}

	return false
}

// SetMskExperReview gets a reference to the given bool and assigns it to the MskExperReview field.
func (o *SignalMutation) SetMskExperReview(v bool) {
	o.MskExperReview = &v
}

// GetMutationStatus returns the MutationStatus field value if set, zero value otherwise.
func (o *SignalMutation) GetMutationStatus() string {
	if o == nil || isNil(o.MutationStatus) {
		var ret string
		return ret
	}
	return *o.MutationStatus
}

// GetMutationStatusOk returns a tuple with the MutationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetMutationStatusOk() (*string, bool) {
	if o == nil || isNil(o.MutationStatus) {
    return nil, false
	}
	return o.MutationStatus, true
}

// HasMutationStatus returns a boolean if a field has been set.
func (o *SignalMutation) HasMutationStatus() bool {
	if o != nil && !isNil(o.MutationStatus) {
		return true
	}

	return false
}

// SetMutationStatus gets a reference to the given string and assigns it to the MutationStatus field.
func (o *SignalMutation) SetMutationStatus(v string) {
	o.MutationStatus = &v
}

// GetOverallNumberOfGermlineHomozygous returns the OverallNumberOfGermlineHomozygous field value if set, zero value otherwise.
func (o *SignalMutation) GetOverallNumberOfGermlineHomozygous() int32 {
	if o == nil || isNil(o.OverallNumberOfGermlineHomozygous) {
		var ret int32
		return ret
	}
	return *o.OverallNumberOfGermlineHomozygous
}

// GetOverallNumberOfGermlineHomozygousOk returns a tuple with the OverallNumberOfGermlineHomozygous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetOverallNumberOfGermlineHomozygousOk() (*int32, bool) {
	if o == nil || isNil(o.OverallNumberOfGermlineHomozygous) {
    return nil, false
	}
	return o.OverallNumberOfGermlineHomozygous, true
}

// HasOverallNumberOfGermlineHomozygous returns a boolean if a field has been set.
func (o *SignalMutation) HasOverallNumberOfGermlineHomozygous() bool {
	if o != nil && !isNil(o.OverallNumberOfGermlineHomozygous) {
		return true
	}

	return false
}

// SetOverallNumberOfGermlineHomozygous gets a reference to the given int32 and assigns it to the OverallNumberOfGermlineHomozygous field.
func (o *SignalMutation) SetOverallNumberOfGermlineHomozygous(v int32) {
	o.OverallNumberOfGermlineHomozygous = &v
}

// GetPathogenic returns the Pathogenic field value if set, zero value otherwise.
func (o *SignalMutation) GetPathogenic() string {
	if o == nil || isNil(o.Pathogenic) {
		var ret string
		return ret
	}
	return *o.Pathogenic
}

// GetPathogenicOk returns a tuple with the Pathogenic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetPathogenicOk() (*string, bool) {
	if o == nil || isNil(o.Pathogenic) {
    return nil, false
	}
	return o.Pathogenic, true
}

// HasPathogenic returns a boolean if a field has been set.
func (o *SignalMutation) HasPathogenic() bool {
	if o != nil && !isNil(o.Pathogenic) {
		return true
	}

	return false
}

// SetPathogenic gets a reference to the given string and assigns it to the Pathogenic field.
func (o *SignalMutation) SetPathogenic(v string) {
	o.Pathogenic = &v
}

// GetPenetrance returns the Penetrance field value if set, zero value otherwise.
func (o *SignalMutation) GetPenetrance() string {
	if o == nil || isNil(o.Penetrance) {
		var ret string
		return ret
	}
	return *o.Penetrance
}

// GetPenetranceOk returns a tuple with the Penetrance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetPenetranceOk() (*string, bool) {
	if o == nil || isNil(o.Penetrance) {
    return nil, false
	}
	return o.Penetrance, true
}

// HasPenetrance returns a boolean if a field has been set.
func (o *SignalMutation) HasPenetrance() bool {
	if o != nil && !isNil(o.Penetrance) {
		return true
	}

	return false
}

// SetPenetrance gets a reference to the given string and assigns it to the Penetrance field.
func (o *SignalMutation) SetPenetrance(v string) {
	o.Penetrance = &v
}

// GetQcPassCountsByTumorType returns the QcPassCountsByTumorType field value if set, zero value otherwise.
func (o *SignalMutation) GetQcPassCountsByTumorType() []CountByTumorType {
	if o == nil || isNil(o.QcPassCountsByTumorType) {
		var ret []CountByTumorType
		return ret
	}
	return o.QcPassCountsByTumorType
}

// GetQcPassCountsByTumorTypeOk returns a tuple with the QcPassCountsByTumorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetQcPassCountsByTumorTypeOk() ([]CountByTumorType, bool) {
	if o == nil || isNil(o.QcPassCountsByTumorType) {
    return nil, false
	}
	return o.QcPassCountsByTumorType, true
}

// HasQcPassCountsByTumorType returns a boolean if a field has been set.
func (o *SignalMutation) HasQcPassCountsByTumorType() bool {
	if o != nil && !isNil(o.QcPassCountsByTumorType) {
		return true
	}

	return false
}

// SetQcPassCountsByTumorType gets a reference to the given []CountByTumorType and assigns it to the QcPassCountsByTumorType field.
func (o *SignalMutation) SetQcPassCountsByTumorType(v []CountByTumorType) {
	o.QcPassCountsByTumorType = v
}

// GetReferenceAllele returns the ReferenceAllele field value if set, zero value otherwise.
func (o *SignalMutation) GetReferenceAllele() string {
	if o == nil || isNil(o.ReferenceAllele) {
		var ret string
		return ret
	}
	return *o.ReferenceAllele
}

// GetReferenceAlleleOk returns a tuple with the ReferenceAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetReferenceAlleleOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceAllele) {
    return nil, false
	}
	return o.ReferenceAllele, true
}

// HasReferenceAllele returns a boolean if a field has been set.
func (o *SignalMutation) HasReferenceAllele() bool {
	if o != nil && !isNil(o.ReferenceAllele) {
		return true
	}

	return false
}

// SetReferenceAllele gets a reference to the given string and assigns it to the ReferenceAllele field.
func (o *SignalMutation) SetReferenceAllele(v string) {
	o.ReferenceAllele = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *SignalMutation) GetStartPosition() int64 {
	if o == nil || isNil(o.StartPosition) {
		var ret int64
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetStartPositionOk() (*int64, bool) {
	if o == nil || isNil(o.StartPosition) {
    return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *SignalMutation) HasStartPosition() bool {
	if o != nil && !isNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given int64 and assigns it to the StartPosition field.
func (o *SignalMutation) SetStartPosition(v int64) {
	o.StartPosition = &v
}

// GetStatsByTumorType returns the StatsByTumorType field value if set, zero value otherwise.
func (o *SignalMutation) GetStatsByTumorType() []StatsByTumorType {
	if o == nil || isNil(o.StatsByTumorType) {
		var ret []StatsByTumorType
		return ret
	}
	return o.StatsByTumorType
}

// GetStatsByTumorTypeOk returns a tuple with the StatsByTumorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetStatsByTumorTypeOk() ([]StatsByTumorType, bool) {
	if o == nil || isNil(o.StatsByTumorType) {
    return nil, false
	}
	return o.StatsByTumorType, true
}

// HasStatsByTumorType returns a boolean if a field has been set.
func (o *SignalMutation) HasStatsByTumorType() bool {
	if o != nil && !isNil(o.StatsByTumorType) {
		return true
	}

	return false
}

// SetStatsByTumorType gets a reference to the given []StatsByTumorType and assigns it to the StatsByTumorType field.
func (o *SignalMutation) SetStatsByTumorType(v []StatsByTumorType) {
	o.StatsByTumorType = v
}

// GetVariantAllele returns the VariantAllele field value if set, zero value otherwise.
func (o *SignalMutation) GetVariantAllele() string {
	if o == nil || isNil(o.VariantAllele) {
		var ret string
		return ret
	}
	return *o.VariantAllele
}

// GetVariantAlleleOk returns a tuple with the VariantAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalMutation) GetVariantAlleleOk() (*string, bool) {
	if o == nil || isNil(o.VariantAllele) {
    return nil, false
	}
	return o.VariantAllele, true
}

// HasVariantAllele returns a boolean if a field has been set.
func (o *SignalMutation) HasVariantAllele() bool {
	if o != nil && !isNil(o.VariantAllele) {
		return true
	}

	return false
}

// SetVariantAllele gets a reference to the given string and assigns it to the VariantAllele field.
func (o *SignalMutation) SetVariantAllele(v string) {
	o.VariantAllele = &v
}

func (o SignalMutation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BiallelicCountsByTumorType) {
		toSerialize["biallelicCountsByTumorType"] = o.BiallelicCountsByTumorType
	}
	if !isNil(o.Chromosome) {
		toSerialize["chromosome"] = o.Chromosome
	}
	if !isNil(o.CountsByTumorType) {
		toSerialize["countsByTumorType"] = o.CountsByTumorType
	}
	if !isNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
	}
	if !isNil(o.GeneralPopulationStats) {
		toSerialize["generalPopulationStats"] = o.GeneralPopulationStats
	}
	if !isNil(o.HugoGeneSymbol) {
		toSerialize["hugoGeneSymbol"] = o.HugoGeneSymbol
	}
	if !isNil(o.MskExperReview) {
		toSerialize["mskExperReview"] = o.MskExperReview
	}
	if !isNil(o.MutationStatus) {
		toSerialize["mutationStatus"] = o.MutationStatus
	}
	if !isNil(o.OverallNumberOfGermlineHomozygous) {
		toSerialize["overallNumberOfGermlineHomozygous"] = o.OverallNumberOfGermlineHomozygous
	}
	if !isNil(o.Pathogenic) {
		toSerialize["pathogenic"] = o.Pathogenic
	}
	if !isNil(o.Penetrance) {
		toSerialize["penetrance"] = o.Penetrance
	}
	if !isNil(o.QcPassCountsByTumorType) {
		toSerialize["qcPassCountsByTumorType"] = o.QcPassCountsByTumorType
	}
	if !isNil(o.ReferenceAllele) {
		toSerialize["referenceAllele"] = o.ReferenceAllele
	}
	if !isNil(o.StartPosition) {
		toSerialize["startPosition"] = o.StartPosition
	}
	if !isNil(o.StatsByTumorType) {
		toSerialize["statsByTumorType"] = o.StatsByTumorType
	}
	if !isNil(o.VariantAllele) {
		toSerialize["variantAllele"] = o.VariantAllele
	}
	return json.Marshal(toSerialize)
}

type NullableSignalMutation struct {
	value *SignalMutation
	isSet bool
}

func (v NullableSignalMutation) Get() *SignalMutation {
	return v.value
}

func (v *NullableSignalMutation) Set(val *SignalMutation) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalMutation) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalMutation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalMutation(val *SignalMutation) *NullableSignalMutation {
	return &NullableSignalMutation{value: val, isSet: true}
}

func (v NullableSignalMutation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalMutation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

