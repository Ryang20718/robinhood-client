package robinhood

import (
	"strings"
	model "github.com/Ryang20718/robinhood-client/models"
)

// GetQuote returns all the latest stock quotes for the list of stocks provided
func (c *Client) GetQuote(stocks ...string) ([]model.Quote, error) {
	url := EPQuotes + "?symbols=" + strings.Join(stocks, ",")
	var r struct{ Results []model.Quote }
	err := c.GetAndDecode(url, &r)
	return r.Results, err
}
