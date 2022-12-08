package checkbox

import "time"

type ReceiptReq struct {
	Id          string `json:"id"`
	CashierName string `json:"cashier_name"`
	Departament string `json:"departament"`
	Goods       []Good `json:"goods"` // required
	Delivery    struct {
		Email  string   `json:"email"`
		Emails []string `json:"emails"`
		Phone  string   `json:"phone"`
	} `json:"delivery"`
	Discounts         []Discount `json:"discounts"`
	Payments          []Payment  `json:"payments"`
	Rounding          bool       `json:"rounding"`
	Header            string     `json:"header"`
	Footer            string     `json:"footer"`
	Barcode           string     `json:"barcode"`
	OrderId           string     `json:"order_id"`
	RelatedReceiptId  string     `json:"related_receipt_id"`
	PreviousReceiptId string     `json:"previous_receipt_id"`
	TechnicalReturn   bool       `json:"technical_return"`
	Context           struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"context"`
	IsPawnshop bool `json:"is_pawnshop"`
	Custom     struct {
		HtmlGlobalHeader    string `json:"html_global_header"`
		HtmlGlobalFooter    string `json:"html_global_footer"`
		HtmlBodyStyle       string `json:"html_body_style"`
		HtmlReceiptStyle    string `json:"html_receipt_style"`
		HtmlRulerStyle      string `json:"html_ruler_style"`
		HtmlLightBlockStyle string `json:"html_light_block_style"`
		TextGlobalHeader    string `json:"text_global_header"`
		TextGlobalFooter    string `json:"text_global_footer"`
	} `json:"custom"`
}

type ReceiptResp struct {
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
	Serial       int        `json:"serial"`
	Goods        []Good     `json:"goods"`
	Payments     []Payment  `json:"payments"`
	TotalSum     int        `json:"total_sum"`
	TotalPayment int        `json:"total_payment"`
	TotalRest    int        `json:"total_rest"`
	RoundSum     int        `json:"round_sum"`
	FiscalCode   string     `json:"fiscal_code"`
	FiscalDate   time.Time  `json:"fiscal_date"`
	DeliveredAt  time.Time  `json:"delivered_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Taxes        []Tax      `json:"taxes"`
	Discounts    []Discount `json:"discounts"`
	OrderId      string     `json:"order_id"`
	Header       string     `json:"header"`
	Footer       string     `json:"footer"`
	Barcode      string     `json:"barcode"`
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
	Shift         ShiftResp `json:"shift"`
	ControlNumber string    `json:"control_number"`
}
