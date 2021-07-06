package types

import (
	"encoding/json"
	"io"

	"github.com/localhost418/accountclient/generated/models"
)

// CreateAccountResponse represents the API response for a POST account ressource request
type CreateAccountResponse struct {
	Data *models.Account `json:"data"`
}

// ReadFrom implements io.ReaderFrom using JSON
func (c *CreateAccountResponse) ReadFrom(r io.Reader) (int64, error) {
	return 0, json.NewDecoder(r).Decode(c)
}
