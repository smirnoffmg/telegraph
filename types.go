package telegraph

import (
	"encoding/json"
	"errors"
)

// Account represents a Telegraph account
// See https://telegra.ph/api#Account
type Account struct {
	ShortName   string `json:"short_name"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	AccessToken string `json:"access_token"`
	PageCount   int    `json:"page_count"`
}

// Node represents a content node which can be a string (text node) or a NodeElement
// See https://telegra.ph/api#Node
type Node interface{}

// NodeElement represents a DOM element node
// See https://telegra.ph/api#NodeElement
type NodeElement struct {
	Tag      string            `json:"tag"`
	Attrs    map[string]string `json:"attrs,omitempty"`
	Children []Node            `json:"children,omitempty"`
}

// nodeWrapper is used to facilitate custom JSON marshaling for the Node type
type nodeWrapper struct {
	Text *string      `json:"text,omitempty"`
	Elem *NodeElement `json:"elem,omitempty"`
}

// MarshalJSON implements custom JSON marshaling for the Node type
func (n *NodeElement) MarshalJSON() ([]byte, error) {
	type nodeElementAlias NodeElement // Define an alias to prevent infinite recursion
	wrapper := struct {
		*nodeElementAlias
	}{
		nodeElementAlias: (*nodeElementAlias)(n),
	}
	return json.Marshal(&wrapper)
}

// UnmarshalJSON implements custom JSON unmarshalling for the Node type
func (n *NodeElement) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	tag, ok := raw["tag"].(string)
	if !ok {
		return errors.New("missing or invalid 'tag' field")
	}
	n.Tag = tag

	if attrsRaw, ok := raw["attrs"].(map[string]interface{}); ok {
		n.Attrs = make(map[string]string)
		for key, value := range attrsRaw {
			if strValue, ok := value.(string); ok {
				n.Attrs[key] = strValue
			}
		}
	}

	if childrenRaw, ok := raw["children"].([]interface{}); ok {
		var children []Node
		for _, childRaw := range childrenRaw {
			if childStr, ok := childRaw.(string); ok {
				children = append(children, childStr)
			}
		}
		n.Children = children
	}

	return nil
}

// Page represents a Telegraph page
// See https://telegra.ph/api#Page
type Page struct {
	Path        string `json:"path"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	ImageURL    string `json:"image_url"`
	Content     []Node `json:"content"`
	Views       int    `json:"views"`
	CanEdit     bool   `json:"can_edit"`
}

// PageList represents a list of Telegraph pages
// See https://telegra.ph/api#PageList
type PageList struct {
	TotalCount int    `json:"total_count"`
	Pages      []Page `json:"pages"`
}

// PageViews represents the views of a Telegraph page
// See https://telegra.ph/api#PageViews
type PageViews struct {
	Path  string `json:"path"`
	Views int    `json:"views"`
}

// CreateAccountResponse represents the response from the createAccount method
// See https://telegra.ph/api#createAccount
type CreateAccountResponse struct {
	Ok     bool    `json:"ok"`
	Result Account `json:"result"`
}

// CreatePageResponse represents the response from the createPage method
// See https://telegra.ph/api#createPage
type CreatePageResponse struct {
	Ok     bool `json:"ok"`
	Result Page `json:"result"`
}

// GetAccountInfoResponse represents the response from the getAccountInfo method
// See https://telegra.ph/api#getAccountInfo
type GetAccountInfoResponse struct {
	Ok     bool    `json:"ok"`
	Result Account `json:"result"`
}

// EditAccountInfoResponse represents the response from the editAccountInfo method
// See https://telegra.ph/api#editAccountInfo
type EditAccountInfoResponse struct {
	Ok     bool    `json:"ok"`
	Result Account `json:"result"`
}

// RevokeAccessTokenResponse represents the response from the revokeAccessToken method
// See https://telegra.ph/api#revokeAccessToken
type RevokeAccessTokenResponse struct {
	Ok     bool    `json:"ok"`
	Result Account `json:"result"`
}

// GetPageResponse represents the response from the getPage method
// See https://telegra.ph/api#getPage
type GetPageResponse struct {
	Ok     bool `json:"ok"`
	Result Page `json:"result"`
}

// GetPageListResponse represents the response from the getPageList method
// See https://telegra.ph/api#getPageList
type GetPageListResponse struct {
	Ok     bool     `json:"ok"`
	Result PageList `json:"result"`
}

// GetViewsResponse represents the response from the getViews method
// See https://telegra.ph/api#getViews
type GetViewsResponse struct {
	Ok     bool      `json:"ok"`
	Result PageViews `json:"result"`
}
