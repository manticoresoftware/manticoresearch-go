# Join

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | Type of the join operation | 
**On** | **interface{}** | List of objects defining joined tables | 
**Query** | Pointer to [**FulltextFilter**](FulltextFilter.md) |  | [optional] 
**Table** | **interface{}** | Basic table of the join operation | 

## Methods

### NewJoin

`func NewJoin(type_ interface{}, on interface{}, table interface{}, ) *Join`

NewJoin instantiates a new Join object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinWithDefaults

`func NewJoinWithDefaults() *Join`

NewJoinWithDefaults instantiates a new Join object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Join) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Join) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Join) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Join) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Join) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetOn

`func (o *Join) GetOn() interface{}`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *Join) GetOnOk() (*interface{}, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *Join) SetOn(v interface{})`

SetOn sets On field to given value.


### SetOnNil

`func (o *Join) SetOnNil(b bool)`

 SetOnNil sets the value for On to be an explicit nil

### UnsetOn
`func (o *Join) UnsetOn()`

UnsetOn ensures that no value is present for On, not even an explicit nil
### GetQuery

`func (o *Join) GetQuery() FulltextFilter`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Join) GetQueryOk() (*FulltextFilter, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Join) SetQuery(v FulltextFilter)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Join) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTable

`func (o *Join) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *Join) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *Join) SetTable(v interface{})`

SetTable sets Table field to given value.


### SetTableNil

`func (o *Join) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *Join) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


