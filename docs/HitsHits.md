# HitsHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the matched document | [optional] 
**Score** | Pointer to **int32** | The score of the matched document | [optional] 
**Source** | Pointer to **map[string]interface{}** | The source data of the matched document | [optional] 
**KnnDist** | Pointer to **float32** | The knn distance of the matched document returned for knn queries | [optional] 
**Highlight** | Pointer to **map[string]interface{}** | The highlighting-related data of the matched document | [optional] 
**Table** | Pointer to **string** | The table name of the matched document returned for percolate queries | [optional] 
**Type** | Pointer to **string** | The type of the matched document returned for percolate queries | [optional] 
**Fields** | Pointer to **map[string]interface{}** | The percolate-related fields of the matched document returned for percolate queries | [optional] 

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

`func (o *HitsHits) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HitsHits) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HitsHits) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HitsHits) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScore

`func (o *HitsHits) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *HitsHits) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *HitsHits) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *HitsHits) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *HitsHits) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HitsHits) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HitsHits) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *HitsHits) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetKnnDist

`func (o *HitsHits) GetKnnDist() float32`

GetKnnDist returns the KnnDist field if non-nil, zero value otherwise.

### GetKnnDistOk

`func (o *HitsHits) GetKnnDistOk() (*float32, bool)`

GetKnnDistOk returns a tuple with the KnnDist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnnDist

`func (o *HitsHits) SetKnnDist(v float32)`

SetKnnDist sets KnnDist field to given value.

### HasKnnDist

`func (o *HitsHits) HasKnnDist() bool`

HasKnnDist returns a boolean if a field has been set.

### GetHighlight

`func (o *HitsHits) GetHighlight() map[string]interface{}`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *HitsHits) GetHighlightOk() (*map[string]interface{}, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *HitsHits) SetHighlight(v map[string]interface{})`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *HitsHits) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.

### GetTable

`func (o *HitsHits) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *HitsHits) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *HitsHits) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *HitsHits) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetType

`func (o *HitsHits) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HitsHits) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HitsHits) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HitsHits) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFields

`func (o *HitsHits) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *HitsHits) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *HitsHits) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *HitsHits) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


