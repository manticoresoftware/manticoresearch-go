# SearchResponseHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxScore** | Pointer to **interface{}** | Maximum score among the matched documents | [optional] 
**Total** | Pointer to **interface{}** | Total number of matched documents | [optional] 
**TotalRelation** | Pointer to **interface{}** | Indicates whether the total number of hits is accurate or an estimate | [optional] 
**Hits** | Pointer to **interface{}** | Array of hit objects, each representing a matched document | [optional] 

## Methods

### NewSearchResponseHits

`func NewSearchResponseHits() *SearchResponseHits`

NewSearchResponseHits instantiates a new SearchResponseHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseHitsWithDefaults

`func NewSearchResponseHitsWithDefaults() *SearchResponseHits`

NewSearchResponseHitsWithDefaults instantiates a new SearchResponseHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxScore

`func (o *SearchResponseHits) GetMaxScore() interface{}`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SearchResponseHits) GetMaxScoreOk() (*interface{}, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *SearchResponseHits) SetMaxScore(v interface{})`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *SearchResponseHits) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### SetMaxScoreNil

`func (o *SearchResponseHits) SetMaxScoreNil(b bool)`

 SetMaxScoreNil sets the value for MaxScore to be an explicit nil

### UnsetMaxScore
`func (o *SearchResponseHits) UnsetMaxScore()`

UnsetMaxScore ensures that no value is present for MaxScore, not even an explicit nil
### GetTotal

`func (o *SearchResponseHits) GetTotal() interface{}`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchResponseHits) GetTotalOk() (*interface{}, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchResponseHits) SetTotal(v interface{})`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchResponseHits) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *SearchResponseHits) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *SearchResponseHits) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetTotalRelation

`func (o *SearchResponseHits) GetTotalRelation() interface{}`

GetTotalRelation returns the TotalRelation field if non-nil, zero value otherwise.

### GetTotalRelationOk

`func (o *SearchResponseHits) GetTotalRelationOk() (*interface{}, bool)`

GetTotalRelationOk returns a tuple with the TotalRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRelation

`func (o *SearchResponseHits) SetTotalRelation(v interface{})`

SetTotalRelation sets TotalRelation field to given value.

### HasTotalRelation

`func (o *SearchResponseHits) HasTotalRelation() bool`

HasTotalRelation returns a boolean if a field has been set.

### SetTotalRelationNil

`func (o *SearchResponseHits) SetTotalRelationNil(b bool)`

 SetTotalRelationNil sets the value for TotalRelation to be an explicit nil

### UnsetTotalRelation
`func (o *SearchResponseHits) UnsetTotalRelation()`

UnsetTotalRelation ensures that no value is present for TotalRelation, not even an explicit nil
### GetHits

`func (o *SearchResponseHits) GetHits() interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchResponseHits) GetHitsOk() (*interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchResponseHits) SetHits(v interface{})`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchResponseHits) HasHits() bool`

HasHits returns a boolean if a field has been set.

### SetHitsNil

`func (o *SearchResponseHits) SetHitsNil(b bool)`

 SetHitsNil sets the value for Hits to be an explicit nil

### UnsetHits
`func (o *SearchResponseHits) UnsetHits()`

UnsetHits ensures that no value is present for Hits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


