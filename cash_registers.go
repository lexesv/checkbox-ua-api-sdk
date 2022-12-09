package checkbox

import "fmt"

// GetCashRegistersInfo
//
// Отримання інформації про пРРО за ключем ліцензії. Необхідно для агенту РРО.
func (ch *Checkbox) GetCashRegistersInfo() (*CashRegistersInfoResp, *Error) {
	resp := &CashRegistersInfoResp{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: true,
		Endpoint:       "/api/v1/cash-registers/info",
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// GetCashRegisters
//
// # Отримання переліку доступних пРРО
//
// InUse - Чи відображати зайняті каси (1 - тільки зайняті, 2 - тільки вільні, 0 - всі)
func (ch *Checkbox) GetCashRegisters(AccessToken string, InUse int, Limit, Offset int) (*CashRegistersResp, *Error) {
	resp := &CashRegistersResp{}
	in_use := ""
	switch InUse {
	case 1:
		in_use = "in_use=true&"
	case 2:
		in_use = "in_use=false&"
	}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       fmt.Sprintf("/api/v1/cash-registers?%slimit=%d&offset=%d", in_use, Limit, Offset),
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// GetCashRegister
//
// Отримання інформації про пРРО
func (ch *Checkbox) GetCashRegister(AccessToken string, CashRegisterId string) (*CashRegisterResp, *Error) {
	resp := &CashRegisterResp{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/cash-registers/" + CashRegisterId,
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}
