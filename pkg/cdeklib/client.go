package cdeklib

import "fmt"

type Client struct {
	ApiURL         string
	IsTestMode     bool
	Auth           Auth
	Account        string
	SecurePassword string
}

func NewClient(isTestMode bool, apiURL string, account string, securePassword string) *Client {
	var auth, err = GetAccessResponse(apiURL, account, securePassword)
	if err != nil {
		fmt.Println("Error getting access token:", err)
	}
	return &Client{Auth: auth, IsTestMode: isTestMode, ApiURL: apiURL, Account: account, SecurePassword: securePassword}
}
