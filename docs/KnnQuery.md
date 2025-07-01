# KnnQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **interface{}** | Field to perform the k-nearest neighbor search on | 
**K** | **interface{}** | The number of nearest neighbors to return | 
**QueryVector** | Pointer to **interface{}** | The vector used as input for the KNN search | [optional] 
**DocId** | Pointer to **interface{}** | The docuemnt ID used as input for the KNN search | [optional] 
**Ef** | Pointer to **interface{}** | Optional parameter controlling the accuracy of the search | [optional] 
**Filter** | Pointer to [**QueryFilter**](QueryFilter.md) |  | [optional] 

## Methods

### NewKnnQuery

`func NewKnnQuery(field interface{}, k interface{}, ) *KnnQuery`

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

`func (o *KnnQuery) GetField() interface{}`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnQuery) GetFieldOk() (*interface{}, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnQuery) SetField(v interface{})`

SetField sets Field field to given value.


### SetFieldNil

`func (o *KnnQuery) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *KnnQuery) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetK

`func (o *KnnQuery) GetK() interface{}`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnQuery) GetKOk() (*interface{}, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnQuery) SetK(v interface{})`

SetK sets K field to given value.


### SetKNil

`func (o *KnnQuery) SetKNil(b bool)`

 SetKNil sets the value for K to be an explicit nil

### UnsetK
`func (o *KnnQuery) UnsetK()`

UnsetK ensures that no value is present for K, not even an explicit nil
### GetQueryVector

`func (o *KnnQuery) GetQueryVector() interface{}`

GetQueryVector returns the QueryVector field if non-nil, zero value otherwise.

### GetQueryVectorOk

`func (o *KnnQuery) GetQueryVectorOk() (*interface{}, bool)`

GetQueryVectorOk returns a tuple with the QueryVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryVector

`func (o *KnnQuery) SetQueryVector(v interface{})`

SetQueryVector sets QueryVector field to given value.

### HasQueryVector

`func (o *KnnQuery) HasQueryVector() bool`

HasQueryVector returns a boolean if a field has been set.

### SetQueryVectorNil

`func (o *KnnQuery) SetQueryVectorNil(b bool)`

 SetQueryVectorNil sets the value for QueryVector to be an explicit nil

### UnsetQueryVector
`func (o *KnnQuery) UnsetQueryVector()`

UnsetQueryVector ensures that no value is present for QueryVector, not even an explicit nil
### GetDocId

`func (o *KnnQuery) GetDocId() interface{}`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *KnnQuery) GetDocIdOk() (*interface{}, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *KnnQuery) SetDocId(v interface{})`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *KnnQuery) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### SetDocIdNil

`func (o *KnnQuery) SetDocIdNil(b bool)`

 SetDocIdNil sets the value for DocId to be an explicit nil

### UnsetDocId
`func (o *KnnQuery) UnsetDocId()`

UnsetDocId ensures that no value is present for DocId, not even an explicit nil
### GetEf

`func (o *KnnQuery) GetEf() interface{}`

GetEf returns the Ef field if non-nil, zero value otherwise.

### GetEfOk

`func (o *KnnQuery) GetEfOk() (*interface{}, bool)`

GetEfOk returns a tuple with the Ef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEf

`func (o *KnnQuery) SetEf(v interface{})`

SetEf sets Ef field to given value.

### HasEf

`func (o *KnnQuery) HasEf() bool`

HasEf returns a boolean if a field has been set.

### SetEfNil

`func (o *KnnQuery) SetEfNil(b bool)`

 SetEfNil sets the value for Ef to be an explicit nil

### UnsetEf
`func (o *KnnQuery) UnsetEf()`

UnsetEf ensures that no value is present for Ef, not even an explicit nil
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


