package checkbox

// SignIn
//
// Авторизація за допомогою логіна і пароля касира
func (ch *Checkbox) SignIn() (*SignIn, *Error) {
	req := &SingInPayload{}
	resp := &SignIn{}
	req.Login = ch.Conf.CashierLogin
	req.Password = ch.Conf.CashierPassword
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/cashier/signin",
		Request:        req,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// SignInPinCode
//
// Авторизація за допомогою ключа ліцензії каси та PIN кода касира
func (ch *Checkbox) SignInPinCode() (*SignIn, *Error) {
	req := &SingInPayload{}
	resp := &SignIn{}
	req.PinCode = ch.Conf.CashierPinCode
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: true,
		Endpoint:       "/api/v1/cashier/signinPinCode",
		Request:        req,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// SignOut
//
// Завершення сесії користувача
func (ch *Checkbox) SignOut(AccessToken string) *Error {
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/cashier/signout",
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       nil,
	}
	err := ch.Request(c)
	return err
}

// GetCashierProfile
//
// Отримання інформації про поточного користувача (касира)
func (ch *Checkbox) GetCashierProfile(AccessToken string) (*CashierProfile, *Error) {
	resp := &CashierProfile{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/cashier/me",
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// GetCashierShift
//
// Отримання інформації про активну зміну користувача (касира)
func (ch *Checkbox) GetCashierShift(AccessToken string) (*Shift, *Error) {
	resp := &Shift{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/cashier/shift",
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// GetCashierSignature
//
// Check Signature
func (ch *Checkbox) GetCashierSignature(AccessToken string) (*CashierSignature, *Error) {
	resp := &CashierSignature{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/cashier/check-signature",
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}
