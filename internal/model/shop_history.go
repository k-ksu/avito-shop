package model

import "time"

// ShopHistory ...
type ShopHistory struct {
	ID        int
	UserID    int64
	ItemID    int
	CreatedAt time.Time
}
