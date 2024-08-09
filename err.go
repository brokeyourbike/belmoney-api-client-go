package belmoney

import (
	"fmt"
)

type UnexpectedResponse struct {
	Status int
	Body   string
}

func (r UnexpectedResponse) Error() string {
	return fmt.Sprintf("Unexpected response from API. Status: %d Body: %s", r.Status, r.Body)
}
