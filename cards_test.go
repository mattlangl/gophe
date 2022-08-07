package gophe

import "testing"

func TestNewCard(t *testing.T) {
	if uint8(NewCardFromId(2)) != 2 {
		t.Errorf("NewCard(2) = %d, wanted 2", uint8(NewCardFromId(2)))
	}
}

func TestNewCardFromString(t *testing.T) {
	if uint8(NewCard("2H")) != 2 {
		t.Errorf("NewCardFromString(\"2H\") = %d, wanted 2", uint8(NewCard("2H")))
	}
	if uint8(NewCard("KS")) != 47 {
		t.Errorf("NewCardFromString(\"KS\") = %d, wanted 47", uint8(NewCard("KS")))
	}
}

func TestCardToString(t *testing.T) {
	if NewCard("KS").ToString() != "Ks" {
		t.Errorf("NewCardFromString(\"KS\").ToString() = %s, wanted \"KS\"", NewCard("KS").ToString())
	}

}
