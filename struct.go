package checkbox

import "time"

type Good struct {
	Good struct {
		Code           string   `json:"code"` // required  (Код товару)
		Name           string   `json:"name"` // required
		Barcode        string   `json:"barcode,omitempty"`
		ExciseBarcode  string   `json:"excise_barcode,omitempty"`
		ExciseBarcodes []string `json:"excise_barcodes,omitempty"`
		Header         string   `json:"header,omitempty"`
		Footer         string   `json:"footer,omitempty"`
		Price          int      `json:"price,omitempty"` //required (Вартість в копійках за quantity = 1000)
		Tax            []int    `json:"tax,omitempty"`
		Uktzed         string   `json:"uktzed,omitempty"`
	} `json:"good"`
	GoodId           string     `json:"good_id,omitempty"`
	Quantity         int        `json:"quantity"` // required (Кількість (Наприклад: 1 шт = 1000, 2.25 кг = 2250))
	IsReturn         bool       `json:"is_return,omitempty"`
	IsWinningsPayout bool       `json:"is_winnings_payout,omitempty"`
	Discounts        []Discount `json:"discounts,omitempty"`
	Sum              int        `json:"sum,omitempty"`
	Taxes            []Tax      `json:"taxes,omitempty"`
}

type Payment struct {
	Type              string `json:"type"` // "CASH" "CARD" "CASHLESS"
	PawnshopIsReturn  bool   `json:"pawnshop_is_return,omitempty"`
	Value             int    `json:"value"` // required
	Label             string `json:"label,omitempty"`
	Code              int    `json:"code,omitempty"`
	Commission        int    `json:"commission,omitempty"`
	CardMask          string `json:"card_mask,omitempty"`
	BankName          string `json:"bank_name,omitempty"`
	AuthCode          string `json:"auth_code,omitempty"`
	Rrn               string `json:"rrn,omitempty"`
	PaymentSystem     string `json:"payment_system,omitempty"`
	OwnerName         string `json:"owner_name,omitempty"`
	Terminal          string `json:"terminal,omitempty"`
	Acquiring         string `json:"acquiring,omitempty"`
	AcquirerAndSeller string `json:"acquirer_and_seller,omitempty"`
	ReceiptNo         string `json:"receipt_no,omitempty"`
	SignatureRequired bool   `json:"signature_required,omitempty"`
}

type Discount struct {
	Type     string `json:"type,omitempty"`
	Mode     string `json:"mode,omitempty"`
	Value    int    `json:"value,omitempty"`
	TaxCode  int    `json:"tax_code,omitempty"`
	TaxCodes []int  `json:"tax_codes,omitempty"`
	Name     string `json:"name,omitempty"`
	Sum      int    `json:"sum,omitempty"`
}

type Tax struct {
	Id           string    `json:"id,omitempty"`
	Code         int       `json:"code,omitempty"`
	Label        string    `json:"label,omitempty"`
	Symbol       string    `json:"symbol,omitempty"`
	Rate         int       `json:"rate,omitempty"`
	ExtraRate    int       `json:"extra_rate,omitempty"`
	Included     bool      `json:"included,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	NoVat        bool      `json:"no_vat,omitempty"`
	AdvancedCode string    `json:"advanced_code,omitempty"`
	Value        int       `json:"value,omitempty"`
	ExtraValue   int       `json:"extra_value,omitempty"`
}
