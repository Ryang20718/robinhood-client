/*
Robinhood API

Robinhood API Documentation

API version: 3.0.1
Contact: austin.millan@protonmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OptionOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionOrder{}

// OptionOrder struct for OptionOrder
type OptionOrder struct {
	CancelUrl *string `json:"cancel_url,omitempty"`
	CanceledQuantity *string `json:"canceled_quantity,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Direction *Direction `json:"direction,omitempty"`
	Id *string `json:"id,omitempty"`
	Legs []OptionOrderLeg `json:"legs,omitempty"`
	PendingQuantity *string `json:"pending_quantity,omitempty"`
	Premium *string `json:"premium,omitempty"`
	ProcessedPremium *string `json:"processed_premium,omitempty"`
	Price *string `json:"price,omitempty"`
	ProcessedQuantity *string `json:"processed_quantity,omitempty"`  // units of 1, 2, 3
	Quantity *string `json:"quantity,omitempty"`
	RefId *string `json:"ref_id,omitempty"`
	State *OrderState `json:"state,omitempty"`
	TimeInForce *TimeInForce `json:"time_in_force,omitempty"`
	Trigger *Trigger `json:"trigger,omitempty"`
	Type *ExecutionType `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	ChainId *string `json:"chain_id,omitempty"`
	ChainSymbol *string `json:"chain_symbol,omitempty"`
	ResponseCategory *string `json:"response_category,omitempty"`
	OpeningStrategy *OpenCloseStrategy `json:"opening_strategy,omitempty"`
	ClosingStrategy *OpenCloseStrategy `json:"closing_strategy,omitempty"`
	StopPrice *string `json:"stop_price,omitempty"`
	// manually added
	ExpirationDate *string  `json:"expiration_date,omitempty"`
	StrikePrice    *string `json:"strike_price,omitempty"`
	Expired        *string    `json:"expired,omitempty"`
	Assigned       *string    `json:"assigned,omitempty"`
	TransactionCode *string `json:"transaction_code,omitempty"` // STO, BTC, BTO, STC
}

// Manually Added
type OptionTransaction struct {
	Ticker          string
	TransactionType string
	Qty             float64
	StrikePrice     float64
	UnitCost        float64
	CreatedAt       string
	ExpirationDate  string
	Status          string // Assigned or Expired or Open
	Tag             string
}
// End

// NewOptionOrder instantiates a new OptionOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionOrder() *OptionOrder {
	this := OptionOrder{}
	return &this
}

// NewOptionOrderWithDefaults instantiates a new OptionOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionOrderWithDefaults() *OptionOrder {
	this := OptionOrder{}
	return &this
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *OptionOrder) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *OptionOrder) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *OptionOrder) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetCanceledQuantity returns the CanceledQuantity field value if set, zero value otherwise.
func (o *OptionOrder) GetCanceledQuantity() string {
	if o == nil || IsNil(o.CanceledQuantity) {
		var ret string
		return ret
	}
	return *o.CanceledQuantity
}

// GetCanceledQuantityOk returns a tuple with the CanceledQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetCanceledQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.CanceledQuantity) {
		return nil, false
	}
	return o.CanceledQuantity, true
}

// HasCanceledQuantity returns a boolean if a field has been set.
func (o *OptionOrder) HasCanceledQuantity() bool {
	if o != nil && !IsNil(o.CanceledQuantity) {
		return true
	}

	return false
}

// SetCanceledQuantity gets a reference to the given string and assigns it to the CanceledQuantity field.
func (o *OptionOrder) SetCanceledQuantity(v string) {
	o.CanceledQuantity = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OptionOrder) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OptionOrder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OptionOrder) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *OptionOrder) GetDirection() Direction {
	if o == nil || IsNil(o.Direction) {
		var ret Direction
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetDirectionOk() (*Direction, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *OptionOrder) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given Direction and assigns it to the Direction field.
func (o *OptionOrder) SetDirection(v Direction) {
	o.Direction = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptionOrder) SetId(v string) {
	o.Id = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *OptionOrder) GetLegs() []OptionOrderLeg {
	if o == nil || IsNil(o.Legs) {
		var ret []OptionOrderLeg
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetLegsOk() ([]OptionOrderLeg, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *OptionOrder) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []OptionOrderLeg and assigns it to the Legs field.
func (o *OptionOrder) SetLegs(v []OptionOrderLeg) {
	o.Legs = v
}

// GetPendingQuantity returns the PendingQuantity field value if set, zero value otherwise.
func (o *OptionOrder) GetPendingQuantity() string {
	if o == nil || IsNil(o.PendingQuantity) {
		var ret string
		return ret
	}
	return *o.PendingQuantity
}

// GetPendingQuantityOk returns a tuple with the PendingQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetPendingQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.PendingQuantity) {
		return nil, false
	}
	return o.PendingQuantity, true
}

// HasPendingQuantity returns a boolean if a field has been set.
func (o *OptionOrder) HasPendingQuantity() bool {
	if o != nil && !IsNil(o.PendingQuantity) {
		return true
	}

	return false
}

// SetPendingQuantity gets a reference to the given string and assigns it to the PendingQuantity field.
func (o *OptionOrder) SetPendingQuantity(v string) {
	o.PendingQuantity = &v
}

// GetPremium returns the Premium field value if set, zero value otherwise.
func (o *OptionOrder) GetPremium() string {
	if o == nil || IsNil(o.Premium) {
		var ret string
		return ret
	}
	return *o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetPremiumOk() (*string, bool) {
	if o == nil || IsNil(o.Premium) {
		return nil, false
	}
	return o.Premium, true
}

// HasPremium returns a boolean if a field has been set.
func (o *OptionOrder) HasPremium() bool {
	if o != nil && !IsNil(o.Premium) {
		return true
	}

	return false
}

// SetPremium gets a reference to the given string and assigns it to the Premium field.
func (o *OptionOrder) SetPremium(v string) {
	o.Premium = &v
}

// GetProcessedPremium returns the ProcessedPremium field value if set, zero value otherwise.
func (o *OptionOrder) GetProcessedPremium() string {
	if o == nil || IsNil(o.ProcessedPremium) {
		var ret string
		return ret
	}
	return *o.ProcessedPremium
}

// GetProcessedPremiumOk returns a tuple with the ProcessedPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetProcessedPremiumOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessedPremium) {
		return nil, false
	}
	return o.ProcessedPremium, true
}

// HasProcessedPremium returns a boolean if a field has been set.
func (o *OptionOrder) HasProcessedPremium() bool {
	if o != nil && !IsNil(o.ProcessedPremium) {
		return true
	}

	return false
}

// SetProcessedPremium gets a reference to the given string and assigns it to the ProcessedPremium field.
func (o *OptionOrder) SetProcessedPremium(v string) {
	o.ProcessedPremium = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OptionOrder) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OptionOrder) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OptionOrder) SetPrice(v string) {
	o.Price = &v
}

// GetProcessedQuantity returns the ProcessedQuantity field value if set, zero value otherwise.
func (o *OptionOrder) GetProcessedQuantity() string {
	if o == nil || IsNil(o.ProcessedQuantity) {
		var ret string
		return ret
	}
	return *o.ProcessedQuantity
}

// GetProcessedQuantityOk returns a tuple with the ProcessedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetProcessedQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessedQuantity) {
		return nil, false
	}
	return o.ProcessedQuantity, true
}

// HasProcessedQuantity returns a boolean if a field has been set.
func (o *OptionOrder) HasProcessedQuantity() bool {
	if o != nil && !IsNil(o.ProcessedQuantity) {
		return true
	}

	return false
}

// SetProcessedQuantity gets a reference to the given string and assigns it to the ProcessedQuantity field.
func (o *OptionOrder) SetProcessedQuantity(v string) {
	o.ProcessedQuantity = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OptionOrder) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OptionOrder) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *OptionOrder) SetQuantity(v string) {
	o.Quantity = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *OptionOrder) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *OptionOrder) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *OptionOrder) SetRefId(v string) {
	o.RefId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OptionOrder) GetState() OrderState {
	if o == nil || IsNil(o.State) {
		var ret OrderState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetStateOk() (*OrderState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OptionOrder) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given OrderState and assigns it to the State field.
func (o *OptionOrder) SetState(v OrderState) {
	o.State = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *OptionOrder) GetTimeInForce() TimeInForce {
	if o == nil || IsNil(o.TimeInForce) {
		var ret TimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetTimeInForceOk() (*TimeInForce, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *OptionOrder) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given TimeInForce and assigns it to the TimeInForce field.
func (o *OptionOrder) SetTimeInForce(v TimeInForce) {
	o.TimeInForce = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *OptionOrder) GetTrigger() Trigger {
	if o == nil || IsNil(o.Trigger) {
		var ret Trigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetTriggerOk() (*Trigger, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *OptionOrder) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given Trigger and assigns it to the Trigger field.
func (o *OptionOrder) SetTrigger(v Trigger) {
	o.Trigger = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OptionOrder) GetType() ExecutionType {
	if o == nil || IsNil(o.Type) {
		var ret ExecutionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetTypeOk() (*ExecutionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OptionOrder) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ExecutionType and assigns it to the Type field.
func (o *OptionOrder) SetType(v ExecutionType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OptionOrder) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OptionOrder) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OptionOrder) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *OptionOrder) GetChainId() string {
	if o == nil || IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetChainIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *OptionOrder) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *OptionOrder) SetChainId(v string) {
	o.ChainId = &v
}

// GetChainSymbol returns the ChainSymbol field value if set, zero value otherwise.
func (o *OptionOrder) GetChainSymbol() string {
	if o == nil || IsNil(o.ChainSymbol) {
		var ret string
		return ret
	}
	return *o.ChainSymbol
}

// GetChainSymbolOk returns a tuple with the ChainSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetChainSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.ChainSymbol) {
		return nil, false
	}
	return o.ChainSymbol, true
}

// HasChainSymbol returns a boolean if a field has been set.
func (o *OptionOrder) HasChainSymbol() bool {
	if o != nil && !IsNil(o.ChainSymbol) {
		return true
	}

	return false
}

// SetChainSymbol gets a reference to the given string and assigns it to the ChainSymbol field.
func (o *OptionOrder) SetChainSymbol(v string) {
	o.ChainSymbol = &v
}

// GetResponseCategory returns the ResponseCategory field value if set, zero value otherwise.
func (o *OptionOrder) GetResponseCategory() string {
	if o == nil || IsNil(o.ResponseCategory) {
		var ret string
		return ret
	}
	return *o.ResponseCategory
}

// GetResponseCategoryOk returns a tuple with the ResponseCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetResponseCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseCategory) {
		return nil, false
	}
	return o.ResponseCategory, true
}

// HasResponseCategory returns a boolean if a field has been set.
func (o *OptionOrder) HasResponseCategory() bool {
	if o != nil && !IsNil(o.ResponseCategory) {
		return true
	}

	return false
}

// SetResponseCategory gets a reference to the given string and assigns it to the ResponseCategory field.
func (o *OptionOrder) SetResponseCategory(v string) {
	o.ResponseCategory = &v
}

// GetOpeningStrategy returns the OpeningStrategy field value if set, zero value otherwise.
func (o *OptionOrder) GetOpeningStrategy() OpenCloseStrategy {
	if o == nil || IsNil(o.OpeningStrategy) {
		var ret OpenCloseStrategy
		return ret
	}
	return *o.OpeningStrategy
}

// GetOpeningStrategyOk returns a tuple with the OpeningStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetOpeningStrategyOk() (*OpenCloseStrategy, bool) {
	if o == nil || IsNil(o.OpeningStrategy) {
		return nil, false
	}
	return o.OpeningStrategy, true
}

// HasOpeningStrategy returns a boolean if a field has been set.
func (o *OptionOrder) HasOpeningStrategy() bool {
	if o != nil && !IsNil(o.OpeningStrategy) {
		return true
	}

	return false
}

// SetOpeningStrategy gets a reference to the given OpenCloseStrategy and assigns it to the OpeningStrategy field.
func (o *OptionOrder) SetOpeningStrategy(v OpenCloseStrategy) {
	o.OpeningStrategy = &v
}

// GetClosingStrategy returns the ClosingStrategy field value if set, zero value otherwise.
func (o *OptionOrder) GetClosingStrategy() OpenCloseStrategy {
	if o == nil || IsNil(o.ClosingStrategy) {
		var ret OpenCloseStrategy
		return ret
	}
	return *o.ClosingStrategy
}

// GetClosingStrategyOk returns a tuple with the ClosingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetClosingStrategyOk() (*OpenCloseStrategy, bool) {
	if o == nil || IsNil(o.ClosingStrategy) {
		return nil, false
	}
	return o.ClosingStrategy, true
}

// HasClosingStrategy returns a boolean if a field has been set.
func (o *OptionOrder) HasClosingStrategy() bool {
	if o != nil && !IsNil(o.ClosingStrategy) {
		return true
	}

	return false
}

// SetClosingStrategy gets a reference to the given OpenCloseStrategy and assigns it to the ClosingStrategy field.
func (o *OptionOrder) SetClosingStrategy(v OpenCloseStrategy) {
	o.ClosingStrategy = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *OptionOrder) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionOrder) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *OptionOrder) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *OptionOrder) SetStopPrice(v string) {
	o.StopPrice = &v
}

func (o OptionOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	if !IsNil(o.CanceledQuantity) {
		toSerialize["canceled_quantity"] = o.CanceledQuantity
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !IsNil(o.PendingQuantity) {
		toSerialize["pending_quantity"] = o.PendingQuantity
	}
	if !IsNil(o.Premium) {
		toSerialize["premium"] = o.Premium
	}
	if !IsNil(o.ProcessedPremium) {
		toSerialize["processed_premium"] = o.ProcessedPremium
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ProcessedQuantity) {
		toSerialize["processed_quantity"] = o.ProcessedQuantity
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["time_in_force"] = o.TimeInForce
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ChainId) {
		toSerialize["chain_id"] = o.ChainId
	}
	if !IsNil(o.ChainSymbol) {
		toSerialize["chain_symbol"] = o.ChainSymbol
	}
	if !IsNil(o.ResponseCategory) {
		toSerialize["response_category"] = o.ResponseCategory
	}
	if !IsNil(o.OpeningStrategy) {
		toSerialize["opening_strategy"] = o.OpeningStrategy
	}
	if !IsNil(o.ClosingStrategy) {
		toSerialize["closing_strategy"] = o.ClosingStrategy
	}
	if !IsNil(o.StopPrice) {
		toSerialize["stop_price"] = o.StopPrice
	}
	return toSerialize, nil
}

type NullableOptionOrder struct {
	value *OptionOrder
	isSet bool
}

func (v NullableOptionOrder) Get() *OptionOrder {
	return v.value
}

func (v *NullableOptionOrder) Set(val *OptionOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionOrder(val *OptionOrder) *NullableOptionOrder {
	return &NullableOptionOrder{value: val, isSet: true}
}

func (v NullableOptionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


