package desc

// InfoResponseInventory ...
type InfoResponseInventory struct {
	// Тип предмета.
	Type string `json:"type,omitempty"`

	// Количество предметов.
	Quantity int32 `json:"quantity,omitempty"`
}
