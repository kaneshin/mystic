package credentials

import "errors"

var (
	// ErrStaticCredentialsEmpty is emitted when static credentials are empty.
	ErrStaticCredentialsEmpty = errors.New("static credentials are empty")
)

// A StaticProvider is a set of credentials which are set pragmatically,
// and will never expire.
type StaticProvider struct {
	Value
}

// NewStaticCredentials returns a pointer to a new Credentials object
// wrapping a static credentials value provider.
func NewStaticCredentials(clientID, clientSecret string) *Credentials {
	return NewCredentials(&StaticProvider{
		Value: Value{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	})
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s *StaticProvider) Retrieve() (Value, error) {
	if s.ClientID == "" || s.ClientSecret == "" {
		return Value{}, ErrStaticCredentialsEmpty
	}

	return s.Value, nil
}
