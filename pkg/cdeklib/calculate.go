package cdeklib

import (
	"bytes"
	"encoding/json"
)

// Calculate returns an array of tariffs based on the given fromLocation, toLocation, and size.
func (c *Client) Calculate(fromLocation, toLocation LocationCalc, size Size) (string, error) {
	if err := c.checkToken(); err != nil {
		return "", err
	}

	req := CalcRequest{
		Type:         CdekType,
		Date:         CdekDate,
		Currency:     CdekCurrency,
		Lang:         CdekLang,
		FromLocation: fromLocation,
		ToLocation:   toLocation,
		Packages:     []Size{size},
	}

	requestBodyBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	endpoint := "v2/calculator/tarifflist"

	jsonData, err := c.generalRequest("POST", endpoint, bytes.NewBuffer(requestBodyBytes))
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
