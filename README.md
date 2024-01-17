# GOPHE: Go-Based Poker Hand Evaluator

GOPHE is a Texas Hold'em Poker Hand Evaluator, inspired by HenryRLee's Poker Hand Evaluator and implemented in Go. This library efficiently handles poker hand evaluations for 5 to 7 card hands using a Perfect Hash Algorithm. The implementation makes use of pre-computed hash tables for rapid and efficient hand evaluation, suited for both casual use and high-performance applications.

## Features
- Evaluates Texas Hold'em hands of 5-7 cards.
- Utilizes a Perfect Hash Algorithm for fast and accurate results.
- Pre-computed hash tables integrated for optimized performance.

## Usage
To use GOPHE in your Go project, import the package and create a new hand with the desired cards. Then, use the `EvaluateHand` or `EvaluateCards` functions to get the rank of the hand.

```go
package main

import (
    "fmt"
    "gophe"
)

func main() {
    // Create a new hand
    hand := gophe.NewHand(cards...)
    // Evaluate the hand
    rank := gophe.EvaluateHand(*hand)

    fmt.Println("Hand Rank:", rank.DescribeRank())
}
```
## Contributing
Contributions to GOPHE are welcome. Please feel free to submit issues and pull requests to the repository.
