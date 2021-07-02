package fivehand

import (
	. "poker/src"
	"strings"
)

type FiveHand struct{}

func (c *FiveHand) Start(hand1, hand2 string) int {
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

		default:
			return c.equalJudgePair(faceCount1, faceCount2, rank1)
		}
	}

}

//getFiveHandRank 得到牌型
func (c *FiveHand) getFiveHandRank(hand string, faceCount FaceCount) (int, string) {
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
func (c *FiveHand) isFlush(hand string) (string, bool) {
	Sb.Reset()
	for i := 0; i < len(hand); i += 2 {
		Sb.WriteString(hand[i : i+1])
	}
	if hand[1] == hand[3] &&
		hand[3] == hand[5] &&
		hand[5] == hand[7] &&
		hand[9] == hand[10] {
		return Sb.String(), true
	} else {
		return Sb.String(), false
	}
}

//  判断无重复字符串是否是顺子，返回排序后的手牌
func (c *FiveHand) isStraight(hand string) (string, bool) {
	hand = Sort(hand)
	isStraight := false
	//  由于是无重复子串，通过头部减去尾部可以直接判断是不是顺子
	head, tail := hand[0:1], hand[4:5]
	if FaceRank[head]-FaceRank[tail] == 4 {
		isStraight = true
	}

	//  A5432的特殊情况
	if strings.Compare(hand, "A5432") == 0 {
		isStraight = true
	}

	return hand, isStraight
}

func (c *FiveHand) equalJudgeHighCard(count1, count2 FaceCount) int {
	m1, m2 := c.getFaceCountMap(count1)[1], c.getFaceCountMap(count2)[1]

	rst := 0
	var v1, v2 int
	for i := 0; i < 5; i++ {
		v1, v2 = m1[i], m2[i]
		if v1 > v2 {
			rst = 1
			break
		} else if v1 < v2 {
			rst = 2
			break
		}
	}
	return rst
}

func (c *FiveHand) equalJudgeFlush(hand1, hand2 string, isGhost1, isGhost2 bool) int {
	if isGhost1 {
		hand1 = c.fillFlush(hand1)
	}

	if isGhost2 {
		hand2 = c.fillFlush(hand2)
	}

	rst := 0
	var v1, v2 int
	for i := 0; i < 5; i++ {
		v1, v2 = FaceRank[hand1[i:i+1]], FaceRank[hand2[i:i+1]]
		if v1 > v2 {
			rst = 1
			break
		} else if v1 < v2 {
			rst = 2
			break
		}
	}
	return rst
}

//fillFlush 4手牌时填充为5手牌同花，5手牌时换掉一个牌成为最大牌(有鬼牌)
//要求传入的手牌是有序的
func (c *FiveHand) fillFlush(hand string) string {
	for i := 14; i >= 2; i-- {
		insertFace := FaceName[i]
		if !strings.Contains(hand, insertFace) {
			sb.Reset()
			for j := 0; j < len(hand); j++ {
				face := hand[j : j+1]
				if FaceRank[insertFace] > FaceRank[face] {
					sb.WriteString(insertFace)
					sb.WriteString(hand[j:])
					break
				} else {
					sb.WriteString(face)
				}
			}
			break
		}
	}
	newHand := sb.String()
	if len(newHand) >= 5 {
		newHand = newHand[0:5]
	}
	return newHand
}

func (c *FiveHand) equalJudgePair(count1, count2 FaceCount, rank int) int {
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

			}
		}

		v1, v2 = v1s[0], v2s[0]

		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = v1s[1], v2s[1]
			if rst, isEqual = Max(v1, v2); isEqual {
				//  可能有3个两对，即AABBCCD牌型
				var v3, v4 int
				if len(v1s) == 3 {
					v3 = v1s[2]
				}
				v4 = map1[1][0]
				if v3 > v4 {
					v1 = v3
				} else {
					v1 = v4
				}

				if len(v2s) == 3 {
					v3 = v2s[2]
				}
				v4 = map2[1][0]
				if v3 > v4 {
					v2 = v3
				} else {
					v2 = v4
				}

				rst, _ = Max(v1, v2)
			}
		}

	case HandRank["三条"]:
		//  鬼牌为三条时将出现两次牌填充至三次即可
		if isGhost1 {
			v1 = map1[2][0]
			map1[2] = map1[2][1:] //去除
			map1[3] = append(map1[3], v1)
		}
		if isGhost2 {
			v1 = map2[2][0]
			map2[2] = map2[2][1:] //去除
			map2[3] = append(map2[3], v1)
		}

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
		//  鬼牌为葫芦时缺少一个出现三次牌，将一个两次牌填充至三次牌即可
		if isGhost1 {
			v1 = map1[2][0]
			map1[2] = map1[2][1:]
			map1[3] = append(map1[3], v1)
		}
		if isGhost2 {
			v1 = map2[2][0]
			map2[2] = map2[2][1:]
			map2[3] = append(map2[3], v1)
		}

		v1, v2 = map1[3][0], map2[3][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			//  7手牌时候葫芦可能没有两次牌，如AAABBBC
			if len(map1[2]) == 0 {
				v1 = map1[3][1]
			} else {
				v1 = map1[2][0]
			}
			if len(map2[2]) == 0 {
				v2 = map2[3][1]
			} else {
				v2 = map2[2][0]
			}
			rst, _ = Max(v1, v2)
		}

	case HandRank["四条"]:
		if isGhost1 && len(map1[4]) == 0 {
			v1 = map1[3][0]
			map1[3] = map1[3][1:] //去除
			map1[4] = append(map1[4], v1)
		}
		if isGhost2 && len(map2[4]) == 0 {
			v1 = map2[3][0]
			map2[3] = map2[3][1:]
			map2[4] = append(map2[4], v1)
		}

		v1, v2 = map1[4][0], map2[4][0]
		if rst, isEqual = Max(v1, v2); isEqual {
			v1, v2 = map1[1][0], map2[1][0]
			rst, _ = Max(v1, v2)
		}
	}
	return rst
}
