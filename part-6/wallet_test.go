package wallet

import (
	"testing"
)

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, 10)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(100)
		assertBalance(t, wallet, 20)
		assertError(t, err, InsufficientFundsError)
	})
}
