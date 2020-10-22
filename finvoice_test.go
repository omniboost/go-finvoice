package finvoice_test

import (
	"encoding/xml"
	"log"
	"testing"

	"github.com/omniboost/go-finvoice"
)

func TestFinvoice(t *testing.T) {
	f := finvoice.NewFinvoice()
	f.InvoiceRow = append(f.InvoiceRow, finvoice.InvoiceRow{
		ArticleName: "Test row",
		InvoicedQuantity: finvoice.InvoicedQuantity{
			QuantityUnitCode: "EUR",
			Amount:           finvoice.Amount(121.0),
		},
		RowVATAmount: finvoice.AmountCurrency{
			AmountCurrencyIdentifier: "EUR",
			Amount:                   finvoice.Amount(21.0),
		},
		RowVATExcludedAmount: finvoice.AmountCurrency{
			AmountCurrencyIdentifier: "EUR",
			Amount:                   finvoice.Amount(100.0),
		},
	})

	b, err := xml.MarshalIndent(f, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(b))
}
