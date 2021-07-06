package account

import "github.com/localhost418/aoustonfra-form3-interview-exercise/account/types"

// Service is the interface for Account ressource operations
type Service interface {
	CreateAccount(request *types.CreateAccountRequest) (*types.CreateAccountResponse, error)
	FetchAccount(request *types.FetchAccountRequest) (*types.FetchAccountResponse, error)
	DeleteAccount(request *types.DeleteAccountRequest) (*types.DeleteAccountResponse, error)
}
