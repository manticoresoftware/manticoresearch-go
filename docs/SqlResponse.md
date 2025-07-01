# SqlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | **map[string]interface{}** |  | 
**Took** | Pointer to **float32** |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 

## Methods

### NewSqlResponse

`func NewSqlResponse(hits map[string]interface{}, ) *SqlResponse`

NewSqlResponse instantiates a new SqlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlResponseWithDefaults

`func NewSqlResponseWithDefaults() *SqlResponse`

NewSqlResponseWithDefaults instantiates a new SqlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *SqlResponse) GetHits() map[string]interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SqlResponse) GetHitsOk() (*map[string]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SqlResponse) SetHits(v map[string]interface{})`

SetHits sets Hits field to given value.


### GetTook

`func (o *SqlResponse) GetTook() float32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SqlResponse) GetTookOk() (*float32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SqlResponse) SetTook(v float32)`

SetTook sets Took field to given value.

### HasTook

`func (o *SqlResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetTimedOut

`func (o *SqlResponse) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SqlResponse) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SqlResponse) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SqlResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


