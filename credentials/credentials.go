// Package credentials provides credential retrieval and management
package credentials

import (
	"sync"
)

// A Value is the OAuth client ID credentials value for individual credential
// fields.
type Value struct {
	// OAuth Client ID
	ClientID string

	// OAuth Client Secret
	ClientSecret string
}

// A Provider is the interface for any component which will provide credentials
// Value.
type Provider interface {
	// Refresh returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Retrieve() (Value, error)
}

// A Credentials provides synchronous safe retrieval of OAuth client ID
// credentials Value.
type Credentials struct {
	creds    Value
	mu       sync.Mutex
	provider Provider
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider Provider) *Credentials {
	return &Credentials{
		provider: provider,
	}
}

// Get returns the credentials value, or error if the credentials Value failed
// to be retrieved.
func (c *Credentials) Get() (Value, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.provider.Retrieve()
}
