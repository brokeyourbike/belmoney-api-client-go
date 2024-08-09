package belmoney

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// request is a wrapper around http.Request.
type request struct {
	req              *http.Request
	expectedStatuses []int
	decodeTo         interface{}
}

func NewRequest(r *http.Request) *request {
	return &request{req: r}
}

func (r *request) ExpectStatus(expected ...int) {
	r.expectedStatuses = expected
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
