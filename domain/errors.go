package domain

import "errors"

var (
	ErrIINIncorect       = errors.New("incorect fomat of IIN")
	ErrZeroAmount        = errors.New("amount should be more than 0")
	ErrIINTransferDenied = errors.New("you do not own the wallet")
)
