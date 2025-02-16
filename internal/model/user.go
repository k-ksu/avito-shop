package model

// User ...
type User struct {
	ID                 int64
	Name               string
	ObfuscatedPassword string
	Coins              int32
}
