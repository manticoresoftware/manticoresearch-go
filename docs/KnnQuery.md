# KnnQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field to perform the k-nearest neighbor search on | 
**K** | **int32** | The number of nearest neighbors to return | 
**QueryVector** | Pointer to **[]float32** | The vector used as input for the KNN search | [optional] 
**DocId** | Pointer to **int64** | The docuemnt ID used as input for the KNN search | [optional] 
**Ef** | Pointer to **int32** | Optional parameter controlling the accuracy of the search | [optional] 
**Filter** | Pointer to [**QueryFilter**](QueryFilter.md) |  | [optional] 

## Methods

### NewKnnQuery

`func NewKnnQuery(field string, k int32, ) *KnnQuery`

NewKnnQuery instantiates a new KnnQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnQueryWithDefaults

`func NewKnnQueryWithDefaults() *KnnQuery`

NewKnnQueryWithDefaults instantiates a new KnnQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *KnnQuery) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnQuery) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnQuery) SetField(v string)`

SetField sets Field field to given value.


### GetK

`func (o *KnnQuery) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnQuery) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnQuery) SetK(v int32)`

SetK sets K field to given value.


### GetQueryVector

`func (o *KnnQuery) GetQueryVector() []float32`

GetQueryVector returns the QueryVector field if non-nil, zero value otherwise.

### GetQueryVectorOk

`func (o *KnnQuery) GetQueryVectorOk() (*[]float32, bool)`

GetQueryVectorOk returns a tuple with the QueryVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVector

`func (o *KnnQuery) SetQueryVector(v []float32)`

SetQueryVector sets QueryVector field to given value.

### HasQueryVector

`func (o *KnnQuery) HasQueryVector() bool`

HasQueryVector returns a boolean if a field has been set.

### GetDocId

`func (o *KnnQuery) GetDocId() int64`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *KnnQuery) GetDocIdOk() (*int64, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *KnnQuery) SetDocId(v int64)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *KnnQuery) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetEf

`func (o *KnnQuery) GetEf() int32`

GetEf returns the Ef field if non-nil, zero value otherwise.

### GetEfOk

`func (o *KnnQuery) GetEfOk() (*int32, bool)`

GetEfOk returns a tuple with the Ef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEf

`func (o *KnnQuery) SetEf(v int32)`

SetEf sets Ef field to given value.

### HasEf

`func (o *KnnQuery) HasEf() bool`

HasEf returns a boolean if a field has been set.

### GetFilter

`func (o *KnnQuery) GetFilter() QueryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *KnnQuery) GetFilterOk() (*QueryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *KnnQuery) SetFilter(v QueryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *KnnQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


