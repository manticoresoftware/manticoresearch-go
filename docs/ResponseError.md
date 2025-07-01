# ResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type or category of the error | 
**Reason** | Pointer to **string** | Detailed explanation of why the error occurred | [optional] 
**Table** | Pointer to **string** | The table related to the error, if applicable | [optional] 

## Methods

### NewResponseError

`func NewResponseError(type_ string, ) *ResponseError`

NewResponseError instantiates a new ResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseErrorWithDefaults

`func NewResponseErrorWithDefaults() *ResponseError`

NewResponseErrorWithDefaults instantiates a new ResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *ResponseError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTable

`func (o *ResponseError) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *ResponseError) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *ResponseError) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *ResponseError) HasTable() bool`

HasTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


