package cdeklib

import "time"

//TODO add types

type DateTime struct {
	Value time.Time
}

type Date struct {
	Value time.Time
}

type UUID [16]byte

type Money struct {
	Value   float64 `json:"value"`
	VatSum  float64 `json:"vat_sum,omitempty"`
	VatRate int     `json:"vat_rate,omitempty"`
}

type Threshold struct {
	Threshold int     `json:"threshold"`
	Sum       float64 `json:"sum"`
	VatSum    float64 `json:"vat_sum,omitempty"`
	VatRate   int     `json:"vat_rate,omitempty"`
}

type Contact struct {
	Company                       string  `json:"company,omitempty"`
	Name                          string  `json:"name"`
	Email                         string  `json:"email,omitempty"`
	Phones                        []Phone `json:"phones"`
	PassportSeries                string  `json:"passportSeries,omitempty"`
	PassportNumber                string  `json:"passport_number,omitempty"`
	PassportDateOfIssue           Date    `json:"passport_date_of_issue,omitempty"`
	PassportOrganization          string  `json:"passport_organization,omitempty"`
	PassportDateOfBirth           Date    `json:"passport_date_of_birth,omitempty"`
	Tin                           string  `json:"tin,omitempty"`
	PassportRequirementsSatisfied bool    `json:"passport_requirements_satisfied,omitempty"`
}

type Phone struct {
	Number     string `json:"number"`
	Additional string `json:"additional,omitempty"`
}

type Seller struct {
	Name          string `json:"name,omitempty"`
	Inn           string `json:"inn,omitempty"`
	Phone         string `json:"phone,omitempty"`
	OwnershipForm int    `json:"ownership_form,omitempty"`
	Address       string `json:"address,omitempty"`
}

type Location struct {
	Code        int     `json:"code,omitempty"`
	FiasGuid    UUID    `json:"fias_guid,omitempty"`
	PostalCode  string  `json:"postal_code,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Region      string  `json:"region,omitempty"`
	SubRegion   string  `json:"sub_region,omitempty"`
	City        string  `json:"city,omitempty"`
	KladrCode   string  `json:"kladr_code,omitempty"`
	Address     string  `json:"address,omitempty"`
}

type Service struct {
	Code      int    `json:"code"`
	Parameter string `json:"parameter,omitempty"`
}

//type Package struct {
//	Number string `json:"number"`
//	Weight int    `json:"weight"`
//	Length int    `json:"length,omitempty"`
//	Width  int    `json:"width,omitempty"`
//	Height int    `json:"height,omitempty"`
//	Items  []Item `json:"items,omitempty"`
//}

//type Item struct {
//	Name        string  `json:"name"`
//	WareKey     string  `json:"ware_key"`
//	Payment     Money   `json:"payment"`
//	Cost        float64 `json:"cost"`
//	Weight      int     `json:"weight"`
//	WeightGross int     `json:"weight_gross,omitempty"`
//	Amount      int     `json:"amount"`
//	NameI18n    string  `json:"name_i18n,omitempty"`
//	Brand       string  `json:"brand,omitempty"`
//	CountryCode string  `json:"country_code,omitempty"`
//	Material    string  `json:"material,omitempty"`
//	WifiGsm     bool    `json:"wifi_gsm,omitempty"`
//	Url         string  `json:"url,omitempty"`
//}

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Status struct {
	Code       string   `json:"code"`
	Name       string   `json:"name"`
	DateTime   DateTime `json:"date_time"`
	ReasonCode string   `json:"reason_code,omitempty"`
	City       string   `json:"city,omitempty"`
}
