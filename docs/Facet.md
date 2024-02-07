# Facet

Query FACET expression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attr** | **string** | The name of an attribute to facet | 
**Alias** | Pointer to **string** | Facet alias | [optional] 
**Limit** | Pointer to **int32** | The number of facet values to return | [optional] 

## Methods

### NewFacet

`func NewFacet(attr string, ) *Facet`

NewFacet instantiates a new Facet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetWithDefaults

`func NewFacetWithDefaults() *Facet`

NewFacetWithDefaults instantiates a new Facet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttr

`func (o *Facet) GetAttr() string`

GetAttr returns the Attr field if non-nil, zero value otherwise.

### GetAttrOk

`func (o *Facet) GetAttrOk() (*string, bool)`

GetAttrOk returns a tuple with the Attr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttr

`func (o *Facet) SetAttr(v string)`

SetAttr sets Attr field to given value.


### GetAlias

`func (o *Facet) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Facet) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Facet) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Facet) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetLimit

`func (o *Facet) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Facet) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Facet) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Facet) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


