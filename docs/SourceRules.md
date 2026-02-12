# SourceRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Includes** | Pointer to **[]string** | List of fields to include in the response | [optional] [default to []]
**Excludes** | Pointer to **[]string** | List of fields to exclude from the response | [optional] [default to [""]]

## Methods

### NewSourceRules

`func NewSourceRules() *SourceRules`

NewSourceRules instantiates a new SourceRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRulesWithDefaults

`func NewSourceRulesWithDefaults() *SourceRules`

NewSourceRulesWithDefaults instantiates a new SourceRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludes

`func (o *SourceRules) GetIncludes() []string`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *SourceRules) GetIncludesOk() (*[]string, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *SourceRules) SetIncludes(v []string)`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *SourceRules) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### GetExcludes

`func (o *SourceRules) GetExcludes() []string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *SourceRules) GetExcludesOk() (*[]string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *SourceRules) SetExcludes(v []string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *SourceRules) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


