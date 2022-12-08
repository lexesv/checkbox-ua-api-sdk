package checkbox

import "time"

type ReceiptReq struct {
	Id          string `json:"id"`
	CashierName string `json:"cashier_name"`
	Departament string `json:"departament"`
	Goods       []struct {
		Good struct {
			Code           string   `json:"code"`
			Name           string   `json:"name"`
			Barcode        string   `json:"barcode"`
			ExciseBarcode  string   `json:"excise_barcode"`
			ExciseBarcodes []string `json:"excise_barcodes"`
			Header         string   `json:"header"`
			Footer         string   `json:"footer"`
			Price          int      `json:"price"`
			Tax            []int    `json:"tax"`
			Uktzed         string   `json:"uktzed"`
		} `json:"good"`
		GoodId           string `json:"good_id"`
		Quantity         int    `json:"quantity"`
		IsReturn         bool   `json:"is_return"`
		IsWinningsPayout bool   `json:"is_winnings_payout"`
		Discounts        []struct {
			Type     string `json:"type"`
			Mode     string `json:"mode"`
			Value    int    `json:"value"`
			TaxCode  int    `json:"tax_code"`
			TaxCodes []int  `json:"tax_codes"`
			Name     string `json:"name"`
			Sum      int    `json:"sum"`
		} `json:"discounts"`
	} `json:"goods"`
	Delivery struct {
		Email  string   `json:"email"`
		Emails []string `json:"emails"`
		Phone  string   `json:"phone"`
	} `json:"delivery"`
	Discounts []struct {
		Type     string `json:"type"`
		Mode     string `json:"mode"`
		Value    int    `json:"value"`
		TaxCode  int    `json:"tax_code"`
		TaxCodes []int  `json:"tax_codes"`
		Name     string `json:"name"`
		Sum      int    `json:"sum"`
	} `json:"discounts"`
	Payments []struct {
		Type              string `json:"type"`
		PawnshopIsReturn  bool   `json:"pawnshop_is_return"`
		Value             int    `json:"value"`
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
	} `json:"payments"`
	Rounding          bool   `json:"rounding"`
	Header            string `json:"header"`
	Footer            string `json:"footer"`
	Barcode           string `json:"barcode"`
	OrderId           string `json:"order_id"`
	RelatedReceiptId  string `json:"related_receipt_id"`
	PreviousReceiptId string `json:"previous_receipt_id"`
	TechnicalReturn   bool   `json:"technical_return"`
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
	Serial int `json:"serial"`
	Goods  []struct {
		Good struct {
			Code           string   `json:"code"`
			Barcode        string   `json:"barcode"`
			Name           string   `json:"name"`
			ExciseBarcodes []string `json:"excise_barcodes"`
			Header         string   `json:"header"`
			Footer         string   `json:"footer"`
			Uktzed         string   `json:"uktzed"`
			Price          int      `json:"price"`
		} `json:"good"`
		GoodId   string `json:"good_id"`
		Sum      int    `json:"sum"`
		Quantity int    `json:"quantity"`
		IsReturn bool   `json:"is_return"`
		Taxes    []struct {
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
		} `json:"taxes"`
		Discounts []struct {
			Type     string `json:"type"`
			Mode     string `json:"mode"`
			Value    int    `json:"value"`
			TaxCode  int    `json:"tax_code"`
			TaxCodes []int  `json:"tax_codes"`
			Name     string `json:"name"`
			Sum      int    `json:"sum"`
		} `json:"discounts"`
	} `json:"goods"`
	Payments []struct {
		Type              string `json:"type"`
		PawnshopIsReturn  bool   `json:"pawnshop_is_return"`
		Value             int    `json:"value"`
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
	} `json:"payments"`
	TotalSum     int       `json:"total_sum"`
	TotalPayment int       `json:"total_payment"`
	TotalRest    int       `json:"total_rest"`
	RoundSum     int       `json:"round_sum"`
	FiscalCode   string    `json:"fiscal_code"`
	FiscalDate   time.Time `json:"fiscal_date"`
	DeliveredAt  time.Time `json:"delivered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Taxes        []struct {
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
	} `json:"taxes"`
	Discounts []struct {
		Type     string `json:"type"`
		Mode     string `json:"mode"`
		Value    int    `json:"value"`
		TaxCode  int    `json:"tax_code"`
		TaxCodes []int  `json:"tax_codes"`
		Name     string `json:"name"`
		Sum      int    `json:"sum"`
	} `json:"discounts"`
	OrderId string `json:"order_id"`
	Header  string `json:"header"`
	Footer  string `json:"footer"`
	Barcode string `json:"barcode"`
	Custom  struct {
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
	Shift struct {
		Id      string `json:"id"`
		Serial  int    `json:"serial"`
		ZReport struct {
			Id        string `json:"id"`
			Serial    int    `json:"serial"`
			IsZReport bool   `json:"is_z_report"`
			Payments  []struct {
				Id                       string `json:"id"`
				Code                     int    `json:"code"`
				Type                     string `json:"type"`
				Label                    string `json:"label"`
				SellSum                  int    `json:"sell_sum"`
				ReturnSum                int    `json:"return_sum"`
				ServiceIn                int    `json:"service_in"`
				ServiceOut               int    `json:"service_out"`
				CashWithdrawal           int    `json:"cash_withdrawal"`
				CashWithdrawalCommission int    `json:"cash_withdrawal_commission"`
			} `json:"payments"`
			Taxes []struct {
				Id              string    `json:"id"`
				Code            int       `json:"code"`
				Label           string    `json:"label"`
				Symbol          string    `json:"symbol"`
				Rate            int       `json:"rate"`
				SellSum         int       `json:"sell_sum"`
				ReturnSum       int       `json:"return_sum"`
				SalesTurnover   int       `json:"sales_turnover"`
				ReturnsTurnover int       `json:"returns_turnover"`
				NoVat           bool      `json:"no_vat"`
				AdvancedCode    string    `json:"advanced_code"`
				CreatedAt       time.Time `json:"created_at"`
				SetupDate       time.Time `json:"setup_date"`
			} `json:"taxes"`
			SellReceiptsCount           int       `json:"sell_receipts_count"`
			ReturnReceiptsCount         int       `json:"return_receipts_count"`
			CashWithdrawalReceiptsCount int       `json:"cash_withdrawal_receipts_count"`
			TransfersCount              int       `json:"transfers_count"`
			TransfersSum                int       `json:"transfers_sum"`
			Balance                     int       `json:"balance"`
			Initial                     int       `json:"initial"`
			CreatedAt                   time.Time `json:"created_at"`
			UpdatedAt                   time.Time `json:"updated_at"`
			DiscountsSum                int       `json:"discounts_sum"`
			ExtraChargeSum              int       `json:"extra_charge_sum"`
		} `json:"z_report"`
		OpenedAt           time.Time `json:"opened_at"`
		ClosedAt           time.Time `json:"closed_at"`
		InitialTransaction struct {
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
		} `json:"initial_transaction"`
		ClosingTransaction struct {
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
		} `json:"closing_transaction"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Balance   struct {
			Initial        int       `json:"initial"`
			Balance        int       `json:"balance"`
			CashSales      int       `json:"cash_sales"`
			CardSales      int       `json:"card_sales"`
			DiscountsSum   int       `json:"discounts_sum"`
			ExtraChargeSum int       `json:"extra_charge_sum"`
			CashReturns    int       `json:"cash_returns"`
			CardReturns    int       `json:"card_returns"`
			ServiceIn      int       `json:"service_in"`
			ServiceOut     int       `json:"service_out"`
			UpdatedAt      time.Time `json:"updated_at"`
		} `json:"balance"`
		Taxes []struct {
			Id              string    `json:"id"`
			Code            int       `json:"code"`
			Label           string    `json:"label"`
			Symbol          string    `json:"symbol"`
			Rate            int       `json:"rate"`
			ExtraRate       int       `json:"extra_rate"`
			Included        bool      `json:"included"`
			CreatedAt       time.Time `json:"created_at"`
			UpdatedAt       time.Time `json:"updated_at"`
			NoVat           bool      `json:"no_vat"`
			AdvancedCode    string    `json:"advanced_code"`
			Sales           int       `json:"sales"`
			Returns         int       `json:"returns"`
			SalesTurnover   int       `json:"sales_turnover"`
			ReturnsTurnover int       `json:"returns_turnover"`
		} `json:"taxes"`
		EmergencyClose        bool   `json:"emergency_close"`
		EmergencyCloseDetails string `json:"emergency_close_details"`
		CashRegister          struct {
			Id           string    `json:"id"`
			FiscalNumber string    `json:"fiscal_number"`
			Active       bool      `json:"active"`
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			Number       string    `json:"number"`
		} `json:"cash_register"`
		Cashier struct {
			Id          string `json:"id"`
			FullName    string `json:"full_name"`
			Nin         string `json:"nin"`
			KeyId       string `json:"key_id"`
			Permissions struct {
				Orders bool `json:"orders"`
			} `json:"permissions"`
			CreatedAt      time.Time `json:"created_at"`
			UpdatedAt      time.Time `json:"updated_at"`
			CertificateEnd time.Time `json:"certificate_end"`
			Blocked        string    `json:"blocked"`
		} `json:"cashier"`
	} `json:"shift"`
	ControlNumber string `json:"control_number"`
}
