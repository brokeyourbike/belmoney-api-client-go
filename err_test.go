package belmoney_test

import (
	"testing"

	"github.com/brokeyourbike/belmoney-api-client-go"
	"github.com/stretchr/testify/assert"
)

func TestUnexpectedResponse(t *testing.T) {
	resp := belmoney.UnexpectedResponse{Status: 500, Body: "I am an error."}
	assert.Equal(t, "Unexpected response from API. Status: 500 Body: I am an error.", resp.Error())
}
