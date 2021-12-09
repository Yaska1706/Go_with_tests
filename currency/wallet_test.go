package currency

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		checkWallet(t, wallet, Bitcoin(10))

	})
	t.Run("Normal Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(5))

		asserNoError(t, err)
		checkWallet(t, wallet, Bitcoin(15))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingbalance := Bitcoin(30)
		Wallet := Wallet{balance: startingbalance}
		err := Wallet.Withdraw(Bitcoin(50))

		assertError(t, err, ErrInsuffientFunds.Error())
		checkWallet(t, Wallet, startingbalance)

	})
}

func checkWallet(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func asserNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Unexpected error")
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()

	if got == nil {
		t.Error("Expected an error")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
