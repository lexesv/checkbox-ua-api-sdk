package checkbox

import "time"

/*
ShiftResp
"id" - унікальний ідентифікатор зміни у форматі UUID
"serial" - порядковий номер зміни
"status" - статус зміни. Може мати значення CREATED/OPENED/CLOSING/CLOSED (СТВОРЕНА/ВІДКРИТА/ЗАКРИВАЄТЬСЯ/ЗАКРИТА)
"z_report" - блок з інформацією по z-звіту касової зміни. Для відкритої зміни це значення завжди буде null
"opened_at" - мітка часу відкриття зміни у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"closed_at" - мітка часу закриття зміни у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm. Для відкритої зміни завжди буде null
"initial_transaction" - блок з інформацією про першу транзакцію обраної зміни
"id" - унікальний ідентифікатор транзакції у форматі UUID
"type" - тип транзакції. Для першої транзакції зміни це значення завжди буде SHIFT_OPEN
"serial" - порядковий номер транзакції
"status" - статус транзакції. Може приймати значення CREATED/PENDING/SIGNED/DELIVERED/DONE(або ERROR) (СТВОРЕНА/НА ПІДПИСАННІ/ПІДПИСАНА/ВІДПРАВЛЕНА/ВИКОНАНА(або ПОМИЛКА))
"request_signed_at" - мітка часу підписання запиту на відкриття зміни у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"request_received_at" - мітка часу отримання запиту на відкриття зміни сервером ДПС у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm. Якщо запит ще не був доставлений у податкову, то дане значення буде null
"response_status" - статус, отриманний від сервера ДПС у відповідь на запит. Успішний запит матиме статус OK
"response_error_message" - текстовий опис помилки, при отриманні її від сервера ДПС у відповідь на запит.
"response_id" - фіскальний номер транзакції відкриття зміни, отриманий від ДПС (null якщо зміна в ДПС ще не відкрилась)
"offline_id" - фіскальний номер транзакції відкриття зміни у випадку, якщо вона відкрита у офлайн режимі (null якщо зміна ще не відкрита або якщо відкрита у онлайн режимі)
"created_at" - мітка часу створення транзакції у у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"updated_at" - мітка часу останнього оновлення даних про транзакцію у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"previous_hash" - хеш XML даних попередньої транзакції
"closing_transaction" - блок даних про останню транзакцію у зміні. Для відкритої зміни прийматиме значення null. Якщо зміна закрита, то даний блок міститиме набір даних стосовно останньої транзакції у зміні аналогічний до "initial_transaction"
"created_at" - мітка часу створення зміни у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"updated_at" - мітка часу останнього оновлення даних зміни у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm (null якщо дані ще ні разу не оновлювались)
"balance" - блок даних по грошовому балансу обраної зміни. Всі значення передаються у копійках.
"initial" - початковий баланс готівкових коштів
"balance" - поточний баланс готівкових коштів
"cash_sales" - сума готівкових продажів
"card_sales" - сума безготівкових продажів
"cash_returns" - сума готівкових повернень
"card_returns" - сума безготівкових повернень
"service_in" - сума службових внесень готівкових
"service_out" - сума службових вилучень готівкових коштів
"updated_at" - мітка часу останнього оновлення даних балансу зміни у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm (null якщо дані ще ні разу не оновлювались)
"taxes" - блок даних з інформацією про податкові ставки, які діють у рамках обраної зміни. Для кожної податкової ставки набір даних складається з наступних пунктів:
"id" - унікальний ідентифікатор податкової ставки у форматі UUID
"code" - цифровий код податкової ставки
"label" - назва податкової ставки
"symbol" - літерний код податкової ставки
"rate" - розмір податкової ставки у відсотках
"extra_rate" - розмір додаткового збору у відсотках
"included" - тип податку вкладений/накладений (true/false). При використанні вкладеного податку сума податку міститься (вкладена) у загальну суму товару. Накладений податок - сума товару не містить податку, податок рахується окремо. Накладений податок використовується вкрай рідко на практиці, тому через веб-інтерфейс сайту Checkbox заблокована можливість створення податкових ставок накладеного податку.
"created_at" - мітка часу створення податкової ставки у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"updated_at" - мітка часу останньої зміни параметрів податкової ставки у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm. На даний момент при використанны методу /api/v1/cashier/shift, значення "updated_at" завжди повертає null
"sales" - сума податку усіх чеків продажу у копійках
"returns" - сума податку усіх чеків повернення у копійках
"sales_turnover" - сума продажів з цією податковою ставкою (оборот) у копійках
"returns_turnover" - сума повернень з цією податковою ставкою (оборот) у копійках "cash_register" - блок даних з інформацією про касу, на якій у даний момент відкрита зміна обраним касиром
"id": - унікальний ідентифікатор каси у форматі UUID
"fiscal_number" - фіскальний номер каси
"active" - статус каси АКТИВНА/НЕАКТИВНА (true/false)
"created_at" - дата реєстрації каси у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"updated_at" - дата останньої зміни даних обраної каси у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"cashier" - блок даних з інформацією про касира, яким було відкрито зміну
"id" - ідентифікатор касира у форматі UUID
"full_name" - ім'я касира
"nin" - індивідуальний податковий номер
"key_id" - ідентифікатор ключа касира, отриманий з його сертифікату
"signature_type" - тип підпису касира. AGENT - якщо використовується програма Чекбокс підпис або підписання відбувається через плагін у браузері, TEST - тестовий касир, CLOUD_SIGNATURE_3 - підпис на захищеному хмарному сервісі Чекбокс, DEPOSITSIGN - хмарний підпис сервісу Депозитсайн.
"permissions" - блок з інформацією про додаткові дозволи для обраного касира. Наприклад, якщо касиру дозволено доступ до API замовлень, то в даному блоці з'явиться ("orders": true)
"created_at" - дата реєстрації касира у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"updated_at" - дата останньої зміни даних обраного касира у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"certificate_end" - дата завершення строку дії ЕЦП касира у форматі ISO 8601 за шаблоном YYYY-MM-DDThh:mm:ss.ssssss±hh:mm
"blocked" - інформація про стан блокування касира. У випадку, якщо транзакційний процесінг отримає помилку при підписанні транзакції, пов'язану із блокуванням ЕЦП касира, в даному полі буде записано цю помилку.
*/
type ShiftResp struct {
	Id      string `json:"id"`
	Serial  int    `json:"serial"`
	Status  string `json:"status,omitempty"`
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
		Type                 string    `json:"type,omitempty"`
		Serial               int       `json:"serial"`
		Status               string    `json:"status,omitempty"`
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
		Type                 string    `json:"type,omitempty"`
		Serial               int       `json:"serial"`
		Status               string    `json:"status,omitempty"`
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
	EmergencyClose        bool   `json:"emergency_close,omitempty"`
	EmergencyCloseDetails string `json:"emergency_close_details,omitempty"`
	CashRegister          struct {
		Id           string    `json:"id"`
		FiscalNumber string    `json:"fiscal_number"`
		Active       bool      `json:"active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	} `json:"cash_register,omitempty"`
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
	} `json:"cashier,omitempty"`
}

type ShiftsResp struct {
	Meta struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"meta"`
	Results []ShiftResp `json:"results"`
}

type CloseShiftReq struct {
	SkipClientNameCheck bool `json:"skip_client_name_check"`
	Report              struct {
		Id                          string           `json:"id"`
		Serial                      int              `json:"serial"`
		Payments                    []ZReportPayment `json:"payments"`              //required
		Taxes                       []ZReportTax     `json:"taxes"`                 //required
		SellReceiptsCount           int              `json:"sell_receipts_count"`   // required
		ReturnReceiptsCount         int              `json:"return_receipts_count"` //required
		CashWithdrawalReceiptsCount int              `json:"cash_withdrawal_receipts_count"`
		LastReceiptId               string           `json:"last_receipt_id"`
		Initial                     int              `json:"initial"` //required
		Balance                     int              `json:"balance"` //required
		SalesRoundUp                int              `json:"sales_round_up"`
		SalesRoundDown              int              `json:"sales_round_down"`
		ReturnsRoundUp              int              `json:"returns_round_up"`
		ReturnsRoundDown            int              `json:"returns_round_down"`
		DiscountsSum                int              `json:"discounts_sum"`
		ExtraChargeSum              int              `json:"extra_charge_sum"`
		CreatedAt                   time.Time        `json:"created_at"`
	} `json:"report"`
	FiscalCode string    `json:"fiscal_code"`
	FiscalDate time.Time `json:"fiscal_date"`
}
type ZReportPayment struct {
	Type                     string `json:"type"` //required
	Code                     int    `json:"code"`
	Label                    string `json:"label"`       //required
	SellSum                  int    `json:"sell_sum"`    //required
	ReturnSum                int    `json:"return_sum"`  //required
	ServiceIn                int    `json:"service_in"`  //required
	ServiceOut               int    `json:"service_out"` //required
	CashWithdrawal           int    `json:"cash_withdrawal"`
	CashWithdrawalCommission int    `json:"cash_withdrawal_commission"`
}

type ZReportTax struct {
	Code            int       `json:"code"`             //required
	Label           string    `json:"label"`            //required
	Symbol          string    `json:"symbol"`           //required
	Rate            int       `json:"rate"`             //required
	ExtraRate       int       `json:"extra_rate"`       //required
	SellSum         int       `json:"sell_sum"`         //required
	ReturnSum       int       `json:"return_sum"`       //required
	SalesTurnover   int       `json:"sales_turnover"`   //required
	ReturnsTurnover int       `json:"returns_turnover"` //required
	SetupDate       time.Time `json:"setup_date"`       //required
	Included        bool      `json:"included"`
	NoVat           bool      `json:"no_vat"`
	AdvancedCode    string    `json:"advanced_code"`
}
