package counterG

import (
	"strconv"
	"strings"
)

//counterG 计算7手牌和鬼牌
type counterG struct{}

func (c *counterG) Start(hand1, hand2 string) int {
	faceCount1, faceCount2 := c.getHandFaceCount(hand1), c.getHandFaceCount(hand2)

	var rank1, rank2 int
	var newHand1, newHand2 string
	rank1, newHand1 = c.getSevenHandFaceRank(hand1, faceCount1)
	rank2, newHand2 = c.getSevenHandFaceRank(hand2, faceCount2)
	if rank1 > rank2 {
		return 1
	} else if rank1 < rank2 {
		return 2
	} else {
		switch rank1 {
		case HandRank["高牌"]:
			fallthrough
		case HandRank["顺子"]:
			fallthrough
		case HandRank["同花"]:
			fallthrough
		case HandRank["同花顺"]:
			if len(newHand1) != 0 {
				hand1 = newHand1
			}
			if len(newHand2) != 0 {
				hand2 = newHand2
			}
			hand1, hand2 = sort(hand1), sort(hand2)
			return c.equalJudgeStraight(sort(hand1), hand2)

		default:
			if len(newHand1) != 0 {
				faceCount1 = c.getHandFaceCount(newHand1)
			}
			if len(newHand2) != 0 {
				faceCount2 = c.getHandFaceCount(newHand2)
			}
			return c.equalJudgePair(faceCount1, faceCount2, rank1)
		}
	}
}

//isRoyalFlush 判断是否可能为皇家同花顺
func (*counterG) isRoyalFlush(hand string) bool {
	return hand[0] == 'A' && hand[2] == 'K' && hand[4] == 'Q' && hand[6] == 'J' && hand[8] == 'T'
	//return strings.Contains(hand, "A") &&
	//	strings.Contains(hand, "K") &&
	//	strings.Contains(hand, "Q") &&
	//	strings.Contains(hand, "J") &&
	//	strings.Contains(hand, "T")
}

//getHandFaceCount 计算手牌中每种牌出现的次数
func (*counterG) getHandFaceCount(hand string) [15]int {
	//  一共有12种牌，最小牌在map中值为2，最大为14，为了方便计算，数组长度为15
	count := [15]int{}
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[hand[i:i+1]]] += 1
	}

	//runes := []rune(hand)
	//for i := 0; i < len(runes); i += 2 {
	//	count[runes[i]]++
	//}
	return count
}

//getHandFaceCountTable 计算出现1次至4次的牌的情况
func (*counterG) getHandFaceCountTable(count [15]int) string {
	var card1, card2, card3, card4 int
	for _, v := range count {
		switch v {
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

	sb := new(strings.Builder)
	writeString(sb, strconv.Itoa(card1), strconv.Itoa(card2), strconv.Itoa(card3), strconv.Itoa(card4))
	return sb.String()
}

func (*counterG) getHandFaceCountMap(count [15]int) map[int]string {
	countMap := make(map[int]string, 5)
	for i, v := range count {
		if v == 0 {
			continue
		} else {
			countMap[v] = FaceName[i] + countMap[v]
		}
	}
	return countMap
}

func (*counterG) getHandFaceCountTableByMap(countMap map[int]string) string {
	table := countMap[1] + countMap[2] + countMap[3] + countMap[4]
	return table
}

//getSevenHandFaceRank 得到牌型，如果有顺子、同花，会返回新的子串
func (c *counterG) getSevenHandFaceRank(hand string, faceCount [15]int) (int, string) {
	table := c.getHandFaceCountTable(faceCount)
	tp := SevenHandCount[table]
	newHand := ""

	//  以下牌型可能是同花或顺子
	if table == "4010" || table == "3200" || table == "5100" || table == "7000" {
		//  同花的点数大于顺子，所以优先判断是不是同花
		if flush, maybe := c.maybeIsFlush(hand, false); maybe == true {
			//  如果是则判断是不是顺子
			if straight, maybe := c.isStraightByNoDuplicate(flush); maybe == true {
				if c.isRoyalFlush(straight) {
					tp = HandRank["皇家同花顺"]
				} else {
					tp = HandRank["同花顺"]
				}
				newHand = straight
			} else {
				tp = HandRank["同花"]
				newHand = flush
			}
		} else if straight, maybe := c.isStraightByCount(faceCount); maybe == true {
			tp = HandRank["顺子"]
			newHand = straight
		}
	}
	return tp, newHand
}

//  maybeIsFlush 判断是否可能是同花，返回可能的序列
func (c *counterG) maybeIsFlush(hand string, isGhost bool) (string, bool) {
	var s1, s2, s3, s4 strings.Builder
	for i := 0; i < len(hand); i += 2 {
		poker := hand[i : i+2]
		switch poker[1:2] {
		case "s":
			s1.WriteString(poker)
		case "h":
			s2.WriteString(poker)
		case "d":
			s3.WriteString(poker)
		case "c":
			s4.WriteString(poker)
		}
	}

	var rst bool
	step := 5
	//  有鬼牌时4个同花色牌即可成同花
	if isGhost {
		step = 4
	}
	if s1.Len() >= 2*step {
		rst = true
		hand = s1.String()
	} else if s2.Len() >= 2*step {
		rst = true
		hand = s2.String()
	} else if s3.Len() >= 2*step {
		rst = true
		hand = s3.String()
	} else if s4.Len() >= 2*step {
		rst = true
		hand = s4.String()
	}
	return hand, rst
}

//isStraight 判断是否有顺子，有的话返回顺子
func (c *counterG) isStraight(hand string) (string, bool) {
	hand = sort(hand)
	s := ""
	//  可以是5、6、7手牌，所以最多循环3次，最少循环1次
	for i := 0; i < (len(hand)-5*2)+2; i += 2 {
		flag := true
		str := hand[i : i+10]
		last := FaceRank[str[0:1]]
		for j := 2; j < 10; j += 2 {
			cur := FaceRank[str[j:j+1]]
			if last-1 != cur {
				flag = false
				break
			}
			last = cur
		}
		if flag {
			//  循环一遍后flag还是true说明是顺子
			s = hand[i : i+10]
			break
		}
	}
	return s, s != ""
}

//  判断无重复字符串是否是顺子
func (c *counterG) isStraightByNoDuplicate(hand string) (string, bool) {
	hand = sort(hand)
	straight := ""
	var head, tail string
	times := 0
	switch len(hand) {
	case 10:
		times = 1
	case 12:
		times = 2
	case 14:
		times = 3
	}
	for i := 0; i < times; i++ {
		head, tail = hand[i*2:i*2+1], hand[i*2+8:i*2+9]
		if FaceRank[head]-FaceRank[tail] == 4 {
			straight = hand[i*2 : i*2+9]
			break
		}
	}
	return straight, len(straight) != 0
}

func (c *counterG) isStraightByCount(count [15]int) (string, bool) {
	sb := strings.Builder{}
	for i := 14; i >= 2; i-- {
		v := count[i]
		if v == 0 {
			continue
		}
		sb.WriteString(FaceName[i])
	}

	hand := sb.String()
	if len(hand) < 5 {
		return "", false
	}

	var head, tail, straight string
	for i := 0; i < len(hand)-5+1; i++ {
		head, tail = hand[i:i+1], hand[i+4:i+5]
		if FaceRank[head]-FaceRank[tail] == 4 {
			straight = hand[i : i+5]
			break
		}
	}
	return straight, len(straight) != 0
}

func (c *counterG) equalJudgeStraight(hand1, hand2 string) int {
	var rst int
	var c1, c2 int
	l1, l2 := len(hand1), len(hand2)
	step1, step2 := 2, 2
	if l1 < 10 {
		step1 = 1
	}
	if l2 < 10 {
		step2 = 1
	}

	for i := 0; i < 5; i++ {
		c1 = FaceRank[hand1[i*step1:i*step1+1]]
		c2 = FaceRank[hand2[i*step2:i*step2+1]]
		if c1 > c2 {
			rst = 1
			break
		} else if c1 < c2 {
			rst = 2
			break
		}
	}
	return rst
}

func (c *counterG) equalJudgePair(count1, count2 [15]int, rank int) int {
	map1 := c.getHandFaceCountMap(count1)
	map2 := c.getHandFaceCountMap(count2)
	var rst int
	var isEqual bool
	switch rank {
	case HandRank["一对"]:
		c1, c2 := FaceRank[map1[2]], FaceRank[map2[2]]
		if rst, isEqual = max(c1, c2); isEqual {
			hand1, hand2 := map1[1], map2[1]
			for i := 0; i < 3; i++ {
				c1 = FaceRank[hand1[i:i+1]]
				c2 = FaceRank[hand2[i:i+1]]
				if c1 > c2 {
					rst = 1
					break
				} else if c1 < c2 {
					rst = 2
					break
				}
			}
		}

	case HandRank["两对"]:
		c1s, c2s := map1[2], map2[2]
		c1, c2 := FaceRank[c1s[0:1]], FaceRank[c2s[0:1]]

		if rst, isEqual = max(c1, c2); isEqual {
			c1, c2 = FaceRank[c1s[1:2]], FaceRank[c2s[1:2]]
			if rst, isEqual = max(c1, c2); isEqual {
				//  可能有3个两对
				var c3, c4 int
				if len(c1s) == 3 {
					c3 = FaceRank[c1s[2:3]]
				}
				c4 = FaceRank[map1[1][0:1]]
				if c3 > c4 {
					c1 = c3
				} else {
					c1 = c4
				}

				if len(c2s) == 3 {
					c3 = FaceRank[c2s[2:3]]
				}
				c4 = FaceRank[map2[1][0:1]]
				if c3 > c4 {
					c2 = c3
				} else {
					c2 = c4
				}

				rst, _ = max(c1, c2)
			}
		}
	case HandRank["三条"]:
		c1, c2 := FaceRank[(map1[3])[0:1]], FaceRank[(map2[3])[0:1]]

		if rst, isEqual = max(c1, c2); isEqual {
			hand1, hand2 := map1[1], map2[1]
			for i := 0; i < 2; i++ {
				c1 = FaceRank[hand1[i:i+1]]
				c2 = FaceRank[hand2[i:i+1]]
				if c1 > c2 {
					rst = 1
					break
				} else if c1 < c2 {
					rst = 2
					break
				}
			}
		}
	case HandRank["葫芦"]:
		c1, c2 := FaceRank[map1[3][0:1]], FaceRank[map2[3][0:1]]

		if rst, isEqual = max(c1, c2); isEqual {
			if len(map1[2]) == 0 {
				c1 = FaceRank[map1[3][1:2]]
			} else {
				c1 = FaceRank[map1[2][0:1]]
			}
			if len(map2[2]) == 0 {
				c2 = FaceRank[map2[3][1:2]]
			} else {
				c2 = FaceRank[map2[2][0:1]]
			}
			rst, _ = max(c1, c2)
		}

	case HandRank["四条"]:
		c1, c2 := FaceRank[map1[4][0:1]], FaceRank[map2[4][0:1]]

		if rst, isEqual = max(c1, c2); isEqual {
			c1, c2 = FaceRank[map1[1][0:1]], FaceRank[map2[1][0:1]]
			rst, _ = max(c1, c2)
		}
	}
	return rst
}

func max(x, y int) (int, bool) {
	if x > y {
		return 1, false
	} else if x < y {
		return 2, false
	} else {
		return 0, true
	}
}

func (c *counterG) getGhostHandFaceRank(hand string, faceCount [15]int) int {
	table := c.getHandFaceCountTable(faceCount)
	rank := GhostHandCount[table]
	//newHand := ""
	if table == "3010" || table == "2200" { // 可能为同花顺
		if flush, maybe := c.maybeIsFlush(hand, true); maybe {

		}
	} else if table == "4100" || table == "6000" { // 可能是顺子、同花、同花顺、皇家同花顺

	}

	return rank

}

//maybeIsStraight 判断无重复字符串，无花色信息，有鬼牌时有没有可能是顺子，有顺子时返回正确的结果
func (c *counterG) maybeIsStraightByNoDuplicate(hand string) {
	hand := sort(hand)
	//  手牌的长度是4、5、6
	for i := 0; i < len(hand)-3; i++ {

		str := hand[i : i*2+1]
	}

	var rst = true
	var flag bool
	//  通过前后数字相减差值判断连续，允许出现一次数字不连续且可补位
	last := FaceRank[hand[0:1]]
	for i := 2; i < len(hand); i += 2 {
		val := FaceRank[hand[i:i+1]]

		if last-1 != val {
			if !flag && last-2 == val {
				flag = true
			} else {
				rst = false
				break
			}
		}
		last = val
	}
}
