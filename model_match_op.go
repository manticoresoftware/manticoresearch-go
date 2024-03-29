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

// checks if the MatchOp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchOp{}

// MatchOp Query match expression with logical operator
type MatchOp struct {
	QueryInfo map[string]interface{} `json:"query_info"`
}

type _MatchOp MatchOp

// NewMatchOp instantiates a new MatchOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchOp(queryInfo map[string]interface{}) *MatchOp {
	this := MatchOp{}
	this.QueryInfo = queryInfo
	return &this
}

// NewMatchOpWithDefaults instantiates a new MatchOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchOpWithDefaults() *MatchOp {
	this := MatchOp{}
	return &this
}

// GetQueryInfo returns the QueryInfo field value
func (o *MatchOp) GetQueryInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.QueryInfo
}

// GetQueryInfoOk returns a tuple with the QueryInfo field value
// and a boolean to check if the value has been set.
func (o *MatchOp) GetQueryInfoOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.QueryInfo, true
}

// SetQueryInfo sets field value
func (o *MatchOp) SetQueryInfo(v map[string]interface{}) {
	o.QueryInfo = v
}

func (o MatchOp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchOp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query_info"] = o.QueryInfo
	return toSerialize, nil
}

func (o *MatchOp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query_info",
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

	varMatchOp := _MatchOp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatchOp)

	if err != nil {
		return err
	}

	*o = MatchOp(varMatchOp)

	return err
}

type NullableMatchOp struct {
	value *MatchOp
	isSet bool
}

func (v NullableMatchOp) Get() *MatchOp {
	return v.value
}

func (v *NullableMatchOp) Set(val *MatchOp) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchOp) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchOp(val *MatchOp) *NullableMatchOp {
	return &NullableMatchOp{value: val, isSet: true}
}

func (v NullableMatchOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


