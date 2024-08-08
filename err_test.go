package belmoney_test

import (
	"testing"

	"github.com/brokeyourbike/belmoney-api-client-go"
	"github.com/stretchr/testify/assert"
)

func TestErrResponse(t *testing.T) {
	err := belmoney.ErrResponse{Status: 200, Err: "damn!", Message: "oh no"}
	assert.Equal(t, "status: 200 error: damn! msg: oh no", err.Error())
}
