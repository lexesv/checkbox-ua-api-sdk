package checkbox

type Checkbox struct {
	Conf Config
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
