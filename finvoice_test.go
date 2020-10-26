package finvoice_test

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-finvoice"
)

func TestFinvoice(t *testing.T) {
	f := finvoice.NewFinvoice()
	f.MessageTransmissionDetails = finvoice.MessageTransmissionDetails{
		MessageSenderDetails: finvoice.MessageSenderDetails{
			FromIdentifier:    "From Identifier",
			FromIntermediator: "From Intermediator",
		},
		MessageReceiverDetails: finvoice.MessageReceiverDetails{
			ToIdentifier:    "To Identifier",
			ToIntermediator: "To Intermediator",
		},
		MessageDetails: finvoice.MessageDetails{
			MessageIdentifier:      "Message Identifier",
			MessageTimeStamp:       finvoice.DateTime{time.Now()},
			RefToMessageIdentifier: "",
		},
	}
	f.SellerPartyDetails = finvoice.SellerPartyDetails{
		SellerPartyIdentifier:  "9876543-0",
		SellerOrganisationName: []string{"Seller Organisation Name"},
		SellerPostalAddressDetails: finvoice.SellerPostalAddressDetails{
			SellerStreetName:         "TEST",
			SellerTownName:           "TEST",
			SellerPostCodeIdentifier: "TEST",
			CountryCode:              "NL",
		},
	}
	f.SellerCommunicationDetails = finvoice.SellerCommunicationDetails{
		SellerEmailaddressIdentifier: "seller2@omniboost.io",
	}
	f.SellerInformationDetails = finvoice.SellerInformationDetails{
		SellerCommonEmailaddressIdentifier: "seller@omniboost.io",
	}
	f.BuyerPartyDetails = finvoice.BuyerPartyDetails{
		BuyerPartyIdentifier:        "Buyer Party Identifier",
		BuyerOrganisationName:       "Omniboost B.V.",
		BuyerOrganisationDepartment: []string{},
		BuyerOrganisationTaxCode:    "",
		BuyerPostalAddressDetails: finvoice.BuyerPostalAddressDetails{
			BuyerStreetName:              "Stadhuisplein 3",
			BuyerTownName:                "Terneuzen",
			BuyerPostCodeIdentifier:      "4531GZ",
			CountryCode:                  "NL",
			CountryName:                  "Netherlands",
			BuyerPostOfficeBoxIdentifier: "",
		},
	}
	f.BuyerCommunicationDetails = finvoice.BuyerCommunicationDetails{
		BuyerEmailaddressIdentifier: "buyer@omniboost.io",
	}
	f.DeliveryDetails = finvoice.DeliveryDetails{
		DeliveryDate: finvoice.Date{
			Format: "CCYYMMDD",
			Date:   "19820412",
		},
	}
	f.InvoiceDetails = finvoice.InvoiceDetails{
		InvoiceTypeCode: "INV01",
		InvoiceTypeText: "Invoice",
		OriginCode:      "Original",
		InvoiceNumber:   "69",
		InvoiceDate: finvoice.Date{
			Format: "CCYYMMDD",
			Date:   "19820412",
		},
		SellerReferenceIdentifier: "Seller Reference Identifier",
		BuyersSellerIdentifier:    "Buyers Seller Identifier",
		SellersBuyerIdentifier:    "Sellers Buyer Identifier",
		OrderIdentifier:           "Order Identifier",
		BuyerReferenceIdentifier:  "Buyer Reference Identifier",
		InvoiceTotalVATExcludedAmount: finvoice.AmountCurrency{
			AmountCurrencyIdentifier: "EUR",
			Amount:                   12.0,
		},
		InvoiceTotalVatAmount: finvoice.AmountCurrency{
			AmountCurrencyIdentifier: "EUR",
			Amount:                   12.0,
		},
		InvoiceTotalVatIncludedAmount: finvoice.AmountCurrency{
			AmountCurrencyIdentifier: "EUR",
			Amount:                   12.0,
		},
		VATSpecificationDetails: finvoice.VATSpecificationDetails{
			VATBaseAmount: finvoice.AmountCurrency{
				AmountCurrencyIdentifier: "EUR",
				Amount:                   12.0,
			},
			VATRatePercent: "12",
			VATRateAmount: finvoice.AmountCurrency{
				AmountCurrencyIdentifier: "EUR",
				Amount:                   12.0,
			},
		},
	}
	f.InvoiceRow = append(f.InvoiceRow, finvoice.InvoiceRow{
		ArticleName:       "Test row",
		ArticleIdentifier: "Article Identifier",
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
	f.EpiDetails = finvoice.EpiDetails{
		EpiIdentificationDetails: finvoice.EpiIdentificationDetails{
			EpiDate: finvoice.Date{
				Format: "CCYYMMDD",
				Date:   "19820412",
			},
		},
		EpiPartyDetails: finvoice.EpiPartyDetails{
			EpiBfiPartyDetails: finvoice.EpiBfiPartyDetails{
				EpiBfiIdentifier: finvoice.EpiBfiIdentifier{
					IdentificationSchemeName: "BIC",
					Value:                    "BANKFIHH",
				},
			},
			EpiBeneficiaryPartyDetails: finvoice.EpiBeneficiaryPartyDetails{
				EpiNameAddressDetails: "TEST",
				EpiBei:                "5647382910",
				EpiAccountID: finvoice.EpiAccountID{
					IdentificationSchemeName: "IBAN",
					Value:                    "FI2757800750155448",
				},
			},
		},
		EpiPaymentInstructionDetails: finvoice.EpiPaymentInstructionDetails{
			EpiRemittanceInfoIdentifier: finvoice.EpiRemittanceInfoIdentifier{
				IdentificationSchemeName: "ISO",
				Value:                    "RF471234567890",
			},
			EpiInstructedAmount: finvoice.AmountCurrency{
				AmountCurrencyIdentifier: "EUR",
				Amount:                   12.0,
			},
			EpiCharge: finvoice.EpiCharge{
				ChargeOption: "SLEV",
			},
			EpiDateOptionDate: finvoice.Date{
				Format: "CCYYMMDD",
				Date:   "19820412",
			},
		},
	}

	b, err := xml.MarshalIndent(f, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
