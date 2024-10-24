/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 5.0.0
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SourceRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceRules{}

// SourceRules Defines which fields to include or exclude in the response for a search query
type SourceRules struct {
	// List of fields to include in the response
	Includes interface{}
	// List of fields to exclude from the response
	Excludes interface{}
	AdditionalProperties map[string]interface{}
}

type _SourceRules SourceRules

// NewSourceRules instantiates a new SourceRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceRules() *SourceRules {
	this := SourceRules{}
	return &this
}

// NewSourceRulesWithDefaults instantiates a new SourceRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceRulesWithDefaults() *SourceRules {
	this := SourceRules{}
	return &this
}

// GetIncludes returns the Includes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRules) GetIncludes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRules) GetIncludesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Includes) {
		return nil, false
	}
	return &o.Includes, true
}

// HasIncludes returns a boolean if a field has been set.
func (o *SourceRules) HasIncludes() bool {
	if o != nil && !IsNil(o.Includes) {
		return true
	}

	return false
}

// SetIncludes gets a reference to the given interface{} and assigns it to the Includes field.
func (o *SourceRules) SetIncludes(v interface{}) {
	o.Includes = v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SourceRules) GetExcludes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceRules) GetExcludesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Excludes) {
		return nil, false
	}
	return &o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *SourceRules) HasExcludes() bool {
	if o != nil && !IsNil(o.Excludes) {
		return true
	}

	return false
}

// SetExcludes gets a reference to the given interface{} and assigns it to the Excludes field.
func (o *SourceRules) SetExcludes(v interface{}) {
	o.Excludes = v
}

func (o SourceRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Includes != nil {
		toSerialize["includes"] = o.Includes
	}
	if o.Excludes != nil {
		toSerialize["excludes"] = o.Excludes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

type NullableSourceRules struct {
	value *SourceRules
	isSet bool
}

func (v NullableSourceRules) Get() *SourceRules {
	return v.value
}

func (v *NullableSourceRules) Set(val *SourceRules) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceRules) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceRules(val *SourceRules) *NullableSourceRules {
	return &NullableSourceRules{value: val, isSet: true}
}

func (v NullableSourceRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


