# BoolFilterBool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Must** | Pointer to [**[]QueryFilter**](QueryFilter.md) |  | [optional] 
**MustNot** | Pointer to [**[]QueryFilter**](QueryFilter.md) |  | [optional] 
**Should** | Pointer to [**[]QueryFilter**](QueryFilter.md) |  | [optional] 

## Methods

### NewBoolFilterBool

`func NewBoolFilterBool() *BoolFilterBool`

NewBoolFilterBool instantiates a new BoolFilterBool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoolFilterBoolWithDefaults

`func NewBoolFilterBoolWithDefaults() *BoolFilterBool`

NewBoolFilterBoolWithDefaults instantiates a new BoolFilterBool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMust

`func (o *BoolFilterBool) GetMust() []QueryFilter`

GetMust returns the Must field if non-nil, zero value otherwise.

### GetMustOk

`func (o *BoolFilterBool) GetMustOk() (*[]QueryFilter, bool)`

GetMustOk returns a tuple with the Must field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMust

`func (o *BoolFilterBool) SetMust(v []QueryFilter)`

SetMust sets Must field to given value.

### HasMust

`func (o *BoolFilterBool) HasMust() bool`

HasMust returns a boolean if a field has been set.

### GetMustNot

`func (o *BoolFilterBool) GetMustNot() []QueryFilter`

GetMustNot returns the MustNot field if non-nil, zero value otherwise.

### GetMustNotOk

`func (o *BoolFilterBool) GetMustNotOk() (*[]QueryFilter, bool)`

GetMustNotOk returns a tuple with the MustNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustNot

`func (o *BoolFilterBool) SetMustNot(v []QueryFilter)`

SetMustNot sets MustNot field to given value.

### HasMustNot

`func (o *BoolFilterBool) HasMustNot() bool`

HasMustNot returns a boolean if a field has been set.

### GetShould

`func (o *BoolFilterBool) GetShould() []QueryFilter`

GetShould returns the Should field if non-nil, zero value otherwise.

### GetShouldOk

`func (o *BoolFilterBool) GetShouldOk() (*[]QueryFilter, bool)`

GetShouldOk returns a tuple with the Should field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShould

`func (o *BoolFilterBool) SetShould(v []QueryFilter)`

SetShould sets Should field to given value.

### HasShould

`func (o *BoolFilterBool) HasShould() bool`

HasShould returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


