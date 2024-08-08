package belmoney_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/belmoney-api-client-go"
	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/IncomingTrx_Response_WithErrors.json
var outResponseWithErrors []byte

//go:embed testdata/IncomingTrx_Response_Success.json
var outResponseSuccess []byte

func TestCreate_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)

	logger, hook := logrustest.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	client := belmoney.NewIncomingClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient), belmoney.WithLogger(logger))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(outResponseSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Create(context.TODO(), belmoney.CreateIncomingTransactionPayload{})
	require.NoError(t, err)

	assert.Equal(t, 1, got.StatusID)
	assert.False(t, got.HasErrors)

	require.Equal(t, 2, len(hook.Entries))
	require.Contains(t, hook.Entries[0].Data, "http.request.method")
	require.Contains(t, hook.Entries[0].Data, "http.request.url")
	require.Contains(t, hook.Entries[0].Data, "http.request.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.status_code")
	require.Contains(t, hook.Entries[1].Data, "http.response.body.content")
	require.Contains(t, hook.Entries[1].Data, "http.response.headers")
}

func TestCreate_Errors(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewIncomingClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(outResponseWithErrors))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Create(context.TODO(), belmoney.CreateIncomingTransactionPayload{})
	require.NoError(t, err)

	assert.Equal(t, 0, got.StatusID)
	assert.True(t, got.HasErrors)
}

func TestCreateReservedAccount_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewIncomingClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.Create(nil, belmoney.CreateIncomingTransactionPayload{}) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}
