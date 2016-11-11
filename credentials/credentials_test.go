package credentials

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubProvider struct {
	creds Value
	err   error
}

func (s *stubProvider) Retrieve() (Value, error) {
	return s.creds, s.err
}

func TestCredentials(t *testing.T) {

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		c := NewCredentials(&stubProvider{
			creds: Value{
				ClientID:     "client-id",
				ClientSecret: "client-secret",
			},
		})

		creds, err := c.Get()
		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, "client-id", creds.ClientID)
		assert.Equal(t, "client-secret", creds.ClientSecret)
	})

	t.Run("GetWithError", func(t *testing.T) {
		t.Parallel()

		c := NewCredentials(&stubProvider{
			err: errors.New("provider error"),
		})

		creds, err := c.Get()
		assert.Error(t, err)
		assert.Equal(t, "", creds.ClientID)
		assert.Equal(t, "", creds.ClientSecret)
	})
}
