package accountclient

// AccountError for account errors
type AccountError struct {
	Message    string
	StatusCode *int
	Error      *error
}

// NewAccountError make a new AccountError from message and status
func NewAccountError(msg string, status int, err *error) *AccountError {
	var s *int
	if status != -1 {
		s = &status
	}
	return &AccountError{
		Message:    msg,
		StatusCode: s,
		Error:      err,
	}
}

const (
	// ErrNoRequest on missing request
	ErrNoRequest = "no request provided"

	// ErrInvalidBody on invalid request body
	ErrInvalidBody = "invalid request body"

	// ErrInvalidRequest on invalid request
	ErrInvalidRequest = "invalid request"

	// ErrDoRequest on request failed
	ErrDoRequest = "request failed"

	// ErrAPIFailure on API failure
	ErrAPIFailure = "api failure"

	// ErrInvalidResponse on invalid response
	ErrInvalidResponse = "invalid response"
)
