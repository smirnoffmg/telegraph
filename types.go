package telegraph

// Account represents a Telegraph account
// See https://telegra.ph/api#Account
type Account struct {
	ShortName   string `json:"short_name"`
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	AccessToken string `json:"access_token"`
	PageCount   int    `json:"page_count"`
}

// Node represents a content node
// See https://telegra.ph/api#Node
type Node struct {
	Tag      string      `json:"tag"`
	Attrs    interface{} `json:"attrs,omitempty"`
	Children []Node      `json:"children,omitempty"`
}

// NodeElement represents a content node element
// See https://telegra.ph/api#NodeElement
type NodeElement struct {
	Tag      string            `json:"tag"`
	Attrs    map[string]string `json:"attrs,omitempty"`
	Children []interface{}     `json:"children,omitempty"`
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
