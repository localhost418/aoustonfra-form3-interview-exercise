package accountclient

import (
	"encoding/json"

	"github.com/localhost418/aoustonfra-form3-interview-exercise/generated/models"
)

// CreateAccountResponse represents the API response for a POST account ressource request
type CreateAccountResponse struct {
	Data *models.Account `json:"data"`
}

// ParseCreateAccountResponse unmarshals the given payload to a CreateAccountResponse struct. returns an error if unmarshal fails.
func ParseCreateAccountResponse(payload []byte) (*CreateAccountResponse, error) {
	response := &CreateAccountResponse{Data: &models.Account{}}
	err := json.Unmarshal(payload, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
