package cdeklib

//TODO describe all the basic structures

type Tariff struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float64 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
}

type Size struct {
	Weight int
	Length int
	Width  int
	Height int
}

type CalcRequest struct {
	Type         int          `json:"type"`
	Date         string       `json:"date"`
	Currency     int          `json:"currency"`
	Lang         string       `json:"lang"`
	FromLocation LocationCalc `json:"from_location"`
	ToLocation   LocationCalc `json:"to_location"`
	Packages     []Size       `json:"packages"`
}

type Package struct {
	Number string `json:"number"`
	Weight int    `json:"weight"`
	Length int    `json:"length,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Items  []Item `json:"items"`
}

//type Item struct {
//	Name    string  `json:"name"`
//	WareKey string  `json:"ware_key"`
//	Payment Money   `json:"payment"`
//	Value   float64 `json:"value"`
//	Cost    float64 `json:"cost"`
//	Weight  int     `json:"weight"`
//	Amount  int     `json:"amount"`
//}

// Recipient minimal order request struct
type Recipient struct {
	Name   string  `json:"name"`
	Phones []Phone `json:"phones"`
}

type LocationCalc struct {
	Code        string `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address,omitempty"`
}

type OrderRequest struct {
	TariffCode   int           `json:"tariff_code"`
	FromLocation OrderLocation `json:"from_location"`
	ToLocation   OrderLocation `json:"to_location"`
	Recipient    Recipient     `json:"recipient"`
	Packages     []Package     `json:"packages"`
}

type OrderLocation struct {
	Code        int    `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"county_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address"`
}

type Item struct {
	Name    string  `json:"name"`
	WareKey string  `json:"ware_key"`
	Payment Money   `json:"payment"`
	Value   int     `json:"value"`
	Cost    float64 `json:"cost"`
	Weight  int     `json:"weight"`
	Amount  int     `json:"amount"`
}
