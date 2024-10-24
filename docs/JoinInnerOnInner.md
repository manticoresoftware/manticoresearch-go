# JoinInnerOnInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Left** | Pointer to [**JoinInnerOnInnerLeft**](JoinInnerOnInnerLeft.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Right** | Pointer to [**JoinBasicCond**](JoinBasicCond.md) |  | [optional] 

## Methods

### NewJoinInnerOnInner

`func NewJoinInnerOnInner() *JoinInnerOnInner`

NewJoinInnerOnInner instantiates a new JoinInnerOnInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinInnerOnInnerWithDefaults

`func NewJoinInnerOnInnerWithDefaults() *JoinInnerOnInner`

NewJoinInnerOnInnerWithDefaults instantiates a new JoinInnerOnInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeft

`func (o *JoinInnerOnInner) GetLeft() JoinInnerOnInnerLeft`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *JoinInnerOnInner) GetLeftOk() (*JoinInnerOnInnerLeft, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *JoinInnerOnInner) SetLeft(v JoinInnerOnInnerLeft)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *JoinInnerOnInner) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetOperator

`func (o *JoinInnerOnInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *JoinInnerOnInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *JoinInnerOnInner) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *JoinInnerOnInner) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetRight

`func (o *JoinInnerOnInner) GetRight() JoinBasicCond`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *JoinInnerOnInner) GetRightOk() (*JoinBasicCond, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *JoinInnerOnInner) SetRight(v JoinBasicCond)`

SetRight sets Right field to given value.

### HasRight

`func (o *JoinInnerOnInner) HasRight() bool`

HasRight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


