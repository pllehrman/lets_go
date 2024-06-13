package wallet

import (
	"errors"
	"fmt"
)

// This method of depositing and then trying to print the balance illustrates the need to reference specific addresses in memory
// Rather than taking a copy of the whole Wallet, we instead take a pointer ot that wallet so that we can change its original values.

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// the "w" which references Wallet does not need to be dereferenced since it's a "struct pointer" and they are automatically dereferenced.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Note the casing on this publically available variable with GLOBAL scope
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}