# RangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Lte** | Pointer to **NullableFloat32** |  | [optional] 
**Gte** | Pointer to **NullableFloat32** |  | [optional] 
**Lt** | Pointer to **NullableFloat32** |  | [optional] 
**Gt** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewRangeFilter

`func NewRangeFilter(field string, ) *RangeFilter`

NewRangeFilter instantiates a new RangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeFilterWithDefaults

`func NewRangeFilterWithDefaults() *RangeFilter`

NewRangeFilterWithDefaults instantiates a new RangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *RangeFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RangeFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RangeFilter) SetField(v string)`

SetField sets Field field to given value.


### GetLte

`func (o *RangeFilter) GetLte() float32`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *RangeFilter) GetLteOk() (*float32, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *RangeFilter) SetLte(v float32)`

SetLte sets Lte field to given value.

### HasLte

`func (o *RangeFilter) HasLte() bool`

HasLte returns a boolean if a field has been set.

### SetLteNil

`func (o *RangeFilter) SetLteNil(b bool)`

 SetLteNil sets the value for Lte to be an explicit nil

### UnsetLte
`func (o *RangeFilter) UnsetLte()`

UnsetLte ensures that no value is present for Lte, not even an explicit nil
### GetGte

`func (o *RangeFilter) GetGte() float32`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *RangeFilter) GetGteOk() (*float32, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *RangeFilter) SetGte(v float32)`

SetGte sets Gte field to given value.

### HasGte

`func (o *RangeFilter) HasGte() bool`

HasGte returns a boolean if a field has been set.

### SetGteNil

`func (o *RangeFilter) SetGteNil(b bool)`

 SetGteNil sets the value for Gte to be an explicit nil

### UnsetGte
`func (o *RangeFilter) UnsetGte()`

UnsetGte ensures that no value is present for Gte, not even an explicit nil
### GetLt

`func (o *RangeFilter) GetLt() float32`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *RangeFilter) GetLtOk() (*float32, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *RangeFilter) SetLt(v float32)`

SetLt sets Lt field to given value.

### HasLt

`func (o *RangeFilter) HasLt() bool`

HasLt returns a boolean if a field has been set.

### SetLtNil

`func (o *RangeFilter) SetLtNil(b bool)`

 SetLtNil sets the value for Lt to be an explicit nil

### UnsetLt
`func (o *RangeFilter) UnsetLt()`

UnsetLt ensures that no value is present for Lt, not even an explicit nil
### GetGt

`func (o *RangeFilter) GetGt() float32`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *RangeFilter) GetGtOk() (*float32, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *RangeFilter) SetGt(v float32)`

SetGt sets Gt field to given value.

### HasGt

`func (o *RangeFilter) HasGt() bool`

HasGt returns a boolean if a field has been set.

### SetGtNil

`func (o *RangeFilter) SetGtNil(b bool)`

 SetGtNil sets the value for Gt to be an explicit nil

### UnsetGt
`func (o *RangeFilter) UnsetGt()`

UnsetGt ensures that no value is present for Gt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


