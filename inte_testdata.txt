package it_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/localhost418/aoustonfra-form3-interview-exercise/pkg/accountclient"
	"github.com/stretchr/testify/assert"
)

const accountAPIURLEnvKey = "ACCOUNT_API_URL"

var accountService = accountclient.New(10, getAccountAPIURL())

func getAccountAPIURL() *url.URL {
	accountAPIURL := os.Getenv(accountAPIURLEnvKey)
	if accountAPIURL != "" {
		url, err := url.Parse(accountAPIURL)
		if err != nil {
			panic("cannot url.Parse env var " + accountAPIURLEnvKey)
		}
		return url
	} else {
		// if env var is not set, use localhost to enable running integration tests locally
		url, _ := url.Parse("http://localhost:8080")
		return url
	}
}
func TestCreateFetchThenDelete(t *testing.T) {
	// CREATE
	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := accountclient.CreateTestAccountRequest(accountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	responseCreate, err := accountService.CreateAccount(createAccountRequest)
	if err == nil {
		t.Logf("Created account with id %s\n", responseCreate.Data.ID)
	} else {
		t.Error(err)
	}
	assert.NotNil(t, responseCreate)
	assert.Nil(t, err)
	assert.EqualValues(t, accountID, *responseCreate.Data.ID)
	assert.EqualValues(t, organisationID, *responseCreate.Data.OrganisationID)
	assert.EqualValues(t, accountType, responseCreate.Data.Type)
	assert.EqualValues(t, bankID, responseCreate.Data.Attributes.BankID)
	assert.EqualValues(t, bankIDCode, responseCreate.Data.Attributes.BankIDCode)
	assert.EqualValues(t, baseCurrency, responseCreate.Data.Attributes.BaseCurrency)
	assert.EqualValues(t, bic, responseCreate.Data.Attributes.Bic)
	assert.EqualValues(t, country, *responseCreate.Data.Attributes.Country)
	assert.EqualValues(t, len(name), len(responseCreate.Data.Attributes.Name))
	assert.EqualValues(t, name[0], responseCreate.Data.Attributes.Name[0])
	assert.EqualValues(t, name[1], responseCreate.Data.Attributes.Name[1])

	// FETCH
	fetchAccountRequest := &accountclient.FetchAccountRequest{AccountID: accountID}
	responseFetch, err := accountService.FetchAccount(fetchAccountRequest)
	if err == nil {
		t.Logf("Fetched account with id %s\n", responseFetch.Data.ID)
	} else {
		t.Error(err)
	}
	assert.NotNil(t, responseFetch)
	assert.Nil(t, err)
	// DELETE
	deleteAccountRequest := &accountclient.DeleteAccountRequest{
		AccountID: accountID,
		Version:   0,
	}
	responseDelete, err := accountService.DeleteAccount(deleteAccountRequest)
	if err == nil {
		t.Logf("Deleted account with id %s\n", accountID)
	} else {
		t.Error(err)
	}
	assert.NotNil(t, responseDelete)
	assert.Nil(t, err)
}

func TestCreateUnvalidUUID(t *testing.T) {
	unvalidAccountID := strfmt.UUID("1234")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := accountclient.CreateTestAccountRequest(unvalidAccountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	responseCreate, err := accountService.CreateAccount(createAccountRequest)
	assert.NotNil(t, err)
	assert.Nil(t, responseCreate)
}

func TestFetchUnvalidUUID(t *testing.T) {
	response, err := accountService.FetchAccount(&accountclient.FetchAccountRequest{AccountID: "1234"})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestDeleteUnvalidUUID(t *testing.T) {
	response, err := accountService.DeleteAccount(&accountclient.DeleteAccountRequest{
		AccountID: "1234",
		Version:   0,
	})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}
func TestFetchFetchUnexistingAccount(t *testing.T) {
	fetchAccountRequest := &accountclient.FetchAccountRequest{AccountID: "ez4bd6f5-c3f5-44b2-b677-acd23cdde73c"}
	response, err := accountService.FetchAccount(fetchAccountRequest)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestDeleteUnexistingAccount(t *testing.T) {
	response, err := accountService.DeleteAccount(&accountclient.DeleteAccountRequest{
		AccountID: "ez4bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Version:   0,
	})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}
