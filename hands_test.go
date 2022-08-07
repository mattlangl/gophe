package gophe

import "testing"

func TestNewHand(t *testing.T) {
	h := NewHand()
	if h.Size() != 0 {
		t.Errorf("h.Size() = %d, wanted 0", h.Size())
	}
	h.ModifyHand(NewCardFromId(1), NewCardFromId(2))
	if h.Size() != 2 {
		t.Errorf("h.Size() = %d, wanted 2", h.Size())
	}
	b := NewHand(NewCardFromId(2))
	if b.Size() != 1 {
		t.Errorf("b.Size() = %d, wanted 2", b.Size())
	}
}
