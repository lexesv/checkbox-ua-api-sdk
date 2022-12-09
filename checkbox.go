package checkbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Checkbox struct {
	Conf         Config
	SuccessCodes []int
}

type Config struct {
	ApiUrl          string
	AccessKey       string
	LicenseKey      string
	ClientName      string
	ClientVersion   string
	CashierLogin    string
	CashierPassword string
	CashierPinCode  string
}

type ReqConfig struct {
	Method         string
	NeedLicenseKey bool
	Endpoint       string
	AccessToken    string //Bearer <токен авторизації>
	Request        interface{}
	Response       interface{}
}

type Error struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message,omitempty"`
	Detail     MessageDetail `json:"detail,omitempty"`
}

// MessageDetail  - У випадку, якщо ваш запит не пройде валідацію формату, ви отримаєте 422 Error:
// Unprocessable Entity зі змістом наступного вигляду, де буде вказано приблизне розташування та опис помилки:
type MessageDetail []struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
	Ctx  struct {
		LimitValue int `json:"limit_value"`
	} `json:"ctx"`
}

func (e *Error) addMsg(err error) {
	e.Message = err.Error()
}
func (e *Error) Error() string {
	return fmt.Sprintf("%d. %s. %#v", e.StatusCode, e.Message, e.Detail)
}

func New(LicenseKey, CashierLogin, CashierPassword, CashierPinCode string) *Checkbox {
	ch := &Checkbox{
		SuccessCodes: []int{200, 201, 202, 205},
	}
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
// Error status codes:
// 403 - Not authenticated
// 422 - Validation Error
func (ch *Checkbox) request(c ReqConfig) (Error *Error) {
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

	if ch.checkStatusCode(resp.StatusCode) {
		if c.Response != nil {
			err = json.Unmarshal(b, c.Response)
		}
	} else {
		err = json.Unmarshal(b, Error)
		if err != nil {
			Error.addMsg(err)
			return Error
		}
		Error.StatusCode = resp.StatusCode
		return Error
	}
	return nil
}

func (ch *Checkbox) checkStatusCode(code int) bool {
	for _, v := range ch.SuccessCodes {
		if v == code {
			return true
		}
	}
	return false
}
