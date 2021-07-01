package counterG

import "strings"

var Faces = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

var FaceRank = map[string]int{
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

var FaceName = map[int]string{
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "T",
	11: "J",
	12: "Q",
	13: "K",
	14: "A",
	15: "X",
}

var HandRank = map[string]int{
	"皇家同花顺": 10,
	"同花顺":   9,
	"四条":    8,
	"葫芦":    7,
	"同花":    6,
	"顺子":    5,
	"三条":    4,
	"两对":    3,
	"一对":    2,
	"高牌":    1,
}

var HandName = map[int]string{
	10: "皇家同花顺",
	9:  "同花顺",
	8:  "四条",
	7:  "葫芦",
	6:  "同花",
	5:  "顺子",
	4:  "三条",
	3:  "两对",
	2:  "一对",
	1:  "高牌",
}

var FiveHandCount = map[string]int{
	"1001": 8,
	"0110": 7,
	"2010": 4,
	"1200": 3,
	"3100": 2,
	"5000": 1,
}

var ForHandCount = map[string]int{
	"0001": 8,
	"1010": 8,
	"0200": 7,
	"2100": 4,
	"4000": 1,
}

var SevenHandCount = map[string]int{
	"0011": HandRank["四条"],
	"1101": HandRank["四条"],
	"3001": HandRank["四条"],
	"1020": HandRank["葫芦"],
	"2110": HandRank["葫芦"],
	"0210": HandRank["葫芦"],
	"1300": HandRank["两对"],
	//  非确定，都可能为同花、顺子。当以上条件不满足时Map值为最大值
	"4010": HandRank["三条"], //  AAABCDE
	"3200": HandRank["两对"], //  AABBCDE
	"5100": HandRank["一对"], //  AABCDEF
	"7000": HandRank["高牌"], //  ABCDEFG
}

var GhostHandCount = map[string]int{
	"0101": HandRank["四条"],
	"2001": HandRank["四条"],
	"0020": HandRank["四条"],
	"1110": HandRank["四条"],
	"0300": HandRank["葫芦"],
	//  非确定
	"3010": HandRank["四条"], //可能为同花顺、皇家同花顺
	"2200": HandRank["葫芦"], //可能为同花顺、皇家同花顺
	"4100": HandRank["三条"], //可能为顺子、同花、同花顺、皇家同花顺
	"6000": HandRank["两对"], //可能为顺子、同花、同花顺、皇家同花顺
}

func writeString(sb *strings.Builder, args ...string) {
	for _, v := range args {
		sb.WriteString(v)
	}
}

func sort(hand string) string {
	runes := []rune(hand)
	////  insert sort
	l := len(hand)
	for i := 2; i < l; i += 2 {
		for v := 0; v < i; v += 2 {
			if FaceRank[string(runes[v])] < FaceRank[string(runes[i])] {
				runes[v], runes[i] = runes[i], runes[v]
				runes[v+1], runes[i+1] = runes[i+1], runes[v+1]
			}
		}
	}

	//quickSort(runes, 0, len(runes)-2)
	return string(runes)
}

func quickSort(hand []rune, low, high int) {
	if low < high {
		pivot := partition(hand, low, high)
		quickSort(hand, low, pivot-2)
		quickSort(hand, pivot+2, high)
	}
}

func partition(hand []rune, low, high int) int {
	pivot := (hand)[low]
	p2 := hand[low+1]
	for low < high {
		for low < high && FaceRank[string(hand[high])] >= FaceRank[string(hand[low])] {
			high -= 2
		}
		(hand)[low], hand[low+1] = (hand)[high], hand[high+1]
		for low < high && FaceRank[string(hand[low])] <= FaceRank[string(hand[high])] {
			low += 2
		}
		(hand)[high], (hand)[high+1] = (hand)[low], (hand)[low+1]
	}
	(hand)[low] = pivot
	hand[low+1] = p2
	return low
}
