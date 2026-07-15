# AggBucketsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to **interface{}** |  | [optional] 
**AfterKey** | Pointer to **map[string]interface{}** | Pagination cursor returned by composite aggregations | [optional] 

## Methods

### NewAggBucketsResult

`func NewAggBucketsResult() *AggBucketsResult`

NewAggBucketsResult instantiates a new AggBucketsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggBucketsResultWithDefaults

`func NewAggBucketsResultWithDefaults() *AggBucketsResult`

NewAggBucketsResultWithDefaults instantiates a new AggBucketsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *AggBucketsResult) GetBuckets() interface{}`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AggBucketsResult) GetBucketsOk() (*interface{}, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AggBucketsResult) SetBuckets(v interface{})`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *AggBucketsResult) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### SetBucketsNil

`func (o *AggBucketsResult) SetBucketsNil(b bool)`

 SetBucketsNil sets the value for Buckets to be an explicit nil

### UnsetBuckets
`func (o *AggBucketsResult) UnsetBuckets()`

UnsetBuckets ensures that no value is present for Buckets, not even an explicit nil
### GetAfterKey

`func (o *AggBucketsResult) GetAfterKey() map[string]interface{}`

GetAfterKey returns the AfterKey field if non-nil, zero value otherwise.

### GetAfterKeyOk

`func (o *AggBucketsResult) GetAfterKeyOk() (*map[string]interface{}, bool)`

GetAfterKeyOk returns a tuple with the AfterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterKey

`func (o *AggBucketsResult) SetAfterKey(v map[string]interface{})`

SetAfterKey sets AfterKey field to given value.

### HasAfterKey

`func (o *AggBucketsResult) HasAfterKey() bool`

HasAfterKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


