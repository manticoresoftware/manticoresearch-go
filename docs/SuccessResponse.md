# SuccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to **interface{}** | Name of the document table | [optional] 
**Id** | Pointer to **interface{}** | ID of the document affected by the request operation | [optional] 
**Created** | Pointer to **interface{}** | Indicates whether the document was created as a result of the operation | [optional] 
**Result** | Pointer to **interface{}** | Result of the operation, typically &#39;created&#39;, &#39;updated&#39;, or &#39;deleted&#39; | [optional] 
**Found** | Pointer to **interface{}** | Indicates whether the document was found in the table | [optional] 
**Status** | Pointer to **interface{}** | HTTP status code representing the result of the operation | [optional] 

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

### GetTable

`func (o *SuccessResponse) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *SuccessResponse) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *SuccessResponse) SetTable(v interface{})`

SetTable sets Table field to given value.

### HasTable

`func (o *SuccessResponse) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTableNil

`func (o *SuccessResponse) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *SuccessResponse) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetId

`func (o *SuccessResponse) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuccessResponse) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuccessResponse) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *SuccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SuccessResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SuccessResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreated

`func (o *SuccessResponse) GetCreated() interface{}`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SuccessResponse) GetCreatedOk() (*interface{}, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SuccessResponse) SetCreated(v interface{})`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SuccessResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *SuccessResponse) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SuccessResponse) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetResult

`func (o *SuccessResponse) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SuccessResponse) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SuccessResponse) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *SuccessResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *SuccessResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *SuccessResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetFound

`func (o *SuccessResponse) GetFound() interface{}`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SuccessResponse) GetFoundOk() (*interface{}, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SuccessResponse) SetFound(v interface{})`

SetFound sets Found field to given value.

### HasFound

`func (o *SuccessResponse) HasFound() bool`

HasFound returns a boolean if a field has been set.

### SetFoundNil

`func (o *SuccessResponse) SetFoundNil(b bool)`

 SetFoundNil sets the value for Found to be an explicit nil

### UnsetFound
`func (o *SuccessResponse) UnsetFound()`

UnsetFound ensures that no value is present for Found, not even an explicit nil
### GetStatus

`func (o *SuccessResponse) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SuccessResponse) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SuccessResponse) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SuccessResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SuccessResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SuccessResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


