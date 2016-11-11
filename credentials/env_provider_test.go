package credentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvProvider(t *testing.T) {

	t.Run("Get", func(t *testing.T) {
		os.Clearenv()
		os.Setenv("GOOGLE_API_CLIENT_ID", "client-id")
		os.Setenv("GOOGLE_API_CLIENT_SECRET", "client-secret")

		c := NewEnvCredentials()

		creds, err := c.Get()
		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, "client-id", creds.ClientID)
		assert.Equal(t, "client-secret", creds.ClientSecret)
	})

	t.Run("GetWithError", func(t *testing.T) {
		candidates := []struct {
			id, secret string
		}{
			{"", ""},
			{"client-id", ""},
			{"", "client-secret"},
		}

		for _, candidate := range candidates {
			candidate := candidate
			os.Clearenv()
			os.Setenv("GOOGLE_API_CLIENT_ID", candidate.id)
			os.Setenv("GOOGLE_API_CLIENT_SECRET", candidate.secret)

			c := NewEnvCredentials()

			creds, err := c.Get()
			assert.Error(t, err)
			assert.Equal(t, "", creds.ClientID)
			assert.Equal(t, "", creds.ClientSecret)
		}
	})
}
