package pointer

// Bitcoin is an OriginalType of int
type Bitcoin int

// Wallet is a struct
type Wallet struct {
	balance Bitcoin
}

// Deposit add money to account
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance return balance of this account
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
