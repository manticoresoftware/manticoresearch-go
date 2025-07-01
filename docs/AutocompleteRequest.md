# AutocompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | **string** | The table to perform the search on | 
**Query** | **string** | The beginning of the string to autocomplete | 
**Options** | Pointer to **map[string]interface{}** | Autocomplete options   - layouts: A comma-separated string of keyboard layout codes to validate and check for spell correction. Available options - us, ru, ua, se, pt, no, it, gr, uk, fr, es, dk, de, ch, br, bg, be. By default, all are enabled.   - fuzziness: (0,1 or 2) Maximum Levenshtein distance for finding typos. Set to 0 to disable fuzzy matching. Default is 2   - prepend: true/false If true, adds an asterisk before the last word for prefix expansion (e.g., *word )   - append:  true/false If true, adds an asterisk after the last word for prefix expansion (e.g., word* )   - expansion_len: Number of characters to expand in the last word. Default is 10.  | [optional] 

## Methods

### NewAutocompleteRequest

`func NewAutocompleteRequest(table string, query string, ) *AutocompleteRequest`

NewAutocompleteRequest instantiates a new AutocompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteRequestWithDefaults

`func NewAutocompleteRequestWithDefaults() *AutocompleteRequest`

NewAutocompleteRequestWithDefaults instantiates a new AutocompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *AutocompleteRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *AutocompleteRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *AutocompleteRequest) SetTable(v string)`

SetTable sets Table field to given value.


### GetQuery

`func (o *AutocompleteRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AutocompleteRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AutocompleteRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetOptions

`func (o *AutocompleteRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AutocompleteRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AutocompleteRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AutocompleteRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


