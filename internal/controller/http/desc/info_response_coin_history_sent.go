package desc

// InfoResponseCoinHistorySent ...
type InfoResponseCoinHistorySent struct {
	// Имя пользователя, которому отправлены монеты.
	ToUser string `json:"toUser,omitempty"` //nolint:tagliatelle

	// Количество отправленных монет.
	Amount int32 `json:"amount,omitempty"`
}
