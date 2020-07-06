package pointer

// Wallet is a struct
type Wallet struct {
	balance int
}

// Deposit add money to account
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance return balance of this account
func (w *Wallet) Balance() int {
	return w.balance
}
