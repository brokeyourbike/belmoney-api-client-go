package belmoney

import (
	"context"
	"fmt"
	"net/http"
)

type ID struct {
	IDType string `json:"IDType"`
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
		PaymentAmount       string `json:"PaymentAmount"`
		OriginalAmount      string `json:"OriginalAmount"`
		Rate                string `json:"Rate"`
		RateID              int    `json:"RateID"`
		PayerCurrencyCode   string `json:"PayerCurrencyCode"`
		PaymentCurrencyCode string `json:"PaymentCurrencyCode"`
	} `json:"AmountAndFees"`
	Payment struct {
		PayerBranchReference string       `json:"PayerBranchReference"`
		PaymentTypeID        string       `json:"PaymentTypeID"`
		BankAccount          *BankAccount `json:"BankAccount,omitempty"`
		CreationDate         string       `json:"CreationDate"`
	} `json:"Payment"`
}

type CreateIncomingTransactionResponse struct {
	Reference string `json:"Reference"`
	StatusID  int    `json:"StatusID"`
	HasErrors bool   `json:"HasErrors"`
	Errors    []struct {
		ErrorCode string `json:"ErrorCode"`
		Message   string `json:"Message"`
	} `json:"Errors"`
}

func (c *incomingClient) Create(ctx context.Context, transactionPayload CreateIncomingTransactionPayload) (data CreateIncomingTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Create", transactionPayload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type IncomingTransactionsStatusesResponse struct {
	HasErrors bool `json:"HasErrors"`
	Errors    []struct {
		ErrorCode string `json:"ErrorCode"`
		Message   string `json:"Message"`
	} `json:"Errors"`
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
	return data, c.do(ctx, req)
}
