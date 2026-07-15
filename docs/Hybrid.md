# Hybrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Search query string | 
**Field** | Pointer to **string** | Vector field name.  | [optional] 

## Methods

### NewHybrid

`func NewHybrid(query string, ) *Hybrid`

NewHybrid instantiates a new Hybrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHybridWithDefaults

`func NewHybridWithDefaults() *Hybrid`

NewHybridWithDefaults instantiates a new Hybrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *Hybrid) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Hybrid) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Hybrid) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetField

`func (o *Hybrid) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Hybrid) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Hybrid) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Hybrid) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


