package cdeklib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) generalRequest(method string, endpoint string, bodyRequest io.Reader) (map[string]interface{}, error) {
	if err := c.checkToken(); err != nil {
		return map[string]interface{}{}, err
	}

	baseURL, err := url.Parse(c.ApiURL)
	if err != nil {
		return map[string]interface{}{}, err
	}

	//endpoint := endpoint

	fullURL := baseURL.ResolveReference(&url.URL{Path: endpoint})

	req, err := http.NewRequest(method, fullURL.String(), bodyRequest)
	if err != nil {
		return map[string]interface{}{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Auth.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return map[string]interface{}{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return jsonData, nil
}
