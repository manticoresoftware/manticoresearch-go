# MatchOpFilter

Query match expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** |  | 
**QueryFields** | **string** |  | 
**Operator** | **string** |  | 

## Methods

### NewMatchOpFilter

`func NewMatchOpFilter(queryString string, queryFields string, operator string, ) *MatchOpFilter`

NewMatchOpFilter instantiates a new MatchOpFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchOpFilterWithDefaults

`func NewMatchOpFilterWithDefaults() *MatchOpFilter`

NewMatchOpFilterWithDefaults instantiates a new MatchOpFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *MatchOpFilter) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *MatchOpFilter) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *MatchOpFilter) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetQueryFields

`func (o *MatchOpFilter) GetQueryFields() string`

GetQueryFields returns the QueryFields field if non-nil, zero value otherwise.

### GetQueryFieldsOk

`func (o *MatchOpFilter) GetQueryFieldsOk() (*string, bool)`

GetQueryFieldsOk returns a tuple with the QueryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFields

`func (o *MatchOpFilter) SetQueryFields(v string)`

SetQueryFields sets QueryFields field to given value.


### GetOperator

`func (o *MatchOpFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MatchOpFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MatchOpFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


