package cycle

import "testing"

func TestCycle(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	if !IsCycle(a) {
		t.Errorf("expected cycle,but not so")
	}
}

// TODO
func TestNotCycle(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b := &link{value: "a"}, &link{value: "b"}
	a.tail = b
	if IsCycle(a) {
		t.Errorf("expected not cycle,but cycle")
	}
}
