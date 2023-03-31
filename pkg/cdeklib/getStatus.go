package cdeklib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) GetStatus(orderID string) (status string, err error) {
	baseURL, err := url.Parse(c.ApiURL)
	if err != nil {
		panic(err)
	}

	endpoint := "v2/orders/" + orderID

	fullURL := baseURL.ResolveReference(&url.URL{Path: endpoint})

	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
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
