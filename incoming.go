package belmoney

import (
	"context"
	"fmt"
	"net/http"
)

type IncomingClient interface {
	Create(ctx context.Context, transactionPayload CreateIncomingTransactionPayload) (CreateIncomingTransactionResponse, error)
	Status(ctx context.Context, reference string) (IncomingTransactionsStatusesResponse, error)
	Receipts(ctx context.Context, reference string) (IncomingTransactionsReceiptsResponse, error)
	RequestCancel(ctx context.Context, reference string) (RequestCancelResponse, error)
	AddSenderDocuments(ctx context.Context, documentsPayload AddSenderDocumentsPayload) (AddSenderDocumentsResponse, error)
	RatesAndFeesList(ctx context.Context) (RatesAndFeesListResponse, error)
	PayerNetworkList(ctx context.Context, payerId int) (PayerNetworkListResponse, error)
}

var _ IncomingClient = (*client)(nil)

type BaseReponse struct {
	HasErrors bool `json:"HasErrors"`
	Errors    []struct {
		ErrorCode string `json:"ErrorCode"`
		Message   string `json:"Message"`
	} `json:"Errors"`
}

type ID struct {
	IDType int `json:"IDType"`
}

type Person struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Address1    string `json:"Address1,omitempty"`
	CountryCode string `json:"CountryCode"`
	StateName   string `json:"StateName,omitempty"`
	CityName    string `json:"CityName,omitempty"`
	ZipCode     string `json:"ZipCode,omitempty"`
	DOB         string `json:"DOB,omitempty"`
	IDs         []ID   `json:"IDs"`
}

type BankAccount struct {
	AccountType int    `json:"AccountType"`
	Code        string `json:"Code"`
	AccountNo   string `json:"AccountNo"`
	Name        string `json:"Name"`
	BranchName  string `json:"BranchName"`
}

type CreateIncomingTransactionPayload struct {
	Reference        string `json:"Reference"`
	TransferReasonID int    `json:"TransferReasonID"`
	Sender           Person `json:"Sender"`
	Beneficiary      Person `json:"Beneficiary"`
	AmountAndFees    struct {
		PaymentAmount       float64 `json:"PaymentAmount"`
		OriginalAmount      float64 `json:"OriginalAmount"`
		Rate                float64 `json:"Rate"`
		RateID              int     `json:"RateID"`
		PayerCurrencyCode   string  `json:"PayerCurrencyCode"`
		PaymentCurrencyCode string  `json:"PaymentCurrencyCode"`
	} `json:"AmountAndFees"`
	Payment struct {
		PayerBranchReference string       `json:"PayerBranchReference"`
		PaymentTypeID        int          `json:"PaymentTypeID"`
		BankAccount          *BankAccount `json:"BankAccount,omitempty"`
		CreationDate         Time         `json:"CreationDate"`
	} `json:"Payment"`
}

type CreateIncomingTransactionResponse struct {
	BaseReponse
	Reference   string `json:"Reference"`
	TransferPIN string `json:"TransferPIN"`
	TransferID  string `json:"TransferID"`
	StatusID    int    `json:"StatusID"`
}

func (c *client) Create(ctx context.Context, transactionPayload CreateIncomingTransactionPayload) (data CreateIncomingTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Create", transactionPayload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type SenderDocument struct {
	TypeID         int    `json:"TypeID"`
	DocumentData   string `json:"DocumentData"`
	Description    string `json:"Description"`
	IDNo           string `json:"IDNo"`
	IssuedDate     string `json:"IssuedDate"`
	ExpirationDate string `json:"ExpirationDate"`
}

type AddSenderDocumentsPayload struct {
	TransferID string           `json:"TransferID"`
	Documents  []SenderDocument `json:"Documents"`
}

type AddSenderDocumentsResponse struct {
	BaseReponse
	TransferID string `json:"TransferID"`
}

func (c *client) AddSenderDocuments(ctx context.Context, documentsPayload AddSenderDocumentsPayload) (data AddSenderDocumentsResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/AddSenderDocuments", documentsPayload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type IncomingTransactionsStatusesResponse struct {
	BaseReponse
	Results []struct {
		Reference     string   `json:"Reference"`
		StatusID      int      `json:"StatusID"`
		HoldReasonIDs []string `json:"HoldReasonIDs"`
	} `json:"Results"`
}

func (c *client) Status(ctx context.Context, reference string) (data IncomingTransactionsStatusesResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Statuses", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddFormParams(map[string]string{"": reference})
	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type IncomingTransactionsReceiptsResponse struct {
	BaseReponse
	Results []struct {
		Reference            string `json:"Reference"`
		Identifier           string `json:"Identifier"`
		Note                 string `json:"Note"`
		DocumentData         string `json:"DocumentData"`
		DocumentDataMimeType string `json:"DocumentDataMimeType"`
		DocumentDataFilename string `json:"DocumentDataFilename"`
	} `json:"Results"`
}

func (c *client) Receipts(ctx context.Context, reference string) (data IncomingTransactionsReceiptsResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Receipts", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddFormParams(map[string]string{"": reference})
	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type RequestCancelResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
}

func (c *client) RequestCancel(ctx context.Context, reference string) (data RequestCancelResponse, err error) {
	type requestCancelPayload struct {
		Reference string `json:"Reference"`
		ReasonID  int    `json:"ReasonID"` // always 0
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/RequestCancel", requestCancelPayload{Reference: reference})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type RatesAndFeesListResponse struct {
	BaseReponse
	Results []struct {
		CountryCode         string  `json:"CountryCode"`
		CountryName         string  `json:"CountryName"`
		PayerID             int     `json:"PayerID"`
		PayerName           string  `json:"PayerName"`
		PayerBranchID       int     `json:"PayerBranchID"`
		PayerBranchName     string  `json:"PayerBranchName"`
		CurrencyCode        string  `json:"CurrencyCode"`
		CurrencyTypeName    string  `json:"CurrencyTypeName"`
		PaymentTypeID       int     `json:"PaymentTypeID"`
		PaymentTypeName     string  `json:"PaymentTypeName"`
		FromAmount          float64 `json:"FromAmount"`
		ToAmount            float64 `json:"ToAmount"`
		PercentageFee       float64 `json:"PercentageFee"`
		FlatFee             float64 `json:"FlatFee"`
		RateTypeID          int     `json:"RateTypeID"`
		RateTypeDescription string  `json:"RateTypeDescription"`
		Rate                float64 `json:"Rate"`
		FromCurrencyCode    string  `json:"FromCurrencyCode"`
	} `json:"Results"`
}

func (c *client) RatesAndFeesList(ctx context.Context) (data RatesAndFeesListResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/RatesAndFeesList", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type PayerNetworkListResponse struct {
	BaseReponse
	Results []struct {
		PayerBranchID   int      `json:"PayerBranchID"`
		PayerBranchName string   `json:"PayerBranchName"`
		Address1        string   `json:"Address1"`
		Address2        string   `json:"Address2"`
		CityName        string   `json:"CityName"`
		StateCode       string   `json:"StateCode"`
		CountryCode     string   `json:"CountryCode"`
		PhoneNumber     string   `json:"PhoneNumber"`
		PayAllCities    bool     `json:"PayAllCities"`
		HasLocations    bool     `json:"HasLocations"`
		PaymentTypes    []int    `json:"PaymentTypes"`
		Currencies      []string `json:"Currencies"`
	} `json:"Results"`
}

func (c *client) PayerNetworkList(ctx context.Context, payerID int) (data PayerNetworkListResponse, err error) {
	type payerNetworkListPayload struct {
		PayerID int `json:"PayerID"`
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/PayerNetworkList", payerNetworkListPayload{PayerID: payerID})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}
