package robinhood

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	rs := []model.OptionOrder{}
	cache := make(map[string]*[]model.OptionAssignment)
	url := EPOptions + "orders/"
	for {
		var results model.GetOptionOrdersResponse
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

	for idx, order := range rs {
		rs[idx].Assigned = new(string)
		rs[idx].Expired = new(string)
		rs[idx].StrikePrice = new(string)
		rs[idx].ExpirationDate = new(string)
		if *order.State != "filled" {
			continue
		}

		for _, leg := range order.Legs {
			instrument, err := c.GetHistoricalOptionsInstrument(ctx, *leg.Option)
			if err != nil {
				return nil, err
			}
			t, _ := time.Parse("2006-01-02", *instrument.ExpirationDate)

			optionExpiryDate := t.Unix()
			today := time.Now().Unix()

			*rs[idx].StrikePrice = *instrument.StrikePrice
			*rs[idx].ExpirationDate = *instrument.ExpirationDate
			// cache recent events API CALL based on ticker --> events
			// *instrument.State --> expired or cancelled
			recentEvents, exists := cache[*order.ChainSymbol]
			if !exists {
				recentEvents, err = c.GetEvents(*order.ChainSymbol)
				if err != nil {
					return nil, err
				}
				cache[*order.ChainSymbol] = recentEvents
			}

			if recentEvents != nil {
				for _, optionAssignment := range *recentEvents {
					strikePrice := *optionAssignment.EquityComponents[0].Price
					if *instrument.StrikePrice == strikePrice && *optionAssignment.EventDate == *instrument.ExpirationDate {
						*rs[idx].Assigned = "true"
					}
				}
			}
			if optionExpiryDate <= today && *rs[idx].Assigned != "true" {
				*rs[idx].Expired = "true"
			}
		}
	}

	return &rs, nil
}
