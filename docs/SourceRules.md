# SourceRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Includes** | Pointer to **interface{}** | List of fields to include in the response | [optional] [default to []]
**Excludes** | Pointer to **interface{}** | List of fields to exclude from the response | [optional] [default to [""]]

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

`func (o *SourceRules) GetIncludes() interface{}`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *SourceRules) GetIncludesOk() (*interface{}, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *SourceRules) SetIncludes(v interface{})`

SetIncludes sets Includes field to given value.

### HasIncludes

`func (o *SourceRules) HasIncludes() bool`

HasIncludes returns a boolean if a field has been set.

### SetIncludesNil

`func (o *SourceRules) SetIncludesNil(b bool)`

 SetIncludesNil sets the value for Includes to be an explicit nil

### UnsetIncludes
`func (o *SourceRules) UnsetIncludes()`

UnsetIncludes ensures that no value is present for Includes, not even an explicit nil
### GetExcludes

`func (o *SourceRules) GetExcludes() interface{}`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *SourceRules) GetExcludesOk() (*interface{}, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *SourceRules) SetExcludes(v interface{})`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *SourceRules) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### SetExcludesNil

`func (o *SourceRules) SetExcludesNil(b bool)`

 SetExcludesNil sets the value for Excludes to be an explicit nil

### UnsetExcludes
`func (o *SourceRules) UnsetExcludes()`

UnsetExcludes ensures that no value is present for Excludes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


