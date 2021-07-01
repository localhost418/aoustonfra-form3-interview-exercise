package accountclient

import (
	"encoding/json"

	"github.com/localhost418/aoustonfra-form3-interview-exercise/generated/models"
)

/*
AccountCreationResponseLinks represents the links provided by the API response when fetching an account ressource
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

// FetchAccountResponse represents the API response for a GET account ressource request
type FetchAccountResponse struct {
	Data  *models.Account               `json:"data,omitempty"`
	Links *AccountCreationResponseLinks `json:"links,omitempty"`
}

// ParseFetchAccountResponse unmarshals the given payload to a FetchAccountResponse struct
func ParseFetchAccountResponse(payload []byte) (*FetchAccountResponse, error) {
	response := &FetchAccountResponse{
		Data:  &models.Account{},
		Links: &AccountCreationResponseLinks{},
	}
	err := json.Unmarshal(payload, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}