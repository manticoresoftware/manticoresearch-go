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
)

// checks if the DeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteResponse{}

// DeleteResponse Success response
type DeleteResponse struct {
	Index *string `json:"_index,omitempty"`
	Deleted *int32 `json:"deleted,omitempty"`
	Id *int64 `json:"_id,omitempty"`
	Result *string `json:"result,omitempty"`
}

// NewDeleteResponse instantiates a new DeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResponse() *DeleteResponse {
	this := DeleteResponse{}
	return &this
}

// NewDeleteResponseWithDefaults instantiates a new DeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResponseWithDefaults() *DeleteResponse {
	this := DeleteResponse{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *DeleteResponse) GetIndex() string {
	if o == nil || IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetIndexOk() (*string, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *DeleteResponse) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *DeleteResponse) SetIndex(v string) {
	o.Index = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DeleteResponse) GetDeleted() int32 {
	if o == nil || IsNil(o.Deleted) {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DeleteResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *DeleteResponse) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DeleteResponse) SetId(v int64) {
	o.Id = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DeleteResponse) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DeleteResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DeleteResponse) SetResult(v string) {
	o.Result = &v
}

func (o DeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["_index"] = o.Index
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableDeleteResponse struct {
	value *DeleteResponse
	isSet bool
}

func (v NullableDeleteResponse) Get() *DeleteResponse {
	return v.value
}

func (v *NullableDeleteResponse) Set(val *DeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResponse(val *DeleteResponse) *NullableDeleteResponse {
	return &NullableDeleteResponse{value: val, isSet: true}
}

func (v NullableDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

