package currency_test

import (
	"testing"

	"github.com/blueboardio/cldr/currency"
)

func TestCode(t *testing.T) {
	for code := range currency.ActiveCurrencies {
		if !code.IsValid() {
			t.Errorf("%s is a known code. IsValid() should return true.", code)
		}
	}
}
