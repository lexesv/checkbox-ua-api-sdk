package checkbox

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func New(LicenseKey, CashierLogin, CashierPassword, CashierPinCode string) *Checkbox {
	ch := &Checkbox{}
	ch.Conf.ApiUrl = "https://api.checkbox.ua"
	ch.Conf.LicenseKey = LicenseKey
	ch.Conf.CashierLogin = CashierLogin
	ch.Conf.CashierPassword = CashierPassword
	ch.Conf.CashierPinCode = CashierPinCode
	return ch
}

func (ch *Checkbox) SetApiUrl(ApiUrl string) {
	ch.Conf.ApiUrl = ApiUrl
}
func (ch *Checkbox) SetAccessKey(AccessKey string) {
	ch.Conf.AccessKey = AccessKey
}
func (ch *Checkbox) SetClientName(ClientName string) {
	ch.Conf.ClientName = ClientName
}
func (ch *Checkbox) SetClientVersion(ClientVersion string) {
	ch.Conf.ClientVersion = ClientVersion
}

// request
// Http запит до сервера API
func (ch *Checkbox) request(c ReqConfig) *Error {
	Error := new(Error)
	body := new(bytes.Reader)
	if c.Request != nil {
		b, err := json.Marshal(c.Request)
		if err != nil {
			Error.addMsg(err)
			return Error
		}
		body = bytes.NewReader(b)
	}
	req, err := http.NewRequest(c.Method, ch.Conf.ApiUrl+c.Endpoint, body)
	if err != nil {
		Error.addMsg(err)
		return Error
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if ch.Conf.ClientName != "" {
		req.Header.Set("X-Client-Name", ch.Conf.ClientName)
	}
	if ch.Conf.ClientVersion != "" {
		req.Header.Set("X-Client-Version", ch.Conf.ClientVersion)
	}
	if ch.Conf.AccessKey != "" {
		req.Header.Set("X-Access-Key", ch.Conf.AccessKey)
	}
	if c.NeedLicenseKey {
		req.Header.Set("X-License-Key", ch.Conf.LicenseKey)
	}
	if c.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Error.addMsg(err)
		return Error
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		Error.addMsg(err)
		return Error
	}
	if resp.StatusCode == 200 || resp.StatusCode == 202 || resp.StatusCode == 201 {
		if c.Response != nil {
			err = json.Unmarshal(b, c.Response)
		}
	} else {
		err = json.Unmarshal(b, Error)
		if err != nil {
			Error.addMsg(err)
			return Error
		}
		return Error
	}
	return nil
}
