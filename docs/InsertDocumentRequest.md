# InsertDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **interface{}** | Name of the table to insert the document into | 
**Cluster** | Pointer to **interface{}** | Name of the cluster to insert the document into | [optional] 
**Id** | Pointer to **interface{}** | Document ID. If not provided, an ID will be auto-generated  | [optional] 
**Doc** | **interface{}** | Object containing document data  | 

## Methods

### NewInsertDocumentRequest

`func NewInsertDocumentRequest(table interface{}, doc interface{}, ) *InsertDocumentRequest`

NewInsertDocumentRequest instantiates a new InsertDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertDocumentRequestWithDefaults

`func NewInsertDocumentRequestWithDefaults() *InsertDocumentRequest`

NewInsertDocumentRequestWithDefaults instantiates a new InsertDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *InsertDocumentRequest) GetTable() interface{}`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *InsertDocumentRequest) GetTableOk() (*interface{}, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *InsertDocumentRequest) SetTable(v interface{})`

SetTable sets Table field to given value.


### SetTableNil

`func (o *InsertDocumentRequest) SetTableNil(b bool)`

 SetTableNil sets the value for Table to be an explicit nil

### UnsetTable
`func (o *InsertDocumentRequest) UnsetTable()`

UnsetTable ensures that no value is present for Table, not even an explicit nil
### GetCluster

`func (o *InsertDocumentRequest) GetCluster() interface{}`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InsertDocumentRequest) GetClusterOk() (*interface{}, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InsertDocumentRequest) SetCluster(v interface{})`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *InsertDocumentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *InsertDocumentRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *InsertDocumentRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetId

`func (o *InsertDocumentRequest) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsertDocumentRequest) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsertDocumentRequest) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *InsertDocumentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InsertDocumentRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InsertDocumentRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDoc

`func (o *InsertDocumentRequest) GetDoc() interface{}`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *InsertDocumentRequest) GetDocOk() (*interface{}, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *InsertDocumentRequest) SetDoc(v interface{})`

SetDoc sets Doc field to given value.


### SetDocNil

`func (o *InsertDocumentRequest) SetDocNil(b bool)`

 SetDocNil sets the value for Doc to be an explicit nil

### UnsetDoc
`func (o *InsertDocumentRequest) UnsetDoc()`

UnsetDoc ensures that no value is present for Doc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


