package robinhood

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlekSi/pointer"
	model "github.com/Ryang20718/robinhood-client/models"
	"github.com/google/uuid"
)

// OptionsOrderOpts encapsulates common Options order choices
type OptionsOrderOpts struct {
	Quantity    float64
	Price       float64
	Direction   model.Direction
	TimeInForce model.TimeInForce
	Type        model.ExecutionType
	Side        model.Side
	Trigger     model.Trigger
}

// OrderOptions places a new order for options
func (c *Client) OrderOptions(q *model.OptionInstrument, o OptionsOrderOpts) (json.RawMessage, error) {
	b := model.OptionOrderInput{
		Account:     c.Account.Url,
		Direction:   &o.Direction,
		TimeInForce: &o.TimeInForce,
		Legs: []model.Leg{{
			Option:         q.Url,
			RatioQuantity:  pointer.ToString("1"),
			Side:           &o.Side,
			PositionEffect: pointer.ToString("open"),
		}},
		Trigger:  &o.Trigger,
		Type:     &o.Type,
		Quantity: pointer.ToString(fmt.Sprintf("%f", o.Quantity)),
		Price:    pointer.ToString(fmt.Sprintf("%f", o.Price)),
		RefId:    pointer.ToString(uuid.New().String()),
	}

	if o.Side != model.BUY {
		b.Legs[0].PositionEffect = pointer.ToString("close")
	}

	bs, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", EPOptions+"orders/", bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var out json.RawMessage
	err = c.DoAndDecode(req, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetOptionsOrders returns all outstanding options orders
func (c *Client) GetOptionsOrders(ctx context.Context) (*[]model.OptionOrder, error) {
	var results model.GetOptionOrdersResponse
	var rs []model.OptionOrder

	url := EPOptions + "orders/"
	for {
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
		if *order.State != "filled" {
			continue
		}

		for _, leg := range order.Legs {
			instrument, err := c.GetHistoricalOptionsInstrument(ctx, *leg.Option)
			if err != nil {
				return nil, err
			}
			fmt.Println("GG", *order.CreatedAt, *instrument.ExpirationDate, *instrument.IssueDate, *instrument.Tradability, *instrument.State, *order.Price, *order.ProcessedQuantity, *order.ChainSymbol, *instrument.StrikePrice, *instrument.Type, *leg.Side)
		}
	}

	return &rs, nil
}
