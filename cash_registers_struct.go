package checkbox

import "time"

/*
CashRegistersInfoResp
"id" - унікальний ідентифікатор каси ПРРО у форматі UUID
"fiscal_number" - фіскальний номер каси
"created_at" - дата створення каси
"updated_at" - дата останнього оновлення (зміни) облікових даних каси
"address" - адреса розміщення каси
"title" - назва
"offline_mode" - офлайн режим. false означає, що каса зараз у онлайні, а true означає, що каса працює у офлайн режимі.
"stay_offline" - якщо даний параметр true, то це означає, що перехід у офлайн був виконаний на стороні клієнта (шляхом відправки запиту go offline, а не завдяки автоматичному офлайн переходу, який відбувається, коли не відповідає сервер ДПС)
"has_shift" - параметр показує, чи має дана каса активну відкриту зміну
"documents_state" - блок даних з інформацією про останній порядковий номер чека, звіта або z-звіта
"last_receipt_code" - останній порядковий номер чека
"last_report_code" - останній порядковий номер звіта, який не є z-звітом
"last_z_report_code" - останній порядковий номер z-звіта
*/
type CashRegistersInfoResp struct {
	Id             string    `json:"id"`
	FiscalNumber   string    `json:"fiscal_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Address        string    `json:"address"`
	Title          string    `json:"title"`
	OfflineMode    bool      `json:"offline_mode"`
	StayOffline    bool      `json:"stay_offline"`
	HasShift       bool      `json:"has_shift"`
	DocumentsState struct {
		LastReceiptCode int `json:"last_receipt_code"`
		LastReportCode  int `json:"last_report_code"`
		LastZReportCode int `json:"last_z_report_code"`
	} `json:"documents_state"`
}

type CashRegistersResp struct {
	Meta struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"meta"`
	Results []struct {
		Id           string    `json:"id"`
		FiscalNumber string    `json:"fiscal_number"`
		Active       bool      `json:"active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Shift        ShiftResp `json:"shift"`
		OfflineMode  bool      `json:"offline_mode"`
		StayOffline  bool      `json:"stay_offline"`
		Branch       struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			Address      string `json:"address"`
			Organization struct {
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
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		} `json:"branch"`
		Address string `json:"address"`
	} `json:"results"`
}

type CashRegisterResp struct {
	Id           string    `json:"id"`
	FiscalNumber string    `json:"fiscal_number"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Shift        ShiftResp `json:"shift"`
	OfflineMode  bool      `json:"offline_mode"`
	StayOffline  bool      `json:"stay_offline"`
	Branch       struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Address      string `json:"address"`
		Organization struct {
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
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"branch"`
	Address string `json:"address"`
}
