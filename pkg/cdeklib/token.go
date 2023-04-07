package cdeklib

import (
	"fmt"
	"time"
)

type Token struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int64
	Scope       string
	Jti         string
}

var token Token

const timeBuffer = 30

func (c *Client) GetAccessToken(apiURL, account, securePassword string) (string, error) {
	apiResponse, err := GetAccessResponse(apiURL, account, securePassword)
	if err != nil {
		return "", err
	}

	token.AccessToken = apiResponse.AccessToken
	token.TokenType = apiResponse.TokenType
	token.ExpiresIn = time.Now().Unix() + int64(apiResponse.ExpiresIn)
	token.Scope = apiResponse.Scope
	token.Jti = apiResponse.Jti

	return apiResponse.AccessToken, nil
}

func (c *Client) checkToken() error {
	if c.isTokenExpires() {
		return c.refreshToken()
	}
	return nil
}

func (c *Client) isTokenExpires() bool {
	if token.ExpiresIn-timeBuffer <= time.Now().Unix() {
		return true
	}
	return false
}

func (c *Client) refreshToken() error {
	if c.isTokenExpires() {
		apiResponse, err := GetAccessResponse(c.ApiURL, c.Account, c.SecurePassword)
		if err != nil {
			fmt.Println(err)
		}

		token.AccessToken = apiResponse.AccessToken
		token.TokenType = apiResponse.TokenType
		token.ExpiresIn = time.Now().Unix() + int64(apiResponse.ExpiresIn)
		token.Scope = apiResponse.Scope
		token.Jti = apiResponse.Jti

		return err
	}
	return nil
}
