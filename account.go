package accountclient

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/localhost418/accountclient/types"
)

const (
	baseAPIPath     = "v1"
	accountsAPIPath = baseAPIPath + "/organisation/accounts"
)

// Client implements account service
type Client struct {
	client *http.Client
	url    url.URL
}

// NewClient creates a new AccountService with the given timeout (in seconds) and accountAPI url (scheme+host+port)
func NewClient(client *http.Client, url url.URL) *Client {
	/*	httpClient: &http.Client{
		Timeout: time.Second * time.Duration(requestTimeout),
	}, */
	return &Client{
		client: client,
		url:    url,
	}
}

// CreateAccount creates an account with the fields declared in the request
func (c *Client) CreateAccount(req *types.CreateAccountRequest) (*types.CreateAccountResponse, error) {
	const method = http.MethodPost
	if req == nil {
		return nil, fail(method, accountsAPIPath, ErrNoRequest)
	}

	body := &bytes.Buffer{}
	_, err := req.WriteTo(body)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrInvalidBody)
	}

	r, err := http.NewRequest(method, buildURL(c.url, []string{accountsAPIPath}), body)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrInvalidRequest)
	}
	r.Header.Add("Accept", "application/vnd.api+json")

	w, err := c.client.Do(r)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrDoRequest)
	}
	defer w.Body.Close()

	if w.StatusCode != http.StatusCreated {
		return nil, failWithStatusCode(method, accountsAPIPath, w.StatusCode, ErrAPIFailure)
	}
	res := &types.CreateAccountResponse{}
	_, err = res.ReadFrom(w.Body)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrInvalidResponse)
	}
	return res, nil
}

// FetchAccount fetch an account by accountID
func (c *Client) FetchAccount(req *types.FetchAccountRequest) (*types.FetchAccountResponse, error) {
	const method = http.MethodPost
	if req == nil {
		return nil, fail(method, accountsAPIPath, ErrNoRequest)
	}

	r, err := http.NewRequest(method, buildURL(c.url, []string{accountsAPIPath, req.AccountID.String()}), nil)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrInvalidRequest)
	}
	r.Header.Add("Accept", "application/vnd.api+json")

	w, err := c.client.Do(r)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrDoRequest)
	}
	defer w.Body.Close()

	if w.StatusCode != http.StatusOK {
		return nil, failWithStatusCode(method, accountsAPIPath, w.StatusCode, ErrAPIFailure)
	}
	res := &types.FetchAccountResponse{}
	_, err = res.ReadFrom(w.Body)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrInvalidResponse)
	}

	return res, nil
}

// DeleteAccount deletes an account by accountID and version
func (c *Client) DeleteAccount(req *types.DeleteAccountRequest) (*types.DeleteAccountResponse, error) {
	const method = http.MethodPost
	if req == nil {
		return nil, fail(method, accountsAPIPath, ErrNoRequest)
	}

	r, err := http.NewRequest(method, buildURL(c.url, []string{accountsAPIPath, req.AccountID.String()}), nil)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrInvalidRequest)
	}

	r.Header.Add("Accept", "application/vnd.api+json")
	q := r.URL.Query()
	q.Add("version", strconv.Itoa(req.Version))
	r.URL.RawQuery = q.Encode()

	w, err := c.client.Do(r)
	if err != nil {
		return nil, fail(method, accountsAPIPath, ErrDoRequest)
	}
	defer w.Body.Close()

	if w.StatusCode != http.StatusNoContent {
		return nil, failWithStatusCode(method, accountsAPIPath, w.StatusCode, ErrAPIFailure)
	}
	// if http.StatusNoContent response is OK
	return &types.DeleteAccountResponse{}, nil
}

// buildURL appends paths segments to url
func buildURL(url url.URL, paths []string) string {
	urlCopy := url
	for _, p := range paths {
		urlCopy.Path = path.Join(urlCopy.Path, p)
	}
	return url.String()
}

// generic error formating
func fail(method, endpoint string, err Error) error {
	return fmt.Errorf("'%s %s': %w", method, endpoint, err)
}

func failWithStatusCode(method, endpoint string, statusCode int, err Error) error {
	return fmt.Errorf("'%s %s' [%d]: %w", method, endpoint, statusCode, err)
}
