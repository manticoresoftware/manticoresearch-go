# SuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | Name of the document index | [optional] 
**Id** | Pointer to **int64** | ID of the document affected by the request operation | [optional] 
**Created** | Pointer to **bool** | Indicates whether the document was created as a result of the operation | [optional] 
**Result** | Pointer to **string** | Result of the operation, typically &#39;created&#39;, &#39;updated&#39;, or &#39;deleted&#39; | [optional] 
**Found** | Pointer to **bool** | Indicates whether the document was found in the index | [optional] 
**Status** | Pointer to **int32** | HTTP status code representing the result of the operation | [optional] 

## Methods

### NewSuccessResponse

`func NewSuccessResponse() *SuccessResponse`

NewSuccessResponse instantiates a new SuccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessResponseWithDefaults

`func NewSuccessResponseWithDefaults() *SuccessResponse`

NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SuccessResponse) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SuccessResponse) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SuccessResponse) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SuccessResponse) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetId

`func (o *SuccessResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuccessResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuccessResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SuccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *SuccessResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SuccessResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SuccessResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SuccessResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetResult

`func (o *SuccessResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SuccessResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SuccessResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *SuccessResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFound

`func (o *SuccessResponse) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SuccessResponse) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SuccessResponse) SetFound(v bool)`

SetFound sets Found field to given value.

### HasFound

`func (o *SuccessResponse) HasFound() bool`

HasFound returns a boolean if a field has been set.

### GetStatus

`func (o *SuccessResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SuccessResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SuccessResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SuccessResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


