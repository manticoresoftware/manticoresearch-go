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

// checks if the ErrorResponseString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseString{}

// ErrorResponseString Error response
type ErrorResponseString struct {
	Error string `json:"error"`
}

type _ErrorResponseString ErrorResponseString

// NewErrorResponseString instantiates a new ErrorResponseString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseString(error_ string) *ErrorResponseString {
	this := ErrorResponseString{}
	this.Error = error_
	return &this
}

// NewErrorResponseStringWithDefaults instantiates a new ErrorResponseString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseStringWithDefaults() *ErrorResponseString {
	this := ErrorResponseString{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorResponseString) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseString) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorResponseString) SetError(v string) {
	o.Error = v
}

func (o ErrorResponseString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *ErrorResponseString) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
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

	varErrorResponseString := _ErrorResponseString{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponseString)

	if err != nil {
		return err
	}

	*o = ErrorResponseString(varErrorResponseString)

	return err
}

type NullableErrorResponseString struct {
	value *ErrorResponseString
	isSet bool
}

func (v NullableErrorResponseString) Get() *ErrorResponseString {
	return v.value
}

func (v *NullableErrorResponseString) Set(val *ErrorResponseString) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseString) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseString(val *ErrorResponseString) *NullableErrorResponseString {
	return &NullableErrorResponseString{value: val, isSet: true}
}

func (v NullableErrorResponseString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


