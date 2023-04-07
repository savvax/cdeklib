package cdeklib

import (
	"bytes"
	"encoding/json"
)

func (c *Client) CreateOrder(addrFrom string, addrTo string, recipient Recipient, packages []Package, tariffCode int) (orderID string, err error) {
	if err := c.checkToken(); err != nil {
		return "", err
	}

	addrFromL := OrderLocation{
		Address: addrFrom,
	}

	addrToL := OrderLocation{
		Address: addrTo,
	}

	endpoint := "v2/orders"

	rqst := OrderRequest{
		TariffCode:   tariffCode,
		FromLocation: addrFromL,
		ToLocation:   addrToL,
		Recipient:    recipient,
		Packages:     packages,
	}

	// Преобразуем JSON-объект в []byte
	requestBodyBytes, err := json.Marshal(rqst)
	if err != nil {
		panic(err)
	}

	jsonData, err := c.generalRequest("POST", endpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}

	uuid := jsonData["entity"].(map[string]interface{})["uuid"].(string)

	return uuid, nil
}
