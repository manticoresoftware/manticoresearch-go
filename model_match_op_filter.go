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

// checks if the MatchOpFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchOpFilter{}

// MatchOpFilter Query match expression
type MatchOpFilter struct {
	QueryString string `json:"query_string"`
	QueryFields string `json:"query_fields"`
	Operator string `json:"operator"`
}

type _MatchOpFilter MatchOpFilter

// NewMatchOpFilter instantiates a new MatchOpFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchOpFilter(queryString string, queryFields string, operator string) *MatchOpFilter {
	this := MatchOpFilter{}
	this.QueryString = queryString
	this.QueryFields = queryFields
	this.Operator = operator
	return &this
}

// NewMatchOpFilterWithDefaults instantiates a new MatchOpFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchOpFilterWithDefaults() *MatchOpFilter {
	this := MatchOpFilter{}
	return &this
}

// GetQueryString returns the QueryString field value
func (o *MatchOpFilter) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *MatchOpFilter) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *MatchOpFilter) SetQueryString(v string) {
	o.QueryString = v
}

// GetQueryFields returns the QueryFields field value
func (o *MatchOpFilter) GetQueryFields() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryFields
}

// GetQueryFieldsOk returns a tuple with the QueryFields field value
// and a boolean to check if the value has been set.
func (o *MatchOpFilter) GetQueryFieldsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryFields, true
}

// SetQueryFields sets field value
func (o *MatchOpFilter) SetQueryFields(v string) {
	o.QueryFields = v
}

// GetOperator returns the Operator field value
func (o *MatchOpFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *MatchOpFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *MatchOpFilter) SetOperator(v string) {
	o.Operator = v
}

func (o MatchOpFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchOpFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query_string"] = o.QueryString
	toSerialize["query_fields"] = o.QueryFields
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *MatchOpFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query_string",
		"query_fields",
		"operator",
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

	varMatchOpFilter := _MatchOpFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatchOpFilter)

	if err != nil {
		return err
	}

	*o = MatchOpFilter(varMatchOpFilter)

	return err
}

type NullableMatchOpFilter struct {
	value *MatchOpFilter
	isSet bool
}

func (v NullableMatchOpFilter) Get() *MatchOpFilter {
	return v.value
}

func (v *NullableMatchOpFilter) Set(val *MatchOpFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchOpFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchOpFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchOpFilter(val *MatchOpFilter) *NullableMatchOpFilter {
	return &NullableMatchOpFilter{value: val, isSet: true}
}

func (v NullableMatchOpFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchOpFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

