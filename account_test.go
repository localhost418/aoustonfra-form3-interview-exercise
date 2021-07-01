package accountclient_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/localhost418/accountclient"
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

// Request errors tests (only units since it doesn't need any response from the server)
func TestClientCreateRequest(t *testing.T) {
	fakeURL, _ := url.Parse("")
	tt := []struct {
		name string
		req  *types.CreateAccountRequest
		res  *types.CreateAccountResponse
		url  url.URL
		msg  string
	}{
		{
			name: "nil request",
			req:  nil,
			msg:  accountclient.ErrNoRequest,
		},
		{
			name: "error do request",
			req:  &types.CreateAccountRequest{},
			msg:  accountclient.ErrDoRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *fakeURL)

			_, err := cli.CreateAccount(tc.req)
			if err == nil {
				t.Fatalf("no error found, expected %s", tc.msg)
			}
			message := err.Message
			if !strings.Contains(message, tc.msg) {
				t.Fatalf("'%s' not found in '%s'", tc.msg, message)
			}
		})
	}

}

func TestClientFetchRequest(t *testing.T) {
	fakeURL, _ := url.Parse("")
	tt := []struct {
		name string
		req  *types.FetchAccountRequest
		res  *types.FetchAccountResponse
		msg  string
	}{
		{
			name: "nil request",
			req:  nil,
			msg:  accountclient.ErrNoRequest,
		},
		{
			name: "error do request",
			req:  &types.FetchAccountRequest{},
			msg:  accountclient.ErrDoRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *fakeURL)

			_, err := cli.FetchAccount(tc.req)
			if err == nil {
				t.Fatalf("no error found, expected %s", tc.msg)
			}
			message := err.Message
			if !strings.Contains(message, tc.msg) {
				t.Fatalf("'%s' not found in '%s'", tc.msg, message)
			}
		})
	}

}

func TestClientDeleteRequest(t *testing.T) {
	fakeURL, _ := url.Parse("")
	tt := []struct {
		name string
		req  *types.DeleteAccountRequest
		res  *types.DeleteAccountResponse
		msg  string
	}{
		{
			name: "nil request",
			req:  nil,
			msg:  accountclient.ErrNoRequest,
		},
		{
			name: "error do request",
			req:  &types.DeleteAccountRequest{},
			msg:  accountclient.ErrDoRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *fakeURL)

			_, err := cli.DeleteAccount(tc.req)
			if err == nil {
				t.Fatalf("no error found, expected %s", tc.msg)
			}
			message := err.Message
			if !strings.Contains(message, tc.msg) {
				t.Fatalf("'%s' not found in '%s'", tc.msg, message)
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
	accountClassification := "Personal"
	accountMatchingOpt := true
	joinAccount := true
	switched := true
	tt := []struct {
		name        string
		req         *types.CreateAccountRequest
		status      int
		res         string
		err         string
		expectedRes *types.CreateAccountResponse
	}{
		{
			name:   "invalid status code ",
			req:    &types.CreateAccountRequest{Data: nil},
			status: 500,
			res:    "",
			err:    accountclient.ErrAPIFailure,
		},
		{
			name:   "invalid json ",
			req:    &types.CreateAccountRequest{Data: nil},
			status: http.StatusCreated,
			res:    `{ "invalid-json': }`,
			err:    accountclient.ErrInvalidResponse,
		},
		{
			name: "response OK",
			req: &types.CreateAccountRequest{
				Data: &models.Account{
					ID:             &accountID,
					OrganisationID: &organisationID,
					Type:           "accounts",
					Attributes: &models.AccountAttributes{
						AccountClassification:   &accountClassification,
						AccountMatchingOptOut:   &accountMatchingOpt,
						AccountNumber:           "41426819",
						AlternativeNames:        []string{"alias1", "alias2", "alias3"},
						BankID:                  "400300",
						BankIDCode:              "GBDSC",
						BaseCurrency:            "GBP",
						Bic:                     "NWBKGB22",
						Country:                 &country,
						CustomerID:              "12345",
						Iban:                    "GB11NWBK40030041426819",
						JointAccount:            &joinAccount,
						Name:                    []string{"name1", "name2"},
						ProcessingService:       "processing_service",
						ReferenceMask:           "4929############",
						SecondaryIdentification: "second_id",
						Status:                  "pending",
						StatusReason:            "transferred",
						Switched:                &switched,
						UserDefinedInformation:  "user_infos",
						ValidationType:          "card",
					},
					Version: &version,
				},
			},
			status: http.StatusCreated,
			res:    `{"data":{"attributes":{"account_classification":"Personal","account_matching_opt_out":true,"account_number":"41426819","alternative_bank_account_names":null,"alternative_names":["alias1","alias2","alias3"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","customer_id":"12345","iban":"GB11NWBK40030041426819","joint_account":true,"name":["name1","name2"],"processing_service":"processing_service","reference_mask":"4929############","secondary_identification":"second_id","status":"pending","status_reason":"transferred","switched":true,"user_defined_information":"user_infos","validation_type":"card"},"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`,
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

			cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *serverURL)

			res, errAcc := cli.CreateAccount(tc.req)
			if (errAcc == nil) != (tc.err == "") {
				t.Fatalf("unexpected error %v ; expected %s", errAcc, tc.err)
			}
			if tc.err != "" {
				return
			}

			if res == nil {
				t.Fatal("unexpected nil response")
			}
			if !accountsDeepEqual(res.Data, tc.req.Data) {
				// Marshal in case of test failure to see clearly which field is not equal
				got, _ := json.Marshal(res.Data)
				want, _ := json.Marshal(tc.req.Data)
				t.Fatalf("wrong response value:\n want %s \n got %s", string(want), string(got))
			}

			if *res.Links.Self != "/v1/organisation/accounts/"+accountID.String() {
				got, _ := json.Marshal(res.Links)
				t.Fatalf("wrong links in create response.\n got %s", string(got))
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
		err    string
	}{
		{
			name:   "invalid status code ",
			req:    &types.FetchAccountRequest{},
			status: http.StatusInternalServerError,
			err:    accountclient.ErrAPIFailure,
		},
		{
			name:   "invalid json ",
			req:    &types.FetchAccountRequest{},
			status: http.StatusOK,
			res:    `{ "invalid-json': }`,
			err:    accountclient.ErrInvalidResponse,
		},
		{
			name:   "response OK",
			req:    &types.FetchAccountRequest{AccountID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"},
			status: http.StatusOK,
			res:    `{"data":{"attributes":{"account_classification":"Personal","account_matching_opt_out":true,"account_number":"41426819","alternative_bank_account_names":null,"alternative_names":["alias1","alias2","alias3"],"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB","customer_id":"12345","iban":"GB11NWBK40030041426819","joint_account":true,"name":["name1","name2"],"processing_service":"processing_service","reference_mask":"4929############","secondary_identification":"second_id","status":"pending","status_reason":"transferred","switched":true,"user_defined_information":"user_infos","validation_type":"card"},"id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0},"links":{"self":"/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"}}`,
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

			cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *serverURL)

			res, errAcc := cli.FetchAccount(tc.req)
			if (errAcc == nil) != (tc.err == "") {
				t.Fatalf("unexpected error %v ; expected %s", errAcc, tc.err)
			}
			if tc.err != "" {
				return
			}

			if res == nil {
				t.Fatal("unexpected nil response")
			}
			accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
			organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
			country := "GB"
			version := int64(0)
			accountClassification := "Personal"
			accountMatchingOpt := true
			joinAccount := true
			switched := true
			expectedAccountRes := &models.Account{
				ID:             &accountID,
				OrganisationID: &organisationID,
				Type:           "accounts",
				Attributes: &models.AccountAttributes{
					//		AcceptanceQualifier:         "same_day",
					AccountClassification:   &accountClassification,
					AccountMatchingOptOut:   &accountMatchingOpt,
					AccountNumber:           "41426819",
					AlternativeNames:        []string{"alias1", "alias2", "alias3"},
					BankID:                  "400300",
					BankIDCode:              "GBDSC",
					BaseCurrency:            "GBP",
					Bic:                     "NWBKGB22",
					Country:                 &country,
					CustomerID:              "12345",
					Iban:                    "GB11NWBK40030041426819",
					JointAccount:            &joinAccount,
					Name:                    []string{"name1", "name2"},
					ProcessingService:       "processing_service",
					ReferenceMask:           "4929############",
					SecondaryIdentification: "second_id",
					Status:                  "pending",
					StatusReason:            "transferred",
					Switched:                &switched,
					UserDefinedInformation:  "user_infos",
					ValidationType:          "card",
				},
				Version: &version,
			}
			if !accountsDeepEqual(res.Data, expectedAccountRes) {
				// Marshal in case of test failure to see clearly which field is not equal
				want, _ := json.Marshal(expectedAccountRes)
				got, _ := json.Marshal(res.Data)
				t.Fatalf("wrong response value:\n want %s \n got %s", string(want), string(got))
			}

			if *res.Links.Self != "/v1/organisation/accounts/"+accountID.String() {
				got, _ := json.Marshal(res.Links)
				t.Fatalf("wrong links in fetch response.\n got %s", string(got))
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
		err    string
	}{
		{
			name:   "invalid status code ",
			req:    &types.DeleteAccountRequest{},
			status: http.StatusInternalServerError,
			err:    accountclient.ErrAPIFailure,
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

			cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *serverURL)

			res, errAcc := cli.DeleteAccount(tc.req)
			if (errAcc == nil) != (tc.err == "") {
				t.Fatalf("unexpected error %v ; expected %s", errAcc, tc.err)
			}
			if tc.err != "" {
				return
			}

			if res == nil {
				t.Fatal("unexpected response")
			}
		})
	}
}

// integration tests for request errors
func TestClientCreateErrorIntegration(t *testing.T) {
	expectedError := accountclient.ErrAPIFailure
	cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *apiURL)

	_, err := cli.CreateAccount(&types.CreateAccountRequest{})
	if err == nil || (!strings.Contains(err.Message, expectedError)) {
		t.Fatalf("unexpected error message '%s': expected '%s'", err.Message, expectedError)
	}
}

func TestClientFetchErrorIntegration(t *testing.T) {
	expectedError := accountclient.ErrAPIFailure
	cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *apiURL)
	_, err := cli.FetchAccount(&types.FetchAccountRequest{AccountID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"})
	if err == nil || (!strings.Contains(err.Message, expectedError)) {
		t.Fatalf("unexpected error message '%s': expected '%s'", err.Message, expectedError)
	}
}

func TestClientDeleteErrorIntegration(t *testing.T) {
	expectedError := accountclient.ErrAPIFailure
	cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *apiURL)
	_, err := cli.DeleteAccount(&types.DeleteAccountRequest{AccountID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"})
	if err == nil || (!strings.Contains(err.Message, expectedError)) {
		t.Fatalf("unexpected error message '%s': expected '%s'", err.Message, expectedError)
	}
}

// integration test running CREATE then FETCH then DELETE account ressource
func TestClientIntegration(t *testing.T) {
	accountID := strfmt.UUID("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	organisationID := strfmt.UUID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	country := "GB"
	version := int64(0)
	accountClassification := "Personal"
	accountMatchingOpt := true
	joinAccount := true
	switched := true
	accountModel := &models.Account{
		ID:             &accountID,
		OrganisationID: &organisationID,
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			//		AcceptanceQualifier:         "same_day",
			AccountClassification: &accountClassification,
			AccountMatchingOptOut: &accountMatchingOpt,
			AccountNumber:         "41426819",
			AlternativeNames:      []string{"alias1", "alias2", "alias3"},
			BankID:                "400300",
			BankIDCode:            "GBDSC",
			BaseCurrency:          "GBP",
			Bic:                   "NWBKGB22",
			Country:               &country,
			//	CustomerID:              "12345",
			Iban:         "GB11NWBK40030041426819",
			JointAccount: &joinAccount,
			Name:         []string{"name1", "name2"},
			//		ProcessingService:       "processing_service",
			//		ReferenceMask:           "4929############",
			SecondaryIdentification: "second_id",
			Status:                  "pending",
			//	StatusReason:            "transferred",
			Switched: &switched,
			//	UserDefinedInformation:  "user_infos",
			//	ValidationType:          "card",
		},
		Version: &version,
	}

	cli := accountclient.NewClient(&http.Client{Timeout: time.Second}, *apiURL)

	resCreate, err := cli.CreateAccount(&types.CreateAccountRequest{Data: accountModel})
	if err != nil {
		t.Fatalf("unexpected error create: %v", err)
	}
	if !accountsDeepEqual(resCreate.Data, accountModel) {
		// Marshal in case of test failure to see clearly which field is not equal
		got, _ := json.Marshal(resCreate.Data)
		want, _ := json.Marshal(accountModel)
		t.Fatalf("wrong create response content:\n want %s \n got %s", string(want), string(got))
	}
	if *resCreate.Links.Self != "/v1/organisation/accounts/"+accountID.String() {
		got, _ := json.Marshal(resCreate.Links)
		t.Fatalf("wrong links in fetch response.\n got %s", string(got))
	}

	resFetch, err := cli.FetchAccount(&types.FetchAccountRequest{AccountID: accountID})
	if err != nil {
		t.Fatalf("unexpected fetch error: %v", err)
	}
	if !accountsDeepEqual(resFetch.Data, accountModel) {
		// Marshal in case of test failure to see clearly which field is not equal
		got, _ := json.Marshal(resFetch.Data)
		want, _ := json.Marshal(accountModel)
		t.Fatalf("wrong fetch response content:\n want %s \n got %s", string(want), string(got))
	}
	if *resFetch.Links.Self != "/v1/organisation/accounts/"+accountID.String() {
		got, _ := json.Marshal(resFetch.Links)
		t.Fatalf("wrong links in fetch response.\n got %s", string(got))
	}

	_, err = cli.DeleteAccount(&types.DeleteAccountRequest{AccountID: accountID, Version: 0})
	if err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}
}

/*
 Implements deepEqual between two *models.Account.
 Since it's only used by tests we marshal both account to JSON and compare the bytes (a bit slow execution but very quick to implement) */
func accountsDeepEqual(x, y *models.Account) bool {
	a, errX := json.Marshal(x)
	b, errY := json.Marshal(y)
	if errX != nil || errY != nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}
