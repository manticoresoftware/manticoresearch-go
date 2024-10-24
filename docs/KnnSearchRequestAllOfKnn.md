# KnnSearchRequestAllOfKnn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocId** | **int64** |  | 
**Field** | **string** |  | 
**K** | Pointer to **int32** |  | [optional] 
**Ef** | Pointer to **int32** |  | [optional] 
**QueryVector** | **[]float32** |  | 

## Methods

### NewKnnSearchRequestAllOfKnn

`func NewKnnSearchRequestAllOfKnn(docId int64, field string, queryVector []float32, ) *KnnSearchRequestAllOfKnn`

NewKnnSearchRequestAllOfKnn instantiates a new KnnSearchRequestAllOfKnn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnSearchRequestAllOfKnnWithDefaults

`func NewKnnSearchRequestAllOfKnnWithDefaults() *KnnSearchRequestAllOfKnn`

NewKnnSearchRequestAllOfKnnWithDefaults instantiates a new KnnSearchRequestAllOfKnn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocId

`func (o *KnnSearchRequestAllOfKnn) GetDocId() int64`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *KnnSearchRequestAllOfKnn) GetDocIdOk() (*int64, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *KnnSearchRequestAllOfKnn) SetDocId(v int64)`

SetDocId sets DocId field to given value.


### GetField

`func (o *KnnSearchRequestAllOfKnn) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnSearchRequestAllOfKnn) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnSearchRequestAllOfKnn) SetField(v string)`

SetField sets Field field to given value.


### GetK

`func (o *KnnSearchRequestAllOfKnn) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnSearchRequestAllOfKnn) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnSearchRequestAllOfKnn) SetK(v int32)`

SetK sets K field to given value.

### HasK

`func (o *KnnSearchRequestAllOfKnn) HasK() bool`

HasK returns a boolean if a field has been set.

### GetEf

`func (o *KnnSearchRequestAllOfKnn) GetEf() int32`

GetEf returns the Ef field if non-nil, zero value otherwise.

### GetEfOk

`func (o *KnnSearchRequestAllOfKnn) GetEfOk() (*int32, bool)`

GetEfOk returns a tuple with the Ef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEf

`func (o *KnnSearchRequestAllOfKnn) SetEf(v int32)`

SetEf sets Ef field to given value.

### HasEf

`func (o *KnnSearchRequestAllOfKnn) HasEf() bool`

HasEf returns a boolean if a field has been set.

### GetQueryVector

`func (o *KnnSearchRequestAllOfKnn) GetQueryVector() []float32`

GetQueryVector returns the QueryVector field if non-nil, zero value otherwise.

### GetQueryVectorOk

`func (o *KnnSearchRequestAllOfKnn) GetQueryVectorOk() (*[]float32, bool)`

GetQueryVectorOk returns a tuple with the QueryVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVector

`func (o *KnnSearchRequestAllOfKnn) SetQueryVector(v []float32)`

SetQueryVector sets QueryVector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


