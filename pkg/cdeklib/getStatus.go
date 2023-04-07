package cdeklib

import (
	"encoding/json"
)

func (c *Client) GetStatus(orderID string) (status string, err error) {
	if err := c.checkToken(); err != nil {
		return "", err
	}

	endpoint := "v2/orders/" + orderID

	jsonData, err := c.generalRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}

	// Output all JSON
	jsonDataBytes, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return "", err
	}

	result := string(jsonDataBytes)
	return result, nil
}
