package belmoney_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/belmoney-api-client-go"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/Outgoing_TransactionsList_Success.json
var outgoingTransactionsListSuccess []byte

//go:embed testdata/Outgoing_Transaction_Success.json
var outgoingTransactionSuccess []byte

func TestTransactionsList_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(outgoingTransactionsListSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.TransactionsList(context.TODO())
	require.NoError(t, err)

	assert.False(t, got.HasErrors)
	require.Equal(t, 0, len(got.Errors))
	require.Equal(t, 3, len(got.References))
}

func TestTransactionsList_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.TransactionsList(nil) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestTransaction_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(outgoingTransactionSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Transaction(context.TODO(), "reference")
	require.NoError(t, err)

	assert.False(t, got.HasErrors)
	require.Equal(t, 0, len(got.Errors))
}

func TestTransaction_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.Transaction(nil, "reference") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}
