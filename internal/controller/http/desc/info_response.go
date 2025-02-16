package desc

// InfoResponse ...
type InfoResponse struct {
	// Количество доступных монет.
	Coins int32 `json:"coins,omitempty"`

	Inventory []InfoResponseInventory `json:"inventory,omitempty"`

	CoinHistory *InfoResponseCoinHistory `json:"coinHistory,omitempty"` //nolint:tagliatelle
}
