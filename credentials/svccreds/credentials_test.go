package svccreds

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
				Type:                    "service_account",
				ProjectID:               "project-id",
				PrivateKeyID:            "some_number",
				PrivateKey:              "-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n",
				ClientEmail:             "visionapi@project-id.iam.gserviceaccount.com",
				ClientID:                "...",
				AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
				TokenURI:                "https://accounts.google.com/o/oauth2/token",
				AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
				ClientX509CertURL:       "https://www.googleapis.com/robot/v1/metadata/x509/visionapi%40project-id.iam.gserviceaccount.com",
			},
		})

		creds, err := c.Get()
		assert.NoError(t, err, "Expected no error")
		assert.Equal(t, "service_account", creds.Type)
		assert.Equal(t, "project-id", creds.ProjectID)
		assert.Equal(t, "some_number", creds.PrivateKeyID)
		assert.Equal(t, "-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n", creds.PrivateKey)
		assert.Equal(t, "visionapi@project-id.iam.gserviceaccount.com", creds.ClientEmail)
		assert.Equal(t, "...", creds.ClientID)
		assert.Equal(t, "https://accounts.google.com/o/oauth2/auth", creds.AuthURI)
		assert.Equal(t, "https://accounts.google.com/o/oauth2/token", creds.TokenURI)
		assert.Equal(t, "https://www.googleapis.com/oauth2/v1/certs", creds.AuthProviderX509CertURL)
		assert.Equal(t, "https://www.googleapis.com/robot/v1/metadata/x509/visionapi%40project-id.iam.gserviceaccount.com", creds.ClientX509CertURL)
	})

	t.Run("GetWithError", func(t *testing.T) {
		t.Parallel()

		c := NewCredentials(&stubProvider{err: errors.New("provider error")})

		creds, err := c.Get()
		assert.Error(t, err)
		assert.Equal(t, "", creds.Type)
		assert.Equal(t, "", creds.ProjectID)
		assert.Equal(t, "", creds.PrivateKeyID)
		assert.Equal(t, "", creds.PrivateKey)
		assert.Equal(t, "", creds.ClientEmail)
		assert.Equal(t, "", creds.ClientID)
		assert.Equal(t, "", creds.AuthURI)
		assert.Equal(t, "", creds.TokenURI)
		assert.Equal(t, "", creds.AuthProviderX509CertURL)
		assert.Equal(t, "", creds.ClientX509CertURL)
	})
}
