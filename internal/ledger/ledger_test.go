package ledger

import "testing"

func TestLedgerRules(t *testing.T) {
	a := NewAccount("100", "Cash", "asset")
	e := NewEntry("1", "100", "sale", 2)
	if a.Validate() != nil || e.Validate() != nil || !Rule1(a, e, Period{Year: 2025, Month: 1}) {
		t.Fatal()
	}
}
