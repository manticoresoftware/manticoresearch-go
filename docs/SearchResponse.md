# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Took** | Pointer to **interface{}** | Time taken to execute the search | [optional] 
**TimedOut** | Pointer to **interface{}** | Indicates whether the search operation timed out | [optional] 
**Aggregations** | Pointer to **interface{}** | Aggregated search results grouped by the specified criteria | [optional] 
**Hits** | Pointer to [**SearchResponseHits**](SearchResponseHits.md) |  | [optional] 
**Profile** | Pointer to **interface{}** | Profile information about the search execution, if profiling is enabled | [optional] 
**Scroll** | Pointer to **interface{}** | Scroll token to be used fo pagination | [optional] 
**Warning** | Pointer to **interface{}** | Warnings encountered during the search operation | [optional] 

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

`func (o *SearchResponse) GetTook() interface{}`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SearchResponse) GetTookOk() (*interface{}, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SearchResponse) SetTook(v interface{})`

SetTook sets Took field to given value.

### HasTook

`func (o *SearchResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### SetTookNil

`func (o *SearchResponse) SetTookNil(b bool)`

 SetTookNil sets the value for Took to be an explicit nil

### UnsetTook
`func (o *SearchResponse) UnsetTook()`

UnsetTook ensures that no value is present for Took, not even an explicit nil
### GetTimedOut

`func (o *SearchResponse) GetTimedOut() interface{}`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SearchResponse) GetTimedOutOk() (*interface{}, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SearchResponse) SetTimedOut(v interface{})`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SearchResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### SetTimedOutNil

`func (o *SearchResponse) SetTimedOutNil(b bool)`

 SetTimedOutNil sets the value for TimedOut to be an explicit nil

### UnsetTimedOut
`func (o *SearchResponse) UnsetTimedOut()`

UnsetTimedOut ensures that no value is present for TimedOut, not even an explicit nil
### GetAggregations

`func (o *SearchResponse) GetAggregations() interface{}`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SearchResponse) GetAggregationsOk() (*interface{}, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SearchResponse) SetAggregations(v interface{})`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *SearchResponse) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### SetAggregationsNil

`func (o *SearchResponse) SetAggregationsNil(b bool)`

 SetAggregationsNil sets the value for Aggregations to be an explicit nil

### UnsetAggregations
`func (o *SearchResponse) UnsetAggregations()`

UnsetAggregations ensures that no value is present for Aggregations, not even an explicit nil
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

`func (o *SearchResponse) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SearchResponse) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SearchResponse) SetProfile(v interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SearchResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *SearchResponse) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *SearchResponse) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetScroll

`func (o *SearchResponse) GetScroll() interface{}`

GetScroll returns the Scroll field if non-nil, zero value otherwise.

### GetScrollOk

`func (o *SearchResponse) GetScrollOk() (*interface{}, bool)`

GetScrollOk returns a tuple with the Scroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScroll

`func (o *SearchResponse) SetScroll(v interface{})`

SetScroll sets Scroll field to given value.

### HasScroll

`func (o *SearchResponse) HasScroll() bool`

HasScroll returns a boolean if a field has been set.

### SetScrollNil

`func (o *SearchResponse) SetScrollNil(b bool)`

 SetScrollNil sets the value for Scroll to be an explicit nil

### UnsetScroll
`func (o *SearchResponse) UnsetScroll()`

UnsetScroll ensures that no value is present for Scroll, not even an explicit nil
### GetWarning

`func (o *SearchResponse) GetWarning() interface{}`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SearchResponse) GetWarningOk() (*interface{}, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SearchResponse) SetWarning(v interface{})`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *SearchResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *SearchResponse) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *SearchResponse) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


