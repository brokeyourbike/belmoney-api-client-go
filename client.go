package belmoney

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	httpClient HttpClient
	logger     *logrus.Logger
}

// ClientOption is a function that configures a Client.
type ClientOption func(*client)

// WithHTTPClient sets the HTTP client for the paystack API client.
func WithHTTPClient(c HttpClient) ClientOption {
	return func(target *client) {
		target.httpClient = c
	}
}

// WithLogger sets the *logrus.Logger for the paystack API client.
func WithLogger(l *logrus.Logger) ClientOption {
	return func(target *client) {
		target.logger = l
	}
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

	if resp.StatusCode >= 500 {
		var errResponse ErrResponse
		if err := json.Unmarshal(b, &errResponse); err != nil {
			return fmt.Errorf("cannot decode err response: %w", err)
		}
		return errResponse
	}

	if req.decodeTo != nil {
		if err := json.Unmarshal(b, req.decodeTo); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
