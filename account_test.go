package accountclient_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	account "github.com/localhost418/accountclient"
	"github.com/localhost418/accountclient/generated/models"
	"github.com/localhost418/accountclient/types"
)

// env var key for integration server account API URL
const accountAPIURLEnvKey = "ACCOUNT_API_URL"

// url for integration server account API
var apiURL = getAccountAPIURL()

func getAccountAPIURL() *url.URL {
	accountAPIURL := os.Getenv(accountAPIURLEnvKey)
	if accountAPIURL != "" {
		url, err := url.Parse(accountAPIURL)
		if err != nil {
			panic("cannot url.Parse env var " + accountAPIURLEnvKey)
		}
		return url
	}
	// if env var is not set, use localhost to enable running integration tests locally
	url, _ := url.Parse("http://localhost:8080")
	return url
}

// Requests errors tests (unit tests only since it fails before sending the request to the server)
func TestClientCreateRequest(t *testing.T) {
	tt := []struct {
		name string
		req  *types.CreateAccountRequest
		res  *types.CreateAccountResponse
		err  error
	}{
		{
			name: "nil request",
			req:  nil,
			err:  account.ErrNoRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cli := account.NewClient(&http.Client{Timeout: time.Second}, *apiURL)

			_, err := cli.CreateAccount(tc.req)
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
		})
	}

}

func TestClientFetchRequest(t *testing.T) {
	tt := []struct {
		name string
		req  *types.FetchAccountRequest
		res  *types.FetchAccountResponse
		err  error
	}{
		{
			name: "nil request",
			req:  nil,
			err:  account.ErrNoRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cli := account.NewClient(&http.Client{Timeout: time.Second}, *apiURL)

			_, err := cli.FetchAccount(tc.req)
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
		})
	}

}

func TestClientDeleteRequest(t *testing.T) {
	tt := []struct {
		name string
		req  *types.DeleteAccountRequest
		res  *types.DeleteAccountResponse
		err  error
	}{
		{
			name: "nil request",
			req:  nil,
			err:  account.ErrNoRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cli := account.NewClient(&http.Client{Timeout: time.Second}, *apiURL)

			_, err := cli.DeleteAccount(tc.req)
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
		})
	}

}

func TestClientCreateResponse(t *testing.T) {
	// consts for request OK (mapped as *string by swagger so need to be declared as variables to get their address)
	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	country := "GB"
	version := int64(0)

	tt := []struct {
		name        string
		req         *types.CreateAccountRequest
		status      int
		res         string
		err         error
		expectedRes *types.CreateAccountResponse
	}{
		{
			name:   "invalid status code ",
			req:    &types.CreateAccountRequest{Data: nil},
			status: 500,
			res:    "",
			err:    account.ErrAPIFailure,
		},
		{
			name:   "invalid json ",
			req:    &types.CreateAccountRequest{Data: nil},
			status: http.StatusCreated,
			res:    `{ "invalid-json': }`,
			err:    account.ErrInvalidResponse,
		},
		{
			name: "response OK",
			req: &types.CreateAccountRequest{
				Data: &models.Account{
					ID:             &accountID,
					OrganisationID: &organisationID,
					Type:           "accounts",
					Attributes: &models.AccountAttributes{
						BankID:       "400300",
						BankIDCode:   "GBDSC",
						BaseCurrency: "GBP",
						Bic:          "NWBKGB22",
						Country:      &country,
						Name:         []string{"name1", "name2"},
					},
					Version: &version,
				},
			},
			status: http.StatusCreated,
			res:    `{"data":{"attributes":{"alternative_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","name":["name1","name2"]},"created_on":"2021-07-04T14:36:29.758Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-07-04T14:36:29.758Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.res))
			})

			srv := httptest.NewServer(handler)
			defer srv.Close()

			serverURL, err := url.Parse(srv.URL)
			if err != nil {
				t.Fatalf("cannot build server url: %s", err)
			}

			cli := account.NewClient(&http.Client{Timeout: time.Second}, *serverURL)

			res, err := cli.CreateAccount(tc.req)
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
			if err != nil {
				return
			}
			if res == nil {
				t.Fatal("unexpected nil response")
			}
			if !accountsDeepEqual(res.Data, tc.req.Data) {
				a, _ := json.Marshal(res.Data)
				b, _ := json.Marshal(tc.req.Data)
				t.Fatalf("wrong response value:\n want %s \n got %s", string(a), string(b))
			}

		})
	}
}

func TestClientFetchResponse(t *testing.T) {
	tt := []struct {
		name   string
		req    *types.FetchAccountRequest
		status int
		res    string
		err    error
	}{
		{
			name:   "invalid status code ",
			req:    &types.FetchAccountRequest{AccountID: ""},
			status: http.StatusInternalServerError,
			err:    account.ErrAPIFailure,
		},
		{
			name:   "invalid json ",
			req:    &types.FetchAccountRequest{AccountID: ""},
			status: http.StatusOK,
			res:    `{ "invalid-json': }`,
			err:    account.ErrInvalidResponse,
		},
		{
			name:   "response OK",
			req:    &types.FetchAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
			status: http.StatusOK,
			res:    `{"data":{"attributes":{"alternative_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","name":["name1","name2"]},"created_on":"2021-07-04T14:36:29.758Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2021-07-04T14:36:29.758Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.res))
			})

			srv := httptest.NewServer(handler)
			defer srv.Close()

			serverURL, err := url.Parse(srv.URL)
			if err != nil {
				t.Fatalf("cannot build server url: %s", err)
			}

			cli := account.NewClient(&http.Client{Timeout: time.Second}, *serverURL)

			res, err := cli.FetchAccount(tc.req)
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
			if err != nil {
				return
			}
			if res == nil {
				t.Fatal("unexpected nil response")
			}
			accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
			organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
			country := "GB"
			version := int64(0)
			expectedAccountRes := &models.Account{
				ID:             &accountID,
				OrganisationID: &organisationID,
				Type:           "accounts",
				Attributes: &models.AccountAttributes{
					BankID:       "400300",
					BankIDCode:   "GBDSC",
					BaseCurrency: "GBP",
					Bic:          "NWBKGB22",
					Country:      &country,
					Name:         []string{"name1", "name2"},
				},
				Version: &version,
			}
			if !accountsDeepEqual(res.Data, expectedAccountRes) {
				a, _ := json.Marshal(res.Data)
				b, _ := json.Marshal(tc.req)
				t.Fatalf("wrong response value:\n want %s \n got %s", string(a), string(b))
			}
		})
	}
}

func TestClientDeleteResponse(t *testing.T) {
	tt := []struct {
		name   string
		req    *types.DeleteAccountRequest
		status int
		res    string
		err    error
	}{
		{
			name:   "invalid status code ",
			req:    &types.DeleteAccountRequest{AccountID: ""},
			status: http.StatusInternalServerError,
			err:    account.ErrAPIFailure,
		},
		{
			name:   "response OK ",
			req:    &types.DeleteAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
			status: http.StatusNoContent,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.res))
			})

			srv := httptest.NewServer(handler)
			defer srv.Close()

			serverURL, err := url.Parse(srv.URL)
			if err != nil {
				t.Fatalf("cannot build server url: %s", err)
			}

			cli := account.NewClient(&http.Client{Timeout: time.Second}, *serverURL)

			res, err := cli.DeleteAccount(tc.req)
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
			if err != nil {
				return
			}
			if res == nil {
				t.Fatal("unexpected response")
			}
		})
	}
}

/*
 Implements deepEqual between two *models.Account.
 Since it's only used by tests we marshal both account to JSON and compare the bytes (a bit slow execution but much faster implementation) */
func accountsDeepEqual(x, y *models.Account) bool {
	a, errX := json.Marshal(x)
	b, errY := json.Marshal(y)
	if errX != nil || errY != nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}

func createTestAccountRequest(accountID, organisationID strfmt.UUID, accountType, bankID, bankIDCode, baseCurrency, bic string, country string, name []string) *types.CreateAccountRequest {
	createAccountRequest := &types.CreateAccountRequest{
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
