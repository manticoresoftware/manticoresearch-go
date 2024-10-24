# KnnDocIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**K** | Pointer to **int32** |  | [optional] 
**Ef** | Pointer to **int32** |  | [optional] 
**DocId** | **int64** |  | 

## Methods

### NewKnnDocIdRequest

`func NewKnnDocIdRequest(field string, docId int64, ) *KnnDocIdRequest`

NewKnnDocIdRequest instantiates a new KnnDocIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnnDocIdRequestWithDefaults

`func NewKnnDocIdRequestWithDefaults() *KnnDocIdRequest`

NewKnnDocIdRequestWithDefaults instantiates a new KnnDocIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *KnnDocIdRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KnnDocIdRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KnnDocIdRequest) SetField(v string)`

SetField sets Field field to given value.


### GetK

`func (o *KnnDocIdRequest) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *KnnDocIdRequest) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *KnnDocIdRequest) SetK(v int32)`

SetK sets K field to given value.

### HasK

`func (o *KnnDocIdRequest) HasK() bool`

HasK returns a boolean if a field has been set.

### GetEf

`func (o *KnnDocIdRequest) GetEf() int32`

GetEf returns the Ef field if non-nil, zero value otherwise.

### GetEfOk

`func (o *KnnDocIdRequest) GetEfOk() (*int32, bool)`

GetEfOk returns a tuple with the Ef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEf

`func (o *KnnDocIdRequest) SetEf(v int32)`

SetEf sets Ef field to given value.

### HasEf

`func (o *KnnDocIdRequest) HasEf() bool`

HasEf returns a boolean if a field has been set.

### GetDocId

`func (o *KnnDocIdRequest) GetDocId() int64`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *KnnDocIdRequest) GetDocIdOk() (*int64, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *KnnDocIdRequest) SetDocId(v int64)`

SetDocId sets DocId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


