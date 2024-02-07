# InFilter

In attribute filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Values** | **[]map[string]interface{}** |  | 

## Methods

### NewInFilter

`func NewInFilter(field string, values []map[string]interface{}, ) *InFilter`

NewInFilter instantiates a new InFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInFilterWithDefaults

`func NewInFilterWithDefaults() *InFilter`

NewInFilterWithDefaults instantiates a new InFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *InFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *InFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *InFilter) SetField(v string)`

SetField sets Field field to given value.


### GetValues

`func (o *InFilter) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InFilter) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InFilter) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


