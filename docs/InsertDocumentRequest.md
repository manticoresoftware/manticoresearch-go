# InsertDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **string** | Name of the table to insert the document into | 
**Cluster** | Pointer to **string** | Name of the cluster to insert the document into | [optional] 
**Id** | Pointer to **int64** | Document ID. If not provided, an ID will be auto-generated  | [optional] 
**Doc** | **map[string]interface{}** | Object containing document data  | 

## Methods

### NewInsertDocumentRequest

`func NewInsertDocumentRequest(table string, doc map[string]interface{}, ) *InsertDocumentRequest`

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

`func (o *InsertDocumentRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *InsertDocumentRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *InsertDocumentRequest) SetTable(v string)`

SetTable sets Table field to given value.


### GetCluster

`func (o *InsertDocumentRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InsertDocumentRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InsertDocumentRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *InsertDocumentRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetId

`func (o *InsertDocumentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InsertDocumentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InsertDocumentRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InsertDocumentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDoc

`func (o *InsertDocumentRequest) GetDoc() map[string]interface{}`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *InsertDocumentRequest) GetDocOk() (*map[string]interface{}, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *InsertDocumentRequest) SetDoc(v map[string]interface{})`

SetDoc sets Doc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


