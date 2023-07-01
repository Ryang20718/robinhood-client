# PaginatedInstrumentSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **string** |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]InstrumentSplit**](InstrumentSplit.md) |  | [optional] 

## Methods

### NewPaginatedInstrumentSplit

`func NewPaginatedInstrumentSplit() *PaginatedInstrumentSplit`

NewPaginatedInstrumentSplit instantiates a new PaginatedInstrumentSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedInstrumentSplitWithDefaults

`func NewPaginatedInstrumentSplitWithDefaults() *PaginatedInstrumentSplit`

NewPaginatedInstrumentSplitWithDefaults instantiates a new PaginatedInstrumentSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedInstrumentSplit) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedInstrumentSplit) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedInstrumentSplit) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedInstrumentSplit) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedInstrumentSplit) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedInstrumentSplit) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedInstrumentSplit) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedInstrumentSplit) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedInstrumentSplit) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedInstrumentSplit) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedInstrumentSplit) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedInstrumentSplit) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedInstrumentSplit) GetResults() []InstrumentSplit`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedInstrumentSplit) GetResultsOk() (*[]InstrumentSplit, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedInstrumentSplit) SetResults(v []InstrumentSplit)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedInstrumentSplit) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


