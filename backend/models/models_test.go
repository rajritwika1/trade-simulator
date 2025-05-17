package models

import "testing"

func TestEstimateSlippage(t *testing.T) {
	got := EstimateSlippage(50000, 0.002, 0.02)
	want := 1.0

	if got != want {
		t.Errorf("Expected slippage %.2f, got %.2f", want, got)

	}
}

func TestEstimateFees(t *testing.T) {
	got := EstimatesFees(50000, 0.002, 0.001)
	want := 0.1
	if got != want {
		t.Errorf("Expected fees %.2f, got %.2f", want, got)
	}
}

func TestEstimateMarketImpact(t *testing.T) {
	got := EstimatesMarketImpact(0.002, 0.02)
	want := 0.00004

	if got != want {
		t.Errorf("Expected impact %.5f, got %.5f", want, got)
	}
}
