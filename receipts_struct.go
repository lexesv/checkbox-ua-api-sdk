package checkbox

import "time"

type ReceiptPayload struct {
	Id          string `json:"id,omitempty"`
	CashierName string `json:"cashier_name,omitempty"`
	Departament string `json:"departament,omitempty"`
	Goods       []Good `json:"goods"` // required
	Delivery    struct {
		Email  string   `json:"email,omitempty"`
		Emails []string `json:"emails,omitempty"`
		Phone  string   `json:"phone,omitempty"`
	} `json:"delivery,omitempty"`
	Discounts         []ReceiptDiscount `json:"discounts,omitempty"`
	Payments          []ReceiptPayment  `json:"payments,omitempty"`
	Rounding          bool              `json:"rounding,omitempty"`
	Header            string            `json:"header,omitempty"`
	Footer            string            `json:"footer,omitempty"`
	Barcode           string            `json:"barcode,omitempty"`
	OrderId           string            `json:"order_id,omitempty"`
	RelatedReceiptId  string            `json:"related_receipt_id,omitempty"`
	PreviousReceiptId string            `json:"previous_receipt_id,omitempty"`
	TechnicalReturn   bool              `json:"technical_return,omitempty"`
	Context           struct {
		AdditionalProp1 string `json:"additionalProp1,omitempty"`
		AdditionalProp2 string `json:"additionalProp2,omitempty"`
		AdditionalProp3 string `json:"additionalProp3,omitempty"`
	} `json:"context,omitempty"`
	IsPawnshop bool `json:"is_pawnshop,omitempty"`
	Custom     struct {
		HtmlGlobalHeader    string `json:"html_global_header,omitempty"`
		HtmlGlobalFooter    string `json:"html_global_footer,omitempty"`
		HtmlBodyStyle       string `json:"html_body_style,omitempty"`
		HtmlReceiptStyle    string `json:"html_receipt_style,omitempty"`
		HtmlRulerStyle      string `json:"html_ruler_style,omitempty"`
		HtmlLightBlockStyle string `json:"html_light_block_style,omitempty"`
		TextGlobalHeader    string `json:"text_global_header,omitempty"`
		TextGlobalFooter    string `json:"text_global_footer,omitempty"`
	} `json:"custom,omitempty"`
}

type Receipt struct {
	Id          string `json:"id"`
	Transaction struct {
		Id                   string    `json:"id"`
		Serial               int       `json:"serial"`
		RequestSignedAt      time.Time `json:"request_signed_at"`
		RequestReceivedAt    time.Time `json:"request_received_at"`
		ResponseStatus       string    `json:"response_status"`
		ResponseErrorMessage string    `json:"response_error_message"`
		ResponseId           string    `json:"response_id"`
		OfflineId            string    `json:"offline_id"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		PreviousHash         string    `json:"previous_hash"`
	} `json:"transaction"`
	Serial       int               `json:"serial"`
	Goods        []Good            `json:"goods"`
	Payments     []ReceiptPayment  `json:"payments"`
	TotalSum     int               `json:"total_sum"`
	TotalPayment int               `json:"total_payment"`
	TotalRest    int               `json:"total_rest"`
	RoundSum     int               `json:"round_sum"`
	FiscalCode   string            `json:"fiscal_code"`
	FiscalDate   time.Time         `json:"fiscal_date"`
	DeliveredAt  time.Time         `json:"delivered_at"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Taxes        []GoodTax         `json:"taxes"`
	Discounts    []ReceiptDiscount `json:"discounts"`
	OrderId      string            `json:"order_id"`
	Header       string            `json:"header"`
	Footer       string            `json:"footer"`
	Barcode      string            `json:"barcode"`
	Custom       struct {
		HtmlGlobalHeader    string `json:"html_global_header"`
		HtmlGlobalFooter    string `json:"html_global_footer"`
		HtmlBodyStyle       string `json:"html_body_style"`
		HtmlReceiptStyle    string `json:"html_receipt_style"`
		HtmlRulerStyle      string `json:"html_ruler_style"`
		HtmlLightBlockStyle string `json:"html_light_block_style"`
		TextGlobalHeader    string `json:"text_global_header"`
		TextGlobalFooter    string `json:"text_global_footer"`
	} `json:"custom"`
	IsCreatedOffline bool      `json:"is_created_offline"`
	IsSentDps        bool      `json:"is_sent_dps"`
	SentDpsAt        time.Time `json:"sent_dps_at"`
	TaxUrl           string    `json:"tax_url"`
	RelatedReceiptId string    `json:"related_receipt_id"`
	TechnicalReturn  bool      `json:"technical_return"`
	CurrencyExchange struct {
		Sell struct {
			Currency string `json:"currency"`
			Value    int    `json:"value"`
			Rate     struct {
				Code       string    `json:"code"`
				SymbolCode string    `json:"symbol_code"`
				Name       string    `json:"name"`
				Sell       int       `json:"sell"`
				Buy        int       `json:"buy"`
				Regulator  int       `json:"regulator"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
				Active     bool      `json:"active"`
				SellSum    int       `json:"sell_sum"`
				BuySum     int       `json:"buy_sum"`
			} `json:"rate"`
		} `json:"sell"`
		Buy struct {
			Currency string `json:"currency"`
			Value    int    `json:"value"`
			Rate     struct {
				Code       string    `json:"code"`
				SymbolCode string    `json:"symbol_code"`
				Name       string    `json:"name"`
				Sell       int       `json:"sell"`
				Buy        int       `json:"buy"`
				Regulator  int       `json:"regulator"`
				CreatedAt  time.Time `json:"created_at"`
				UpdatedAt  time.Time `json:"updated_at"`
				Active     bool      `json:"active"`
				SellSum    int       `json:"sell_sum"`
				BuySum     int       `json:"buy_sum"`
			} `json:"rate"`
		} `json:"buy"`
		Type       string `json:"type"`
		Reversal   bool   `json:"reversal"`
		ClientInfo string `json:"client_info"`
		Commission int    `json:"commission"`
		Cross      int    `json:"cross"`
	} `json:"currency_exchange"`
	Shift         Shift  `json:"shift"`
	ControlNumber string `json:"control_number"`
}

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
	GoodId           string            `json:"good_id,omitempty"`
	Quantity         int               `json:"quantity"` // required (Кількість (Наприклад: 1 шт = 1000, 2.25 кг = 2250))
	IsReturn         bool              `json:"is_return,omitempty"`
	IsWinningsPayout bool              `json:"is_winnings_payout,omitempty"`
	Discounts        []ReceiptDiscount `json:"discounts,omitempty"`
	Sum              int               `json:"sum,omitempty"`
	Taxes            []GoodTax         `json:"taxes,omitempty"`
}

type GoodTax struct {
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

type ReceiptPayment struct {
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

type ReceiptDiscount struct {
	Type     string `json:"type,omitempty"`
	Mode     string `json:"mode,omitempty"`
	Value    int    `json:"value,omitempty"`
	TaxCode  int    `json:"tax_code,omitempty"`
	TaxCodes []int  `json:"tax_codes,omitempty"`
	Name     string `json:"name,omitempty"`
	Sum      int    `json:"sum,omitempty"`
}
