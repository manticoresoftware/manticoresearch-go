# SqlObjResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | **interface{}** |  | 
**Took** | Pointer to **interface{}** |  | [optional] 
**TimedOut** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSqlObjResponse

`func NewSqlObjResponse(hits interface{}, ) *SqlObjResponse`

NewSqlObjResponse instantiates a new SqlObjResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlObjResponseWithDefaults

`func NewSqlObjResponseWithDefaults() *SqlObjResponse`

NewSqlObjResponseWithDefaults instantiates a new SqlObjResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *SqlObjResponse) GetHits() interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SqlObjResponse) GetHitsOk() (*interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SqlObjResponse) SetHits(v interface{})`

SetHits sets Hits field to given value.


### SetHitsNil

`func (o *SqlObjResponse) SetHitsNil(b bool)`

 SetHitsNil sets the value for Hits to be an explicit nil

### UnsetHits
`func (o *SqlObjResponse) UnsetHits()`

UnsetHits ensures that no value is present for Hits, not even an explicit nil
### GetTook

`func (o *SqlObjResponse) GetTook() interface{}`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *SqlObjResponse) GetTookOk() (*interface{}, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *SqlObjResponse) SetTook(v interface{})`

SetTook sets Took field to given value.

### HasTook

`func (o *SqlObjResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.

### SetTookNil

`func (o *SqlObjResponse) SetTookNil(b bool)`

 SetTookNil sets the value for Took to be an explicit nil

### UnsetTook
`func (o *SqlObjResponse) UnsetTook()`

UnsetTook ensures that no value is present for Took, not even an explicit nil
### GetTimedOut

`func (o *SqlObjResponse) GetTimedOut() interface{}`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *SqlObjResponse) GetTimedOutOk() (*interface{}, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *SqlObjResponse) SetTimedOut(v interface{})`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *SqlObjResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### SetTimedOutNil

`func (o *SqlObjResponse) SetTimedOutNil(b bool)`

 SetTimedOutNil sets the value for TimedOut to be an explicit nil

### UnsetTimedOut
`func (o *SqlObjResponse) UnsetTimedOut()`

UnsetTimedOut ensures that no value is present for TimedOut, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


