# UpdateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **interface{}** | Name of the document table | 
**Cluster** | Pointer to **interface{}** | Name of the document cluster | [optional] 
**Doc** | **interface{}** | Object containing the document fields to update | 
**Id** | Pointer to **interface{}** | Document ID | [optional] 
**Query** | Pointer to [**UpdateDocumentRequestQuery**](UpdateDocumentRequestQuery.md) |  | [optional] 

## Methods

### NewUpdateDocumentRequest

`func NewUpdateDocumentRequest(table interface{}, doc interface{}, ) *UpdateDocumentRequest`

NewUpdateDocumentRequest instantiates a new UpdateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDocumentRequestWithDefaults

`func NewUpdateDocumentRequestWithDefaults() *UpdateDocumentRequest`

NewUpdateDocumentRequestWithDefaults instantiates a new UpdateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *UpdateDocumentRequest) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *UpdateDocumentRequest) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *UpdateDocumentRequest) SetTable(v interface{})`

SetTable sets Table field to given value.


### SetTableNil

`func (o *UpdateDocumentRequest) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *UpdateDocumentRequest) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetCluster

`func (o *UpdateDocumentRequest) GetCluster() interface{}`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *UpdateDocumentRequest) GetClusterOk() (*interface{}, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *UpdateDocumentRequest) SetCluster(v interface{})`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *UpdateDocumentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *UpdateDocumentRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *UpdateDocumentRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDoc

`func (o *UpdateDocumentRequest) GetDoc() interface{}`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *UpdateDocumentRequest) GetDocOk() (*interface{}, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *UpdateDocumentRequest) SetDoc(v interface{})`

SetDoc sets Doc field to given value.


### SetDocNil

`func (o *UpdateDocumentRequest) SetDocNil(b bool)`

 SetDocNil sets the value for Doc to be an explicit nil

### UnsetDoc
`func (o *UpdateDocumentRequest) UnsetDoc()`

UnsetDoc ensures that no value is present for Doc, not even an explicit nil
### GetId

`func (o *UpdateDocumentRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDocumentRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDocumentRequest) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *UpdateDocumentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdateDocumentRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateDocumentRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuery

`func (o *UpdateDocumentRequest) GetQuery() UpdateDocumentRequestQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UpdateDocumentRequest) GetQueryOk() (*UpdateDocumentRequestQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UpdateDocumentRequest) SetQuery(v UpdateDocumentRequestQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *UpdateDocumentRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


