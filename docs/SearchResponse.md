# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Took** | Pointer to **int32** | Time taken to execute the search | [optional] 
**TimedOut** | Pointer to **bool** | Indicates whether the search operation timed out | [optional] 
**Aggregations** | Pointer to **map[string]interface{}** | Aggregated search results grouped by the specified criteria | [optional] 
**Hits** | Pointer to [**SearchResponseHits**](SearchResponseHits.md) |  | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Profile information about the search execution, if profiling is enabled | [optional] 
**Warning** | Pointer to **map[string]interface{}** | Warnings encountered during the search operation | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTook

`func (o *SearchResponse) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchResponse) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchResponse) SetTook(v int32)`

SetTook sets Took field to given value.

### HasTook

`func (o *SearchResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetTimedOut

`func (o *SearchResponse) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SearchResponse) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SearchResponse) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SearchResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetAggregations

`func (o *SearchResponse) GetAggregations() map[string]interface{}`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SearchResponse) GetAggregationsOk() (*map[string]interface{}, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SearchResponse) SetAggregations(v map[string]interface{})`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *SearchResponse) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetHits

`func (o *SearchResponse) GetHits() SearchResponseHits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchResponse) GetHitsOk() (*SearchResponseHits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchResponse) SetHits(v SearchResponseHits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetProfile

`func (o *SearchResponse) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SearchResponse) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SearchResponse) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SearchResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetWarning

`func (o *SearchResponse) GetWarning() map[string]interface{}`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SearchResponse) GetWarningOk() (*map[string]interface{}, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SearchResponse) SetWarning(v map[string]interface{})`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SearchResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


