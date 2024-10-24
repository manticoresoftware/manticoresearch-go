# ResponseErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type or category of the error | 
**Reason** | Pointer to **NullableString** | Detailed explanation of why the error occurred | [optional] 
**Index** | Pointer to **NullableString** | The index related to the error, if applicable | [optional] 

## Methods

### NewResponseErrorDetails

`func NewResponseErrorDetails(type_ string, ) *ResponseErrorDetails`

NewResponseErrorDetails instantiates a new ResponseErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseErrorDetailsWithDefaults

`func NewResponseErrorDetailsWithDefaults() *ResponseErrorDetails`

NewResponseErrorDetailsWithDefaults instantiates a new ResponseErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseErrorDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseErrorDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseErrorDetails) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *ResponseErrorDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResponseErrorDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResponseErrorDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResponseErrorDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ResponseErrorDetails) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ResponseErrorDetails) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetIndex

`func (o *ResponseErrorDetails) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ResponseErrorDetails) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ResponseErrorDetails) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ResponseErrorDetails) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *ResponseErrorDetails) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *ResponseErrorDetails) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


