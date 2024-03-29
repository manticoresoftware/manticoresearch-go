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

// checks if the SortMVA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortMVA{}

// SortMVA Query sort expression for MVA attributes
type SortMVA struct {
	Attr string `json:"attr"`
	Order string `json:"order"`
	Mode string `json:"mode"`
}

type _SortMVA SortMVA

// NewSortMVA instantiates a new SortMVA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortMVA(attr string, order string, mode string) *SortMVA {
	this := SortMVA{}
	this.Attr = attr
	this.Order = order
	this.Mode = mode
	return &this
}

// NewSortMVAWithDefaults instantiates a new SortMVA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortMVAWithDefaults() *SortMVA {
	this := SortMVA{}
	return &this
}

// GetAttr returns the Attr field value
func (o *SortMVA) GetAttr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attr
}

// GetAttrOk returns a tuple with the Attr field value
// and a boolean to check if the value has been set.
func (o *SortMVA) GetAttrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attr, true
}

// SetAttr sets field value
func (o *SortMVA) SetAttr(v string) {
	o.Attr = v
}

// GetOrder returns the Order field value
func (o *SortMVA) GetOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *SortMVA) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *SortMVA) SetOrder(v string) {
	o.Order = v
}

// GetMode returns the Mode field value
func (o *SortMVA) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *SortMVA) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *SortMVA) SetMode(v string) {
	o.Mode = v
}

func (o SortMVA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortMVA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attr"] = o.Attr
	toSerialize["order"] = o.Order
	toSerialize["mode"] = o.Mode
	return toSerialize, nil
}

func (o *SortMVA) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attr",
		"order",
		"mode",
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

	varSortMVA := _SortMVA{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSortMVA)

	if err != nil {
		return err
	}

	*o = SortMVA(varSortMVA)

	return err
}

type NullableSortMVA struct {
	value *SortMVA
	isSet bool
}

func (v NullableSortMVA) Get() *SortMVA {
	return v.value
}

func (v *NullableSortMVA) Set(val *SortMVA) {
	v.value = val
	v.isSet = true
}

func (v NullableSortMVA) IsSet() bool {
	return v.isSet
}

func (v *NullableSortMVA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortMVA(val *SortMVA) *NullableSortMVA {
	return &NullableSortMVA{value: val, isSet: true}
}

func (v NullableSortMVA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortMVA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


