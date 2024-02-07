# MatchPhraseFilter

Query match expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryPhrase** | **string** |  | 
**QueryFields** | **string** |  | 

## Methods

### NewMatchPhraseFilter

`func NewMatchPhraseFilter(queryPhrase string, queryFields string, ) *MatchPhraseFilter`

NewMatchPhraseFilter instantiates a new MatchPhraseFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchPhraseFilterWithDefaults

`func NewMatchPhraseFilterWithDefaults() *MatchPhraseFilter`

NewMatchPhraseFilterWithDefaults instantiates a new MatchPhraseFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryPhrase

`func (o *MatchPhraseFilter) GetQueryPhrase() string`

GetQueryPhrase returns the QueryPhrase field if non-nil, zero value otherwise.

### GetQueryPhraseOk

`func (o *MatchPhraseFilter) GetQueryPhraseOk() (*string, bool)`

GetQueryPhraseOk returns a tuple with the QueryPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPhrase

`func (o *MatchPhraseFilter) SetQueryPhrase(v string)`

SetQueryPhrase sets QueryPhrase field to given value.


### GetQueryFields

`func (o *MatchPhraseFilter) GetQueryFields() string`

GetQueryFields returns the QueryFields field if non-nil, zero value otherwise.

### GetQueryFieldsOk

`func (o *MatchPhraseFilter) GetQueryFieldsOk() (*string, bool)`

GetQueryFieldsOk returns a tuple with the QueryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFields

`func (o *MatchPhraseFilter) SetQueryFields(v string)`

SetQueryFields sets QueryFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


