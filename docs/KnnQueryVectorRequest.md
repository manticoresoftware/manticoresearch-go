# KnnQueryVectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**K** | Pointer to **int32** |  | [optional] 
**Ef** | Pointer to **int32** |  | [optional] 
**QueryVector** | **[]float32** |  | 

## Methods

### NewKnnQueryVectorRequest

`func NewKnnQueryVectorRequest(field string, queryVector []float32, ) *KnnQueryVectorRequest`

NewKnnQueryVectorRequest instantiates a new KnnQueryVectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnQueryVectorRequestWithDefaults

`func NewKnnQueryVectorRequestWithDefaults() *KnnQueryVectorRequest`

NewKnnQueryVectorRequestWithDefaults instantiates a new KnnQueryVectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *KnnQueryVectorRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnQueryVectorRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnQueryVectorRequest) SetField(v string)`

SetField sets Field field to given value.


### GetK

`func (o *KnnQueryVectorRequest) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnQueryVectorRequest) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnQueryVectorRequest) SetK(v int32)`

SetK sets K field to given value.

### HasK

`func (o *KnnQueryVectorRequest) HasK() bool`

HasK returns a boolean if a field has been set.

### GetEf

`func (o *KnnQueryVectorRequest) GetEf() int32`

GetEf returns the Ef field if non-nil, zero value otherwise.

### GetEfOk

`func (o *KnnQueryVectorRequest) GetEfOk() (*int32, bool)`

GetEfOk returns a tuple with the Ef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEf

`func (o *KnnQueryVectorRequest) SetEf(v int32)`

SetEf sets Ef field to given value.

### HasEf

`func (o *KnnQueryVectorRequest) HasEf() bool`

HasEf returns a boolean if a field has been set.

### GetQueryVector

`func (o *KnnQueryVectorRequest) GetQueryVector() []float32`

GetQueryVector returns the QueryVector field if non-nil, zero value otherwise.

### GetQueryVectorOk

`func (o *KnnQueryVectorRequest) GetQueryVectorOk() (*[]float32, bool)`

GetQueryVectorOk returns a tuple with the QueryVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVector

`func (o *KnnQueryVectorRequest) SetQueryVector(v []float32)`

SetQueryVector sets QueryVector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


