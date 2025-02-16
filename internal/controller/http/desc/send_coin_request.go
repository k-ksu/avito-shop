package desc

// SendCoinRequest ...
type SendCoinRequest struct {
	// Имя пользователя, которому нужно отправить монеты.
	ToUser string `json:"toUser"` //nolint:tagliatelle

	// Количество монет, которые необходимо отправить.
	Amount int32 `json:"amount"`
}
