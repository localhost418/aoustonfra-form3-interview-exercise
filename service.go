package accountclient

import "github.com/localhost418/accountclient/types"

// Service is the interface for Account ressource operations
type Service interface {
	CreateAccount(request *types.CreateAccountRequest) (*types.CreateAccountResponse, error)
	FetchAccount(request *types.FetchAccountRequest) (*types.FetchAccountResponse, error)
	DeleteAccount(request *types.DeleteAccountRequest) (*types.DeleteAccountResponse, error)
}
