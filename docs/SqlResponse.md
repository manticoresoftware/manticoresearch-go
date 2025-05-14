# SqlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | **interface{}** |  | 
**Took** | Pointer to **interface{}** |  | [optional] 
**TimedOut** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSqlResponse

`func NewSqlResponse(hits interface{}, ) *SqlResponse`

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

`func (o *SqlResponse) GetHits() interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SqlResponse) GetHitsOk() (*interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SqlResponse) SetHits(v interface{})`

SetHits sets Hits field to given value.


### SetHitsNil

`func (o *SqlResponse) SetHitsNil(b bool)`

 SetHitsNil sets the value for Hits to be an explicit nil

### UnsetHits
`func (o *SqlResponse) UnsetHits()`

UnsetHits ensures that no value is present for Hits, not even an explicit nil
### GetTook

`func (o *SqlResponse) GetTook() interface{}`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SqlResponse) GetTookOk() (*interface{}, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SqlResponse) SetTook(v interface{})`

SetTook sets Took field to given value.

### HasTook

`func (o *SqlResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### SetTookNil

`func (o *SqlResponse) SetTookNil(b bool)`

 SetTookNil sets the value for Took to be an explicit nil

### UnsetTook
`func (o *SqlResponse) UnsetTook()`

UnsetTook ensures that no value is present for Took, not even an explicit nil
### GetTimedOut

`func (o *SqlResponse) GetTimedOut() interface{}`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SqlResponse) GetTimedOutOk() (*interface{}, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SqlResponse) SetTimedOut(v interface{})`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SqlResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### SetTimedOutNil

`func (o *SqlResponse) SetTimedOutNil(b bool)`

 SetTimedOutNil sets the value for TimedOut to be an explicit nil

### UnsetTimedOut
`func (o *SqlResponse) UnsetTimedOut()`

UnsetTimedOut ensures that no value is present for TimedOut, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


