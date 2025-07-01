# BulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **interface{}** | List of results | [optional] 
**Errors** | Pointer to **interface{}** | Errors occurred during the bulk operation | [optional] 
**Error** | Pointer to **interface{}** | Error message describing an error if such occurred | [optional] 
**CurrentLine** | Pointer to **interface{}** | Number of the row returned in the response | [optional] 
**SkippedLines** | Pointer to **interface{}** | Number of rows skipped in the response | [optional] 

## Methods

### NewBulkResponse

`func NewBulkResponse() *BulkResponse`

NewBulkResponse instantiates a new BulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWithDefaults

`func NewBulkResponseWithDefaults() *BulkResponse`

NewBulkResponseWithDefaults instantiates a new BulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BulkResponse) GetItems() interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkResponse) GetItemsOk() (*interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkResponse) SetItems(v interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *BulkResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BulkResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BulkResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetErrors

`func (o *BulkResponse) GetErrors() interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkResponse) GetErrorsOk() (*interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkResponse) SetErrors(v interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *BulkResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *BulkResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetError

`func (o *BulkResponse) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BulkResponse) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BulkResponse) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *BulkResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BulkResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BulkResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetCurrentLine

`func (o *BulkResponse) GetCurrentLine() interface{}`

GetCurrentLine returns the CurrentLine field if non-nil, zero value otherwise.

### GetCurrentLineOk

`func (o *BulkResponse) GetCurrentLineOk() (*interface{}, bool)`

GetCurrentLineOk returns a tuple with the CurrentLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLine

`func (o *BulkResponse) SetCurrentLine(v interface{})`

SetCurrentLine sets CurrentLine field to given value.

### HasCurrentLine

`func (o *BulkResponse) HasCurrentLine() bool`

HasCurrentLine returns a boolean if a field has been set.

### SetCurrentLineNil

`func (o *BulkResponse) SetCurrentLineNil(b bool)`

 SetCurrentLineNil sets the value for CurrentLine to be an explicit nil

### UnsetCurrentLine
`func (o *BulkResponse) UnsetCurrentLine()`

UnsetCurrentLine ensures that no value is present for CurrentLine, not even an explicit nil
### GetSkippedLines

`func (o *BulkResponse) GetSkippedLines() interface{}`

GetSkippedLines returns the SkippedLines field if non-nil, zero value otherwise.

### GetSkippedLinesOk

`func (o *BulkResponse) GetSkippedLinesOk() (*interface{}, bool)`

GetSkippedLinesOk returns a tuple with the SkippedLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedLines

`func (o *BulkResponse) SetSkippedLines(v interface{})`

SetSkippedLines sets SkippedLines field to given value.

### HasSkippedLines

`func (o *BulkResponse) HasSkippedLines() bool`

HasSkippedLines returns a boolean if a field has been set.

### SetSkippedLinesNil

`func (o *BulkResponse) SetSkippedLinesNil(b bool)`

 SetSkippedLinesNil sets the value for SkippedLines to be an explicit nil

### UnsetSkippedLines
`func (o *BulkResponse) UnsetSkippedLines()`

UnsetSkippedLines ensures that no value is present for SkippedLines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


