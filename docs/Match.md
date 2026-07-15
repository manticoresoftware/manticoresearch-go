# Match

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Operator** | Pointer to **string** |  | [optional] 
**Boost** | Pointer to **float32** |  | [optional] 

## Methods

### NewMatch

`func NewMatch(query string, ) *Match`

NewMatch instantiates a new Match object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchWithDefaults

`func NewMatchWithDefaults() *Match`

NewMatchWithDefaults instantiates a new Match object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *Match) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Match) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Match) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetOperator

`func (o *Match) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Match) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Match) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Match) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetBoost

`func (o *Match) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *Match) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *Match) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *Match) HasBoost() bool`

HasBoost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


