# JoinInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | [**[]JoinInnerOnInner**](JoinInnerOnInner.md) |  | 
**Query** | Pointer to [**FulltextFilter**](FulltextFilter.md) |  | [optional] 
**Table** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewJoinInner

`func NewJoinInner(on []JoinInnerOnInner, table string, type_ string, ) *JoinInner`

NewJoinInner instantiates a new JoinInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinInnerWithDefaults

`func NewJoinInnerWithDefaults() *JoinInner`

NewJoinInnerWithDefaults instantiates a new JoinInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOn

`func (o *JoinInner) GetOn() []JoinInnerOnInner`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *JoinInner) GetOnOk() (*[]JoinInnerOnInner, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *JoinInner) SetOn(v []JoinInnerOnInner)`

SetOn sets On field to given value.


### GetQuery

`func (o *JoinInner) GetQuery() FulltextFilter`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *JoinInner) GetQueryOk() (*FulltextFilter, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *JoinInner) SetQuery(v FulltextFilter)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *JoinInner) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTable

`func (o *JoinInner) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *JoinInner) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *JoinInner) SetTable(v string)`

SetTable sets Table field to given value.


### GetType

`func (o *JoinInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JoinInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JoinInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


