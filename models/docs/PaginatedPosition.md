# PaginatedPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **string** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]Position**](Position.md) |  | [optional] 

## Methods

### NewPaginatedPosition

`func NewPaginatedPosition() *PaginatedPosition`

NewPaginatedPosition instantiates a new PaginatedPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedPositionWithDefaults

`func NewPaginatedPositionWithDefaults() *PaginatedPosition`

NewPaginatedPositionWithDefaults instantiates a new PaginatedPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedPosition) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedPosition) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedPosition) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedPosition) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedPosition) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedPosition) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedPosition) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedPosition) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedPosition) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedPosition) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedPosition) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedPosition) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedPosition) GetResults() []Position`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedPosition) GetResultsOk() (*[]Position, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedPosition) SetResults(v []Position)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedPosition) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


