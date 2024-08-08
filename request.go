package belmoney

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// request is a wrapper around http.Request.
type request struct {
	req      *http.Request
	decodeTo interface{}
}

func NewRequest(r *http.Request) *request {
	return &request{req: r}
}

func (r *request) DecodeTo(to interface{}) {
	r.decodeTo = to
}

// AddFormParams adds multiple form parameters to the request.
func (r *request) AddFormParams(params map[string]string) {
	formData := url.Values{}
	for _, k := range params {
		formData.Add(k, params[k])
	}

	r.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.req.Body = io.NopCloser(strings.NewReader(formData.Encode()))
}

// AddQueryParam adds a query parameter to the request.
func (r *request) AddQueryParam(key, value string) {
	r.AddQueryParams(map[string]string{key: value})
}

// AddQueryParams adds multiple query parameters to the request.
func (r *request) AddQueryParams(params map[string]string) {
	q := r.req.URL.Query()
	for _, k := range params {
		q.Add(k, params[k])
	}
	r.req.URL.RawQuery = q.Encode()
}
