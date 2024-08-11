package belmoney

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/brokeyourbike/belmoney-api-client-go/hash"
	"github.com/sirupsen/logrus"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	baseURL    string
	token      string
	secret     string
	hasher     hash.Hasher
	httpClient HttpClient
	logger     *logrus.Logger
}

func NewClient(baseURL, token, secret string, options ...ClientOption) *client {
	c := &client{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		token:   token,
		secret:  secret,
		hasher:  hash.NewHasher(token, secret),
	}

	c.httpClient = http.DefaultClient

	for _, option := range options {
		option(c)
	}

	return c
}

// ClientOption is a function that configures a Client.
type ClientOption func(*client)

// WithHTTPClient sets the HTTP client for the API client.
func WithHTTPClient(c HttpClient) ClientOption {
	return func(target *client) {
		target.httpClient = c
	}
}

// WithLogger sets the *logrus.Logger for the API client.
func WithLogger(l *logrus.Logger) ClientOption {
	return func(target *client) {
		target.logger = l
	}
}

func (c *client) newRequest(ctx context.Context, method, url string, body interface{}) (*request, error) {
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
	req.Header.Set("Authorization", fmt.Sprintf("APIAuth %s", c.hasher.Generate(time.Now())))
	return NewRequest(req), nil
}

func (c *client) do(ctx context.Context, req *request) error {
	resp, err := c.httpClient.Do(req.req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if c.logger != nil {
		c.logger.WithContext(ctx).WithFields(logrus.Fields{
			"http.response.status_code":  resp.StatusCode,
			"http.response.body.content": string(b),
			"http.response.headers":      resp.Header,
		}).Debug("belmoney.client -> response")
	}

	if !slices.Contains(req.expectedStatuses, resp.StatusCode) {
		return UnexpectedResponse{Status: resp.StatusCode, Body: string(b)}
	}

	if req.decodeTo != nil {
		if err := json.Unmarshal(b, req.decodeTo); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
