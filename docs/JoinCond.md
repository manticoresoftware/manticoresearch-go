# JoinCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **interface{}** | Field to join on | 
**Table** | **interface{}** | Joined table | 
**Type** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewJoinCond

`func NewJoinCond(field interface{}, table interface{}, ) *JoinCond`

NewJoinCond instantiates a new JoinCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinCondWithDefaults

`func NewJoinCondWithDefaults() *JoinCond`

NewJoinCondWithDefaults instantiates a new JoinCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *JoinCond) GetField() interface{}`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *JoinCond) GetFieldOk() (*interface{}, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *JoinCond) SetField(v interface{})`

SetField sets Field field to given value.


### SetFieldNil

`func (o *JoinCond) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *JoinCond) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetTable

`func (o *JoinCond) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *JoinCond) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *JoinCond) SetTable(v interface{})`

SetTable sets Table field to given value.


### SetTableNil

`func (o *JoinCond) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *JoinCond) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetType

`func (o *JoinCond) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JoinCond) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JoinCond) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *JoinCond) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *JoinCond) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *JoinCond) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


