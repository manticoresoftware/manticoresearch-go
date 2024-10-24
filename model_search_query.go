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

// checks if the SearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchQuery{}

// SearchQuery Defines a query structure for performing search operations
type SearchQuery struct {
	// Filter object defining a query string
	QueryString interface{}
	// Filter object defining a match keyword passed as a string or in a Match object
	Match interface{}
	// Filter object defining a match phrase
	MatchPhrase interface{}
	// Filter object to select all documents
	MatchAll interface{}
	Bool *BoolFilter
	Equals interface{}
	// Filter to match a given set of attribute values.
	In interface{}
	// Filter to match a given range of attribute values passed in Range objects
	Range interface{}
	GeoDistance *GeoDistance
	Highlight *Highlight
}

// NewSearchQuery instantiates a new SearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchQuery() *SearchQuery {
	this := SearchQuery{}
	return &this
}

// NewSearchQueryWithDefaults instantiates a new SearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchQueryWithDefaults() *SearchQuery {
	this := SearchQuery{}
	return &this
}

// GetQueryString returns the QueryString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetQueryString() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetQueryStringOk() (*interface{}, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return &o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *SearchQuery) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given interface{} and assigns it to the QueryString field.
func (o *SearchQuery) SetQueryString(v interface{}) {
	o.QueryString = v
}

// GetMatch returns the Match field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetMatch() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetMatchOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return &o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *SearchQuery) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given interface{} and assigns it to the Match field.
func (o *SearchQuery) SetMatch(v interface{}) {
	o.Match = v
}

// GetMatchPhrase returns the MatchPhrase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetMatchPhrase() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MatchPhrase
}

// GetMatchPhraseOk returns a tuple with the MatchPhrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetMatchPhraseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MatchPhrase) {
		return nil, false
	}
	return &o.MatchPhrase, true
}

// HasMatchPhrase returns a boolean if a field has been set.
func (o *SearchQuery) HasMatchPhrase() bool {
	if o != nil && !IsNil(o.MatchPhrase) {
		return true
	}

	return false
}

// SetMatchPhrase gets a reference to the given interface{} and assigns it to the MatchPhrase field.
func (o *SearchQuery) SetMatchPhrase(v interface{}) {
	o.MatchPhrase = v
}

// GetMatchAll returns the MatchAll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetMatchAll() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MatchAll
}

// GetMatchAllOk returns a tuple with the MatchAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetMatchAllOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MatchAll) {
		return nil, false
	}
	return &o.MatchAll, true
}

// HasMatchAll returns a boolean if a field has been set.
func (o *SearchQuery) HasMatchAll() bool {
	if o != nil && !IsNil(o.MatchAll) {
		return true
	}

	return false
}

// SetMatchAll gets a reference to the given interface{} and assigns it to the MatchAll field.
func (o *SearchQuery) SetMatchAll(v interface{}) {
	o.MatchAll = v
}

// GetBool returns the Bool field value if set, zero value otherwise.
func (o *SearchQuery) GetBool() BoolFilter {
	if o == nil || IsNil(o.Bool) {
		var ret BoolFilter
		return ret
	}
	return *o.Bool
}

// GetBoolOk returns a tuple with the Bool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetBoolOk() (*BoolFilter, bool) {
	if o == nil || IsNil(o.Bool) {
		return nil, false
	}
	return o.Bool, true
}

// HasBool returns a boolean if a field has been set.
func (o *SearchQuery) HasBool() bool {
	if o != nil && !IsNil(o.Bool) {
		return true
	}

	return false
}

// SetBool gets a reference to the given BoolFilter and assigns it to the Bool field.
func (o *SearchQuery) SetBool(v BoolFilter) {
	o.Bool = &v
}

// GetEquals returns the Equals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetEquals() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Equals
}

// GetEqualsOk returns a tuple with the Equals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetEqualsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Equals) {
		return nil, false
	}
	return &o.Equals, true
}

// HasEquals returns a boolean if a field has been set.
func (o *SearchQuery) HasEquals() bool {
	if o != nil && !IsNil(o.Equals) {
		return true
	}

	return false
}

// SetEquals gets a reference to the given interface{} and assigns it to the Equals field.
func (o *SearchQuery) SetEquals(v interface{}) {
	o.Equals = v
}

// GetIn returns the In field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetIn() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetInOk() (*interface{}, bool) {
	if o == nil || IsNil(o.In) {
		return nil, false
	}
	return &o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *SearchQuery) HasIn() bool {
	if o != nil && !IsNil(o.In) {
		return true
	}

	return false
}

// SetIn gets a reference to the given interface{} and assigns it to the In field.
func (o *SearchQuery) SetIn(v interface{}) {
	o.In = v
}

// GetRange returns the Range field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchQuery) GetRange() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchQuery) GetRangeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return &o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *SearchQuery) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given interface{} and assigns it to the Range field.
func (o *SearchQuery) SetRange(v interface{}) {
	o.Range = v
}

// GetGeoDistance returns the GeoDistance field value if set, zero value otherwise.
func (o *SearchQuery) GetGeoDistance() GeoDistance {
	if o == nil || IsNil(o.GeoDistance) {
		var ret GeoDistance
		return ret
	}
	return *o.GeoDistance
}

// GetGeoDistanceOk returns a tuple with the GeoDistance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetGeoDistanceOk() (*GeoDistance, bool) {
	if o == nil || IsNil(o.GeoDistance) {
		return nil, false
	}
	return o.GeoDistance, true
}

// HasGeoDistance returns a boolean if a field has been set.
func (o *SearchQuery) HasGeoDistance() bool {
	if o != nil && !IsNil(o.GeoDistance) {
		return true
	}

	return false
}

// SetGeoDistance gets a reference to the given GeoDistance and assigns it to the GeoDistance field.
func (o *SearchQuery) SetGeoDistance(v GeoDistance) {
	o.GeoDistance = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *SearchQuery) GetHighlight() Highlight {
	if o == nil || IsNil(o.Highlight) {
		var ret Highlight
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQuery) GetHighlightOk() (*Highlight, bool) {
	if o == nil || IsNil(o.Highlight) {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *SearchQuery) HasHighlight() bool {
	if o != nil && !IsNil(o.Highlight) {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given Highlight and assigns it to the Highlight field.
func (o *SearchQuery) SetHighlight(v Highlight) {
	o.Highlight = &v
}

func (o SearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.QueryString != nil {
		toSerialize["query_string"] = o.QueryString
	}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}
	if o.MatchPhrase != nil {
		toSerialize["match_phrase"] = o.MatchPhrase
	}
	if o.MatchAll != nil {
		toSerialize["match_all"] = o.MatchAll
	}
	if !IsNil(o.Bool) {
		toSerialize["bool"] = o.Bool
	}
	if o.Equals != nil {
		toSerialize["equals"] = o.Equals
	}
	if o.In != nil {
		toSerialize["in"] = o.In
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.GeoDistance) {
		toSerialize["geo_distance"] = o.GeoDistance
	}
	if !IsNil(o.Highlight) {
		toSerialize["highlight"] = o.Highlight
	}
	return toSerialize, nil
}

type NullableSearchQuery struct {
	value *SearchQuery
	isSet bool
}

func (v NullableSearchQuery) Get() *SearchQuery {
	return v.value
}

func (v *NullableSearchQuery) Set(val *SearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchQuery(val *SearchQuery) *NullableSearchQuery {
	return &NullableSearchQuery{value: val, isSet: true}
}

func (v NullableSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


