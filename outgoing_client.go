package belmoney

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/brokeyourbike/belmoney-api-client-go/hash"
	"github.com/sirupsen/logrus"
)

type OutClient interface {
	Create(ctx context.Context, transactionPayload CreateOutTransactionPayload) (CreateOutTransactionResponse, error)
}

var _ OutClient = (*outClient)(nil)

type outClient struct {
	client
	baseURL string
	token   string
	secret  string
	hasher  hash.Hasher
}

func NewOutClient(baseURL, token, secret string, options ...ClientOption) *outClient {
	c := &outClient{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		token:   token,
		secret:  secret,
		hasher:  hash.NewHasher(token, secret),
	}

	c.httpClient = http.DefaultClient

	for _, option := range options {
		option(&c.client)
	}

	return c
}

func (c *outClient) newRequest(ctx context.Context, method, url string, body interface{}) (*request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var b []byte

	if body != nil {
		b, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}

		req.Body = io.NopCloser(bytes.NewReader(b))
		req.ContentLength = int64(len(b))
		req.Header.Set("Content-Type", "application/json")
	}

	if c.logger != nil {
		c.logger.WithContext(ctx).WithFields(logrus.Fields{
			"http.request.method":       req.Method,
			"http.request.url":          req.URL.String(),
			"http.request.body.content": string(b),
		}).Debug("belmoney.client -> request")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("APIAuth %s", c.hasher.Generate()))
	return NewRequest(req), nil
}
