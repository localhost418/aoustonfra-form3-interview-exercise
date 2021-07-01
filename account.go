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

// NewClient creates a new Client (*http.Client and api URL)
func NewClient(client *http.Client, url url.URL) *Client {
	return &Client{
		client: client,
		url:    url,
	}
}

// CreateAccount creates an account with the fields declared in the request
func (c *Client) CreateAccount(req *types.CreateAccountRequest) (*types.CreateAccountResponse, *AccountError) {
	const method = http.MethodPost
	if req == nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrNoRequest), -1, nil)
	}

	body := &bytes.Buffer{}
	_, err := req.WriteTo(body)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrInvalidBody), -1, &err)
	}

	r, err := http.NewRequest(method, buildURL(c.url, []string{accountsAPIPath}), body)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrInvalidRequest), -1, &err)
	}
	r.Header.Add("Accept", "application/vnd.api+json")

	w, err := c.client.Do(r)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrDoRequest), -1, &err)
	}
	defer w.Body.Close()

	status := w.StatusCode
	if status != http.StatusCreated {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrAPIFailure), status, nil)
	}
	res := &types.CreateAccountResponse{}
	_, err = res.ReadFrom(w.Body)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrInvalidResponse), status, &err)
	}
	return res, nil
}

// FetchAccount fetch an account by accountID
func (c *Client) FetchAccount(req *types.FetchAccountRequest) (*types.FetchAccountResponse, *AccountError) {
	const method = http.MethodGet
	if req == nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrNoRequest), -1, nil)
	}

	r, err := http.NewRequest(method, buildURL(c.url, []string{accountsAPIPath, req.AccountID.String()}), nil)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrInvalidRequest), -1, &err)
	}
	r.Header.Add("Accept", "application/vnd.api+json")

	w, err := c.client.Do(r)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrDoRequest), -1, &err)
	}
	defer w.Body.Close()

	status := w.StatusCode
	if status != http.StatusOK {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrAPIFailure), status, nil)
	}
	res := &types.FetchAccountResponse{}
	_, err = res.ReadFrom(w.Body)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrInvalidResponse), status, &err)
	}

	return res, nil
}

// DeleteAccount deletes an account by accountID and version
func (c *Client) DeleteAccount(req *types.DeleteAccountRequest) (*types.DeleteAccountResponse, *AccountError) {
	const method = http.MethodDelete
	if req == nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrNoRequest), -1, nil)
	}

	r, err := http.NewRequest(method, buildURL(c.url, []string{accountsAPIPath, req.AccountID.String()}), nil)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrInvalidRequest), -1, &err)
	}

	r.Header.Add("Accept", "application/vnd.api+json")
	q := r.URL.Query()
	q.Add("version", strconv.Itoa(req.Version))
	r.URL.RawQuery = q.Encode()

	w, err := c.client.Do(r)
	if err != nil {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrDoRequest), -1, &err)
	}
	defer w.Body.Close()

	status := w.StatusCode
	if status != http.StatusNoContent {
		return nil, NewAccountError(fail(method, accountsAPIPath, ErrAPIFailure), status, nil)
	}
	return &types.DeleteAccountResponse{}, nil
}

// buildURL appends paths segments to url
func buildURL(url url.URL, paths []string) string {
	for _, p := range paths {
		url.Path = path.Join(url.Path, p)
	}
	return url.String()
}

// generic error formating
func fail(method, endpoint string, msg string) string {
	return fmt.Sprintf("'%s %s': %s", method, endpoint, msg)
}
