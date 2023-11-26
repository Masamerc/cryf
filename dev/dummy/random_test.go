package dummy

import "testing"

func TestRandomDate(t *testing.T) {
	dateRange := "all"
	date := RandomDate(dateRange)
	if len(date) != 10 {
		t.Errorf("Expected date length of 10, got %d", len(date))
	}
}

func TestRandomAmount(t *testing.T) {
	amount := RandomAmount(SALES)
	// if it's SALES, the amount should be between 1 and 100000
	if amount < 1 || amount > 100000 {
		t.Errorf("Expected amount between 1 and 100000, got %d", amount)
	}

}
