package accountclient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/localhost418/aoustonfra-form3-interview-exercise/pkg/util/mocks"
	"github.com/stretchr/testify/assert"
)

// tested service
var accountService AccountService

func init() {
	// inner http client is mocked but we still need a valid url to test url merging with paths in the tested service
	url, _ := url.Parse("http://localhost:8080")
	accountService = New(10, url)
	accountService.setHTTPClient(&mocks.MockClient{})
}

// mock for http client Do func. Returns request response with the given 'statusCodeResponse' and 'jsonResponseStr' as body
func customHTTPDoFuncOk(statusCodeResponse int, jsonResponseStr string) (*http.Response, error) {
	return &http.Response{
		StatusCode: statusCodeResponse,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonResponseStr))),
	}, nil
}

// mock for http client Do func. Returns an error response
func customHTTPDoFuncError() (*http.Response, error) {
	return nil, errors.New("error sending http request")
}

func TestCreateOK(t *testing.T) {
	var jsonResponseStr = `{"data":{"attributes":{"alternative_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","name":["name1","name2"]},"created_on":"2021-07-04T14:36:29.758Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-07-04T14:36:29.758Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(201, jsonResponseStr)
	}

	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := CreateTestAccountRequest(accountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	response, err := accountService.CreateAccount(createAccountRequest)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, accountID, *response.Data.ID)
	assert.EqualValues(t, organisationID, *response.Data.OrganisationID)
	assert.EqualValues(t, accountType, response.Data.Type)
	assert.EqualValues(t, bankID, response.Data.Attributes.BankID)
	assert.EqualValues(t, bankIDCode, response.Data.Attributes.BankIDCode)
	assert.EqualValues(t, baseCurrency, response.Data.Attributes.BaseCurrency)
	assert.EqualValues(t, bic, response.Data.Attributes.Bic)
	assert.EqualValues(t, country, *response.Data.Attributes.Country)
	assert.EqualValues(t, len(name), len(response.Data.Attributes.Name))
	assert.EqualValues(t, name[0], response.Data.Attributes.Name[0])
	assert.EqualValues(t, name[1], response.Data.Attributes.Name[1])

}

func TestCreateNOK(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(500, "")
	}
	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := CreateTestAccountRequest(accountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	response, err := accountService.CreateAccount(createAccountRequest)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestCreateCannotUnmarshalResponse(t *testing.T) {
	var jsonResponseStr = `{"data":{"attribute}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(201, jsonResponseStr)
	}

	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := CreateTestAccountRequest(accountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	response, err := accountService.CreateAccount(createAccountRequest)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestCreateNilRequest(t *testing.T) {
	var jsonResponseStr = `{"data":{"attributes":{"alternative_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","name":["name1","name2"]},"created_on":"2021-07-04T14:36:29.758Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-07-04T14:36:29.758Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(201, jsonResponseStr)
	}
	response, err := accountService.CreateAccount(nil)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestCreateHttpDoError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncError()
	}
	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := CreateTestAccountRequest(accountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	response, err := accountService.CreateAccount(createAccountRequest)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestDeleteOK(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(204, "")
	}
	response, err := accountService.DeleteAccount(&DeleteAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", Version: 0})
	assert.NotNil(t, response)
	assert.Nil(t, err)
}

func TestDeleteNOK(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(500, "")
	}
	response, err := accountService.DeleteAccount(&DeleteAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", Version: 0})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestDeleteNilRequest(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(204, "")
	}
	response, err := accountService.DeleteAccount(nil)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestDeleteHttpDoError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncError()
	}
	response, err := accountService.DeleteAccount(&DeleteAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", Version: 0})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestDeleteUnparsableAccountID(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(200, "")
	}
	response, err := accountService.DeleteAccount(&DeleteAccountRequest{AccountID: "#$%^&*"})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestFetchOK(t *testing.T) {
	jsonResponseStr := `{"data":{"attributes":{"alternative_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","name":["name1","name2"]},"created_on":"2021-07-04T14:58:34.543Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-07-04T14:58:34.543Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(200, jsonResponseStr)
	}

	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	fetchAccountRequest := &FetchAccountRequest{AccountID: accountID}
	response, err := accountService.FetchAccount(fetchAccountRequest)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, accountID, *response.Data.ID)
	assert.EqualValues(t, organisationID, *response.Data.OrganisationID)
	assert.EqualValues(t, accountType, response.Data.Type)
	assert.EqualValues(t, bankID, response.Data.Attributes.BankID)
	assert.EqualValues(t, bankIDCode, response.Data.Attributes.BankIDCode)
	assert.EqualValues(t, baseCurrency, response.Data.Attributes.BaseCurrency)
	assert.EqualValues(t, bic, response.Data.Attributes.Bic)
	assert.EqualValues(t, country, *response.Data.Attributes.Country)
	assert.EqualValues(t, len(name), len(response.Data.Attributes.Name))
	assert.EqualValues(t, name[0], response.Data.Attributes.Name[0])
	assert.EqualValues(t, name[1], response.Data.Attributes.Name[1])
}

func TestFetchCannotUnmarshalResponse(t *testing.T) {
	jsonResponseStr := `{"data":{"attr"}`
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(200, jsonResponseStr)
	}
	fetchAccountRequest := &FetchAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}
	response, err := accountService.FetchAccount(fetchAccountRequest)

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestFetchNOK(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(500, "kreltmkjekl")
	}
	response, err := accountService.FetchAccount(&FetchAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestFetchNilRequest(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(200, "")
	}
	response, err := accountService.FetchAccount(nil)
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestFetchHttpDoError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncError()
	}
	response, err := accountService.FetchAccount(&FetchAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestFetchUnparsableAccountID(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return customHTTPDoFuncOk(200, "")
	}
	response, err := accountService.FetchAccount(&FetchAccountRequest{AccountID: "#$%^&*"})
	assert.NotNil(t, err)
	assert.Nil(t, response)
}

// TODO cover c.newRequest returning error
/*
func TestCreateEmptyRequestUrl(t *testing.T) {
}
*/

/*

// TODO cover json.Marshal(request) returning err
func TestCreateCannotMarshalRequestBody(t *testing.T) {

	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	accountType := "accounts"
	bankID := "400300"
	bankIDCode := "GBDSC"
	baseCurrency := "GBP"
	bic := "NWBKGB22"
	country := "GB"
	name := []string{"name1", "name2"}
	createAccountRequest := CreateTestAccountRequest(accountID, organisationID, accountType, bankID,
		bankIDCode, baseCurrency, bic, country, name)
	response, err := accountService.CreateAccount(createAccountRequest)
}
*/
