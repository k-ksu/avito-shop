package model

type (
	// UserInfo ...
	UserInfo struct {
		CoinHistory CoinHistory
		Coins       int32
		Inventory   []Inventory
	}

	// CoinHistory ...
	CoinHistory struct {
		Received []ReceivedCoins
		Sent     []SentCoins
	}

	// ReceivedCoins ...
	ReceivedCoins struct {
		Amount   int32
		FromUser string
	}

	// SentCoins ...
	SentCoins struct {
		Amount int32
		ToUser string
	}

	// Inventory ...
	Inventory struct {
		Quantity int32
		Type     string
	}
)
