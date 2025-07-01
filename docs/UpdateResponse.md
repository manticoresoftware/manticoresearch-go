# UpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to **interface{}** | Name of the document table | [optional] 
**Updated** | Pointer to **interface{}** | Number of documents updated | [optional] 
**Id** | Pointer to **interface{}** | Document ID | [optional] 
**Result** | Pointer to **interface{}** | Result of the update operation, typically &#39;updated&#39; | [optional] 

## Methods

### NewUpdateResponse

`func NewUpdateResponse() *UpdateResponse`

NewUpdateResponse instantiates a new UpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResponseWithDefaults

`func NewUpdateResponseWithDefaults() *UpdateResponse`

NewUpdateResponseWithDefaults instantiates a new UpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *UpdateResponse) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *UpdateResponse) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *UpdateResponse) SetTable(v interface{})`

SetTable sets Table field to given value.

### HasTable

`func (o *UpdateResponse) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTableNil

`func (o *UpdateResponse) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *UpdateResponse) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetUpdated

`func (o *UpdateResponse) GetUpdated() interface{}`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpdateResponse) GetUpdatedOk() (*interface{}, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpdateResponse) SetUpdated(v interface{})`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpdateResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *UpdateResponse) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *UpdateResponse) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetId

`func (o *UpdateResponse) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateResponse) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateResponse) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *UpdateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdateResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResult

`func (o *UpdateResponse) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateResponse) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateResponse) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *UpdateResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *UpdateResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


