/*
Package gophe implements a Texas-Holdem Poker Hand Evaluator based
on HenryRLee's implementation using a Perfect Hash Algoritihim. It
currently handles hands of 5-7 cards, and uses pre-computed hash tables
(I reformatted HenryRLee's into go files) to evaluate hands extremely fast.
*/
package gophe

func hashQuinary(q [13]byte, size uint8) int {
	sum := 0
	for i, v := range q {
		sum += int(dp[v][12-i][size])
		size -= v
		if size == 0 {
			break
		}
	}
	return sum
}

func EvaluateHand(h Hand) Rank {
	if suits[h.getSuitHash()] > 0 {
		return Rank{flush[h.getSuitBinary()[suits[h.getSuitHash()]-1]]}
	}
	hash := hashQuinary(h.getQuinary(), h.Size())
	switch h.Size() {
	case 5:
		return Rank{noflush5[hash]}
	case 6:
		return Rank{noflush6[hash]}
	case 7:
		return Rank{noflush7[hash]}
	}
	return Rank{noflush5[hash]}
}

func EvaluateCards(cards ...Card) Rank {
	h := NewHand()
	for _, v := range cards {
		h.ModifyHand(v)
	}
	return EvaluateHand(*h)
}
