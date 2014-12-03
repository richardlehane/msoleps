package types

import "testing"

func TestCurrency(t *testing.T) {
	var c Currency = 52500
	if c.String() != "$5.2500" {
		t.Errorf("Currency: expecting $5.2500, got %s", c.String())
	}
}
