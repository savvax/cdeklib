package cdeklib

import "fmt"

type Client struct {
	Token      string
	IsTestMode bool
	ApiURL     string
}

func NewClient(isTestMode bool, apiURL string, account string, securePassword string) *Client {
	var token, err = GetAccessToken(apiURL, account, securePassword)
	if err != nil {
		fmt.Println("Error getting access token:", err)
	}
	return &Client{Token: token, IsTestMode: isTestMode, ApiURL: apiURL}
}
