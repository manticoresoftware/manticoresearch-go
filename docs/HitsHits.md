# HitsHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** | The ID of the matched document | [optional] 
**Score** | Pointer to **interface{}** | The score of the matched document | [optional] 
**Source** | Pointer to **interface{}** | The source data of the matched document | [optional] 
**KnnDist** | Pointer to **interface{}** | The knn distance of the matched document returned for knn queries | [optional] 
**Highlight** | Pointer to **interface{}** | The highlighting-related data of the matched document | [optional] 
**Table** | Pointer to **interface{}** | The table name of the matched document returned for percolate queries | [optional] 
**Type** | Pointer to **interface{}** | The type of the matched document returned for percolate queries | [optional] 
**Fields** | Pointer to **interface{}** | The percolate-related fields of the matched document returned for percolate queries | [optional] 

## Methods

### NewHitsHits

`func NewHitsHits() *HitsHits`

NewHitsHits instantiates a new HitsHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHitsHitsWithDefaults

`func NewHitsHitsWithDefaults() *HitsHits`

NewHitsHitsWithDefaults instantiates a new HitsHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HitsHits) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HitsHits) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HitsHits) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *HitsHits) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HitsHits) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HitsHits) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScore

`func (o *HitsHits) GetScore() interface{}`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *HitsHits) GetScoreOk() (*interface{}, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *HitsHits) SetScore(v interface{})`

SetScore sets Score field to given value.

### HasScore

`func (o *HitsHits) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *HitsHits) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *HitsHits) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetSource

`func (o *HitsHits) GetSource() interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HitsHits) GetSourceOk() (*interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HitsHits) SetSource(v interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *HitsHits) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *HitsHits) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *HitsHits) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetKnnDist

`func (o *HitsHits) GetKnnDist() interface{}`

GetKnnDist returns the KnnDist field if non-nil, zero value otherwise.

### GetKnnDistOk

`func (o *HitsHits) GetKnnDistOk() (*interface{}, bool)`

GetKnnDistOk returns a tuple with the KnnDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnnDist

`func (o *HitsHits) SetKnnDist(v interface{})`

SetKnnDist sets KnnDist field to given value.

### HasKnnDist

`func (o *HitsHits) HasKnnDist() bool`

HasKnnDist returns a boolean if a field has been set.

### SetKnnDistNil

`func (o *HitsHits) SetKnnDistNil(b bool)`

 SetKnnDistNil sets the value for KnnDist to be an explicit nil

### UnsetKnnDist
`func (o *HitsHits) UnsetKnnDist()`

UnsetKnnDist ensures that no value is present for KnnDist, not even an explicit nil
### GetHighlight

`func (o *HitsHits) GetHighlight() interface{}`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *HitsHits) GetHighlightOk() (*interface{}, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *HitsHits) SetHighlight(v interface{})`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *HitsHits) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### SetHighlightNil

`func (o *HitsHits) SetHighlightNil(b bool)`

 SetHighlightNil sets the value for Highlight to be an explicit nil

### UnsetHighlight
`func (o *HitsHits) UnsetHighlight()`

UnsetHighlight ensures that no value is present for Highlight, not even an explicit nil
### GetTable

`func (o *HitsHits) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *HitsHits) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *HitsHits) SetTable(v interface{})`

SetTable sets Table field to given value.

### HasTable

`func (o *HitsHits) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTableNil

`func (o *HitsHits) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *HitsHits) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetType

`func (o *HitsHits) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HitsHits) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HitsHits) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *HitsHits) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HitsHits) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HitsHits) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFields

`func (o *HitsHits) GetFields() interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *HitsHits) GetFieldsOk() (*interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *HitsHits) SetFields(v interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *HitsHits) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *HitsHits) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *HitsHits) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


