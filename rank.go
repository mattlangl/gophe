package gophe

type RankCategory uint8

const (
	StraightFlush RankCategory = iota + 1
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

var RankCategoryDescription = [11]string{
	"s",
	"Straight Flush",
	"Four Of a Kind",
	"Full House",
	"Flush",
	"Straight",
	"Three Of a Kind",
	"Two Pair",
	"One Pair",
	"High Card",
}

func GetRankCategory(rank uint16) RankCategory {
	switch true {
	case rank > 6185:
		return HighCard
	case rank > 3325:
		return OnePair
	case rank > 2467:
		return TwoPair
	case rank > 1609:
		return ThreeOfAKind
	case rank > 1599:
		return Straight
	case rank > 322:
		return Flush
	case rank > 166:
		return FullHouse
	case rank > 10:
		return FourOfAKind
	default:
		return StraightFlush
	}
}

func DescribeRankCategory(category RankCategory) string {
	return RankCategoryDescription[category]
}

func DescribeRank(rank uint16) string {
	return RankDescription[rank][1]
}

func DescribeSampleHand(rank uint16) string {
	return RankDescription[rank][0]
}

func IsFlush(rank uint16) bool {
	switch GetRankCategory(rank) {
	case StraightFlush:
		return true
	case Flush:
		return true
	default:
		return false
	}
}

type Rank struct {
	value uint16
}

func (a Rank) Compare(b Rank) int {
	return int(a.value) - int(b.value)
}

func (r Rank) GetValue() uint16 {
	return r.value
}

func (r Rank) GetCategory() RankCategory {
	return GetRankCategory(r.GetValue())
}

func (r Rank) DescribeCategory() string {
	return DescribeRankCategory(r.GetCategory())
}

func (r Rank) DescribeRank() string {
	return DescribeRank(r.GetValue())
}

func (r Rank) DescribeSampleHand() string {
	return DescribeSampleHand(r.GetValue())
}

func (r Rank) IsFlush() bool {
	return IsFlush(r.GetValue())
}
