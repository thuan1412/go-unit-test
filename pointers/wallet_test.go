package pointers

import (
	"testing"
)

func TestWallet_Withdraw(t *testing.T) {
	tests := []struct {
		name   string
		wallet Wallet
		want   float64
	}{
		{
			name:   "Withdraw",
			wallet: Wallet{balance: 100},
			want:   90,
		},
	}
}

func TestWallet_Balance(t *testing.T) {
	tests := []struct {
		name   string
		wallet Wallet
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet := Wallet{}
			if got := wallet.Balance(); got != tt.want {
				t.Errorf("Wallet.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet_Deposit(t *testing.T) {
	tests := []struct {
		name   string
		wallet Wallet
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallet := Wallet{}
			if got := wallet.Deposit(); got != tt.want {
				t.Errorf("Wallet.Deposit() = %v, want %v", got, tt.want)
			}
		})
	}
}
