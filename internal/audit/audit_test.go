package audit

import "testing"

func TestTrail(t *testing.T) {
	tr := Trail{}
	if e := tr.Append(NewEvent("1", "d", "create", "u", "x")); e != nil || len(tr.ForEntity("d")) != 1 {
		t.Fatal(e)
	}
}
