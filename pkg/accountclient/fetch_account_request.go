package accountclient

import "github.com/go-openapi/strfmt"

// FetchAccountRequest contains all the parameters to GET an Account ressource through the account API
type FetchAccountRequest struct {
	AccountID strfmt.UUID
}
