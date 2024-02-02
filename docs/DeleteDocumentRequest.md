# DeleteDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** | Index name | 
**Cluster** | Pointer to **string** | cluster name | [optional] 
**Id** | Pointer to **int64** | Document ID | [optional] 
**Query** | Pointer to **map[string]interface{}** | Query tree object | [optional] 

## Methods

### NewDeleteDocumentRequest

`func NewDeleteDocumentRequest(index string, ) *DeleteDocumentRequest`

NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentRequestWithDefaults

`func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest`

NewDeleteDocumentRequestWithDefaults instantiates a new DeleteDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *DeleteDocumentRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DeleteDocumentRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DeleteDocumentRequest) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetCluster

`func (o *DeleteDocumentRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeleteDocumentRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeleteDocumentRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DeleteDocumentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetId

`func (o *DeleteDocumentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteDocumentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteDocumentRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteDocumentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuery

`func (o *DeleteDocumentRequest) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DeleteDocumentRequest) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DeleteDocumentRequest) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DeleteDocumentRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


