package belmoney_test

import (
	"testing"

	"github.com/brokeyourbike/belmoney-api-client-go"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
	}{
		{"transaction", "0001-01-01T00:00:00", false},
		{"transaction", "3000-01-01T02:00:00", false},
		{"transaction", "1000-01-01T00:00:00", false},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var d belmoney.Time

			err := d.UnmarshalJSON([]byte(test.value))
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
