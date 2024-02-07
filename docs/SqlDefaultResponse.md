# SqlDefaultResponse

Response object from sql request.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSqlDefaultResponse

`func NewSqlDefaultResponse(error_ string, ) *SqlDefaultResponse`

NewSqlDefaultResponse instantiates a new SqlDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlDefaultResponseWithDefaults

`func NewSqlDefaultResponseWithDefaults() *SqlDefaultResponse`

NewSqlDefaultResponseWithDefaults instantiates a new SqlDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SqlDefaultResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SqlDefaultResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SqlDefaultResponse) SetError(v string)`

SetError sets Error field to given value.


### GetStatus

`func (o *SqlDefaultResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlDefaultResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlDefaultResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlDefaultResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


