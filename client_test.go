package belmoney

import (
	"testing"

	logrustest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	httpClient := NewMockHttpClient(t)
	logger, _ := logrustest.NewNullLogger()

	client := NewOutClient("", "", "", WithHTTPClient(httpClient), WithLogger(logger))

	assert.Same(t, httpClient, client.httpClient)
	assert.Same(t, logger, client.logger)
}
