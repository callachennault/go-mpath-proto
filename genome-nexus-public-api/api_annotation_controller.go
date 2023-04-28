/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_public_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// AnnotationControllerApiService AnnotationControllerApi service
type AnnotationControllerApiService service

type ApiFetchVariantAnnotationByGenomicLocationGETRequest struct {
	ctx context.Context
	ApiService *AnnotationControllerApiService
	genomicLocation string
	isoformOverrideSource *string
	token *string
	fields *[]string
}

// Isoform override source. For example uniprot
func (r ApiFetchVariantAnnotationByGenomicLocationGETRequest) IsoformOverrideSource(isoformOverrideSource string) ApiFetchVariantAnnotationByGenomicLocationGETRequest {
	r.isoformOverrideSource = &isoformOverrideSource
	return r
}

// Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;}
func (r ApiFetchVariantAnnotationByGenomicLocationGETRequest) Token(token string) ApiFetchVariantAnnotationByGenomicLocationGETRequest {
	r.token = &token
	return r
}

// Comma separated list of fields to include (case-sensitive!). For example: hotspots
func (r ApiFetchVariantAnnotationByGenomicLocationGETRequest) Fields(fields []string) ApiFetchVariantAnnotationByGenomicLocationGETRequest {
	r.fields = &fields
	return r
}

func (r ApiFetchVariantAnnotationByGenomicLocationGETRequest) Execute() (*VariantAnnotation, *http.Response, error) {
	return r.ApiService.FetchVariantAnnotationByGenomicLocationGETExecute(r)
}

/*
FetchVariantAnnotationByGenomicLocationGET Retrieves VEP annotation for the provided genomic location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param genomicLocation A genomic location. For example 7,140453136,140453136,A,T
 @return ApiFetchVariantAnnotationByGenomicLocationGETRequest
*/
func (a *AnnotationControllerApiService) FetchVariantAnnotationByGenomicLocationGET(ctx context.Context, genomicLocation string) ApiFetchVariantAnnotationByGenomicLocationGETRequest {
	return ApiFetchVariantAnnotationByGenomicLocationGETRequest{
		ApiService: a,
		ctx: ctx,
		genomicLocation: genomicLocation,
	}
}

// Execute executes the request
//  @return VariantAnnotation
func (a *AnnotationControllerApiService) FetchVariantAnnotationByGenomicLocationGETExecute(r ApiFetchVariantAnnotationByGenomicLocationGETRequest) (*VariantAnnotation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariantAnnotation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnotationControllerApiService.FetchVariantAnnotationByGenomicLocationGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/annotation/genomic/{genomicLocation}"
	localVarPath = strings.Replace(localVarPath, "{"+"genomicLocation"+"}", url.PathEscape(parameterToString(r.genomicLocation, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isoformOverrideSource != nil {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(*r.isoformOverrideSource, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchVariantAnnotationByGenomicLocationPOSTRequest struct {
	ctx context.Context
	ApiService *AnnotationControllerApiService
	genomicLocations *[]GenomicLocation
	isoformOverrideSource *string
	token *string
	fields *[]string
}

// List of Genomic Locations
func (r ApiFetchVariantAnnotationByGenomicLocationPOSTRequest) GenomicLocations(genomicLocations []GenomicLocation) ApiFetchVariantAnnotationByGenomicLocationPOSTRequest {
	r.genomicLocations = &genomicLocations
	return r
}

// Isoform override source. For example uniprot
func (r ApiFetchVariantAnnotationByGenomicLocationPOSTRequest) IsoformOverrideSource(isoformOverrideSource string) ApiFetchVariantAnnotationByGenomicLocationPOSTRequest {
	r.isoformOverrideSource = &isoformOverrideSource
	return r
}

// Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;}
func (r ApiFetchVariantAnnotationByGenomicLocationPOSTRequest) Token(token string) ApiFetchVariantAnnotationByGenomicLocationPOSTRequest {
	r.token = &token
	return r
}

// Comma separated list of fields to include (case-sensitive!). For example: hotspots
func (r ApiFetchVariantAnnotationByGenomicLocationPOSTRequest) Fields(fields []string) ApiFetchVariantAnnotationByGenomicLocationPOSTRequest {
	r.fields = &fields
	return r
}

func (r ApiFetchVariantAnnotationByGenomicLocationPOSTRequest) Execute() ([]VariantAnnotation, *http.Response, error) {
	return r.ApiService.FetchVariantAnnotationByGenomicLocationPOSTExecute(r)
}

/*
FetchVariantAnnotationByGenomicLocationPOST Retrieves VEP annotation for the provided list of genomic locations

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchVariantAnnotationByGenomicLocationPOSTRequest
*/
func (a *AnnotationControllerApiService) FetchVariantAnnotationByGenomicLocationPOST(ctx context.Context) ApiFetchVariantAnnotationByGenomicLocationPOSTRequest {
	return ApiFetchVariantAnnotationByGenomicLocationPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []VariantAnnotation
func (a *AnnotationControllerApiService) FetchVariantAnnotationByGenomicLocationPOSTExecute(r ApiFetchVariantAnnotationByGenomicLocationPOSTRequest) ([]VariantAnnotation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []VariantAnnotation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnotationControllerApiService.FetchVariantAnnotationByGenomicLocationPOST")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/annotation/genomic"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.genomicLocations == nil {
		return localVarReturnValue, nil, reportError("genomicLocations is required and must be specified")
	}

	if r.isoformOverrideSource != nil {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(*r.isoformOverrideSource, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.genomicLocations
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchVariantAnnotationByIdGETRequest struct {
	ctx context.Context
	ApiService *AnnotationControllerApiService
	variantId string
	isoformOverrideSource *string
	token *string
	fields *[]string
}

// Isoform override source. For example uniprot
func (r ApiFetchVariantAnnotationByIdGETRequest) IsoformOverrideSource(isoformOverrideSource string) ApiFetchVariantAnnotationByIdGETRequest {
	r.isoformOverrideSource = &isoformOverrideSource
	return r
}

// Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;}
func (r ApiFetchVariantAnnotationByIdGETRequest) Token(token string) ApiFetchVariantAnnotationByIdGETRequest {
	r.token = &token
	return r
}

// Comma separated list of fields to include (case-sensitive!). For example: annotation_summary
func (r ApiFetchVariantAnnotationByIdGETRequest) Fields(fields []string) ApiFetchVariantAnnotationByIdGETRequest {
	r.fields = &fields
	return r
}

func (r ApiFetchVariantAnnotationByIdGETRequest) Execute() (*VariantAnnotation, *http.Response, error) {
	return r.ApiService.FetchVariantAnnotationByIdGETExecute(r)
}

/*
FetchVariantAnnotationByIdGET Retrieves VEP annotation for the give dbSNP id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param variantId dbSNP id. For example rs116035550.
 @return ApiFetchVariantAnnotationByIdGETRequest
*/
func (a *AnnotationControllerApiService) FetchVariantAnnotationByIdGET(ctx context.Context, variantId string) ApiFetchVariantAnnotationByIdGETRequest {
	return ApiFetchVariantAnnotationByIdGETRequest{
		ApiService: a,
		ctx: ctx,
		variantId: variantId,
	}
}

// Execute executes the request
//  @return VariantAnnotation
func (a *AnnotationControllerApiService) FetchVariantAnnotationByIdGETExecute(r ApiFetchVariantAnnotationByIdGETRequest) (*VariantAnnotation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariantAnnotation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnotationControllerApiService.FetchVariantAnnotationByIdGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/annotation/dbsnp/{variantId}"
	localVarPath = strings.Replace(localVarPath, "{"+"variantId"+"}", url.PathEscape(parameterToString(r.variantId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isoformOverrideSource != nil {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(*r.isoformOverrideSource, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchVariantAnnotationByIdPOSTRequest struct {
	ctx context.Context
	ApiService *AnnotationControllerApiService
	variantIds *[]string
	isoformOverrideSource *string
	token *string
	fields *[]string
}

// List of variant IDs. For example [\&quot;rs116035550\&quot;]
func (r ApiFetchVariantAnnotationByIdPOSTRequest) VariantIds(variantIds []string) ApiFetchVariantAnnotationByIdPOSTRequest {
	r.variantIds = &variantIds
	return r
}

// Isoform override source. For example uniprot
func (r ApiFetchVariantAnnotationByIdPOSTRequest) IsoformOverrideSource(isoformOverrideSource string) ApiFetchVariantAnnotationByIdPOSTRequest {
	r.isoformOverrideSource = &isoformOverrideSource
	return r
}

// Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;}
func (r ApiFetchVariantAnnotationByIdPOSTRequest) Token(token string) ApiFetchVariantAnnotationByIdPOSTRequest {
	r.token = &token
	return r
}

// Comma separated list of fields to include (case-sensitive!). For example: annotation_summary
func (r ApiFetchVariantAnnotationByIdPOSTRequest) Fields(fields []string) ApiFetchVariantAnnotationByIdPOSTRequest {
	r.fields = &fields
	return r
}

func (r ApiFetchVariantAnnotationByIdPOSTRequest) Execute() ([]VariantAnnotation, *http.Response, error) {
	return r.ApiService.FetchVariantAnnotationByIdPOSTExecute(r)
}

/*
FetchVariantAnnotationByIdPOST Retrieves VEP annotation for the provided list of dbSNP ids

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchVariantAnnotationByIdPOSTRequest
*/
func (a *AnnotationControllerApiService) FetchVariantAnnotationByIdPOST(ctx context.Context) ApiFetchVariantAnnotationByIdPOSTRequest {
	return ApiFetchVariantAnnotationByIdPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []VariantAnnotation
func (a *AnnotationControllerApiService) FetchVariantAnnotationByIdPOSTExecute(r ApiFetchVariantAnnotationByIdPOSTRequest) ([]VariantAnnotation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []VariantAnnotation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnotationControllerApiService.FetchVariantAnnotationByIdPOST")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/annotation/dbsnp/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.variantIds == nil {
		return localVarReturnValue, nil, reportError("variantIds is required and must be specified")
	}

	if r.isoformOverrideSource != nil {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(*r.isoformOverrideSource, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.variantIds
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchVariantAnnotationGETRequest struct {
	ctx context.Context
	ApiService *AnnotationControllerApiService
	variant string
	isoformOverrideSource *string
	token *string
	fields *[]string
}

// Isoform override source. For example uniprot
func (r ApiFetchVariantAnnotationGETRequest) IsoformOverrideSource(isoformOverrideSource string) ApiFetchVariantAnnotationGETRequest {
	r.isoformOverrideSource = &isoformOverrideSource
	return r
}

// Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;}
func (r ApiFetchVariantAnnotationGETRequest) Token(token string) ApiFetchVariantAnnotationGETRequest {
	r.token = &token
	return r
}

// Comma separated list of fields to include (case-sensitive!). For example: hotspots
func (r ApiFetchVariantAnnotationGETRequest) Fields(fields []string) ApiFetchVariantAnnotationGETRequest {
	r.fields = &fields
	return r
}

func (r ApiFetchVariantAnnotationGETRequest) Execute() (*VariantAnnotation, *http.Response, error) {
	return r.ApiService.FetchVariantAnnotationGETExecute(r)
}

/*
FetchVariantAnnotationGET Retrieves VEP annotation for the provided variant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param variant Variant. For example 17:g.41242962_41242963insGA
 @return ApiFetchVariantAnnotationGETRequest
*/
func (a *AnnotationControllerApiService) FetchVariantAnnotationGET(ctx context.Context, variant string) ApiFetchVariantAnnotationGETRequest {
	return ApiFetchVariantAnnotationGETRequest{
		ApiService: a,
		ctx: ctx,
		variant: variant,
	}
}

// Execute executes the request
//  @return VariantAnnotation
func (a *AnnotationControllerApiService) FetchVariantAnnotationGETExecute(r ApiFetchVariantAnnotationGETRequest) (*VariantAnnotation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VariantAnnotation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnotationControllerApiService.FetchVariantAnnotationGET")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/annotation/{variant}"
	localVarPath = strings.Replace(localVarPath, "{"+"variant"+"}", url.PathEscape(parameterToString(r.variant, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isoformOverrideSource != nil {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(*r.isoformOverrideSource, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFetchVariantAnnotationPOSTRequest struct {
	ctx context.Context
	ApiService *AnnotationControllerApiService
	variants *[]string
	isoformOverrideSource *string
	token *string
	fields *[]string
}

// List of variants. For example [\&quot;X:g.66937331T&gt;A\&quot;,\&quot;17:g.41242962_41242963insGA\&quot;] (GRCh37) or [\&quot;1:g.182712A&gt;C\&quot;, \&quot;2:g.265023C&gt;T\&quot;, \&quot;3:g.319781del\&quot;, \&quot;19:g.110753dup\&quot;, \&quot;1:g.1385015_1387562del\&quot;] (GRCh38)
func (r ApiFetchVariantAnnotationPOSTRequest) Variants(variants []string) ApiFetchVariantAnnotationPOSTRequest {
	r.variants = &variants
	return r
}

// Isoform override source. For example uniprot
func (r ApiFetchVariantAnnotationPOSTRequest) IsoformOverrideSource(isoformOverrideSource string) ApiFetchVariantAnnotationPOSTRequest {
	r.isoformOverrideSource = &isoformOverrideSource
	return r
}

// Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;}
func (r ApiFetchVariantAnnotationPOSTRequest) Token(token string) ApiFetchVariantAnnotationPOSTRequest {
	r.token = &token
	return r
}

// Comma separated list of fields to include (case-sensitive!). For example: hotspots
func (r ApiFetchVariantAnnotationPOSTRequest) Fields(fields []string) ApiFetchVariantAnnotationPOSTRequest {
	r.fields = &fields
	return r
}

func (r ApiFetchVariantAnnotationPOSTRequest) Execute() ([]VariantAnnotation, *http.Response, error) {
	return r.ApiService.FetchVariantAnnotationPOSTExecute(r)
}

/*
FetchVariantAnnotationPOST Retrieves VEP annotation for the provided list of variants

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFetchVariantAnnotationPOSTRequest
*/
func (a *AnnotationControllerApiService) FetchVariantAnnotationPOST(ctx context.Context) ApiFetchVariantAnnotationPOSTRequest {
	return ApiFetchVariantAnnotationPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []VariantAnnotation
func (a *AnnotationControllerApiService) FetchVariantAnnotationPOSTExecute(r ApiFetchVariantAnnotationPOSTRequest) ([]VariantAnnotation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []VariantAnnotation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnnotationControllerApiService.FetchVariantAnnotationPOST")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/annotation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.variants == nil {
		return localVarReturnValue, nil, reportError("variants is required and must be specified")
	}

	if r.isoformOverrideSource != nil {
		localVarQueryParams.Add("isoformOverrideSource", parameterToString(*r.isoformOverrideSource, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.variants
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
