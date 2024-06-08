package telegraph

import "fmt"

// CreatePage creates a new page on Telegraph
// See https://telegra.ph/api#createPage
func (c *Client) CreatePage(accessToken, title string, content []Node, authorName, authorURL string) (*Page, error) {
	body := map[string]interface{}{
		"access_token": accessToken,
		"title":        title,
		"content":      content,
		"author_name":  authorName,
		"author_url":   authorURL,
	}

	var result CreatePageResponse
	if err := c.doRequest("POST", "createPage", body, &result); err != nil {
		return nil, fmt.Errorf("failed to create page: %w", err)
	}

	if !result.Ok {
		return nil, ErrCreatePageFailed
	}

	return &result.Result, nil
}

// EditPage edits an existing page on Telegraph
// See https://telegra.ph/api#editPage
func (c *Client) EditPage(accessToken, path, title string, content []Node, authorName, authorURL string) (*Page, error) {
	body := map[string]interface{}{
		"access_token": accessToken,
		"path":         path,
		"title":        title,
		"content":      content,
		"author_name":  authorName,
		"author_url":   authorURL,
	}

	var result CreatePageResponse
	if err := c.doRequest("POST", "editPage/"+path, body, &result); err != nil {
		return nil, fmt.Errorf("failed to edit page: %w", err)
	}

	if !result.Ok {
		return nil, ErrEditPageFailed
	}

	return &result.Result, nil
}

// GetPage retrieves a page from Telegraph
// See https://telegra.ph/api#getPage
func (c *Client) GetPage(path string, returnContent bool) (*Page, error) {
	body := map[string]interface{}{
		"return_content": returnContent,
	}

	var result GetPageResponse
	if err := c.doRequest("GET", "getPage/"+path, body, &result); err != nil {
		return nil, fmt.Errorf("failed to get page: %w", err)
	}

	if !result.Ok {
		return nil, ErrGetPageFailed
	}

	return &result.Result, nil
}

// GetPageList retrieves a list of pages for a Telegraph account
// See https://telegra.ph/api#getPageList
func (c *Client) GetPageList(accessToken string, offset, limit int) (*PageList, error) {
	body := map[string]interface{}{
		"access_token": accessToken,
		"offset":       offset,
		"limit":        limit,
	}

	var result GetPageListResponse
	if err := c.doRequest("POST", "getPageList", body, &result); err != nil {
		return nil, fmt.Errorf("failed to get page list: %w", err)
	}

	if !result.Ok {
		return nil, ErrGetPageListFailed
	}

	return &result.Result, nil
}

// GetViews retrieves the number of views for a page on Telegraph
// See https://telegra.ph/api#getViews
func (c *Client) GetViews(path string, year, month, day int) (*PageViews, error) {
	body := map[string]interface{}{
		"path":  path,
		"year":  year,
		"month": month,
		"day":   day,
	}

	var result GetViewsResponse
	if err := c.doRequest("POST", "getViews", body, &result); err != nil {
		return nil, fmt.Errorf("failed to get views: %w", err)
	}

	if !result.Ok {
		return nil, ErrGetViewsFailed
	}

	return &result.Result, nil
}
