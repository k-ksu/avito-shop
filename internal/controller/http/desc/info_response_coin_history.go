package desc

// InfoResponseCoinHistory ...
type InfoResponseCoinHistory struct {
	Received []InfoResponseCoinHistoryReceived `json:"received,omitempty"`

	Sent []InfoResponseCoinHistorySent `json:"sent,omitempty"`
}
