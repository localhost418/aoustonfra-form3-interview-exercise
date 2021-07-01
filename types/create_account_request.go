package types

import (
	"encoding/json"
	"io"

	"github.com/localhost418/accountclient/generated/models"
)

// CreateAccountRequest contains all the parameters to POST an Account ressource through the account API
type CreateAccountRequest struct {
	Data *models.Account `json:"data"`
}

// WriteTo implements io.WriterTo using JSON
func (c *CreateAccountRequest) WriteTo(w io.Writer) (int64, error) {
	return 0, json.NewEncoder(w).Encode(c)
}
