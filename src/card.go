package src

var Face = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
	"X": 15,
}

var Suit = map[rune]int{
	's': 1,
	'h': 2,
	'd': 3,
	'c': 4,
	'n': 5,
}

var HandRank = map[string]int{
	"皇家同花顺": 10,
	"同花顺":   9,
	"4条":    8,
	"葫芦":    7,
	"同花":    6,
	"顺子":    5,
	"3条":    4,
	"两对":    3,
	"一对":    2,
	"高牌":    1,
}

var HandName = map[int]string{
	10: "皇家同花顺",
	9:  "同花顺",
	8:  "4条",
	7:  "葫芦",
	6:  "同花",
	5:  "顺子",
	4:  "3条",
	3:  "两对",
	2:  "一对",
	1:  "高牌",
}
