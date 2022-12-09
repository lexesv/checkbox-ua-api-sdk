package checkbox

import "fmt"

// CreateShift
//
// Відкриття нової зміни касиром.
//
// Для створення зміни необхідно вказати ключ ліцензії конкретного пРРО розташованого в тій же торговій точці, що знаходиться касир.
//
// Створюється об'єкт зміни в стані "CREATED" та транзакція відкриття зміни (поле "initial_transaction").
//
// Для переведення зміни в статус OPENED необхідно щоб транзакція була підписана за допомогою КЕП та доставлена в ДПС, це як правило триває декілька секунд.
//
// Після створення запиту необхідно відслідковувати статус зміни доки він не змінить на OPENED або CLOSED.
//
// Статус зміни можна відслідковувати за допомогою GET запиту по шляху /api/v1/shifts/{shift_id}, де {shift_id} - ідентифікатор зміни. Або /api/v1/cashier/shift
//
// У разі, якщо зміна перейшла у статус CLOSED - це значить, що зміна не може бути відкрита. Деталізація причини відмови у створенні зміни знаходиться в полі initial_transaciton.
//
// Після того як робочу зміну буде успішно відкрито можна її закрити або створювати чеки.
//
// AccessToken: Bearer <токен авторизації>
func (ch *Checkbox) CreateShift(AccessToken string) (*ShiftResp, *Error) {
	resp := &ShiftResp{}
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: true,
		Endpoint:       "/api/v1/shifts",
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// Shifts
//
// Отримання змін поточного касира
func (ch *Checkbox) Shifts(AccessToken string, Desc bool, Limit, Offset int) (*ShiftsResp, *Error) {
	resp := &ShiftsResp{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       fmt.Sprintf("/api/v1/shifts?desc=%t&limit=%d&offset=%d", Desc, Limit, Offset),
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// GetShift
//
// Отримання інформації про зміну
func (ch *Checkbox) GetShift(AccessToken, ShiftId string) (*ShiftResp, *Error) {
	resp := &ShiftResp{}
	c := ReqConfig{
		Method:         "GET",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/shifts/" + ShiftId,
		AccessToken:    AccessToken,
		Request:        nil,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}

// CloseShift
//
// Приклад: CloseShift(AccessToken, nil)
//
// Якщо аргумент req = nil то всі дані по Z-звіту будуть сформовані на боці серверу
//
// Створення Z-Звіту та закриття поточної зміни користувачем (касиром).
//
// Стан зміни встановлюється як "CLOSING" та створюється транзакція закриття зміни (поле "closing_transaction").
//
// Для переведення зміни в статус CLOSED необхідно щоб транзакція була підписана за допомогою КЕП та доставлена в ДПС.
//
// Статус зміни можна відслідковувати за допомогою GET запиту по шляху /api/v1/shifts/{shift_id}, де {shift_id} - ідентифікатор зміни.
//
// Після закриття зміни в її рамках більше не можливо буде виконувати дії. Для продовження роботи потрібно створити нову зміну.
//
// # Опціонально Z-Звіт може бути сформований на стороні клієнта та переданий в тілі цього запиту
//
// Увага! При формуванні звіту на стороні клієнту перевірка коректності розрахунку оборотів та сум продажу не виконується!
func (ch *Checkbox) CloseShift(AccessToken string, req *CloseShiftReq) (*ShiftResp, *Error) {
	resp := &ShiftResp{}
	c := ReqConfig{
		Method:         "POST",
		NeedLicenseKey: false,
		Endpoint:       "/api/v1/shifts/close",
		AccessToken:    AccessToken,
		Request:        req,
		Response:       resp,
	}
	err := ch.Request(c)
	return resp, err
}
