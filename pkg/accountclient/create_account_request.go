package accountclient

import (
	"github.com/go-openapi/strfmt"
	"github.com/localhost418/aoustonfra-form3-interview-exercise/generated/models"
)

// CreateAccountRequest contains all the parameters to POST an Account ressource through the account API
type CreateAccountRequest struct {
	Data *models.Account `json:"data"`
}

// CreateTestAccountRequest creates a CreateAccountRequest with the given fields (used by unit and integration tests)
func CreateTestAccountRequest(accountID, organisationID strfmt.UUID, accountType, bankID, bankIDCode, baseCurrency, bic string, country string, name []string) *CreateAccountRequest {
	createAccountRequest := &CreateAccountRequest{
		Data: &models.Account{
			ID:             &accountID,
			OrganisationID: &organisationID,
			Type:           accountType,
			Attributes: &models.AccountAttributes{
				BankID:       bankID,
				BankIDCode:   bankIDCode,
				BaseCurrency: baseCurrency,
				Bic:          bic,
				Country:      &country,
				Name:         name,
			},
		},
	}
	return createAccountRequest
}
