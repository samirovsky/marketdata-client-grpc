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

## Try biject on your own copy

1. **Use this template** (top right) → create the repository in your own account.
2. Open Tickerline's [migration guide](https://biject-demo.vercel.app/migrate/samirovsky/tickerline-marketdata) → **Migrate my client** → install biject
   on your copy only.
3. biject opens a setup PR adding its `verify` workflow, then runs the migration **in your
   own GitHub Actions** (your code never leaves your CI) and opens a pull request: every
   changed call site labelled, your tests run against Tickerline's v2 contract mock, and the
   patch re-checked by `api-migrations verify`.
