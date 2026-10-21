// Package quotes reads live quotes from the Tickerline market data gRPC API.
package quotes

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"time"

	marketdatav1 "github.com/samirovsky/marketdata-client-grpc/gen/tickerline/marketdata/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// Address of the Tickerline gRPC endpoint.
const Address = "tickerline-grpc.fly.dev:443"

// Client wraps the generated stub with the API key.
type Client struct {
	conn *grpc.ClientConn
	stub marketdatav1.MarketDataClient
	key  string
}

// Dial connects over TLS.
func Dial() (*Client, error) {
	conn, err := grpc.NewClient(Address, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12})))
	if err != nil {
		return nil, err
	}
	key := os.Getenv("TICKERLINE_KEY")
	if key == "" {
		key = "demo"
	}
	return &Client{conn: conn, stub: marketdatav1.NewMarketDataClient(conn), key: key}, nil
}

// Close releases the connection.
func (c *Client) Close() error { return c.conn.Close() }

func (c *Client) ctx(parent context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, 10*time.Second)
	return metadata.AppendToOutgoingContext(ctx, "x-tickerline-key", c.key), cancel
}

// Mid is a formatted quote line.
type Mid struct {
	Symbol   string
	Price    float64
	Currency string
	Volume   int32
	At       time.Time
}

// Latest returns the latest quote of a symbol.
func (c *Client) Latest(parent context.Context, symbol string) (Mid, error) {
	ctx, cancel := c.ctx(parent)
	defer cancel()
	q, err := c.stub.GetQuote(ctx, &marketdatav1.GetQuoteRequest{Symbol: symbol})
	if err != nil {
		return Mid{}, fmt.Errorf("quote %s: %w", symbol, err)
	}
	return Mid{Symbol: q.GetSymbol(), Price: q.GetPrice(), Currency: q.GetCurrency(), Volume: q.GetVolume(), At: q.GetTimestamp().AsTime()}, nil
}

// HourlyCloses returns the last n hourly closes.
func (c *Client) HourlyCloses(parent context.Context, symbol string, n int32) ([]float64, error) {
	ctx, cancel := c.ctx(parent)
	defer cancel()
	res, err := c.stub.ListCandles(ctx, &marketdatav1.ListCandlesRequest{Symbol: symbol, Interval: marketdatav1.Interval_INTERVAL_1H, Limit: n})
	if err != nil {
		return nil, err
	}
	out := make([]float64, 0, len(res.GetCandles()))
	for _, k := range res.GetCandles() {
		out = append(out, k.GetC())
	}
	return out, nil
}
