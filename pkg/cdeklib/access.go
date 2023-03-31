package cdeklib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GetAccessToken retrieves an access token for the CDEK API using the client credentials grant.
// The apiURL parameter specifies the URL of the CDEK API.
// The account and securePassword parameters are used to authenticate with the API.
func GetAccessToken(apiURL, account, securePassword string) (string, error) {
	// Input validation to ensure that all parameters are non-empty
	if apiURL == "" {
		return "", errors.New("apiURL is required")
	}
	if account == "" {
		return "", errors.New("account is required")
	}
	if securePassword == "" {
		return "", errors.New("securePassword is required")
	}
	//
	endpoint := "v2/oauth/token?parameters"
	u := apiURL + endpoint

	// Create request body
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", account)
	data.Set("client_secret", securePassword)

	// Create HTTP request
	request, err := http.NewRequest("POST", u, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	// Set the Content-Type header to indicate that the request body is in URL-encoded form
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the HTTP request using the default client and read the response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check if the response status code is not OK (200)
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get response from CDEK API: %d", response.StatusCode)
	}

	// Parse the JSON response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var apiResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
		Jti         string `json:"jti"`
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", err
	}

	return apiResponse.AccessToken, nil
}
