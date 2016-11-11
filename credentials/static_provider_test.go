package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticProvider(t *testing.T) {

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		c := NewStaticCredentials("client-id", "client-secret")

		creds, err := c.Get()
		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, "client-id", creds.ClientID)
		assert.Equal(t, "client-secret", creds.ClientSecret)
	})

	t.Run("GetWithError", func(t *testing.T) {
		t.Parallel()

		c := NewStaticCredentials("", "")

		creds, err := c.Get()
		assert.Error(t, err)
		assert.Equal(t, "", creds.ClientID)
		assert.Equal(t, "", creds.ClientSecret)
	})
}
