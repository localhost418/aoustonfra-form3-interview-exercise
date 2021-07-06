package account

// Error for account errors
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrNoRequest on missing request
	ErrNoRequest = Error("no request provided")

	// ErrInvalidBody
	ErrInvalidBody = Error("invalid request body")

	// ErrInvalidRequest
	ErrInvalidRequest = Error("invalid request")

	// ErrDoRequest
	ErrDoRequest = Error("request failed")

	// ErrAPIFailure
	ErrAPIFailure = Error("api failure")

	// ErrInvalidResponse
	ErrInvalidResponse = Error("invalid response")
)
