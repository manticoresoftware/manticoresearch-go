# SourceByRules

Query fields to be included/excluded to/from response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Includes** | **[]string** |  | [default to []]
**Excludes** | **[]string** |  | [default to [""]]

## Methods

### NewSourceByRules

`func NewSourceByRules(includes []string, excludes []string, ) *SourceByRules`

NewSourceByRules instantiates a new SourceByRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceByRulesWithDefaults

`func NewSourceByRulesWithDefaults() *SourceByRules`

NewSourceByRulesWithDefaults instantiates a new SourceByRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludes

`func (o *SourceByRules) GetIncludes() []string`

GetIncludes returns the Includes field if non-nil, zero value otherwise.

### GetIncludesOk

`func (o *SourceByRules) GetIncludesOk() (*[]string, bool)`

GetIncludesOk returns a tuple with the Includes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludes

`func (o *SourceByRules) SetIncludes(v []string)`

SetIncludes sets Includes field to given value.


### GetExcludes

`func (o *SourceByRules) GetExcludes() []string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *SourceByRules) GetExcludesOk() (*[]string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *SourceByRules) SetExcludes(v []string)`

SetExcludes sets Excludes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


