# DeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | The name of the index from which the document was deleted | [optional] 
**Deleted** | Pointer to **int32** | Number of documents deleted | [optional] 
**Id** | Pointer to **int64** | The ID of the deleted document. If multiple documents are deleted, the ID of the first deleted document is returned | [optional] 
**Found** | Pointer to **bool** | Indicates whether any documents to be deleted were found | [optional] 
**Result** | Pointer to **string** | Result of the delete operation, typically &#39;deleted&#39; | [optional] 

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

### GetIndex

`func (o *DeleteResponse) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DeleteResponse) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DeleteResponse) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DeleteResponse) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetDeleted

`func (o *DeleteResponse) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteResponse) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteResponse) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetId

`func (o *DeleteResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFound

`func (o *DeleteResponse) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *DeleteResponse) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *DeleteResponse) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *DeleteResponse) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetResult

`func (o *DeleteResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DeleteResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


