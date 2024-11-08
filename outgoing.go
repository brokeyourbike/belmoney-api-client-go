package belmoney

import (
	"context"
	"fmt"
	"net/http"
)

type OutgoingClient interface {
	TransactionsList(ctx context.Context) (TransactionsListResponse, error)
	Transaction(ctx context.Context, reference string) (TransactionResponse, error)
	Processing(ctx context.Context, reference string) (ProcessingResponse, error)
	Paid(ctx context.Context, reference, note string) (PaidResponse, error)
	Cancel(ctx context.Context, reference, note string) (CancelResponse, error)
	UpdateRate(ctx context.Context, reference, note string, rate float64) (UpdateRateResponse, error)
}

var _ OutgoingClient = (*client)(nil)

type TransactionsListResponse struct {
	BaseReponse
	References []string `json:"References"`
}

func (c *client) TransactionsList(ctx context.Context) (data TransactionsListResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, "/NewList", nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type TransactionResponse struct {
	BaseReponse
	Transfer struct {
		TransferID        string           `json:"TransferID"`
		Reference         string           `json:"Reference"`
		TransferPIN       string           `json:"TransferPIN"`
		TransConfirmation string           `json:"TransConfirmation"`
		TransferReasonID  TransferReasonId `json:"TransferReasonID"`
		Sender            struct {
			AgencyReference        string `json:"AgencyReference"`
			FirstName              string `json:"FirstName"`
			MiddleName             string `json:"MiddleName"`
			LastName               string `json:"LastName"`
			SecondLastName         string `json:"SecondLastName"`
			Address1               string `json:"Address1"`
			CountryCode            string `json:"CountryCode"`
			StateCode              string `json:"StateCode"`
			StateName              string `json:"StateName"`
			CityCode               int    `json:"CityCode"`
			CityName               string `json:"CityName"`
			ZipCode                string `json:"ZipCode"`
			DOB                    *Time  `json:"DOB"`
			PhoneNumber            string `json:"PhoneNumber"`
			CellPhoneNumber        string `json:"CellPhoneNumber"`
			PrimaryPhoneNumber     string `json:"PrimaryPhoneNumber"`
			PhoneNumberCountryCode string `json:"PhoneNumberCountryCode"`
			Email                  string `json:"Email"`
			Sex                    string `json:"Sex"`
			BirthCityName          string `json:"BirthCityName"`
			BirthCountryCode       string `json:"BirthCountryCode"`
			CitizenshipCountryCode string `json:"CitizenshipCountryCode"`
			ProfessionID           int    `json:"ProfessionID"`
			PEPTypeID              int    `json:"PEPTypeID"`
			SenderTypeID           int    `json:"SenderTypeID"`
			IDs                    []struct {
				IDIssuedDate      *Time        `json:"IDIssuedDate"`
				IDExpirationDate  *Time        `json:"IDExpirationDate"`
				IssuedCountryCode string       `json:"IssuedCountryCode"`
				Authority         string       `json:"Authority"`
				IDCopy            bool         `json:"IDCopy"`
				IDType            PersonIdType `json:"IDType"`
				IDNo              string       `json:"IDNo"`
			} `json:"IDs"`
		} `json:"Sender"`
		Beneficiary struct {
			AgencyReference        string         `json:"AgencyReference"`
			FirstName              string         `json:"FirstName"`
			MiddleName             string         `json:"MiddleName"`
			LastName               string         `json:"LastName"`
			SecondLastName         string         `json:"SecondLastName"`
			Address1               string         `json:"Address1"`
			CountryCode            string         `json:"CountryCode"`
			StateCode              string         `json:"StateCode"`
			CityCode               int            `json:"CityCode"`
			CityName               string         `json:"CityName"`
			ZipCode                string         `json:"ZipCode"`
			PhoneNumber            string         `json:"PhoneNumber"`
			CellPhoneNumber        string         `json:"CellPhoneNumber"`
			PrimaryPhoneNumber     string         `json:"PrimaryPhoneNumber"`
			Email                  string         `json:"Email"`
			RelationshipToSenderID RelationTypeId `json:"RelationshipToSenderID"`
			IDs                    []struct {
				IDCopy bool         `json:"IDCopy"`
				IDType PersonIdType `json:"IDType"`
				IDNo   string       `json:"IDNo"`
			} `json:"IDs"`
		} `json:"Beneficiary"`
		AmountAndFees struct {
			PaymentAmount       float64 `json:"PaymentAmount"`
			PaymentCurrencyCode string  `json:"PaymentCurrencyCode"`
			OriginalAmount      float64 `json:"OriginalAmount"`
			Rate                float64 `json:"Rate"`
			RateID              int     `json:"RateID"`
			PayerCurrencyCode   string  `json:"PayerCurrencyCode"`
			PercentFee          float64 `json:"PercentFee"`
			FlatFee             float64 `json:"FlatFee"`
			OtherFee            float64 `json:"OtherFee"`
			Tax                 float64 `json:"Tax"`
			FeesTax             float64 `json:"FeesTax"`
			Discount            float64 `json:"Discount"`
		} `json:"AmountAndFees"`
		Payment struct {
			PayerBranchReference string        `json:"PayerBranchReference"`
			PaymentTypeID        PaymentTypeId `json:"PaymentTypeID"`
			LocationCode         string        `json:"LocationCode"`
			BankAccount          struct {
				Code        string        `json:"Code"`
				Name        string        `json:"Name"`
				BranchCode  string        `json:"BranchCode"`
				BranchName  string        `json:"BranchName"`
				AccountType AccountTypeId `json:"AccountType"`
				AccountNo   string        `json:"AccountNo"`
			} `json:"BankAccount"`
		} `json:"Payment"`
		Notes        string `json:"Notes"`
		CreationDate Time   `json:"CreationDate"`
	} `json:"Transfer"`
}

func (c *client) Transaction(ctx context.Context, reference string) (data TransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/New/%s", reference), nil)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type ProcessingResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
}

func (c *client) Processing(ctx context.Context, reference string) (data ProcessingResponse, err error) {
	type processingPayload struct {
		Reference string `json:"Reference"`
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/Processing", processingPayload{Reference: reference})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type PaidResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
}

func (c *client) Paid(ctx context.Context, reference, note string) (data PaidResponse, err error) {
	type paidPayload struct {
		Reference string `json:"Reference"`
		Note      string `json:"Note"`
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/Paid", paidPayload{Reference: reference, Note: note})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type CancelResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
}

func (c *client) Cancel(ctx context.Context, reference, note string) (data CancelResponse, err error) {
	type cancelPayload struct {
		Reference string `json:"Reference"`
		Note      string `json:"Note"`
		ReasonID  int    `json:"ReasonID"`
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/Cancel", cancelPayload{Reference: reference, Note: note})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}

type UpdateRateResponse struct {
	BaseReponse
	Reference string `json:"Reference"`
}

func (c *client) UpdateRate(ctx context.Context, reference, note string, rate float64) (data UpdateRateResponse, err error) {
	type cancelPayload struct {
		Reference string  `json:"Reference"`
		Note      string  `json:"Note"`
		Rate      float64 `json:"Rate"`
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/UpdateRate", cancelPayload{Reference: reference, Note: note, Rate: rate})
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	req.ExpectStatus(http.StatusOK)
	return data, c.do(ctx, req)
}
