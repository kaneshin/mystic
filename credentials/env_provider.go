package credentials

import (
	"os"

	"errors"
)

var (
	// ErrClientIDNotFound is returned when the Client ID can't be
	// found in the process's environment.
	ErrClientIDNotFound = errors.New("GOOGLE_API_CLIENT_ID not found in environment")

	// ErrClientSecretNotFound is returned when the Client Secret can't be
	// found in the process's environment.
	ErrClientSecretNotFound = errors.New("GOOGLE_API_CLIENT_SECRET not found in environment")
)

// A EnvProvider retrieves credentials from the environment variables of the
// running process.
//
// Environment variables used:
//
// * Client ID:     GOOGLE_API_CLIENT_ID
// * Client Secret: GOOGLE_API_CLIENT_SECRET
type EnvProvider struct {
	retrieved bool
}

// NewEnvCredentials returns a pointer to a new Credentials object
// wrapping the environment variable provider.
func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	e.retrieved = false

	id := os.Getenv("GOOGLE_API_CLIENT_ID")
	if id == "" {
		return Value{}, ErrClientIDNotFound
	}

	secret := os.Getenv("GOOGLE_API_CLIENT_SECRET")
	if secret == "" {
		return Value{}, ErrClientSecretNotFound
	}

	e.retrieved = true
	return Value{
		ClientID:     id,
		ClientSecret: secret,
	}, nil
}
