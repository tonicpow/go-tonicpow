package tonicpow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testAPIKey         = "TestAPIKey12345678987654321"
	testAdvertiserName = "TonicPow Test"
)

// TestVersion will test the method Version()
func TestVersion(t *testing.T) {
	t.Parallel()

	t.Run("get version", func(t *testing.T) {
		ver := Version()
		assert.Equal(t, version, ver)
	})
}

// TestUserAgent will test the method UserAgent()
func TestUserAgent(t *testing.T) {
	t.Parallel()

	t.Run("get user agent", func(t *testing.T) {
		agent := UserAgent()
		assert.Equal(t, defaultUserAgent, agent)
	})
}
