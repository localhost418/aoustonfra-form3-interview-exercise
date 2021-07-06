package accountclient

// Error for account errors
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNoRequest on missing request
	ErrNoRequest = Error("no request provided")

	// ErrInvalidBody on invalid request body
	ErrInvalidBody = Error("invalid request body")

	// ErrInvalidRequest on invalid request
	ErrInvalidRequest = Error("invalid request")

	// ErrDoRequest on request failed
	ErrDoRequest = Error("request failed")

	// ErrAPIFailure on API failure
	ErrAPIFailure = Error("api failure")

	// ErrInvalidResponse on invalid response
	ErrInvalidResponse = Error("invalid response")
)
