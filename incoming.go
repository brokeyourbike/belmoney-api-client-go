package belmoney

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

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
		CreationDate         time.Time    `json:"CreationDate"`
	} `json:"Payment"`
}

type CreateIncomingTransactionResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
	StatusID  int    `json:"StatusID"`
}

func (c *incomingClient) Create(ctx context.Context, transactionPayload CreateIncomingTransactionPayload) (data CreateIncomingTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Create", transactionPayload)
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

func (c *incomingClient) Status(ctx context.Context, reference string) (data IncomingTransactionsStatusesResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Statuses", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddFormParams(map[string]string{"": reference})
	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type requestCancelPayload struct {
	Reference string `json:"Reference"`
	ReasonID  int    `json:"ReasonID"` // always 0
}

type RequestCancelResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
}

func (c *incomingClient) RequestCancel(ctx context.Context, reference string) (data RequestCancelResponse, err error) {
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

func (c *incomingClient) RatesAndFeesList(ctx context.Context) (data RatesAndFeesListResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/RatesAndFeesList", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type payerNetworkListPayload struct {
	PayerID int `json:"PayerID"`
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
		LocationPoints  []string `json:"LocationPoints"`
	} `json:"Results"`
}

func (c *incomingClient) PayerNetworkList(ctx context.Context, payerID int) (data PayerNetworkListResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/PayerNetworkList", payerNetworkListPayload{PayerID: payerID})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}
