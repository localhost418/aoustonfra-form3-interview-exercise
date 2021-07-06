package types

import "github.com/go-openapi/strfmt"

// DeleteAccountRequest contains all the parameters to DELETE an Account ressource through the account API
type DeleteAccountRequest struct {
	AccountID strfmt.UUID
	Version   int
}
