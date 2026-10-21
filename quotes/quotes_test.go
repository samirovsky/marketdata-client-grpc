package quotes

import (
	"context"
	"testing"
)

func TestLatestAndHourly(t *testing.T) {
	c, err := Dial()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	q, err := c.Latest(context.Background(), "NOVA")
	if err != nil {
		t.Fatal(err)
	}
	if q.Symbol != "NOVA" || q.Price <= 0 || q.Volume <= 0 {
		t.Fatalf("unexpected quote %+v", q)
	}
	closes, err := c.HourlyCloses(context.Background(), "NOVA", 6)
	if err != nil {
		t.Fatal(err)
	}
	if len(closes) != 6 {
		t.Fatalf("got %d closes", len(closes))
	}
}
