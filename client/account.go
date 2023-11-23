package robinhood

import (
	model "github.com/Ryang20718/robinhood-client/models"
)

// GetAccounts returns all the accounts associated with a login/client.
func (c *Client) GetAccounts() ([]model.AccountInfo, error) {
	var r struct{ Results []model.AccountInfo }
	err := c.GetAndDecode(EPAccounts, &r)
	if err != nil {
		return nil, err
	}

	return r.Results, err
}

// GetCryptoAccounts will return associated cryto account
func (c *Client) GetCryptoAccounts() ([]model.CryptoAccount, error) {
	var r struct{ Results []model.CryptoAccount }
	err := c.GetAndDecode(EPCryptoAccount, &r)
	if err != nil {
		return nil, err
	}

	return r.Results, err
}
