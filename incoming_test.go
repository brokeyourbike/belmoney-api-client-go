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

//go:embed testdata/Incoming_Create_Errors.json
var incomingCreateErrors []byte

//go:embed testdata/Incoming_Create_Success.json
var incomingCreateSuccess []byte

//go:embed testdata/Incoming_AddSenderDocuments_Success.json
var incomingAddSenderDocumentsSuccess []byte

//go:embed testdata/Incoming_Statuses_Success.json
var incomingStatusesSuccess []byte

//go:embed testdata/Incoming_Receipts_Success.json
var incomingReceiptsSuccess []byte

//go:embed testdata/Incoming_RatesAndFeesList_Success.json
var incomingRatesAndFeesListSuccess []byte

//go:embed testdata/Incoming_PayerNetworkList_Success.json
var incomingPayerNetworkListSuccess []byte

//go:embed testdata/Incoming_PayerNetworkList_Errors.json
var incomingPayerNetworkListErrors []byte

func TestCreate_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)

	logger, hook := logrustest.NewNullLogger()
	logger.SetLevel(logrus.DebugLevel)

	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient), belmoney.WithLogger(logger))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingCreateSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Create(context.TODO(), belmoney.CreateIncomingTransactionPayload{})
	require.NoError(t, err)

	assert.Equal(t, belmoney.StatusIdCreated, got.StatusID)
	assert.Equal(t, "901", got.TransferPIN)
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
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingCreateErrors))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Create(context.TODO(), belmoney.CreateIncomingTransactionPayload{})
	require.NoError(t, err)

	assert.Equal(t, belmoney.StatusId(0), got.StatusID)
	assert.True(t, got.HasErrors)
}

func TestCreate_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.Create(nil, belmoney.CreateIncomingTransactionPayload{}) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestAddSenderDocuments_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingAddSenderDocumentsSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.AddSenderDocuments(context.TODO(), belmoney.AddSenderDocumentsPayload{})
	require.NoError(t, err)

	assert.Equal(t, "12345", got.TransferID)
	assert.False(t, got.HasErrors)
}

func TestAddSenderDocuments_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.AddSenderDocuments(nil, belmoney.AddSenderDocumentsPayload{}) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestStatus_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingStatusesSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Status(context.TODO(), "12345")
	require.NoError(t, err)

	assert.False(t, got.HasErrors)
	require.Equal(t, 1, len(got.Results))
}

func TestStatus_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.Status(nil, "") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestReceipts_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingReceiptsSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.Receipts(context.TODO(), "12345")
	require.NoError(t, err)

	assert.False(t, got.HasErrors)
	require.Equal(t, 1, len(got.Results))
}

func TestReceipts_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.Receipts(nil, "") //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestRatesAndFeesList_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingRatesAndFeesListSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.RatesAndFeesList(context.TODO())
	require.NoError(t, err)

	assert.False(t, got.HasErrors)
	require.Equal(t, 1, len(got.Results))
}

func TestRatesAndFeesList_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.RatesAndFeesList(nil) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}

func TestPayerNetworkList_Success(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingPayerNetworkListSuccess))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.PayerNetworkList(context.TODO(), 0)
	require.NoError(t, err)

	assert.False(t, got.HasErrors)
	require.Equal(t, 1, len(got.Results))
}

func TestPayerNetworkList_Errors(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	resp := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(incomingPayerNetworkListErrors))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.PayerNetworkList(context.TODO(), 0)
	require.NoError(t, err)

	assert.True(t, got.HasErrors)
	require.Equal(t, 1, len(got.Errors))
	require.Equal(t, 0, len(got.Results))
}

func TestPayerNetworkList_RequestErr(t *testing.T) {
	mockHttpClient := belmoney.NewMockHttpClient(t)
	client := belmoney.NewClient("baseurl", "client_id", "client_secret", belmoney.WithHTTPClient(mockHttpClient))

	_, err := client.PayerNetworkList(nil, 0) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to create request")
}
