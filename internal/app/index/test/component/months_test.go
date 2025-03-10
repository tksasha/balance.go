package component_test

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/tksasha/balance/internal/app/index/component"
	"gotest.tools/v3/golden"
)

func TestMonths(t *testing.T) {
	component := component.New()

	testdata := map[string]string{
		"":    "months.html",
		"uah": "months-uah.html",
		"usd": "months-usd.html",
		"eur": "months-eur.html",
	}

	for currency, filename := range testdata {
		writer := bytes.NewBuffer([]byte{})

		values := url.Values{}

		if currency != "" {
			values.Add("currency", currency)
		}

		if err := component.Months(values).Render(writer); err != nil {
			t.Fatalf("failed to render months: %v", err)
		}

		golden.Assert(t, writer.String(), filename)
	}
}
