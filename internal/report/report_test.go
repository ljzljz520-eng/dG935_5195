package report

import "testing"

func TestReport(t *testing.T) {
	r := New("r", "Summary")
	r.Add("a", 2)
	r.Add("b", 3)
	if r.Total() != 5 || len(r.Render()) == 0 {
		t.Fatal()
	}
}
