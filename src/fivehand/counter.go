package fivehand

import (
	. "poker/src"
	"strings"
)

type Counter struct{}

func (c *Counter) Start(hand1, hand2 string) int {
	faceCount1, faceCount2 := GetFaceCount(hand1), GetFaceCount(hand2)

	rank1, newHand1 := c.getFiveHandRank(hand1, faceCount1)
	rank2, newHand2 := c.getFiveHandRank(hand2, faceCount2)

	if rank1 > rank2 {
		return 1
	} else if rank1 < rank2 {
		return 2
	} else {
		switch rank1 {

		//  高牌和同花需要依次判断face
		case HandRank["高牌"]:
			fallthrough
		case HandRank["同花"]:
			for i := 0; i < 5; i++ {
				v1, v2 := FaceRank[newHand1[i:i+1]], FaceRank[newHand2[i:i+1]]
				if v1 > v2 {
					return 1
				} else if v1 < v2 {
					return 2
				}
			}
			return 0

		//  顺子和同花顺只需要判断头牌
		case HandRank["顺子"]:
			fallthrough
		case HandRank["同花顺"]:
			rst, _ := Max(FaceRank[newHand1[0:1]], FaceRank[newHand2[0:1]])
			return rst

		case HandRank["皇家同花顺"]:
			return 0

		default:
			return c.equalJudgePair(faceCount1, faceCount2, rank1)
		}
	}
}

//getFiveHandRank 得到牌型
func (c *Counter) getFiveHandRank(hand string, faceCount FaceCount) (int, string) {
	code := GetFaceCountCode(faceCount)
	rank := FiveHandCount[code]
	newHand := ""

	//  该牌型不确定，可能是顺子、同花、同花顺、皇家同花顺
	if code == "5000" {
		var straight string
		//  同花的点数大于顺子，所以优先判断是不是同花
		if flush, is := c.isFlush(hand); is {
			//  如果是则判断是不是顺子
			if straight, is = c.isStraight(flush); is {
				if IsRoyalFlush(straight) {
					rank = HandRank["皇家同花顺"]
				} else {
					rank = HandRank["同花顺"]
				}
			} else {
				rank = HandRank["同花"]
			}
		} else if straight, is = c.isStraight(flush); is {
			rank = HandRank["顺子"]
		}
		newHand = straight
	}
	return rank, newHand
}

//isFlush 判断是否可能是同花，顺便把花色去了
func (c *Counter) isFlush(hand string) (string, bool) {
	Sb.Reset()
	for i := 0; i < len(hand); i += 2 {
		Sb.WriteString(hand[i : i+1])
	}
	if hand[1] == hand[3] &&
		hand[3] == hand[5] &&
		hand[5] == hand[7] &&
		hand[7] == hand[9] {
		return Sb.String(), true
	} else {
		return Sb.String(), false
	}
}

//  判断无重复字符串是否是顺子，返回排序后的手牌
func (c *Counter) isStraight(hand string) (string, bool) {
	hand = Sort(hand)
	isStraight := false
	//  由于是无重复子串，通过头部减去尾部可以直接判断是不是顺子
	head, tail := hand[0:1], hand[4:5]
	if FaceRank[head]-FaceRank[tail] == 4 {
		isStraight = true
	}

	//  A5432的特殊情况
	if strings.Compare(hand, "A5432") == 0 {
		hand = "5432A"
		isStraight = true
	}

	return hand, isStraight
}

func (c *Counter) equalJudgePair(count1, count2 FaceCount, rank int) int {
	map1 := GetFaceCountMap(count1)
	map2 := GetFaceCountMap(count2)
	var rst, v1, v2 int
	var isEqual bool
	switch rank {
	case HandRank["一对"]:
		v1, v2 = map1[2][0], map2[2][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			vs1, vs2 := map1[1], map2[1]
			for i := 0; i < 3; i++ {
				v1 = vs1[i]
				v2 = vs2[i]
				if v1 > v2 {
					rst = 1
					break
				} else if v1 < v2 {
					rst = 2
					break
				}
			}
		}

	case HandRank["两对"]:
		v1, v2 = map1[2][0], map2[2][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[2][1], map2[2][1]
			if rst, isEqual = Max(v1, v2); isEqual {
				v1, v2 = map1[1][0], map1[2][0]
				rst, _ = Max(v1, v2)
			}
		}

	case HandRank["三条"]:
		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1s, v2s := map1[1], map2[1]
			for i := 0; i < 2; i++ {
				v1, v2 = v1s[i], v2s[i]
				if v1 > v2 {
					rst = 1
					break
				} else if v1 < v2 {
					rst = 2
					break
				}
			}
		}

	case HandRank["葫芦"]:
		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1 = map1[2][0]
			v2 = map2[2][0]
			rst, _ = Max(v1, v2)
		}

	case HandRank["四条"]:
		v1, v2 = map1[4][0], map2[4][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[1][0], map2[1][0]
			rst, _ = Max(v1, v2)
		}
	}
	return rst
}
