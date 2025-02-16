package errs

import "errors"

// ErrNoRows ...
var ErrNoRows = errors.New("no rows in result set")

// ErrInvalidPassword ...
var ErrInvalidPassword = errors.New("invalid password")

// ErrNotEnoughMoney ...
var ErrNotEnoughMoney = errors.New("not enough money")

// ErrUserNotExists ...
var ErrUserNotExists = errors.New("user does not exist")

// ErrNoSuchMerch ...
var ErrNoSuchMerch = errors.New("no such merch")
