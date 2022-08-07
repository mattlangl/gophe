package gophe

var rankMap = map[byte]uint8{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}
var suitMap = map[byte]uint8{
	'C': 0,
	'c': 0,
	'D': 1,
	'd': 1,
	'H': 2,
	'h': 2,
	'S': 3,
	's': 3,
}

var rankReverseArray = [13]byte{
	'2', '3', '4', '5', '6', '7', '8',
	'9', 'T', 'J', 'Q', 'K', 'A',
}

var suitReverseArray = [4]byte{'c', 'd', 'h', 's'}

type Card uint8

func NewCardFromId(id uint8) Card {
	return Card(id)
}

func NewCard(name string) Card {
	return NewCardFromId((rankMap[name[0]] << 2) + suitMap[name[1]])
}

func (c Card) Rank() byte {
	return rankReverseArray[c>>2]
}

func (c Card) Suit() byte {
	return suitReverseArray[c&3]
}

func (c Card) ToString() string {
	return string([]byte{c.Rank(), c.Suit()})
}

func (c Card) ID() uint8 {
	return uint8(c)
}
