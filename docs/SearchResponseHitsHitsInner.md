# SearchResponseHitsHitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the matched document | [optional] 
**Score** | Pointer to **int32** | The score of the matched document | [optional] 
**Source** | Pointer to **map[string]interface{}** | The source data of the matched document | [optional] 

## Methods

### NewSearchResponseHitsHitsInner

`func NewSearchResponseHitsHitsInner() *SearchResponseHitsHitsInner`

NewSearchResponseHitsHitsInner instantiates a new SearchResponseHitsHitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseHitsHitsInnerWithDefaults

`func NewSearchResponseHitsHitsInnerWithDefaults() *SearchResponseHitsHitsInner`

NewSearchResponseHitsHitsInnerWithDefaults instantiates a new SearchResponseHitsHitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchResponseHitsHitsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchResponseHitsHitsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchResponseHitsHitsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SearchResponseHitsHitsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScore

`func (o *SearchResponseHitsHitsInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchResponseHitsHitsInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchResponseHitsHitsInner) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchResponseHitsHitsInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *SearchResponseHitsHitsInner) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SearchResponseHitsHitsInner) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SearchResponseHitsHitsInner) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *SearchResponseHitsHitsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


