package finvoice

import (
	"encoding/xml"

	"github.com/cydev/zero"
	"github.com/omniboost/go-finvoice/omitempty"
)

type Finvoice struct {
	XMLName                    xml.Name                   `xml:"Finvoice"`
	Version                    string                     `xml:"Version,attr"`
	Xsi                        string                     `xml:"xsi,attr"`
	NoNamespaceSchemaLocation  string                     `xml:"noNamespaceSchemaLocation,attr"`
	MessageTransmissionDetails MessageTransmissionDetails `xml:"MessageTransmissionDetails,omitempty"`
	SellerPartyDetails         SellerPartyDetails         `xml:"SellerPartyDetails,omitempty"`
	SellerInformationDetails   SellerInformationDetails   `xml:"SellerInformationDetails,omitempty"`
	BuyerPartyDetails          BuyerPartyDetails          `xml:"BuyerPartyDetails,omitempty"`
	DeliveryDetails            DeliveryDetails            `xml:"DeliveryDetails,omitempty"`
	InvoiceDetails             InvoiceDetails             `xml:"InvoiceDetails,omitempty"`
	PaymentStatusDetails       PaymentStatusDetails       `xml:"PaymentStatusDetails,omitempty"`
	InvoiceRow                 []InvoiceRow               `xml:"InvoiceRow"`
	EpiDetails                 EpiDetails                 `xml:"EpiDetails,omitempty"`
}

func (f Finvoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func NewFinvoice() *Finvoice {
	return &Finvoice{
		Xsi:                       "http://www.w3.org/2001/XMLSchema-instance",
		NoNamespaceSchemaLocation: "Finvoice2.01.xsd",
	}
}

type MessageTransmissionDetails struct {
	MessageSenderDetails   MessageSenderDetails   `xml:"MessageSenderDetails"`
	MessageReceiverDetails MessageReceiverDetails `xml:"MessageReceiverDetails"`
	MessageDetails         MessageDetails         `xml:"MessageDetails"`
}

type MessageSenderDetails struct {
	FromIdentifier    string `xml:"FromIdentifier"`
	FromIntermediator string `xml:"FromIntermediator"`
}

type MessageReceiverDetails struct {
	ToIdentifier    string `xml:"ToIdentifier"`
	ToIntermediator string `xml:"ToIntermediator"`
}

type MessageDetails struct {
	MessageIdentifier      string `xml:"MessageIdentifier"`
	MessageTimeStamp       string `xml:"MessageTimeStamp"`
	RefToMessageIdentifier string `xml:"RefToMessageIdentifier"`
}

type SellerPartyDetails struct {
	SellerPartyIdentifier        string                     `xml:"SellerPartyIdentifier"`
	SellerPartyIdentifierUrlText string                     `xml:"SellerPartyIdentifierUrlText"`
	SellerOrganisationName       []string                   `xml:"SellerOrganisationName"`
	SellerOrganisationDepartment []string                   `xml:"SellerOrganisationDepartment"`
	SellerOrganisationTaxCode    string                     `xml:"SellerOrganisationTaxCode"`
	SellerPostalAddressDetails   SellerPostalAddressDetails `xml:"SellerPostalAddressDetails"`
}

type SellerPostalAddressDetails struct {
	SellerStreetName         string `xml:"SellerStreetName"`
	SellerTownName           string `xml:"SellerTownName"`
	SellerPostCodeIdentifier string `xml:"SellerPostCodeIdentifier"`
	CountryCode              string `xml:"CountryCode"`
}

type SellerInformationDetails struct {
	SellerAccountDetails []struct {
		SellerAccountID struct {
			IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
		} `xml:"SellerAccountID"`
		SellerBic struct {
			IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
		} `xml:"SellerBic"`
	} `xml:"SellerAccountDetails"`
}

type BuyerPartyDetails struct {
	BuyerPartyIdentifier        string                    `xml:"BuyerPartyIdentifier"`
	BuyerOrganisationName       string                    `xml:"BuyerOrganisationName"`
	BuyerOrganisationDepartment []string                  `xml:"BuyerOrganisationDepartment"`
	BuyerOrganisationTaxCode    string                    `xml:"BuyerOrganisationTaxCode"`
	BuyerPostalAddressDetails   BuyerPostalAddressDetails `xml:"BuyerPostalAddressDetails"`
}

func (b BuyerPartyDetails) IsEmpty() bool {
	return zero.IsZero(b)
}

type BuyerPostalAddressDetails struct {
	BuyerStreetName              string `xml:"BuyerStreetName"`
	BuyerTownName                string `xml:"BuyerTownName"`
	BuyerPostCodeIdentifier      string `xml:"BuyerPostCodeIdentifier"`
	CountryCode                  string `xml:"CountryCode"`
	CountryName                  string `xml:"CountryName"`
	BuyerPostOfficeBoxIdentifier string `xml:"BuyerPostOfficeBoxIdentifier"`
}

type DeliveryDetails struct {
	DeliveryDate struct {
		Format string `xml:"Format,attr"`
	} `xml:"DeliveryDate"`
}

type InvoiceDetails struct {
	InvoiceTypeCode string `xml:"InvoiceTypeCode"`
	InvoiceTypeText string `xml:"InvoiceTypeText"`
	OriginCode      string `xml:"OriginCode"`
	InvoiceNumber   string `xml:"InvoiceNumber"`
	InvoiceDate     struct {
		Format string `xml:"Format,attr"`
	} `xml:"InvoiceDate"`
	OrderIdentifier               string         `xml:"OrderIdentifier"`
	InvoiceTotalVATExcludedAmount AmountCurrency `xml:"InvoiceTotalVatExcludedAmount"`
	InvoiceTotalVatAmount         AmountCurrency `xml:"InvoiceTotalVatAmount"`
	InvoiceTotalVatIncludedAmount AmountCurrency `xml:"InvoiceTotalVatIncludedAmount"`
	VATSpecificationDetails       struct {
		VATBaseAmount  AmountCurrency `xml:"VatBaseAmount"`
		VATRatePercent string         `xml:"VatRatePercent"`
		VATRateAmount  AmountCurrency `xml:"VatRateAmount"`
	} `xml:"VatSpecificationDetails"`
	PaymentTermsDetails struct {
		PaymentTermsFreeText string `xml:"PaymentTermsFreeText"`
		InvoiceDueDate       struct {
			Format string `xml:"Format,attr"`
		} `xml:"InvoiceDueDate"`
		PaymentOverDueFineDetails struct {
			PaymentOverDueFineFreeText string `xml:"PaymentOverDueFineFreeText"`
			PaymentOverDueFinePercent  string `xml:"PaymentOverDueFinePercent"`
		} `xml:"PaymentOverDueFineDetails"`
	} `xml:"PaymentTermsDetails"`
}

func (i InvoiceDetails) IsEmpty() bool {
	return zero.IsZero(i)
}

type PaymentStatusDetails struct {
	PaymentStatusCode string `xml:"PaymentStatusCode"`
}

func (p PaymentStatusDetails) IsEmpty() bool {
	return p.PaymentStatusCode == ""
}

type InvoiceRow struct {
	ArticleIdentifier string           `xml:"ArticleIdentifier,omitempty"`
	ArticleName       string           `xml:"ArticleName"`
	DeliveredQuantity string           `xml:"DeliveredQuantity,omitempty"`
	OrderedQuantity   string           `xml:"OrderedQuantity,omitempty"`
	InvoicedQuantity  InvoicedQuantity `xml:"InvoicedQuantity,omitempty"`

	UnitPriceAmount      AmountCurrency `xml:"UnitPriceAmount,omitempty"`
	RowFreeText          string         `xml:"RowFreeText"`
	RowVATAmount         AmountCurrency `xml:"RowVatAmount"`
	RowVATExcludedAmount AmountCurrency `xml:"RowVatExcludedAmount"`
}

type InvoicedQuantity struct {
	QuantityUnitCode string `xml:"QuantityUnitCode,attr"`
	Amount           Amount `xml:",attr"`
}

type EpiDetails struct {
	EpiIdentificationDetails struct {
		EpiDate struct {
			Format string `xml:"Format,attr"`
		} `xml:"EpiDate"`
		EpiReference string `xml:"EpiReference"`
	} `xml:"EpiIdentificationDetails"`
	EpiPartyDetails struct {
		EpiBfiPartyDetails struct {
			EpiBfiIdentifier struct {
				IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
			} `xml:"EpiBfiIdentifier"`
		} `xml:"EpiBfiPartyDetails"`
		EpiBeneficiaryPartyDetails struct {
			EpiNameAddressDetails string `xml:"EpiNameAddressDetails"`
			EpiBei                string `xml:"EpiBei"`
			EpiAccountID          struct {
				IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
			} `xml:"EpiAccountID"`
		} `xml:"EpiBeneficiaryPartyDetails"`
	} `xml:"EpiPartyDetails"`
	EpiPaymentInstructionDetails struct {
		EpiPaymentInstructionId     string `xml:"EpiPaymentInstructionId"`
		EpiRemittanceInfoIdentifier struct {
			IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
		} `xml:"EpiRemittanceInfoIdentifier"`
		EpiInstructedAmount AmountCurrency `xml:"EpiInstructedAmount"`
		EpiCharge           struct {
			ChargeOption string `xml:"ChargeOption,attr"`
		} `xml:"EpiCharge"`
		EpiDateOptionDate struct {
			Format string `xml:"Format,attr"`
		} `xml:"EpiDateOptionDate"`
	} `xml:"EpiPaymentInstructionDetails"`
}

type Amount float64

type AmountCurrency struct {
	AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
	Amount                   Amount `xml:",attr"`
}
