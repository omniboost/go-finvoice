package finvoice

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/cydev/zero"
	"github.com/omniboost/go-finvoice/omitempty"
)

type Finvoice struct {
	XMLName                    xml.Name                   `xml:"Finvoice"`
	Version                    string                     `xml:"Version,attr"`
	Xsi                        string                     `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation  string                     `xml:"xsi:noNamespaceSchemaLocation,attr"`
	MessageTransmissionDetails MessageTransmissionDetails `xml:"MessageTransmissionDetails"`
	SellerPartyDetails         SellerPartyDetails         `xml:"SellerPartyDetails"`
	SellerCommunicationDetails SellerCommunicationDetails `xml:"SellerCommunicationDetails"`
	SellerInformationDetails   SellerInformationDetails   `xml:"SellerInformationDetails"`
	BuyerPartyDetails          BuyerPartyDetails          `xml:"BuyerPartyDetails"`
	BuyerCommunicationDetails  BuyerCommunicationDetails  `xml:"BuyerCommunicationDetails"`
	DeliveryDetails            DeliveryDetails            `xml:"DeliveryDetails"`
	InvoiceDetails             InvoiceDetails             `xml:"InvoiceDetails"`
	PaymentStatusDetails       PaymentStatusDetails       `xml:"PaymentStatusDetails"`
	InvoiceRows                []InvoiceRow               `xml:"InvoiceRow"`
	EpiDetails                 EpiDetails                 `xml:"EpiDetails"`
}

func (f Finvoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(f, e, start)
}

func NewFinvoice() *Finvoice {
	return &Finvoice{
		Version:                   "3.0",
		Xsi:                       "http://www.w3.org/2001/XMLSchema-instance",
		NoNamespaceSchemaLocation: "Finvoice3.0.xsd",
	}
}

type MessageTransmissionDetails struct {
	MessageSenderDetails   MessageSenderDetails   `xml:"MessageSenderDetails"`
	MessageReceiverDetails MessageReceiverDetails `xml:"MessageReceiverDetails"`
	MessageDetails         MessageDetails         `xml:"MessageDetails"`
}

func (m MessageTransmissionDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(m, e, start)
}

type MessageSenderDetails struct {
	FromIdentifier    string `xml:"FromIdentifier"`
	FromIntermediator string `xml:"FromIntermediator"`
}

func (m MessageSenderDetails) IsEmpty() bool {
	return zero.IsZero(m)
}

type MessageReceiverDetails struct {
	ToIdentifier    string `xml:"ToIdentifier"`
	ToIntermediator string `xml:"ToIntermediator"`
}

type MessageDetails struct {
	MessageIdentifier      string   `xml:"MessageIdentifier"`
	MessageTimeStamp       DateTime `xml:"MessageTimeStamp"`
	RefToMessageIdentifier string   `xml:"RefToMessageIdentifier"`
}

type SellerPartyDetails struct {
	SellerPartyIdentifier        string                     `xml:"SellerPartyIdentifier"`
	SellerPartyIdentifierUrlText string                     `xml:"SellerPartyIdentifierUrlText"`
	SellerOrganisationName       string                     `xml:"SellerOrganisationName"`
	SellerOrganisationDepartment string                     `xml:"SellerOrganisationDepartment"`
	SellerOrganisationTaxCode    string                     `xml:"SellerOrganisationTaxCode"`
	SellerPostalAddressDetails   SellerPostalAddressDetails `xml:"SellerPostalAddressDetails"`
}

type SellerPostalAddressDetails struct {
	SellerStreetName         string `xml:"SellerStreetName"`
	SellerTownName           string `xml:"SellerTownName"`
	SellerPostCodeIdentifier string `xml:"SellerPostCodeIdentifier"`
	CountryCode              string `xml:"CountryCode"`
}

type SellerCommunicationDetails struct {
	SellerEmailaddressIdentifier string `xml:"SellerEmailaddressIdentifier"`
}

type SellerInformationDetails struct {
	SellerCommonEmailaddressIdentifier string `xml:"SellerCommonEmailaddressIdentifier"`
	SellerAccountDetails               []struct {
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

type BuyerCommunicationDetails struct {
	BuyerEmailaddressIdentifier string `xml:"BuyerEmailaddressIdentifier"`
}

type DeliveryDetails struct {
	DeliveryDate Date `xml:"DeliveryDate"`
}

type InvoiceDetails struct {
	InvoiceTypeCode               string                    `xml:"InvoiceTypeCode"`
	InvoiceTypeText               string                    `xml:"InvoiceTypeText"`
	OriginCode                    string                    `xml:"OriginCode"`
	InvoiceNumber                 string                    `xml:"InvoiceNumber"`
	InvoiceDate                   Date                      `xml:"InvoiceDate"`
	SellerReferenceIdentifier     string                    `xml:"SellerReferenceIdentifier"`
	BuyersSellerIdentifier        string                    `xml:"BuyersSellerIdentifier"`
	SellersBuyerIdentifier        string                    `xml:"SellersBuyerIdentifier"`
	OrderIdentifier               string                    `xml:"OrderIdentifier"`
	BuyerReferenceIdentifier      string                    `xml:"BuyerReferenceIdentifier"`
	ProjectReferenceIdentifier    string                    `xml:"ProjectReferenceIdentifier"`
	DefinitionDetails             DefinitionDetails         `xml:"DefinitionDetails"`
	InvoiceTotalVATExcludedAmount AmountCurrency            `xml:"InvoiceTotalVatExcludedAmount"`
	InvoiceTotalVatAmount         AmountCurrency            `xml:"InvoiceTotalVatAmount"`
	InvoiceTotalVatIncludedAmount AmountCurrency            `xml:"InvoiceTotalVatIncludedAmount"`
	VATSpecificationDetails       []VATSpecificationDetails `xml:"VatSpecificationDetails"`
	PaymentTermsDetails           PaymentTermsDetails       `xml:"PaymentTermsDetails"`
}

type DefinitionDetails struct {
	DefinitionHeaderText string `xml:"DefinitionHeaderText"`
}

func (i InvoiceDetails) IsEmpty() bool {
	return zero.IsZero(i)
}

type PaymentStatusDetails struct {
	PaymentStatusCode string `xml:"PaymentStatusCode,omitempty"`
}

func (p PaymentStatusDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(p, e, start)
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

func (i InvoiceRow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(i, e, start)
}

type InvoicedQuantity struct {
	QuantityUnitCode string `xml:"QuantityUnitCode,attr"`
	Amount           Amount `xml:",chardata"`
}

func (i InvoicedQuantity) IsEmpty() bool {
	return zero.IsZero(i)
}

type EpiDetails struct {
	EpiIdentificationDetails     EpiIdentificationDetails     `xml:"EpiIdentificationDetails"`
	EpiPartyDetails              EpiPartyDetails              `xml:"EpiPartyDetails"`
	EpiPaymentInstructionDetails EpiPaymentInstructionDetails `xml:"EpiPaymentInstructionDetails"`
}

func (d EpiDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

type EpiIdentificationDetails struct {
	EpiDate      Date   `xml:"EpiDate"`
	EpiReference string `xml:"EpiReference"`
}

func (d EpiIdentificationDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

type EpiPartyDetails struct {
	EpiBfiPartyDetails         EpiBfiPartyDetails         `xml:"EpiBfiPartyDetails"`
	EpiBeneficiaryPartyDetails EpiBeneficiaryPartyDetails `xml:"EpiBeneficiaryPartyDetails"`
}

type EpiBfiPartyDetails struct {
	EpiBfiIdentifier EpiBfiIdentifier `xml:"EpiBfiIdentifier"`
}

type EpiBfiIdentifier struct {
	IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
	Value                    string `xml:",chardata"`
}

type EpiBeneficiaryPartyDetails struct {
	EpiNameAddressDetails string       `xml:"EpiNameAddressDetails"`
	EpiBei                string       `xml:"EpiBei"`
	EpiAccountID          EpiAccountID `xml:"EpiAccountID"`
}

type EpiAccountID struct {
	IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
	Value                    string `xml:",chardata"`
}

type EpiPaymentInstructionDetails struct {
	EpiPaymentInstructionId     string                      `xml:"EpiPaymentInstructionId"`
	EpiRemittanceInfoIdentifier EpiRemittanceInfoIdentifier `xml:"EpiRemittanceInfoIdentifier"`
	EpiInstructedAmount         AmountCurrency              `xml:"EpiInstructedAmount"`
	EpiCharge                   EpiCharge                   `xml:"EpiCharge"`
	EpiDateOptionDate           Date                        `xml:"EpiDateOptionDate"`
}

type EpiRemittanceInfoIdentifier struct {
	IdentificationSchemeName string `xml:"IdentificationSchemeName,attr"`
	Value                    string `xml:",chardata"`
}

type EpiCharge struct {
	ChargeOption string `xml:"ChargeOption,attr"`
}

type Amount float64

func (a Amount) MarshalText() ([]byte, error) {
	s := fmt.Sprintf("%.2f", float64(a))
	s = strings.Replace(s, ".", ",", -1)
	return []byte(s), nil
}

type AmountCurrency struct {
	AmountCurrencyIdentifier string `xml:"AmountCurrencyIdentifier,attr"`
	Amount                   Amount `xml:",chardata"`
}

func (d AmountCurrency) IsEmpty() bool {
	return d.AmountCurrencyIdentifier == ""
}

type Date struct {
	Format string `xml:"Format,attr"`
	Date   string `xml:",chardata"`
}

func (d Date) IsEmpty() bool {
	return zero.IsZero(d)
}

type Number float64

func (n Number) IsEmpty() bool {
	return n == Number(0.0)
}

type VATSpecificationDetails struct {
	VATBaseAmount  AmountCurrency `xml:"VatBaseAmount"`
	VATRatePercent Number         `xml:"VatRatePercent"`
	VATRateAmount  AmountCurrency `xml:"VatRateAmount"`
}

type PaymentTermsDetails struct {
	PaymentTermsFreeText      string `xml:"PaymentTermsFreeText"`
	InvoiceDueDate            Date   `xml:"InvoiceDueDate,omitempty"`
	PaymentOverDueFineDetails struct {
		PaymentOverDueFineFreeText string `xml:"PaymentOverDueFineFreeText"`
		PaymentOverDueFinePercent  Number `xml:"PaymentOverDueFinePercent,omitempty"`
	} `xml:"PaymentOverDueFineDetails"`
}

func (p PaymentTermsDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(p, e, start)
}

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time.Format("2006-01-02T15:04:05-07:00"))
}

func (d *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try untill date format
	d.Time, err = time.Parse("2006-01-02T15:04:05-07:00", value)
	return
}

func (d DateTime) MarshalText() ([]byte, error) {
	return []byte(d.Format("2006-01-02T15:04:05-07:00")), nil
}
