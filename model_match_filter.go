/*
Manticore Search Client

Сlient for Manticore Search. 

API version: 3.3.1
Contact: info@manticoresearch.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MatchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchFilter{}

// MatchFilter Query match filter
type MatchFilter struct {
	QueryString string `json:"query_string"`
	QueryFields string `json:"query_fields"`
}

type _MatchFilter MatchFilter

// NewMatchFilter instantiates a new MatchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchFilter(queryString string, queryFields string) *MatchFilter {
	this := MatchFilter{}
	this.QueryString = queryString
	this.QueryFields = queryFields
	return &this
}

// NewMatchFilterWithDefaults instantiates a new MatchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchFilterWithDefaults() *MatchFilter {
	this := MatchFilter{}
	var queryString string = ""
	this.QueryString = queryString
	var queryFields string = "*"
	this.QueryFields = queryFields
	return &this
}

// GetQueryString returns the QueryString field value
func (o *MatchFilter) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *MatchFilter) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *MatchFilter) SetQueryString(v string) {
	o.QueryString = v
}

// GetQueryFields returns the QueryFields field value
func (o *MatchFilter) GetQueryFields() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryFields
}

// GetQueryFieldsOk returns a tuple with the QueryFields field value
// and a boolean to check if the value has been set.
func (o *MatchFilter) GetQueryFieldsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryFields, true
}

// SetQueryFields sets field value
func (o *MatchFilter) SetQueryFields(v string) {
	o.QueryFields = v
}

func (o MatchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query_string"] = o.QueryString
	toSerialize["query_fields"] = o.QueryFields
	return toSerialize, nil
}

func (o *MatchFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query_string",
		"query_fields",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMatchFilter := _MatchFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatchFilter)

	if err != nil {
		return err
	}

	*o = MatchFilter(varMatchFilter)

	return err
}

type NullableMatchFilter struct {
	value *MatchFilter
	isSet bool
}

func (v NullableMatchFilter) Get() *MatchFilter {
	return v.value
}

func (v *NullableMatchFilter) Set(val *MatchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchFilter(val *MatchFilter) *NullableMatchFilter {
	return &NullableMatchFilter{value: val, isSet: true}
}

func (v NullableMatchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

