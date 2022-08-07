package main

import (
	"fmt"

	"github.com/mattlangl/gophe"
)

func main() {
	fmt.Printf("Input card 1: ")

	c1 := gophe.NewCard("Qc")
	fmt.Printf("Input card 2: ")
	c2 := gophe.NewCard("Qd")

	dealt := 0
	for i := 0; i < dealt; i++ {
		fmt.Printf("Communitty card %d:", i)
	}
	var a uint8
	var o1 uint8
	var wins int = 0
	var losses int = 0
	var ties int = 0
	var total int = 0
	for a = 0; a < 48; a++ {
		if a == c1.ID() || a == c2.ID() {
			continue
		}
		for b := a + 1; b < 49; b++ {
			if b == c1.ID() || b == c2.ID() {
				continue
			}
			for c := b + 1; c < 50; c++ {
				if c == c1.ID() || c == c2.ID() {
					continue
				}
				for d := c + 1; d < 51; d++ {
					if d == c1.ID() || d == c2.ID() {
						continue
					}
					for e := d + 1; e < 52; e++ {
						if e == c1.ID() || e == c2.ID() {
							continue
						}
						cc := gophe.NewHand(
							gophe.NewCardFromId(a),
							gophe.NewCardFromId(b),
							gophe.NewCardFromId(c),
							gophe.NewCardFromId(d),
							gophe.NewCardFromId(e),
						)
						for o1 = 0; o1 < 51; o1++ {
							if o1 == a || o1 == b || o1 == c || o1 == d || o1 == e || o1 == c1.ID() || o1 == c2.ID() {
								continue
							}
							for o2 := o1 + 1; o2 < 52; o2++ {
								if o2 == a || o2 == b || o2 == c || o2 == d || o2 == e || o2 == c1.ID() || o2 == c2.ID() {
									continue
								}
								h1 := cc.AddCards(c1, c2)
								h2 := cc.AddCards(gophe.NewCardFromId(o1), gophe.NewCardFromId(o2))
								r1 := gophe.EvaluateHand(h1)
								r2 := gophe.EvaluateHand(h2)
								total++
								if r1.Compare(r2) > 0 {
									wins++
								} else if r2.Compare(r1) > 0 {
									losses++
								} else {
									ties++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("Wins: %d\n", wins)
	fmt.Printf("Losses: %d\n", losses)
	fmt.Printf("Ties: %d\n", ties)
	fmt.Printf("Total: %d\n", total)
}
