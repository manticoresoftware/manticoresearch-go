# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | User question to send to the chat model | 
**Table** | **string** | Vectorized table to retrieve context from | 
**ModelName** | **string** | Name of the chat model | 
**ConversationUuid** | Pointer to **string** | Existing conversation id to continue the dialog, or an empty string to start a new conversation. If omitted, a new id is generated.  | [optional] 
**VectorField** | Pointer to **string** | A specific vector field to search by. If omitted, Buddy uses the first &#x60;FLOAT_VECTOR&#x60; field from &#x60;SHOW CREATE TABLE&#x60;.  | [optional] 
**Fields** | Pointer to **string** | Legacy alias for &#x60;vector_field&#x60;. A request must not include both &#x60;vector_field&#x60; and &#x60;fields&#x60;.  | [optional] 

## Methods

### NewChat

`func NewChat(query string, table string, modelName string, ) *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *Chat) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Chat) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Chat) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTable

`func (o *Chat) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *Chat) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *Chat) SetTable(v string)`

SetTable sets Table field to given value.


### GetModelName

`func (o *Chat) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Chat) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Chat) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetConversationUuid

`func (o *Chat) GetConversationUuid() string`

GetConversationUuid returns the ConversationUuid field if non-nil, zero value otherwise.

### GetConversationUuidOk

`func (o *Chat) GetConversationUuidOk() (*string, bool)`

GetConversationUuidOk returns a tuple with the ConversationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationUuid

`func (o *Chat) SetConversationUuid(v string)`

SetConversationUuid sets ConversationUuid field to given value.

### HasConversationUuid

`func (o *Chat) HasConversationUuid() bool`

HasConversationUuid returns a boolean if a field has been set.

### GetVectorField

`func (o *Chat) GetVectorField() string`

GetVectorField returns the VectorField field if non-nil, zero value otherwise.

### GetVectorFieldOk

`func (o *Chat) GetVectorFieldOk() (*string, bool)`

GetVectorFieldOk returns a tuple with the VectorField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorField

`func (o *Chat) SetVectorField(v string)`

SetVectorField sets VectorField field to given value.

### HasVectorField

`func (o *Chat) HasVectorField() bool`

HasVectorField returns a boolean if a field has been set.

### GetFields

`func (o *Chat) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Chat) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Chat) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Chat) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


