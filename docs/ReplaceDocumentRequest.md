# ReplaceDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doc** | **interface{}** | Object containing the new document data to replace the existing one. | 

## Methods

### NewReplaceDocumentRequest

`func NewReplaceDocumentRequest(doc interface{}, ) *ReplaceDocumentRequest`

NewReplaceDocumentRequest instantiates a new ReplaceDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceDocumentRequestWithDefaults

`func NewReplaceDocumentRequestWithDefaults() *ReplaceDocumentRequest`

NewReplaceDocumentRequestWithDefaults instantiates a new ReplaceDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoc

`func (o *ReplaceDocumentRequest) GetDoc() interface{}`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *ReplaceDocumentRequest) GetDocOk() (*interface{}, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *ReplaceDocumentRequest) SetDoc(v interface{})`

SetDoc sets Doc field to given value.


### SetDocNil

`func (o *ReplaceDocumentRequest) SetDocNil(b bool)`

 SetDocNil sets the value for Doc to be an explicit nil

### UnsetDoc
`func (o *ReplaceDocumentRequest) UnsetDoc()`

UnsetDoc ensures that no value is present for Doc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


