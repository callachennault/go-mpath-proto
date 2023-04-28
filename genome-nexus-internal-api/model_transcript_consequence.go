/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_internal_api

import (
	"encoding/json"
)

// TranscriptConsequence struct for TranscriptConsequence
type TranscriptConsequence struct {
	// Amino acids
	AminoAcids *string `json:"amino_acids,omitempty"`
	// Canonical transcript indicator
	Canonical *string `json:"canonical,omitempty"`
	// Codons
	Codons *string `json:"codons,omitempty"`
	// List of consequence terms
	ConsequenceTerms []string `json:"consequence_terms,omitempty"`
	Exon *string `json:"exon,omitempty"`
	// Ensembl gene id
	GeneId *string `json:"gene_id,omitempty"`
	// Hugo gene symbol
	GeneSymbol *string `json:"gene_symbol,omitempty"`
	// HGNC id
	HgncId *string `json:"hgnc_id,omitempty"`
	// HGVSc
	Hgvsc *string `json:"hgvsc,omitempty"`
	// HGVSg
	Hgvsg *string `json:"hgvsg,omitempty"`
	// HGVSp
	Hgvsp *string `json:"hgvsp,omitempty"`
	// Polyphen Prediction
	PolyphenPrediction *string `json:"polyphen_prediction,omitempty"`
	// Polyphen Score
	PolyphenScore *float64 `json:"polyphen_score,omitempty"`
	// Protein end position
	ProteinEnd *int32 `json:"protein_end,omitempty"`
	// Ensembl protein id
	ProteinId *string `json:"protein_id,omitempty"`
	// Protein start position
	ProteinStart *int32 `json:"protein_start,omitempty"`
	// List of RefSeq transcript ids
	RefseqTranscriptIds []string `json:"refseq_transcript_ids,omitempty"`
	// Sift Prediction
	SiftPrediction *string `json:"sift_prediction,omitempty"`
	// Sift Score
	SiftScore *float64 `json:"sift_score,omitempty"`
	// Ensembl transcript id
	TranscriptId string `json:"transcript_id"`
	UniprotId *string `json:"uniprotId,omitempty"`
	// Variant allele
	VariantAllele *string `json:"variant_allele,omitempty"`
}

// NewTranscriptConsequence instantiates a new TranscriptConsequence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscriptConsequence(transcriptId string) *TranscriptConsequence {
	this := TranscriptConsequence{}
	this.TranscriptId = transcriptId
	return &this
}

// NewTranscriptConsequenceWithDefaults instantiates a new TranscriptConsequence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscriptConsequenceWithDefaults() *TranscriptConsequence {
	this := TranscriptConsequence{}
	return &this
}

// GetAminoAcids returns the AminoAcids field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetAminoAcids() string {
	if o == nil || isNil(o.AminoAcids) {
		var ret string
		return ret
	}
	return *o.AminoAcids
}

// GetAminoAcidsOk returns a tuple with the AminoAcids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetAminoAcidsOk() (*string, bool) {
	if o == nil || isNil(o.AminoAcids) {
    return nil, false
	}
	return o.AminoAcids, true
}

// HasAminoAcids returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasAminoAcids() bool {
	if o != nil && !isNil(o.AminoAcids) {
		return true
	}

	return false
}

// SetAminoAcids gets a reference to the given string and assigns it to the AminoAcids field.
func (o *TranscriptConsequence) SetAminoAcids(v string) {
	o.AminoAcids = &v
}

// GetCanonical returns the Canonical field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetCanonical() string {
	if o == nil || isNil(o.Canonical) {
		var ret string
		return ret
	}
	return *o.Canonical
}

// GetCanonicalOk returns a tuple with the Canonical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetCanonicalOk() (*string, bool) {
	if o == nil || isNil(o.Canonical) {
    return nil, false
	}
	return o.Canonical, true
}

// HasCanonical returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasCanonical() bool {
	if o != nil && !isNil(o.Canonical) {
		return true
	}

	return false
}

// SetCanonical gets a reference to the given string and assigns it to the Canonical field.
func (o *TranscriptConsequence) SetCanonical(v string) {
	o.Canonical = &v
}

// GetCodons returns the Codons field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetCodons() string {
	if o == nil || isNil(o.Codons) {
		var ret string
		return ret
	}
	return *o.Codons
}

// GetCodonsOk returns a tuple with the Codons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetCodonsOk() (*string, bool) {
	if o == nil || isNil(o.Codons) {
    return nil, false
	}
	return o.Codons, true
}

// HasCodons returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasCodons() bool {
	if o != nil && !isNil(o.Codons) {
		return true
	}

	return false
}

// SetCodons gets a reference to the given string and assigns it to the Codons field.
func (o *TranscriptConsequence) SetCodons(v string) {
	o.Codons = &v
}

// GetConsequenceTerms returns the ConsequenceTerms field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetConsequenceTerms() []string {
	if o == nil || isNil(o.ConsequenceTerms) {
		var ret []string
		return ret
	}
	return o.ConsequenceTerms
}

// GetConsequenceTermsOk returns a tuple with the ConsequenceTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetConsequenceTermsOk() ([]string, bool) {
	if o == nil || isNil(o.ConsequenceTerms) {
    return nil, false
	}
	return o.ConsequenceTerms, true
}

// HasConsequenceTerms returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasConsequenceTerms() bool {
	if o != nil && !isNil(o.ConsequenceTerms) {
		return true
	}

	return false
}

// SetConsequenceTerms gets a reference to the given []string and assigns it to the ConsequenceTerms field.
func (o *TranscriptConsequence) SetConsequenceTerms(v []string) {
	o.ConsequenceTerms = v
}

// GetExon returns the Exon field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetExon() string {
	if o == nil || isNil(o.Exon) {
		var ret string
		return ret
	}
	return *o.Exon
}

// GetExonOk returns a tuple with the Exon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetExonOk() (*string, bool) {
	if o == nil || isNil(o.Exon) {
    return nil, false
	}
	return o.Exon, true
}

// HasExon returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasExon() bool {
	if o != nil && !isNil(o.Exon) {
		return true
	}

	return false
}

// SetExon gets a reference to the given string and assigns it to the Exon field.
func (o *TranscriptConsequence) SetExon(v string) {
	o.Exon = &v
}

// GetGeneId returns the GeneId field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetGeneId() string {
	if o == nil || isNil(o.GeneId) {
		var ret string
		return ret
	}
	return *o.GeneId
}

// GetGeneIdOk returns a tuple with the GeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetGeneIdOk() (*string, bool) {
	if o == nil || isNil(o.GeneId) {
    return nil, false
	}
	return o.GeneId, true
}

// HasGeneId returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasGeneId() bool {
	if o != nil && !isNil(o.GeneId) {
		return true
	}

	return false
}

// SetGeneId gets a reference to the given string and assigns it to the GeneId field.
func (o *TranscriptConsequence) SetGeneId(v string) {
	o.GeneId = &v
}

// GetGeneSymbol returns the GeneSymbol field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetGeneSymbol() string {
	if o == nil || isNil(o.GeneSymbol) {
		var ret string
		return ret
	}
	return *o.GeneSymbol
}

// GetGeneSymbolOk returns a tuple with the GeneSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetGeneSymbolOk() (*string, bool) {
	if o == nil || isNil(o.GeneSymbol) {
    return nil, false
	}
	return o.GeneSymbol, true
}

// HasGeneSymbol returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasGeneSymbol() bool {
	if o != nil && !isNil(o.GeneSymbol) {
		return true
	}

	return false
}

// SetGeneSymbol gets a reference to the given string and assigns it to the GeneSymbol field.
func (o *TranscriptConsequence) SetGeneSymbol(v string) {
	o.GeneSymbol = &v
}

// GetHgncId returns the HgncId field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetHgncId() string {
	if o == nil || isNil(o.HgncId) {
		var ret string
		return ret
	}
	return *o.HgncId
}

// GetHgncIdOk returns a tuple with the HgncId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetHgncIdOk() (*string, bool) {
	if o == nil || isNil(o.HgncId) {
    return nil, false
	}
	return o.HgncId, true
}

// HasHgncId returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasHgncId() bool {
	if o != nil && !isNil(o.HgncId) {
		return true
	}

	return false
}

// SetHgncId gets a reference to the given string and assigns it to the HgncId field.
func (o *TranscriptConsequence) SetHgncId(v string) {
	o.HgncId = &v
}

// GetHgvsc returns the Hgvsc field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetHgvsc() string {
	if o == nil || isNil(o.Hgvsc) {
		var ret string
		return ret
	}
	return *o.Hgvsc
}

// GetHgvscOk returns a tuple with the Hgvsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetHgvscOk() (*string, bool) {
	if o == nil || isNil(o.Hgvsc) {
    return nil, false
	}
	return o.Hgvsc, true
}

// HasHgvsc returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasHgvsc() bool {
	if o != nil && !isNil(o.Hgvsc) {
		return true
	}

	return false
}

// SetHgvsc gets a reference to the given string and assigns it to the Hgvsc field.
func (o *TranscriptConsequence) SetHgvsc(v string) {
	o.Hgvsc = &v
}

// GetHgvsg returns the Hgvsg field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetHgvsg() string {
	if o == nil || isNil(o.Hgvsg) {
		var ret string
		return ret
	}
	return *o.Hgvsg
}

// GetHgvsgOk returns a tuple with the Hgvsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetHgvsgOk() (*string, bool) {
	if o == nil || isNil(o.Hgvsg) {
    return nil, false
	}
	return o.Hgvsg, true
}

// HasHgvsg returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasHgvsg() bool {
	if o != nil && !isNil(o.Hgvsg) {
		return true
	}

	return false
}

// SetHgvsg gets a reference to the given string and assigns it to the Hgvsg field.
func (o *TranscriptConsequence) SetHgvsg(v string) {
	o.Hgvsg = &v
}

// GetHgvsp returns the Hgvsp field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetHgvsp() string {
	if o == nil || isNil(o.Hgvsp) {
		var ret string
		return ret
	}
	return *o.Hgvsp
}

// GetHgvspOk returns a tuple with the Hgvsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetHgvspOk() (*string, bool) {
	if o == nil || isNil(o.Hgvsp) {
    return nil, false
	}
	return o.Hgvsp, true
}

// HasHgvsp returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasHgvsp() bool {
	if o != nil && !isNil(o.Hgvsp) {
		return true
	}

	return false
}

// SetHgvsp gets a reference to the given string and assigns it to the Hgvsp field.
func (o *TranscriptConsequence) SetHgvsp(v string) {
	o.Hgvsp = &v
}

// GetPolyphenPrediction returns the PolyphenPrediction field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetPolyphenPrediction() string {
	if o == nil || isNil(o.PolyphenPrediction) {
		var ret string
		return ret
	}
	return *o.PolyphenPrediction
}

// GetPolyphenPredictionOk returns a tuple with the PolyphenPrediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetPolyphenPredictionOk() (*string, bool) {
	if o == nil || isNil(o.PolyphenPrediction) {
    return nil, false
	}
	return o.PolyphenPrediction, true
}

// HasPolyphenPrediction returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasPolyphenPrediction() bool {
	if o != nil && !isNil(o.PolyphenPrediction) {
		return true
	}

	return false
}

// SetPolyphenPrediction gets a reference to the given string and assigns it to the PolyphenPrediction field.
func (o *TranscriptConsequence) SetPolyphenPrediction(v string) {
	o.PolyphenPrediction = &v
}

// GetPolyphenScore returns the PolyphenScore field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetPolyphenScore() float64 {
	if o == nil || isNil(o.PolyphenScore) {
		var ret float64
		return ret
	}
	return *o.PolyphenScore
}

// GetPolyphenScoreOk returns a tuple with the PolyphenScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetPolyphenScoreOk() (*float64, bool) {
	if o == nil || isNil(o.PolyphenScore) {
    return nil, false
	}
	return o.PolyphenScore, true
}

// HasPolyphenScore returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasPolyphenScore() bool {
	if o != nil && !isNil(o.PolyphenScore) {
		return true
	}

	return false
}

// SetPolyphenScore gets a reference to the given float64 and assigns it to the PolyphenScore field.
func (o *TranscriptConsequence) SetPolyphenScore(v float64) {
	o.PolyphenScore = &v
}

// GetProteinEnd returns the ProteinEnd field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetProteinEnd() int32 {
	if o == nil || isNil(o.ProteinEnd) {
		var ret int32
		return ret
	}
	return *o.ProteinEnd
}

// GetProteinEndOk returns a tuple with the ProteinEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetProteinEndOk() (*int32, bool) {
	if o == nil || isNil(o.ProteinEnd) {
    return nil, false
	}
	return o.ProteinEnd, true
}

// HasProteinEnd returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasProteinEnd() bool {
	if o != nil && !isNil(o.ProteinEnd) {
		return true
	}

	return false
}

// SetProteinEnd gets a reference to the given int32 and assigns it to the ProteinEnd field.
func (o *TranscriptConsequence) SetProteinEnd(v int32) {
	o.ProteinEnd = &v
}

// GetProteinId returns the ProteinId field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetProteinId() string {
	if o == nil || isNil(o.ProteinId) {
		var ret string
		return ret
	}
	return *o.ProteinId
}

// GetProteinIdOk returns a tuple with the ProteinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetProteinIdOk() (*string, bool) {
	if o == nil || isNil(o.ProteinId) {
    return nil, false
	}
	return o.ProteinId, true
}

// HasProteinId returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasProteinId() bool {
	if o != nil && !isNil(o.ProteinId) {
		return true
	}

	return false
}

// SetProteinId gets a reference to the given string and assigns it to the ProteinId field.
func (o *TranscriptConsequence) SetProteinId(v string) {
	o.ProteinId = &v
}

// GetProteinStart returns the ProteinStart field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetProteinStart() int32 {
	if o == nil || isNil(o.ProteinStart) {
		var ret int32
		return ret
	}
	return *o.ProteinStart
}

// GetProteinStartOk returns a tuple with the ProteinStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetProteinStartOk() (*int32, bool) {
	if o == nil || isNil(o.ProteinStart) {
    return nil, false
	}
	return o.ProteinStart, true
}

// HasProteinStart returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasProteinStart() bool {
	if o != nil && !isNil(o.ProteinStart) {
		return true
	}

	return false
}

// SetProteinStart gets a reference to the given int32 and assigns it to the ProteinStart field.
func (o *TranscriptConsequence) SetProteinStart(v int32) {
	o.ProteinStart = &v
}

// GetRefseqTranscriptIds returns the RefseqTranscriptIds field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetRefseqTranscriptIds() []string {
	if o == nil || isNil(o.RefseqTranscriptIds) {
		var ret []string
		return ret
	}
	return o.RefseqTranscriptIds
}

// GetRefseqTranscriptIdsOk returns a tuple with the RefseqTranscriptIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetRefseqTranscriptIdsOk() ([]string, bool) {
	if o == nil || isNil(o.RefseqTranscriptIds) {
    return nil, false
	}
	return o.RefseqTranscriptIds, true
}

// HasRefseqTranscriptIds returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasRefseqTranscriptIds() bool {
	if o != nil && !isNil(o.RefseqTranscriptIds) {
		return true
	}

	return false
}

// SetRefseqTranscriptIds gets a reference to the given []string and assigns it to the RefseqTranscriptIds field.
func (o *TranscriptConsequence) SetRefseqTranscriptIds(v []string) {
	o.RefseqTranscriptIds = v
}

// GetSiftPrediction returns the SiftPrediction field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetSiftPrediction() string {
	if o == nil || isNil(o.SiftPrediction) {
		var ret string
		return ret
	}
	return *o.SiftPrediction
}

// GetSiftPredictionOk returns a tuple with the SiftPrediction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetSiftPredictionOk() (*string, bool) {
	if o == nil || isNil(o.SiftPrediction) {
    return nil, false
	}
	return o.SiftPrediction, true
}

// HasSiftPrediction returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasSiftPrediction() bool {
	if o != nil && !isNil(o.SiftPrediction) {
		return true
	}

	return false
}

// SetSiftPrediction gets a reference to the given string and assigns it to the SiftPrediction field.
func (o *TranscriptConsequence) SetSiftPrediction(v string) {
	o.SiftPrediction = &v
}

// GetSiftScore returns the SiftScore field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetSiftScore() float64 {
	if o == nil || isNil(o.SiftScore) {
		var ret float64
		return ret
	}
	return *o.SiftScore
}

// GetSiftScoreOk returns a tuple with the SiftScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetSiftScoreOk() (*float64, bool) {
	if o == nil || isNil(o.SiftScore) {
    return nil, false
	}
	return o.SiftScore, true
}

// HasSiftScore returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasSiftScore() bool {
	if o != nil && !isNil(o.SiftScore) {
		return true
	}

	return false
}

// SetSiftScore gets a reference to the given float64 and assigns it to the SiftScore field.
func (o *TranscriptConsequence) SetSiftScore(v float64) {
	o.SiftScore = &v
}

// GetTranscriptId returns the TranscriptId field value
func (o *TranscriptConsequence) GetTranscriptId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TranscriptId
}

// GetTranscriptIdOk returns a tuple with the TranscriptId field value
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetTranscriptIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TranscriptId, true
}

// SetTranscriptId sets field value
func (o *TranscriptConsequence) SetTranscriptId(v string) {
	o.TranscriptId = v
}

// GetUniprotId returns the UniprotId field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetUniprotId() string {
	if o == nil || isNil(o.UniprotId) {
		var ret string
		return ret
	}
	return *o.UniprotId
}

// GetUniprotIdOk returns a tuple with the UniprotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetUniprotIdOk() (*string, bool) {
	if o == nil || isNil(o.UniprotId) {
    return nil, false
	}
	return o.UniprotId, true
}

// HasUniprotId returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasUniprotId() bool {
	if o != nil && !isNil(o.UniprotId) {
		return true
	}

	return false
}

// SetUniprotId gets a reference to the given string and assigns it to the UniprotId field.
func (o *TranscriptConsequence) SetUniprotId(v string) {
	o.UniprotId = &v
}

// GetVariantAllele returns the VariantAllele field value if set, zero value otherwise.
func (o *TranscriptConsequence) GetVariantAllele() string {
	if o == nil || isNil(o.VariantAllele) {
		var ret string
		return ret
	}
	return *o.VariantAllele
}

// GetVariantAlleleOk returns a tuple with the VariantAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptConsequence) GetVariantAlleleOk() (*string, bool) {
	if o == nil || isNil(o.VariantAllele) {
    return nil, false
	}
	return o.VariantAllele, true
}

// HasVariantAllele returns a boolean if a field has been set.
func (o *TranscriptConsequence) HasVariantAllele() bool {
	if o != nil && !isNil(o.VariantAllele) {
		return true
	}

	return false
}

// SetVariantAllele gets a reference to the given string and assigns it to the VariantAllele field.
func (o *TranscriptConsequence) SetVariantAllele(v string) {
	o.VariantAllele = &v
}

func (o TranscriptConsequence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AminoAcids) {
		toSerialize["amino_acids"] = o.AminoAcids
	}
	if !isNil(o.Canonical) {
		toSerialize["canonical"] = o.Canonical
	}
	if !isNil(o.Codons) {
		toSerialize["codons"] = o.Codons
	}
	if !isNil(o.ConsequenceTerms) {
		toSerialize["consequence_terms"] = o.ConsequenceTerms
	}
	if !isNil(o.Exon) {
		toSerialize["exon"] = o.Exon
	}
	if !isNil(o.GeneId) {
		toSerialize["gene_id"] = o.GeneId
	}
	if !isNil(o.GeneSymbol) {
		toSerialize["gene_symbol"] = o.GeneSymbol
	}
	if !isNil(o.HgncId) {
		toSerialize["hgnc_id"] = o.HgncId
	}
	if !isNil(o.Hgvsc) {
		toSerialize["hgvsc"] = o.Hgvsc
	}
	if !isNil(o.Hgvsg) {
		toSerialize["hgvsg"] = o.Hgvsg
	}
	if !isNil(o.Hgvsp) {
		toSerialize["hgvsp"] = o.Hgvsp
	}
	if !isNil(o.PolyphenPrediction) {
		toSerialize["polyphen_prediction"] = o.PolyphenPrediction
	}
	if !isNil(o.PolyphenScore) {
		toSerialize["polyphen_score"] = o.PolyphenScore
	}
	if !isNil(o.ProteinEnd) {
		toSerialize["protein_end"] = o.ProteinEnd
	}
	if !isNil(o.ProteinId) {
		toSerialize["protein_id"] = o.ProteinId
	}
	if !isNil(o.ProteinStart) {
		toSerialize["protein_start"] = o.ProteinStart
	}
	if !isNil(o.RefseqTranscriptIds) {
		toSerialize["refseq_transcript_ids"] = o.RefseqTranscriptIds
	}
	if !isNil(o.SiftPrediction) {
		toSerialize["sift_prediction"] = o.SiftPrediction
	}
	if !isNil(o.SiftScore) {
		toSerialize["sift_score"] = o.SiftScore
	}
	if true {
		toSerialize["transcript_id"] = o.TranscriptId
	}
	if !isNil(o.UniprotId) {
		toSerialize["uniprotId"] = o.UniprotId
	}
	if !isNil(o.VariantAllele) {
		toSerialize["variant_allele"] = o.VariantAllele
	}
	return json.Marshal(toSerialize)
}

type NullableTranscriptConsequence struct {
	value *TranscriptConsequence
	isSet bool
}

func (v NullableTranscriptConsequence) Get() *TranscriptConsequence {
	return v.value
}

func (v *NullableTranscriptConsequence) Set(val *TranscriptConsequence) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscriptConsequence) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscriptConsequence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscriptConsequence(val *TranscriptConsequence) *NullableTranscriptConsequence {
	return &NullableTranscriptConsequence{value: val, isSet: true}
}

func (v NullableTranscriptConsequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscriptConsequence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

