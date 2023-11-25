package robinhood

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	model "github.com/Ryang20718/robinhood-client/models"
	"github.com/pkg/errors"
)

// OrderSide is which side of the trade an order is on
type OrderSide int

// OrderOpts encapsulates differences between order types
type OrderOpts struct {
	Side          model.Side
	Type          model.ExecutionType
	Quantity      uint64
	Price         float64
	TimeInForce   model.TimeInForce
	ExtendedHours bool
	Stop, Force   bool
}

// Order places an order for a given instrument
func (c *Client) Order(i *model.InstrumentData, o OrderOpts) (*model.Order, error) {
	trigger := model.IMMEDIATE
	a := model.Order{
		Account:       c.Account.Url,
		Instrument:    i.Url,
		Symbol:        i.Symbol,
		Type:          &o.Type,
		TimeInForce:   &o.TimeInForce,
		Quantity:      pointer.ToString(string(o.Quantity)),
		Side:          &o.Side,
		ExtendedHours: &o.ExtendedHours,
		Price:         pointer.ToString(fmt.Sprintf("%f", o.Price)),
		Trigger:       &trigger,
	}

	if o.Stop {
		a.StopPrice = pointer.ToString(fmt.Sprintf("%f", o.Price))
		trigger := model.STOP
		a.Trigger = &trigger
	}

	bs, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	post, err := http.NewRequest("POST", EPOrders, bytes.NewReader(bs))
	post.Header.Add("Content-Type", "application/json")

	var out model.Order
	err = c.DoAndDecode(post, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// UpdateOrder returns any errors and updates the item with any recent changes.
func (c *Client) UpdateOrder(o model.Order) error {
	return c.GetAndDecode(pointer.GetString(o.Url), o)
}

// CancelOrder attempts to cancel an odrer
func (c *Client) CancelOrder(o *model.Order) error {
	post, err := http.NewRequest("POST", pointer.GetString(o.CancelUrl), nil)
	if err != nil {
		return err
	}

	var o2 model.Order
	err = c.DoAndDecode(post, &o2)
	if err != nil {
		return errors.Wrap(err, "could not decode response")
	}

	if pointer.GetString(o2.RejectReason) != "" {
		return errors.New(pointer.GetString(o2.RejectReason))
	}
	return nil
}

// type PaginatedOrder2 struct {
// 	Count    *string       `json:"count,omitempty"`
// 	Next     *string       `json:"next,omitempty"`
// 	Previous *string       `json:"previous,omitempty"`
// 	Results  []interface{} `json:"results,omitempty"`
// }

// GetStockOrders returns orders made by this client.
func (c *Client) GetStockOrders() ([]model.Transaction, error) {
	rs := make([]model.Order, 0)

	transactionList := []model.Transaction{}
	// err := c.GetAndDecode(EPBase+"/orders?page_size=200", &results)

	url := EPOrders
	for {
		var results model.PaginatedOrder
		err := c.GetAndDecode(url, &results)
		if err != nil {
			return nil, err
		}
		rs = append(rs, results.Results...)
		if results.Next == nil {
			break
		}
		url = *results.Next
	}

	for _, order := range rs {
		instrumentData, err := c.GetInstrument(*order.Instrument)
		if err != nil {
			fmt.Println("ERR", "No data", err.Error())
			return transactionList, nil
		}
		order.Symbol = instrumentData.Symbol
		if (*order.State == "cancelled" && len(order.Executions) > 0) || *order.State == "filled" {
			fmt.Println("-1", order.Price, "0", *order.State, "1", order.AveragePrice, "2", *order.Quantity, "3", *instrumentData.Symbol, "4", *order.CreatedAt, "5", *order.Side)
			unitCost, _ := strconv.ParseFloat(*order.AveragePrice, 64)
			qty, _ := strconv.ParseFloat(*order.Quantity, 64)
			transaction := model.Transaction{
				Ticker:          *instrumentData.Symbol,
				TransactionType: fmt.Sprintf("%s", *order.Side), // Buy. Sell
				Qty:             qty,
				UnitCost:        unitCost,
				CreatedAt:       (*order.CreatedAt).Format("2006-01-02 15:04:05"),
				Tag:             fmt.Sprintf("%s", *order.Side),
			}
			transactionList = append(transactionList, transaction)
		}

	}

	return transactionList, nil
}

// RecentOrders returns any recent orders made by this client.
func (c *Client) RecentOrders() ([]model.Order, error) {
	var o struct {
		Results []model.Order
	}
	err := c.GetAndDecode(EPOrders, &o)
	if err != nil {
		return nil, err
	}
	for _, order := range o.Results {
		instrumentData, err := c.GetInstrument(pointer.GetString(order.Instrument))
		if err != nil {
			break
		}
		order.Symbol = instrumentData.Symbol
	}

	return o.Results, nil
}

// RecentOrders returns any events
/*
Returns the events related to a options all events shown here will be assigned
*/
func (c *Client) GetEvents(sym string) (*[]model.OptionAssignment, error) {
	var rs []model.OptionAssignment
	var results model.GetEventsResponse
	instrument, err := c.GetInstrumentForSymbol(sym)
	if err != nil {
		// some symbols may be nil
		return nil, nil
	}
	url := EPEvents + "?equity_instrument_id=" + *instrument.Id
	for {
		err := c.GetAndDecode(url, &results)
		if err != nil {
			return nil, err
		}

		rs = append(rs, *results.Results...)
		if results.Next == nil {
			break
		}

		url = *results.Next
	}
	return &rs, nil
}

/*
Returns the stock splits (THIS DOESN'T WORK)
*/
func (c *Client) GetStockSplits(sym string) (*[]interface{}, error) {
	var rs []interface{}
	var results model.GetStockSplitResponse //model.GetEventsResponse
	instrument, err := c.GetInstrumentForSymbol(sym)
	if err != nil {
		return nil, err
	}
	url := EPInstruments + *instrument.Id + "/splits/"
	for {
		fmt.Println(url)
		err := c.GetAndDecode(url, &results)
		if err != nil {
			return nil, err
		}

		rs = append(rs, *results.Results...)
		if results.Next == nil {
			break
		}

		url = *results.Next
	}
	return &rs, nil
}
