package checkbox

// NewSimpleReceipt
// Code: Код товару
// Name: Назва товару
// Price: Вартість в копійках за quantity = 1000
// Quantity: Кількість (Наприклад: 1 шт = 1000, 2.25 кг = 2250)
// PayType: форма оплати "CASH" "CARD" "CASHLESS"
// PayVal: Оплата в копійках
func (ch *Checkbox) NewSimpleReceipt(Code, Name string, Price, Quantity int, PayType string, PayVal int) *ReceiptReq {
	var goods []Good
	good := Good{}
	good.Good.Code = Code
	good.Good.Name = Name
	good.Good.Price = Price
	good.Quantity = Quantity
	goods = append(goods, good)
	var payments []Payment
	payment := Payment{
		Type:  PayType,
		Value: PayVal,
	}
	payments = append(payments, payment)
	return &ReceiptReq{
		Goods:    goods,
		Payments: payments,
	}
}

// CreateReceipt
// Створення чеку продажу/повернення, його фіскалізація та доставка клієнту по email.
func (ch *Checkbox) CreateReceipt(AccessToken string, req *ReceiptReq) (resp *ReceiptResp, err *Error) {
	resp = &ReceiptResp{}
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/receipts/sell",
		AccessToken:    AccessToken,
		Request:        req,
		Response:       resp,
	}
	err = ch.request(c)
	return resp, err
}
