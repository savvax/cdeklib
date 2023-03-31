package cdeklib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Calculate returns an array of tariffs based on the given fromLocation, toLocation, and size.
func (c *Client) Calculate(fromLocation, toLocation LocationCalc, size Size) (string, error) {
	req := CalcRequest{
		Type:         CdekType,
		Date:         CdekDate,
		Currency:     CdekCurrency,
		Lang:         CdekLang,
		FromLocation: fromLocation,
		ToLocation:   toLocation,
		Packages:     []Size{size},
	}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	baseURL, err := url.Parse(c.ApiURL)
	if err != nil {
		panic(err)
	}

	endpoint := "v2/calculator/tarifflist"

	fullURL := baseURL.ResolveReference(&url.URL{Path: endpoint})

	// Create a new HTTP request with the built request body.
	request, err := http.NewRequest("POST", fullURL.String(), strings.NewReader(string(requestBody)))
	if err != nil {
		return "", err
	}

	// Set the request headers.
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.Token)

	// Send the request to the CDEK API.
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response status code is not OK.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get a response from CDEK API: %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
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
