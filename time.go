package belmoney

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	formats := []string{
		time.RFC3339,
		time.DateOnly,
		"2006-01-02T15:04:05",       // Your format was slightly incorrect
		"2006-01-02T15:04:05Z07:00", // Example with timezone offset
	}

	for _, f := range formats {
		parsed, err := time.Parse(f, s)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("cannot parse time string %s", s)
}
