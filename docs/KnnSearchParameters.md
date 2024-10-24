# KnnSearchParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**K** | Pointer to **int32** |  | [optional] 
**Ef** | Pointer to **int32** |  | [optional] 

## Methods

### NewKnnSearchParameters

`func NewKnnSearchParameters(field string, ) *KnnSearchParameters`

NewKnnSearchParameters instantiates a new KnnSearchParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnSearchParametersWithDefaults

`func NewKnnSearchParametersWithDefaults() *KnnSearchParameters`

NewKnnSearchParametersWithDefaults instantiates a new KnnSearchParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *KnnSearchParameters) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnSearchParameters) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnSearchParameters) SetField(v string)`

SetField sets Field field to given value.


### GetK

`func (o *KnnSearchParameters) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnSearchParameters) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnSearchParameters) SetK(v int32)`

SetK sets K field to given value.

### HasK

`func (o *KnnSearchParameters) HasK() bool`

HasK returns a boolean if a field has been set.

### GetEf

`func (o *KnnSearchParameters) GetEf() int32`

GetEf returns the Ef field if non-nil, zero value otherwise.

### GetEfOk

`func (o *KnnSearchParameters) GetEfOk() (*int32, bool)`

GetEfOk returns a tuple with the Ef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEf

`func (o *KnnSearchParameters) SetEf(v int32)`

SetEf sets Ef field to given value.

### HasEf

`func (o *KnnSearchParameters) HasEf() bool`

HasEf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


