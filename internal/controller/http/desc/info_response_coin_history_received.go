package desc

// InfoResponseCoinHistoryReceived ...
type InfoResponseCoinHistoryReceived struct {
	// Имя пользователя, который отправил монеты.
	FromUser string `json:"fromUser,omitempty"` //nolint:tagliatelle

	// Количество полученных монет.
	Amount int32 `json:"amount,omitempty"`
}
