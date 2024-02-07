# BoolFilter

Boolean attribute filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Should** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Must** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MustNot** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewBoolFilter

`func NewBoolFilter() *BoolFilter`

NewBoolFilter instantiates a new BoolFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoolFilterWithDefaults

`func NewBoolFilterWithDefaults() *BoolFilter`

NewBoolFilterWithDefaults instantiates a new BoolFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShould

`func (o *BoolFilter) GetShould() []map[string]interface{}`

GetShould returns the Should field if non-nil, zero value otherwise.

### GetShouldOk

`func (o *BoolFilter) GetShouldOk() (*[]map[string]interface{}, bool)`

GetShouldOk returns a tuple with the Should field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShould

`func (o *BoolFilter) SetShould(v []map[string]interface{})`

SetShould sets Should field to given value.

### HasShould

`func (o *BoolFilter) HasShould() bool`

HasShould returns a boolean if a field has been set.

### GetMust

`func (o *BoolFilter) GetMust() []map[string]interface{}`

GetMust returns the Must field if non-nil, zero value otherwise.

### GetMustOk

`func (o *BoolFilter) GetMustOk() (*[]map[string]interface{}, bool)`

GetMustOk returns a tuple with the Must field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMust

`func (o *BoolFilter) SetMust(v []map[string]interface{})`

SetMust sets Must field to given value.

### HasMust

`func (o *BoolFilter) HasMust() bool`

HasMust returns a boolean if a field has been set.

### GetMustNot

`func (o *BoolFilter) GetMustNot() []map[string]interface{}`

GetMustNot returns the MustNot field if non-nil, zero value otherwise.

### GetMustNotOk

`func (o *BoolFilter) GetMustNotOk() (*[]map[string]interface{}, bool)`

GetMustNotOk returns a tuple with the MustNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustNot

`func (o *BoolFilter) SetMustNot(v []map[string]interface{})`

SetMustNot sets MustNot field to given value.

### HasMustNot

`func (o *BoolFilter) HasMustNot() bool`

HasMustNot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


