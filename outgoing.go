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
	References []string `json:"References"`
}

func (c *client) Transaction(ctx context.Context, reference string) (data TransactionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/NewList/%s", reference), nil)
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
