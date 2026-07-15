# AggBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **interface{}** |  | [optional] 
**DocCount** | **int32** | Number of documents in the bucket | 
**Status** | Pointer to [**FacetBucketStatus**](FacetBucketStatus.md) |  | [optional] 

## Methods

### NewAggBucket

`func NewAggBucket(docCount int32, ) *AggBucket`

NewAggBucket instantiates a new AggBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggBucketWithDefaults

`func NewAggBucketWithDefaults() *AggBucket`

NewAggBucketWithDefaults instantiates a new AggBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AggBucket) GetKey() interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AggBucket) GetKeyOk() (*interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AggBucket) SetKey(v interface{})`

SetKey sets Key field to given value.

### HasKey

`func (o *AggBucket) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *AggBucket) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *AggBucket) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetDocCount

`func (o *AggBucket) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *AggBucket) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *AggBucket) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.


### GetStatus

`func (o *AggBucket) GetStatus() FacetBucketStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggBucket) GetStatusOk() (*FacetBucketStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggBucket) SetStatus(v FacetBucketStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AggBucket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


