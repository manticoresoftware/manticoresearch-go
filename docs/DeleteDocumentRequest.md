# DeleteDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **interface{}** | Table name | 
**Cluster** | Pointer to **interface{}** | Cluster name | [optional] 
**Id** | Pointer to **interface{}** | The ID of document for deletion | [optional] 
**Query** | Pointer to **interface{}** | Defines the criteria to match documents for deletion | [optional] 

## Methods

### NewDeleteDocumentRequest

`func NewDeleteDocumentRequest(table interface{}, ) *DeleteDocumentRequest`

NewDeleteDocumentRequest instantiates a new DeleteDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentRequestWithDefaults

`func NewDeleteDocumentRequestWithDefaults() *DeleteDocumentRequest`

NewDeleteDocumentRequestWithDefaults instantiates a new DeleteDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *DeleteDocumentRequest) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *DeleteDocumentRequest) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *DeleteDocumentRequest) SetTable(v interface{})`

SetTable sets Table field to given value.


### SetTableNil

`func (o *DeleteDocumentRequest) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *DeleteDocumentRequest) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetCluster

`func (o *DeleteDocumentRequest) GetCluster() interface{}`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeleteDocumentRequest) GetClusterOk() (*interface{}, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeleteDocumentRequest) SetCluster(v interface{})`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DeleteDocumentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *DeleteDocumentRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *DeleteDocumentRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetId

`func (o *DeleteDocumentRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteDocumentRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteDocumentRequest) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *DeleteDocumentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeleteDocumentRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeleteDocumentRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuery

`func (o *DeleteDocumentRequest) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DeleteDocumentRequest) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DeleteDocumentRequest) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DeleteDocumentRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *DeleteDocumentRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *DeleteDocumentRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


