package checkbox

import "time"

type SingInReq struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
	PinCode  string `json:"pin_code,omitempty"`
}

type SignInResp struct {
	Type        string `json:"type"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type CashierProfileResp struct {
	Id            string `json:"id"`
	FullName      string `json:"full_name"`
	Nin           string `json:"nin"`
	KeyId         string `json:"key_id"`
	SignatureType string `json:"signature_type"`
	Permissions   struct {
		Orders bool `json:"orders"`
	} `json:"permissions"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CertificateEnd time.Time `json:"certificate_end"`
	Blocked        string    `json:"blocked"`
	Organization   struct {
		Id                        string    `json:"id"`
		Title                     string    `json:"title"`
		Edrpou                    string    `json:"edrpou"`
		TaxNumber                 string    `json:"tax_number"`
		CreatedAt                 time.Time `json:"created_at"`
		UpdatedAt                 time.Time `json:"updated_at"`
		SubscriptionExp           string    `json:"subscription_exp"`
		ConcordLogin              string    `json:"concord_login"`
		ConcordUid                string    `json:"concord_uid"`
		ReceiptsRatelimitCount    int       `json:"receipts_ratelimit_count"`
		ReceiptsRatelimitInterval int       `json:"receipts_ratelimit_interval"`
		CanSendSms                bool      `json:"can_send_sms"`
	} `json:"organization"`
}

type CashierSignatureResp struct {
	Online               bool   `json:"online"`
	Type                 string `json:"type"`
	ShiftOpenPossibility bool   `json:"shift_open_possibility"`
}

/*
type CashierShiftResp struct {
	Id      string `json:"id"`
	Serial  int    `json:"serial"`
	Status  string `json:"status"`
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
		Type                 string    `json:"type"`
		Serial               int       `json:"serial"`
		Status               string    `json:"status"`
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
		Type                 string    `json:"type"`
		Serial               int       `json:"serial"`
		Status               string    `json:"status"`
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
	CashRegister struct {
		Id           string    `json:"id"`
		FiscalNumber string    `json:"fiscal_number"`
		Active       bool      `json:"active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	} `json:"cash_register"`
}
*/
