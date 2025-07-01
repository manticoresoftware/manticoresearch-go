# ResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | Type or category of the error | 
**Reason** | Pointer to **interface{}** | Detailed explanation of why the error occurred | [optional] 
**Table** | Pointer to **interface{}** | The table related to the error, if applicable | [optional] 

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
### GetTable

`func (o *ResponseError) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *ResponseError) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *ResponseError) SetTable(v interface{})`

SetTable sets Table field to given value.

### HasTable

`func (o *ResponseError) HasTable() bool`

HasTable returns a boolean if a field has been set.

### SetTableNil

`func (o *ResponseError) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *ResponseError) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


