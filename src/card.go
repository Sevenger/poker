package src

import (
	"strconv"
	"strings"
)

//FaceCount 用于记录手牌每个牌面出现的次数。下标为牌面，值为次数
//一共有13种牌(含鬼牌)，最小牌在map中值为2，最大为15，为了方便计算，数组长度为16
type FaceCount [16]int

type FaceCountCode string

type FaceCountMap map[int][]int

var Sb, S1, S2, S3, S4 strings.Builder

func init() {
	Sb.Grow(5)
	S1.Grow(5)
	S2.Grow(5)
	S3.Grow(5)
	S4.Grow(5)
}

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
	"1001": HandRank["四条"], // 8
	"0110": HandRank["葫芦"], // 7
	"2010": HandRank["三条"], // 4
	"1200": HandRank["两对"], // 3
	"3100": HandRank["一对"], // 2
	//  非确定，可能为顺子、同花、同花顺、皇家同花顺
	"5000": HandRank["高牌"], // 1
}

var SevenHandCountCode = map[string]int{
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

var GhostHandCountCode = map[string]int{
	"0101": HandRank["四条"],
	"2001": HandRank["四条"],
	"0020": HandRank["四条"],
	"1110": HandRank["四条"],
	"0300": HandRank["葫芦"],
	//  非确定
	"3010": HandRank["四条"], //可能为同花顺、皇家同花顺
	"2200": HandRank["葫芦"], //可能为同花顺、皇家同花顺
	"4100": HandRank["三条"], //可能为顺子、同花、同花顺、皇家同花顺
	"6000": HandRank["一对"], //可能为顺子、同花、同花顺、皇家同花顺
}

// Sort 对手牌进行排序，由于手牌最少长4，最多长7，插入排序性能更好
func Sort(hand string) string {
	runes := []rune(hand)
	//  insert sort
	//l := len(hand)
	//for i := 2; i < l; i += 2 {
	//	for v := 0; v < i; v += 2 {
	//		if FaceRank[string(runes[v])] < FaceRank[string(runes[i])] {
	//			runes[v], runes[i] = runes[i], runes[v]
	//			runes[v+1], runes[i+1] = runes[i+1], runes[v+1]
	//		}
	//	}
	//}

	l := len(hand)
	for i := 1; i < l; i++ {
		for v := 0; v < i; v++ {
			if FaceRank[string(runes[v])] < FaceRank[string(runes[i])] {
				runes[v], runes[i] = runes[i], runes[v]
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

func Max(x, y int) (int, bool) {
	if x > y {
		return 1, false
	} else if x < y {
		return 2, false
	} else {
		return 0, true
	}
}

//GetFaceCount 计算手牌中每种牌出现的次数
func GetFaceCount(hand string) FaceCount {
	var count FaceCount
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[hand[i:i+1]]]++
	}
	return count
}

// GetFaceCountCode getFaceCountCode 计算出现1次至4次的牌的code
func GetFaceCountCode(count FaceCount) string {
	Sb.Reset()
	var card1, card2, card3, card4 int
	for i, v := range count {
		//  不记录鬼牌
		if i == 15 {
			continue
		}
		switch v {
		case 0:
			continue
		case 1:
			card1++
		case 2:
			card2++
		case 3:
			card3++
		case 4:
			card4++
		}
	}

	Sb.WriteString(strconv.Itoa(card1))
	Sb.WriteString(strconv.Itoa(card2))
	Sb.WriteString(strconv.Itoa(card3))
	Sb.WriteString(strconv.Itoa(card4))
	return Sb.String()
}

//GetFaceCountMap 返回一个map，该map的key为出现的次数，v是一个[]int，存储了牌面值
func GetFaceCountMap(count FaceCount) map[int][]int {
	countMap := make(map[int][]int, 5)
	countMap[1] = make([]int, 0, 7)
	countMap[2] = make([]int, 0, 3)
	countMap[3] = make([]int, 0, 2)
	countMap[4] = make([]int, 0, 1)
	//  此处i是牌面的值，v是牌面出现的次数，i从14开始，不记录鬼牌
	var i, v int
	for i = 14; i >= 2; i-- {
		v = count[i]
		if v == 0 {
			continue
		} else {
			countMap[v] = append(countMap[v], i)
		}
	}
	return countMap
}

// IsRoyalFlush isRoyalFlush 在保证同花的前提下判断是否可能为皇家同花顺
func IsRoyalFlush(hand string) bool {
	return strings.Compare(hand, "AKQJT") == 0
}
