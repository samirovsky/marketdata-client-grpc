package main

import (
	"context"
	"fmt"
	"log"

	"github.com/samirovsky/marketdata-client-grpc/quotes"
)

func main() {
	c, err := quotes.Dial()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for _, s := range []string{"TLX", "ORBT", "KELP", "NOVA"} {
		q, err := c.Latest(context.Background(), s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-5s %10.2f %s  vol %d  %s\n", q.Symbol, q.Price, q.Currency, q.Volume, q.At.Format("15:04"))
	}
}
