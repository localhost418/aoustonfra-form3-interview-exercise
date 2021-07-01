package types

import (
	"encoding/json"
	"io"

	"github.com/localhost418/accountclient/generated/models"
)

// FetchAccountResponse represents the API response for a GET account ressource request
type FetchAccountResponse struct {
	Data  *models.Account               `json:"data,omitempty"`
	Links *AccountCreationResponseLinks `json:"links,omitempty"`
}

// ReadFrom implements io.ReaderFrom using JSON
func (c *FetchAccountResponse) ReadFrom(r io.Reader) (int64, error) {
	return 0, json.NewDecoder(r).Decode(c)
}
