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

// checks if the JoinOn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoinOn{}

// JoinOn struct for JoinOn
type JoinOn struct {
	Right *JoinCond
	Left *JoinCond
	Operator *string
}

// NewJoinOn instantiates a new JoinOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoinOn() *JoinOn {
	this := JoinOn{}
	return &this
}

// NewJoinOnWithDefaults instantiates a new JoinOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoinOnWithDefaults() *JoinOn {
	this := JoinOn{}
	return &this
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *JoinOn) GetRight() JoinCond {
	if o == nil || IsNil(o.Right) {
		var ret JoinCond
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinOn) GetRightOk() (*JoinCond, bool) {
	if o == nil || IsNil(o.Right) {
		return nil, false
	}
	return o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *JoinOn) HasRight() bool {
	if o != nil && !IsNil(o.Right) {
		return true
	}

	return false
}

// SetRight gets a reference to the given JoinCond and assigns it to the Right field.
func (o *JoinOn) SetRight(v JoinCond) {
	o.Right = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *JoinOn) GetLeft() JoinCond {
	if o == nil || IsNil(o.Left) {
		var ret JoinCond
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinOn) GetLeftOk() (*JoinCond, bool) {
	if o == nil || IsNil(o.Left) {
		return nil, false
	}
	return o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *JoinOn) HasLeft() bool {
	if o != nil && !IsNil(o.Left) {
		return true
	}

	return false
}

// SetLeft gets a reference to the given JoinCond and assigns it to the Left field.
func (o *JoinOn) SetLeft(v JoinCond) {
	o.Left = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *JoinOn) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoinOn) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *JoinOn) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *JoinOn) SetOperator(v string) {
	o.Operator = &v
}

func (o JoinOn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoinOn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Right) {
		toSerialize["right"] = o.Right
	}
	if !IsNil(o.Left) {
		toSerialize["left"] = o.Left
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	return toSerialize, nil
}

type NullableJoinOn struct {
	value *JoinOn
	isSet bool
}

func (v NullableJoinOn) Get() *JoinOn {
	return v.value
}

func (v *NullableJoinOn) Set(val *JoinOn) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinOn) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinOn(val *JoinOn) *NullableJoinOn {
	return &NullableJoinOn{value: val, isSet: true}
}

func (v NullableJoinOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

