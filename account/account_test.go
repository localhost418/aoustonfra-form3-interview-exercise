package account_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/localhost418/aoustonfra-form3-interview-exercise/account"
	"github.com/localhost418/aoustonfra-form3-interview-exercise/account/types"
)

const realURL = "http://todo.url"

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
		{
			name: "invalid request",
			req: &types.CreateAccountRequest{
				Data: nil,
			},
			err: account.ErrInvalidBody,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			serverURL, err := url.Parse(realURL)
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

			// TODO: check response
		})
	}

}

func TestClientCreateResponse(t *testing.T) {
	tt := []struct {
		name string

		req *types.CreateAccountRequest

		status int
		res    string

		err error
	}{
		{
			name: "invalid status code ",
			req:  &types.CreateAccountRequest{Data: nil},
			err:  account.ErrAPIFailure,
		},
		{
			name: "invalid json ",
			req:  &types.CreateAccountRequest{Data: nil},
			res:  `{ "invalid-json': }`,
			err:  account.ErrInvalidResponse,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

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

			_, err = cli.CreateAccount(tc.req)
			if err == nil {
				t.Fatalf("expected an error")
			}
			if !errors.Is(err, tc.err) {
				t.Fatalf("unexpected error %v ; expected %v", err, tc.err)
			}
		})
	}
}
