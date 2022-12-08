package checkbox

func (ch *Checkbox) NewReceipt() *ReceiptReq {
	return new(ReceiptReq)
}

// CreateReceipt
// Створення чеку продажу/повернення, його фіскалізація та доставка клієнту по email.
func (ch *Checkbox) CreateReceipt(AccessToken string, req *ReceiptReq) (*ReceiptResp, *Error) {
	resp := &ReceiptResp{}
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/receipts/sel",
		AccessToken:    AccessToken,
		Request:        req,
		Response:       resp,
	}
	err := ch.request(c)
	return resp, err
}
