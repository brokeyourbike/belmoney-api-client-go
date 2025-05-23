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
	PreRegister(ctx context.Context, payload PreRegisterIncomingTransactionPayload) (PreRegisterIncomingTransactionResponse, error)
	Confirm(ctx context.Context, payload ConfirmIncomingTransactionPayload) (ConfirmIncomingTransactionResponse, error)
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
	IDType            PersonIdType `json:"IDType"`
	IDCopy            bool         `json:"IDCopy,omitempty"`
	IDNo              string       `json:"IDNo,omitempty"`
	Authority         string       `json:"Authority,omitempty"`
	IssuedCountryCode string       `json:"IssuedCountryCode,omitempty"`
	IDIssuedDate      string       `json:"IDIssuedDate,omitempty"`
	IDExpirationDate  string       `json:"IDExpirationDate,omitempty"`
}

type Sex string

const (
	SexMale   Sex = "M"
	SexFemale Sex = "F"
)

type Person struct {
	FirstName              string `json:"FirstName"`
	MiddleName             string `json:"MiddleName,omitempty"`
	LastName               string `json:"LastName"`
	SecondLastName         string `json:"SecondLastName,omitempty"`
	Address1               string `json:"Address1,omitempty"`
	CountryCode            string `json:"CountryCode"`
	StateName              string `json:"StateName,omitempty"`
	CityName               string `json:"CityName,omitempty"`
	ZipCode                string `json:"ZipCode,omitempty"`
	DOB                    string `json:"DOB,omitempty"`
	PhoneNumber            string `json:"PhoneNumber,omitempty"`
	PrimaryPhoneNumber     string `json:"PrimaryPhoneNumber,omitempty"`
	PhoneNumberCountryCode string `json:"PhoneNumberCountryCode,omitempty"`
	Email                  string `json:"Email,omitempty"`
	Sex                    Sex    `json:"Sex,omitempty"`
	BirthCityName          string `json:"BirthCityName,omitempty"`
	BirthCountryCode       string `json:"BirthCountryCode,omitempty"`
	CitizenshipCountryCode string `json:"CitizenshipCountryCode,omitempty"`
	AgencyReference        string `json:"AgencyReference,omitempty"`
	ProfessionID           int    `json:"ProfessionID,omitempty"`
	PEPTypeID              int    `json:"PEPTypeID,omitempty"`
	SenderTypeID           int    `json:"SenderTypeID,omitempty"`
	RelationshipToSenderID int    `json:"RelationshipToSenderID,omitempty"`
	IDs                    []ID   `json:"IDs"`
}

type BankAccount struct {
	AccountType AccountTypeId `json:"AccountType"`
	AccountNo   string        `json:"AccountNo"`
	Code        string        `json:"Code"`
	Name        string        `json:"Name"`
	BranchCode  string        `json:"BranchCode"`
	BranchName  string        `json:"BranchName"`
}

type CreateIncomingTransactionPayload struct {
	Reference        string           `json:"Reference"`
	TransferReasonID TransferReasonId `json:"TransferReasonID"`
	Sender           Person           `json:"Sender"`
	Beneficiary      Person           `json:"Beneficiary"`
	AmountAndFees    struct {
		PaymentAmount       float64    `json:"PaymentAmount"`
		OriginalAmount      float64    `json:"OriginalAmount"`
		Rate                float64    `json:"Rate"`
		RateID              RateTypeId `json:"RateID"`
		PayerCurrencyCode   string     `json:"PayerCurrencyCode"`
		PaymentCurrencyCode string     `json:"PaymentCurrencyCode"`
		PercentFee          float64    `json:"PercentFee"`
		FlatFee             float64    `json:"FlatFee"`
		OtherFee            float64    `json:"OtherFee"`
		Tax                 float64    `json:"Tax"`
		FeesTax             float64    `json:"FeesTax"`
		Discount            float64    `json:"Discount"`
	} `json:"AmountAndFees"`
	Payment struct {
		PayerBranchReference    string        `json:"PayerBranchReference"`
		PaymentTypeID           PaymentTypeId `json:"PaymentTypeID"`
		PaymentProcessorCode    string        `json:"PaymentProcessorCode,omitempty"`
		PaymentConfirmationCode string        `json:"PaymentConfirmationCode,omitempty"`
		BankAccount             *BankAccount  `json:"BankAccount,omitempty"`
		CreationDate            Time          `json:"CreationDate"`
	} `json:"Payment"`
}

type CreateIncomingTransactionResponse struct {
	BaseReponse
	Reference   string   `json:"Reference"`
	TransferPIN string   `json:"TransferPIN"`
	TransferID  string   `json:"TransferID"`
	StatusID    StatusId `json:"StatusID"`
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
	TypeID         DocumentTypeId `json:"TypeID"`
	DocumentData   string         `json:"DocumentData"`
	Description    string         `json:"Description"`
	IDNo           string         `json:"IDNo"`
	IssuedDate     string         `json:"IssuedDate"`
	ExpirationDate string         `json:"ExpirationDate"`
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
		StatusID      StatusId `json:"StatusID"`
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
		Reference string         `json:"Reference"`
		ReasonID  CancelReasonId `json:"ReasonID"` // always 0
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

type PreRegisterIncomingTransactionPayload struct {
	Reference        string           `json:"Reference"`
	TransferID       string           `json:"TransferID"`
	TransferPIN      string           `json:"TransferPIN"`
	TransferReasonID TransferReasonId `json:"TransferReasonID"`
	Sender           Person           `json:"Sender"`
	Beneficiary      Person           `json:"Beneficiary"`
	AmountAndFees    struct {
		PaymentAmount       float64    `json:"PaymentAmount"`
		OriginalAmount      float64    `json:"OriginalAmount"`
		Rate                float64    `json:"Rate"`
		RateID              RateTypeId `json:"RateID"`
		PayerCurrencyCode   string     `json:"PayerCurrencyCode"`
		PaymentCurrencyCode string     `json:"PaymentCurrencyCode"`
		PercentFee          float64    `json:"PercentFee"`
		FlatFee             float64    `json:"FlatFee"`
		OtherFee            float64    `json:"OtherFee"`
		Tax                 float64    `json:"Tax"`
		FeesTax             float64    `json:"FeesTax"`
		Discount            float64    `json:"Discount"`
	} `json:"AmountAndFees"`
	Payment struct {
		PayerBranchReference    string        `json:"PayerBranchReference"`
		PaymentTypeID           PaymentTypeId `json:"PaymentTypeID"`
		PaymentProcessorCode    string        `json:"PaymentProcessorCode,omitempty"`
		PaymentConfirmationCode string        `json:"PaymentConfirmationCode,omitempty"`
		BankAccount             *BankAccount  `json:"BankAccount,omitempty"`
		CreationDate            Time          `json:"CreationDate"`
	} `json:"Payment"`
}

type PreRegisterIncomingTransactionResponse struct {
	BaseReponse
	Reference   string   `json:"Reference"`
	TransferPIN string   `json:"TransferPIN"`
	TransferID  string   `json:"TransferID"`
	StatusID    StatusId `json:"StatusID"`
}

func (c *client) PreRegister(ctx context.Context, payload PreRegisterIncomingTransactionPayload) (data PreRegisterIncomingTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/PreRegister", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type ConfirmIncomingTransactionPayload struct {
	Reference       string `json:"Reference"`
	TransferPayment struct {
		PaymentTypeID           PaymentTypeId `json:"PaymentTypeID"`
		PaymentProcessorCode    string        `json:"PaymentProcessorCode"`
		PaymentConfirmationCode string        `json:"PaymentConfirmationCode"`
	} `json:"TransferPayment"`
}

type ConfirmIncomingTransactionResponse struct {
	BaseReponse
	Reference   string   `json:"Reference"`
	TransferPIN string   `json:"TransferPIN"`
	TransferID  string   `json:"TransferID"`
	StatusID    StatusId `json:"StatusID"`
}

func (c *client) Confirm(ctx context.Context, payload ConfirmIncomingTransactionPayload) (data ConfirmIncomingTransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/Confirm", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}
