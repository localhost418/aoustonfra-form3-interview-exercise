package types

import (
	"encoding/json"
	"io"

	"github.com/localhost418/accountclient/generated/models"
)

// CreateAccountResponse represents the API response for a POST account ressource request
type CreateAccountResponse struct {
	Data  *models.Account               `json:"data"`
	Links *AccountCreationResponseLinks `json:"links,omitempty"`
}

/*
AccountCreationResponseLinks represents the links provided by the API response when creating/fetching an account ressource
(since we only generated Account ressource this struct was missing)
*/
type AccountCreationResponseLinks struct {

	// Link to the first resource in the list
	// Example: https://api.test.form3.tech/v1/api_name/resource_type
	First *string `json:"first,omitempty"`

	// Link to the last resource in the list
	// Example: https://api.test.form3.tech/v1/api_name/resource_type
	Last *string `json:"last,omitempty"`

	// Link to the next resource in the list
	// Example: https://api.test.form3.tech/v1/api_name/resource_type
	Next *string `json:"next,omitempty"`

	// Link to the previous resource in the list
	// Example: https://api.test.form3.tech/v1/api_name/resource_type
	Prev *string `json:"prev,omitempty"`

	// Link to this resource type
	// Example: https://api.test.form3.tech/v1/api_name/resource_type
	// Required: true
	Self *string `json:"self"`
}

// ReadFrom implements io.ReaderFrom using JSON
func (c *CreateAccountResponse) ReadFrom(r io.Reader) (int64, error) {
	return 0, json.NewDecoder(r).Decode(c)
}
