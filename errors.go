package telegraph

import "errors"

// Define package-specific errors
var (
	ErrUnexpectedStatusCode    = errors.New("unexpected status code")
	ErrCreateAccountFailed     = errors.New("failed to create account")
	ErrGetAccountInfoFailed    = errors.New("failed to get account info")
	ErrEditAccountInfoFailed   = errors.New("failed to edit account info")
	ErrRevokeAccessTokenFailed = errors.New("failed to revoke access token")
	ErrCreatePageFailed        = errors.New("failed to create page")
	ErrEditPageFailed          = errors.New("failed to edit page")
	ErrGetPageFailed           = errors.New("failed to get page")
	ErrGetPageListFailed       = errors.New("failed to get page list")
	ErrGetViewsFailed          = errors.New("failed to get views")
)
