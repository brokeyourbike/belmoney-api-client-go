package hash_test

import (
	"testing"
	"time"

	"github.com/brokeyourbike/belmoney-api-client-go/hash"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	hasher := hash.NewHasher("token", "secret")

	want := "dG9rZW46MjAwNjAxMDJfMTUwNDA1OlJjN0lla3JicUNYTjJOcGpxMGFTRWFwbjRaNXUwQ0Uxc2M1S21ESmRQekdXQnBiRE9YYlhvMzJ1ajFDb0JTcUNOUmVySmtlS0UvVGswdGFIdUxuQzhBPT0="
	assert.Equal(t, want, hasher.Generate(time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)))
}
