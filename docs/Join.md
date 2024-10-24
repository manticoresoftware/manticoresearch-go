# Join

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the join operation | 
**On** | [**[]JoinOn**](JoinOn.md) | List of objects defining joined tables | 
**Query** | Pointer to [**FulltextFilter**](FulltextFilter.md) |  | [optional] 
**Table** | **string** | Basic table of the join operation | 

## Methods

### NewJoin

`func NewJoin(type_ string, on []JoinOn, table string, ) *Join`

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

`func (o *Join) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Join) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Join) SetType(v string)`

SetType sets Type field to given value.


### GetOn

`func (o *Join) GetOn() []JoinOn`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *Join) GetOnOk() (*[]JoinOn, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *Join) SetOn(v []JoinOn)`

SetOn sets On field to given value.


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

`func (o *Join) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *Join) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *Join) SetTable(v string)`

SetTable sets Table field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


