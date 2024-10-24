# ErrorResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponseError

`func NewErrorResponseError(type_ string, ) *ErrorResponseError`

NewErrorResponseError instantiates a new ErrorResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseErrorWithDefaults

`func NewErrorResponseErrorWithDefaults() *ErrorResponseError`

NewErrorResponseErrorWithDefaults instantiates a new ErrorResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *ErrorResponseError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorResponseError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorResponseError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorResponseError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetIndex

`func (o *ErrorResponseError) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ErrorResponseError) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ErrorResponseError) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ErrorResponseError) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


