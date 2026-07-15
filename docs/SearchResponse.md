# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Took** | Pointer to **int32** | Time taken to execute the search | [optional] 
**TimedOut** | Pointer to **bool** | Indicates whether the search operation timed out | [optional] 
**Aggregations** | Pointer to [**map[string]AggBucketsResult**](AggBucketsResult.md) | Aggregated search results grouped by the specified criteria. Each named aggregation typically contains a &#x60;buckets&#x60; array (or keyed map) of bucket objects with &#x60;key&#x60;, &#x60;doc_count&#x60;, and optional &#x60;status&#x60;.  | [optional] 
**Hits** | Pointer to [**SearchResponseHits**](SearchResponseHits.md) |  | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Profile information about the search execution, if profiling is enabled | [optional] 
**Scroll** | Pointer to **string** | Scroll token to be used fo pagination | [optional] 
**Warning** | Pointer to **map[string]interface{}** | Warnings encountered during the search operation | [optional] 
**ConversationUuid** | Pointer to **string** | Existing or generated conversation id (conversational search) | [optional] 
**UserQuery** | Pointer to **string** | Original user query (conversational search) | [optional] 
**SearchQuery** | Pointer to **string** | Standalone search query used for KNN retrieval (conversational search) | [optional] 
**Response** | Pointer to **string** | LLM answer as generated (conversational search) | [optional] 
**Sources** | Pointer to **string** | JSON string containing retrieved source rows used as LLM context (conversational search).  | [optional] 

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

`func (o *SearchResponse) GetAggregations() map[string]AggBucketsResult`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *SearchResponse) GetAggregationsOk() (*map[string]AggBucketsResult, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *SearchResponse) SetAggregations(v map[string]AggBucketsResult)`

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

### GetScroll

`func (o *SearchResponse) GetScroll() string`

GetScroll returns the Scroll field if non-nil, zero value otherwise.

### GetScrollOk

`func (o *SearchResponse) GetScrollOk() (*string, bool)`

GetScrollOk returns a tuple with the Scroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScroll

`func (o *SearchResponse) SetScroll(v string)`

SetScroll sets Scroll field to given value.

### HasScroll

`func (o *SearchResponse) HasScroll() bool`

HasScroll returns a boolean if a field has been set.

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

### GetConversationUuid

`func (o *SearchResponse) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *SearchResponse) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *SearchResponse) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.

### HasConversationUuid

`func (o *SearchResponse) HasConversationUuid() bool`

HasConversationUuid returns a boolean if a field has been set.

### GetUserQuery

`func (o *SearchResponse) GetUserQuery() string`

GetUserQuery returns the UserQuery field if non-nil, zero value otherwise.

### GetUserQueryOk

`func (o *SearchResponse) GetUserQueryOk() (*string, bool)`

GetUserQueryOk returns a tuple with the UserQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuery

`func (o *SearchResponse) SetUserQuery(v string)`

SetUserQuery sets UserQuery field to given value.

### HasUserQuery

`func (o *SearchResponse) HasUserQuery() bool`

HasUserQuery returns a boolean if a field has been set.

### GetSearchQuery

`func (o *SearchResponse) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *SearchResponse) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *SearchResponse) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *SearchResponse) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetResponse

`func (o *SearchResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SearchResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SearchResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SearchResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSources

`func (o *SearchResponse) GetSources() string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SearchResponse) GetSourcesOk() (*string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SearchResponse) SetSources(v string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SearchResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


