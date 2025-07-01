# JoinOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Right** | Pointer to [**JoinCond**](JoinCond.md) |  | [optional] 
**Left** | Pointer to [**JoinCond**](JoinCond.md) |  | [optional] 
**Operator** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewJoinOn

`func NewJoinOn() *JoinOn`

NewJoinOn instantiates a new JoinOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinOnWithDefaults

`func NewJoinOnWithDefaults() *JoinOn`

NewJoinOnWithDefaults instantiates a new JoinOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRight

`func (o *JoinOn) GetRight() JoinCond`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *JoinOn) GetRightOk() (*JoinCond, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *JoinOn) SetRight(v JoinCond)`

SetRight sets Right field to given value.

### HasRight

`func (o *JoinOn) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetLeft

`func (o *JoinOn) GetLeft() JoinCond`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *JoinOn) GetLeftOk() (*JoinCond, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *JoinOn) SetLeft(v JoinCond)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *JoinOn) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetOperator

`func (o *JoinOn) GetOperator() interface{}`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JoinOn) GetOperatorOk() (*interface{}, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JoinOn) SetOperator(v interface{})`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *JoinOn) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *JoinOn) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *JoinOn) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


