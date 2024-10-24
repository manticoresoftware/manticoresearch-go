# UpdateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **string** | Name of the document index | 
**Cluster** | Pointer to **string** | Name of the document cluster | [optional] 
**Doc** | **map[string]interface{}** | Object containing the document fields to update | 
**Id** | Pointer to **int64** | Document ID | [optional] 
**Query** | Pointer to [**NullableQueryFilter**](QueryFilter.md) |  | [optional] 

## Methods

### NewUpdateDocumentRequest

`func NewUpdateDocumentRequest(index string, doc map[string]interface{}, ) *UpdateDocumentRequest`

NewUpdateDocumentRequest instantiates a new UpdateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentRequestWithDefaults

`func NewUpdateDocumentRequestWithDefaults() *UpdateDocumentRequest`

NewUpdateDocumentRequestWithDefaults instantiates a new UpdateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *UpdateDocumentRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *UpdateDocumentRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *UpdateDocumentRequest) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetCluster

`func (o *UpdateDocumentRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateDocumentRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateDocumentRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *UpdateDocumentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDoc

`func (o *UpdateDocumentRequest) GetDoc() map[string]interface{}`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *UpdateDocumentRequest) GetDocOk() (*map[string]interface{}, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *UpdateDocumentRequest) SetDoc(v map[string]interface{})`

SetDoc sets Doc field to given value.


### GetId

`func (o *UpdateDocumentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDocumentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDocumentRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateDocumentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuery

`func (o *UpdateDocumentRequest) GetQuery() QueryFilter`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateDocumentRequest) GetQueryOk() (*QueryFilter, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateDocumentRequest) SetQuery(v QueryFilter)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *UpdateDocumentRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *UpdateDocumentRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *UpdateDocumentRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


