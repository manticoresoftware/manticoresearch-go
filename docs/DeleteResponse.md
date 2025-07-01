# DeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to **interface{}** | The name of the table from which the document was deleted | [optional] 
**Deleted** | Pointer to **interface{}** | Number of documents deleted | [optional] 
**Id** | Pointer to **interface{}** | The ID of the deleted document. If multiple documents are deleted, the ID of the first deleted document is returned | [optional] 
**Found** | Pointer to **interface{}** | Indicates whether any documents to be deleted were found | [optional] 
**Result** | Pointer to **interface{}** | Result of the delete operation, typically &#39;deleted&#39; | [optional] 

## Methods

### NewDeleteResponse

`func NewDeleteResponse() *DeleteResponse`

NewDeleteResponse instantiates a new DeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteResponseWithDefaults

`func NewDeleteResponseWithDefaults() *DeleteResponse`

NewDeleteResponseWithDefaults instantiates a new DeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *DeleteResponse) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *DeleteResponse) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *DeleteResponse) SetTable(v interface{})`

SetTable sets Table field to given value.

### HasTable

`func (o *DeleteResponse) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTableNil

`func (o *DeleteResponse) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *DeleteResponse) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetDeleted

`func (o *DeleteResponse) GetDeleted() interface{}`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteResponse) GetDeletedOk() (*interface{}, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteResponse) SetDeleted(v interface{})`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *DeleteResponse) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *DeleteResponse) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetId

`func (o *DeleteResponse) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteResponse) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteResponse) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *DeleteResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeleteResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeleteResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetFound

`func (o *DeleteResponse) GetFound() interface{}`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *DeleteResponse) GetFoundOk() (*interface{}, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *DeleteResponse) SetFound(v interface{})`

SetFound sets Found field to given value.

### HasFound

`func (o *DeleteResponse) HasFound() bool`

HasFound returns a boolean if a field has been set.

### SetFoundNil

`func (o *DeleteResponse) SetFoundNil(b bool)`

 SetFoundNil sets the value for Found to be an explicit nil

### UnsetFound
`func (o *DeleteResponse) UnsetFound()`

UnsetFound ensures that no value is present for Found, not even an explicit nil
### GetResult

`func (o *DeleteResponse) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteResponse) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteResponse) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *DeleteResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *DeleteResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *DeleteResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


