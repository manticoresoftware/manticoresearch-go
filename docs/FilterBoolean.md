# FilterBoolean

Query filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterField** | **string** |  | 
**Operation** | **string** |  | 
**FilterValue** | **bool** |  | 

## Methods

### NewFilterBoolean

`func NewFilterBoolean(filterField string, operation string, filterValue bool, ) *FilterBoolean`

NewFilterBoolean instantiates a new FilterBoolean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterBooleanWithDefaults

`func NewFilterBooleanWithDefaults() *FilterBoolean`

NewFilterBooleanWithDefaults instantiates a new FilterBoolean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterField

`func (o *FilterBoolean) GetFilterField() string`

GetFilterField returns the FilterField field if non-nil, zero value otherwise.

### GetFilterFieldOk

`func (o *FilterBoolean) GetFilterFieldOk() (*string, bool)`

GetFilterFieldOk returns a tuple with the FilterField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterField

`func (o *FilterBoolean) SetFilterField(v string)`

SetFilterField sets FilterField field to given value.


### GetOperation

`func (o *FilterBoolean) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *FilterBoolean) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *FilterBoolean) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetFilterValue

`func (o *FilterBoolean) GetFilterValue() bool`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *FilterBoolean) GetFilterValueOk() (*bool, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *FilterBoolean) SetFilterValue(v bool)`

SetFilterValue sets FilterValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


