package accountclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	urlutil "github.com/localhost418/aoustonfra-form3-interview-exercise/pkg/util/url"
)

const (
	baseAPIPath     = "v1"
	accountsAPIPath = baseAPIPath + "/organisation/accounts"
)

// New creates a new AccountService with the given timeout (in seconds) and accountAPI url (scheme+host+port)
func New(requestTimeout int, url *url.URL) AccountService {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * time.Duration(requestTimeout),
		},
		url: url,
	}
}

// Client contains the inner http.Client and API url
type Client struct {
	httpClient HTTPClient
	url        *url.URL
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// AccountService is the interface for Account ressource operations
type AccountService interface {
	CreateAccount(request *CreateAccountRequest) (*CreateAccountResponse, error)

	FetchAccount(request *FetchAccountRequest) (*FetchAccountResponse, error)

	DeleteAccount(request *DeleteAccountRequest) (*DeleteAccountResponse, error)

	setHTTPClient(client HTTPClient)
}

// setHTTPClient to override the inner http.Client (used by unit tests)
func (c *Client) setHTTPClient(client HTTPClient) {
	c.httpClient = client
}

// newRequest creates an http.Request with the given method and c.url+paths url
func (c *Client) newRequest(method string, paths ...string) (*http.Request, error) {
	return http.NewRequest(method, urlutil.JoinURLAndPaths(c.url, paths...).String(), nil)
}

// generic error formating for unexpected status code
func errFormatWrongStatusCode(method, endpoint string, statusCode int, payload []byte) error {
	return fmt.Errorf("[%s %s error\nStatus [%d]\nPayload [%s]]", method, endpoint, statusCode, string(payload))
}

// CreateAccount creates an account with the fields declared in the request
func (c *Client) CreateAccount(request *CreateAccountRequest) (*CreateAccountResponse, error) {
	method := "POST"
	if request == nil {
		return nil, fmt.Errorf("error: received nil request on %s %s endpoint", method, accountsAPIPath)
	}
	req, err := c.newRequest(method, accountsAPIPath)
	if err != nil {
		return nil, err
	}
	requestBodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
	defer req.Body.Close()
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	payload, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 201 {
		return nil, errFormatWrongStatusCode(method, accountsAPIPath, response.StatusCode, payload)
	}
	createAccountResponse, err := ParseCreateAccountResponse(payload)
	if err != nil {
		return nil, err
	}
	return createAccountResponse, nil
}

// FetchAccount fetch an account by accountID
func (c *Client) FetchAccount(request *FetchAccountRequest) (*FetchAccountResponse, error) {
	method := "GET"
	if request == nil {
		return nil, fmt.Errorf("error: received nil request on %s %s endpoint", method, accountsAPIPath)
	}

	req, err := c.newRequest(method, accountsAPIPath, request.AccountID.String())
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.api+json")
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	payload, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return nil, errFormatWrongStatusCode(method, path.Join(accountsAPIPath, request.AccountID.String()), response.StatusCode, payload)
	}
	fetchAccountResponse, err := ParseFetchAccountResponse(payload)
	if err != nil {
		return nil, err
	}
	return fetchAccountResponse, nil
}

// DeleteAccount deletes an account by accountID and version
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	method := "DELETE"
	if request == nil {
		return nil, fmt.Errorf("error: received nil request on %s %s endpoint", method, accountsAPIPath)
	}
	req, err := c.newRequest(method, accountsAPIPath, request.AccountID.String())
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("version", strconv.Itoa(request.Version))
	req.URL.RawQuery = q.Encode()
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	payload, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 204 {
		return nil, errFormatWrongStatusCode(method, path.Join(accountsAPIPath, request.AccountID.String()), response.StatusCode, payload)
	}
	return &DeleteAccountResponse{}, nil
}
