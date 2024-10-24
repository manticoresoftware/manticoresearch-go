# ResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | Type or category of the error | 
**Reason** | Pointer to **interface{}** | Detailed explanation of why the error occurred | [optional] 
**Index** | Pointer to **interface{}** | The index related to the error, if applicable | [optional] 

## Methods

### NewResponseError

`func NewResponseError(type_ interface{}, ) *ResponseError`

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

`func (o *ResponseError) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseError) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseError) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ResponseError) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ResponseError) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetReason

`func (o *ResponseError) GetReason() interface{}`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseError) GetReasonOk() (*interface{}, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseError) SetReason(v interface{})`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ResponseError) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ResponseError) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetIndex

`func (o *ResponseError) GetIndex() interface{}`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ResponseError) GetIndexOk() (*interface{}, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ResponseError) SetIndex(v interface{})`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ResponseError) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *ResponseError) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *ResponseError) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


