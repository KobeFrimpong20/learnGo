package main

import (
    "testing"
)

func TestWallet(t *testing.T) {

    t.Run("Testing the Deposit function", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        assertBalance(t, wallet, Bitcoin(10))
    })
    t.Run("Testing the Withdraw function", func(t *testing.T) {
        wallet := Wallet{balance:  Bitcoin(10)}
        err := wallet.Withdraw(Bitcoin(5))

        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(5))
    })
    t.Run("Insufficient funds", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(0)}
        err :=  wallet.Withdraw(Bitcoin(5))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, Bitcoin(0))
    })
}

func assertNoError(t *testing.T, e error) {
    t.Helper()

    if e != nil {
        t.Fatal("expected no error, but received one")
    }
}

func assertError(t *testing.T, got error, want error) {
    t.Helper()

    if got == nil {
        t.Fatal("expected an error, but did not get one")
    }
    if got != want {
        t.Errorf("got %q, but wanted %q", got, want)
    }
}

func assertBalance(t *testing.T, w Wallet, b Bitcoin) {
    t.Helper()
    got := w.Balance()

    if got != b {
        t.Errorf("got %s, but expected %s", got, b)
    }
}
