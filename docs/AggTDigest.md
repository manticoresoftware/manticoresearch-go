# AggTDigest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compression** | Pointer to **float32** | Accuracy/memory trade-off. Default is 200. | [optional] 

## Methods

### NewAggTDigest

`func NewAggTDigest() *AggTDigest`

NewAggTDigest instantiates a new AggTDigest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggTDigestWithDefaults

`func NewAggTDigestWithDefaults() *AggTDigest`

NewAggTDigestWithDefaults instantiates a new AggTDigest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompression

`func (o *AggTDigest) GetCompression() float32`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *AggTDigest) GetCompressionOk() (*float32, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *AggTDigest) SetCompression(v float32)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *AggTDigest) HasCompression() bool`

HasCompression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


