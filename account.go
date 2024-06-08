package telegraph

import "fmt"

// CreateAccount creates a new Telegraph account
// See https://telegra.ph/api#createAccount
func (c *Client) CreateAccount(shortName, authorName, authorURL string) (*Account, error) {
	account := &Account{
		ShortName:  shortName,
		AuthorName: authorName,
		AuthorURL:  authorURL,
	}

	var result CreateAccountResponse
	if err := c.doRequest("POST", "createAccount", account, &result); err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	if !result.Ok {
		return nil, ErrCreateAccountFailed
	}

	return &result.Result, nil
}

// GetAccountInfo retrieves information about a Telegraph account
// See https://telegra.ph/api#getAccountInfo
func (c *Client) GetAccountInfo(accessToken string, fields []string) (*Account, error) {
	body := map[string]interface{}{
		"access_token": accessToken,
		"fields":       fields,
	}

	var result GetAccountInfoResponse
	if err := c.doRequest("POST", "getAccountInfo", body, &result); err != nil {
		return nil, fmt.Errorf("failed to get account info: %w", err)
	}

	if !result.Ok {
		return nil, ErrGetAccountInfoFailed
	}

	return &result.Result, nil
}

// EditAccountInfo edits information of a Telegraph account
// See https://telegra.ph/api#editAccountInfo
func (c *Client) EditAccountInfo(accessToken, shortName, authorName, authorURL string) (*Account, error) {
	body := map[string]interface{}{
		"access_token": accessToken,
		"short_name":   shortName,
		"author_name":  authorName,
		"author_url":   authorURL,
	}

	var result EditAccountInfoResponse
	if err := c.doRequest("POST", "editAccountInfo", body, &result); err != nil {
		return nil, fmt.Errorf("failed to edit account info: %w", err)
	}

	if !result.Ok {
		return nil, ErrEditAccountInfoFailed
	}

	return &result.Result, nil
}

// RevokeAccessToken revokes an access token for a Telegraph account
// See https://telegra.ph/api#revokeAccessToken
func (c *Client) RevokeAccessToken(accessToken string) (*Account, error) {
	body := map[string]interface{}{
		"access_token": accessToken,
	}

	var result RevokeAccessTokenResponse
	if err := c.doRequest("POST", "revokeAccessToken", body, &result); err != nil {
		return nil, fmt.Errorf("failed to revoke access token: %w", err)
	}

	if !result.Ok {
		return nil, ErrRevokeAccessTokenFailed
	}

	return &result.Result, nil
}
