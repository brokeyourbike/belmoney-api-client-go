package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"
)

type Hasher interface {
	// Generate, generates a new authentication token.
	Generate() string
}

type hasher struct {
	token  string
	secret string
}

func NewHasher(token, secret string) *hasher {
	return &hasher{token: token, secret: secret}
}

func (s *hasher) Generate() string {
	// Get the current timestamp in the format YYYYMMDD_HHmmss in UTC
	timestamp := time.Now().UTC().Format("20060102_150405")

	// Create the HMAC SHA512 hash of the timestamp using the secret
	h := hmac.New(sha512.New, []byte(s.secret))
	h.Write([]byte(timestamp))
	timestampHash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Create the checksum input string
	checksumInput := fmt.Sprintf("%s:%s:%s", s.token, timestamp, timestampHash)

	// Generate the final auth hash (Base64 encoded)
	authHash := base64.StdEncoding.EncodeToString([]byte(checksumInput))

	return authHash
}
