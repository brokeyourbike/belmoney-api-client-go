# belmoney-api-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/brokeyourbike/belmoney-api-client-go.svg)](https://pkg.go.dev/github.com/brokeyourbike/belmoney-api-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/brokeyourbike/belmoney-api-client-go)](https://goreportcard.com/report/github.com/brokeyourbike/belmoney-api-client-go)

Belmoney Bank API Client for Go

## Installation

```bash
go get -u github.com/brokeyourbike/belmoney-api-client-go
```

## Usage

### Outgoing

```go
outClient := belmoney.NewOutClient("base_url", "client_id", "client_secret")
outClient.Create(context.TODO(), belmoney.CreateOutTransactionPayload{})
```

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## License
[BSD-3-Clause License](https://github.com/brokeyourbike/belmoney-api-client-go/blob/main/LICENSE)
