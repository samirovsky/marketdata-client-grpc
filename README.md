# marketdata-client-grpc

A quotes service on the Tickerline market data **gRPC** API (v1), in Go. The `.proto`
files are vendored in `proto/`; stubs are generated at build time into `gen/` (not
committed):

```sh
go generate ./...   # buf + protoc-gen-go + protoc-gen-go-grpc, via go run
go test ./...
go run ./cmd/quotes
```

When Tickerline releases v2, biject replaces the vendored proto with v2, moves the imports
to `tickerline/marketdata/v2`, and adapts the renamed generated identifiers
(`GetQuote` → `FetchQuote`, `Price` → `Last`, `Limit` → `Count`, `int64` volume) before
opening a PR checked against the contract mock.
