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
