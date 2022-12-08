package checkbox

import "time"

type Good struct {
	Good struct {
		Code           string   `json:"code"` // required  (Код товару)
		Name           string   `json:"name"` // required
		Barcode        string   `json:"barcode"`
		ExciseBarcode  string   `json:"excise_barcode"`
		ExciseBarcodes []string `json:"excise_barcodes"`
		Header         string   `json:"header"`
		Footer         string   `json:"footer"`
		Price          int      `json:"price"` //required (Вартість в копійках за quantity = 1000)
		Tax            []int    `json:"tax"`
		Uktzed         string   `json:"uktzed"`
	} `json:"good"`
	GoodId           string     `json:"good_id"`
	Quantity         int        `json:"quantity"` // required (Кількість (Наприклад: 1 шт = 1000, 2.25 кг = 2250))
	IsReturn         bool       `json:"is_return"`
	IsWinningsPayout bool       `json:"is_winnings_payout"`
	Discounts        []Discount `json:"discounts"`
	Sum              int        `json:"sum"`
	Taxes            []Tax      `json:"taxes"`
}

type Payment struct {
	Type              string `json:"type"` // "CASH" "CARD" "CASHLESS"
	PawnshopIsReturn  bool   `json:"pawnshop_is_return"`
	Value             int    `json:"value"` // required
	Label             string `json:"label"`
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
	Type     string `json:"type"`
	Mode     string `json:"mode"`
	Value    int    `json:"value"`
	TaxCode  int    `json:"tax_code"`
	TaxCodes []int  `json:"tax_codes"`
	Name     string `json:"name"`
	Sum      int    `json:"sum"`
}

type Tax struct {
	Id           string    `json:"id"`
	Code         int       `json:"code"`
	Label        string    `json:"label"`
	Symbol       string    `json:"symbol"`
	Rate         int       `json:"rate"`
	ExtraRate    int       `json:"extra_rate"`
	Included     bool      `json:"included"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	NoVat        bool      `json:"no_vat"`
	AdvancedCode string    `json:"advanced_code"`
	Value        int       `json:"value"`
	ExtraValue   int       `json:"extra_value"`
}
