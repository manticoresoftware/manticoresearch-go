# MatchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** |  | [default to ""]
**QueryFields** | **string** |  | [default to "*"]

## Methods

### NewMatchFilter

`func NewMatchFilter(queryString string, queryFields string, ) *MatchFilter`

NewMatchFilter instantiates a new MatchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchFilterWithDefaults

`func NewMatchFilterWithDefaults() *MatchFilter`

NewMatchFilterWithDefaults instantiates a new MatchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *MatchFilter) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *MatchFilter) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *MatchFilter) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetQueryFields

`func (o *MatchFilter) GetQueryFields() string`

GetQueryFields returns the QueryFields field if non-nil, zero value otherwise.

### GetQueryFieldsOk

`func (o *MatchFilter) GetQueryFieldsOk() (*string, bool)`

GetQueryFieldsOk returns a tuple with the QueryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFields

`func (o *MatchFilter) SetQueryFields(v string)`

SetQueryFields sets QueryFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


