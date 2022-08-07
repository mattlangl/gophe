package gophe

import (
	"testing"
)

func TestEvalOne(t *testing.T) {
	h := NewHand(NewCard("Ah"), NewCard("Kh"), NewCard("Qh"), NewCard("Jh"), NewCard("Td"))
	r1 := EvaluateHand(*h)
	c1 := GetRankCategory(r1.GetValue())
	if DescribeRankCategory(c1) != "Straight" {
		t.Errorf("DescribeRankCategory(c) = %s, wanted Straight", DescribeRankCategory(c1))
	}
	h.ModifyHand(NewCard("Th"))
	r2 := EvaluateHand(*h)
	c2 := GetRankCategory(r2.GetValue())
	if DescribeRankCategory(c2) != "Straight Flush" {
		t.Errorf("DescribeRankCategory(c) = %s, wanted Straight", DescribeRankCategory(c2))
	}
}

func BenchmarkEvaluateAllFiveCards(b *testing.B) {
	var a uint8
	for a = 0; a < 48; a++ {
		for q := a + 1; q < 49; q++ {
			for c := q + 1; c < 50; c++ {
				for d := c + 1; d < 51; d++ {
					for e := d + 1; e < 52; e++ {
						EvaluateCards(NewCardFromId(a), NewCardFromId(q), NewCardFromId(c), NewCardFromId(d), NewCardFromId(e))
					}
				}
			}
		}
	}
}

func BenchmarkEvaluateAllSixCards(b *testing.B) {
	var a uint8
	for a = 0; a < 47; a++ {
		for q := a + 1; q < 48; q++ {
			for c := q + 1; c < 49; c++ {
				for d := c + 1; d < 50; d++ {
					for e := d + 1; e < 51; e++ {
						for f := e + 1; f < 52; f++ {
							EvaluateCards(NewCardFromId(a), NewCardFromId(q), NewCardFromId(c), NewCardFromId(d), NewCardFromId(e), NewCardFromId(f))
						}
					}
				}
			}
		}
	}
}

func BenchmarkEvaluateAllSevenCards(b *testing.B) {
	var a uint8
	for a = 0; a < 46; a++ {
		for q := a + 1; q < 47; q++ {
			for c := q + 1; c < 48; c++ {
				for d := c + 1; d < 49; d++ {
					for e := d + 1; e < 50; e++ {
						for f := e + 1; f < 51; f++ {
							for g := f + 1; g < 52; g++ {
								EvaluateCards(NewCardFromId(a), NewCardFromId(q), NewCardFromId(c), NewCardFromId(d), NewCardFromId(e), NewCardFromId(f), NewCardFromId(g))
							}
						}
					}
				}
			}
		}
	}
}
