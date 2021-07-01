package counterG

import (
	"strconv"
	"strings"
)

//counterG 计算7手牌和鬼牌
type counterG struct{}

//FaceCount 用于记录手牌每个牌面出现的次数。下标为牌面，值为次数
//一共有13种牌(含鬼牌)，最小牌在map中值为2，最大为15，为了方便计算，数组长度为16
type FaceCount [16]int

func (c *counterG) Start(hand1, hand2 string) int {
	faceCount1, faceCount2 := c.getFaceCount(hand1), c.getFaceCount(hand2)

	rank1, newHand1 := c.getHandRank(hand1, faceCount1)
	rank2, newHand2 := c.getHandRank(hand1, faceCount1)

	if rank1 > rank2 {
		return 1
	} else if rank1 < rank2 {
		return 2
	} else {
		switch rank1 {

		case HandRank["高牌"]:
			return c.equalJudgeHighCard(faceCount1, faceCount2)

		case HandRank["同花"]:
			return c.equalJudgeFlush(hand1, hand2)

		//  顺子和同花顺只需要判断头牌
		case HandRank["顺子"]:
			fallthrough
		case HandRank["同花顺"]:
			if len(newHand1) != 0 {
				hand1 = newHand1
			}
			if len(newHand2) != 0 {
				hand2 = newHand2
			}
			rst, _ := max(FaceRank[hand1[0:1]], FaceRank[hand2[0:1]])
			return rst

		default:
			if len(newHand1) != 0 {
				faceCount1 = c.getFaceCount(newHand1)
			}
			if len(newHand2) != 0 {
				faceCount2 = c.getFaceCount(newHand2)
			}
			return 0
			//return c.equalJudgePair(faceCount1, faceCount2, rank1)
		}
	}
}

//isRoyalFlush 在保证同花的前提下判断是否可能为皇家同花顺
func (*counterG) isRoyalFlush(hand string) bool {
	return strings.Compare(hand, "AKQJT") == 0
}

//getFaceCount 计算手牌中每种牌出现的次数
func (*counterG) getFaceCount(hand string) FaceCount {
	var count FaceCount
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[hand[i:i+1]]]++
	}
	return count
}

//getFaceCountCode 计算出现1次至4次的牌的code
func (*counterG) getFaceCountCode(count FaceCount) string {
	sb.Reset()
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

	sb.WriteString(strconv.Itoa(card1))
	sb.WriteString(strconv.Itoa(card2))
	sb.WriteString(strconv.Itoa(card3))
	sb.WriteString(strconv.Itoa(card4))
	return sb.String()
}

//getFaceCountMap 返回一个map，该map的key为出现的次数，v是一个[]int，存储了牌面值
func (*counterG) getFaceCountMap(count FaceCount) map[int][]int {
	countMap := make(map[int][]int, 5)
	//  此处i是牌面的值，v是牌面出现的次数
	for i, v := range count {
		//  不记录鬼牌
		if i == 15 {
			continue
		}
		if v == 0 {
			continue
		} else {
			countMap[v] = append(countMap[v], i)
		}
	}
	return countMap
}

func (c *counterG) getHandRank(hand string, count FaceCount) (int, string) {
	//  鬼牌
	if count[15] == 1 {
		return c.getGhostHandRank(hand, count)
	} else {
		return c.getSevenHandRank(hand, count)
	}
}

func (c *counterG) getGhostHandRank(hand string, faceCount FaceCount) (int, string) {
	code := c.getFaceCountCode(faceCount)
	rank := GhostHandCountCode[code]
	newHand := ""

	//  以下牌型不确定，需要继续判断
	if code == "3010" || code == "2200" || code == "4100" || code == "6000" {
		//  判断是否是同花
		if flush, maybe := c.maybeIsFlush(hand, true); maybe {
			if straight, maybe := c.maybeIsStraightByNoDuplicate(flush, true); maybe {
				if c.isRoyalFlush(straight) {
					rank = HandRank["皇家同花顺"]
				} else {
					rank = HandRank["同花顺"]
				}
				newHand = straight
			} else if code == "4100" || code == "6000" { //  3010,2200最低是四条，所以不需要是同花
				rank = HandRank["同花"]
				newHand = flush
			}
		} else if code == "4100" || code == "6000" {
			//  这两种牌型还要再判断是不是顺子
			if straight, maybe := c.maybeIsStraightByFaceCount(faceCount); maybe {
				rank = HandRank["顺子"]
				newHand = straight
			}
		}
	}
	return rank, newHand
}

//getSevenHandRank 得到牌型，如果有顺子、同花，会返回新的子串
func (c *counterG) getSevenHandRank(hand string, faceCount FaceCount) (int, string) {
	code := c.getFaceCountCode(faceCount)
	rank := SevenHandCountCode[code]
	newHand := ""

	//  以下牌型可能是同花或顺子
	if code == "4010" || code == "3200" || code == "5100" || code == "7000" {
		//  同花的点数大于顺子，所以优先判断是不是同花
		if flush, maybe := c.maybeIsFlush(hand, false); maybe {
			//  如果是则判断是不是顺子
			if straight, maybe := c.isStraightByNoDuplicate(flush, true); maybe {
				if c.isRoyalFlush(straight) {
					rank = HandRank["皇家同花顺"]
				} else {
					rank = HandRank["同花顺"]
				}
				newHand = straight
			} else {
				rank = HandRank["同花"]
				newHand = flush
			}
		} else if straight, maybe := c.isStraightByCount(faceCount); maybe {
			rank = HandRank["顺子"]
			newHand = straight
		}
	}
	return rank, newHand
}

//maybeIsFlush 判断是否可能是同花，可能时返回最大的同花牌
//传入的字符串要求包含花色信息。返回的字符串无花色信息
func (c *counterG) maybeIsFlush(hand string, isGhost bool) (string, bool) {
	s1.Reset()
	s2.Reset()
	s3.Reset()
	s4.Reset()
	for i := 0; i < len(hand); i += 2 {
		face := hand[i : i+1]
		suit := hand[i+1 : i+2]
		switch suit {
		case "s":
			s1.WriteString(face)
		case "h":
			s2.WriteString(face)
		case "d":
			s3.WriteString(face)
		case "c":
			s4.WriteString(face)
		}
	}

	flush := ""
	length := 5
	//  有鬼牌时4个同花色牌即可成同花
	if isGhost {
		length = 4
	}
	if s1.Len() >= length {
		flush = s1.String()
	} else if s2.Len() >= length {
		flush = s2.String()
	} else if s3.Len() >= length {
		flush = s3.String()
	} else if s4.Len() >= length {
		flush = s4.String()
	}
	return flush, len(flush) > 0
}

//  判断无重复字符串是否是顺子
func (c *counterG) isStraightByNoDuplicate(hand string, needSort bool) (string, bool) {
	if needSort {
		hand = sort(hand)
	}
	straight := ""
	//  由于是无重复子串，通过头部减去尾部可以直接判断是不是顺子
	var head, tail string
	for i := 0; i < len(hand)-4; i++ {
		head, tail = hand[i:i+1], hand[i+4:i+5]
		if FaceRank[head]-FaceRank[tail] == 4 {
			straight = hand[i : i+5]
			break
		}
	}
	return straight, len(straight) > 0
}

func (c *counterG) isStraightByCount(count FaceCount) (string, bool) {
	sb.Reset()
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

	return c.isStraightByNoDuplicate(hand, false)
}

//  maybeIsStraight 用于判断有鬼牌手牌是不是顺子，是顺子时返回最大顺子手牌
func (c *counterG) maybeIsStraightByNoDuplicate(hand string, needSort bool) (string, bool) {
	if needSort {
		hand = sort(hand)
	}
	sb.Reset()
	var insertFlag bool
	var val, lastVal, curVal int
	var str, curKey string
	//  有鬼牌时手牌长度是4、5、6，所以至多判断3次
	//  有鬼牌时通过前后数字相减差值判断连续，允许出现一次数字不连续且可补位
	for i := 0; i < len(hand)-3; i++ {
		insertFlag = false
		str = hand[i : i+4]

		lastVal = FaceRank[str[0:1]]
		sb.WriteString(str[0:1])
		for j := 1; j < len(str); j++ {
			curKey = str[j : j+1]
			curVal = FaceRank[curKey]
			val = lastVal - curVal
			if val == 1 {
				sb.WriteString(curKey)
			} else if val == 2 && insertFlag == false {
				insertFlag = true
				//  将缺的键插入
				sb.WriteString(FaceName[lastVal-1])
				sb.WriteString(curKey)
			} else {
				sb.Reset()
				break
			}
			lastVal = curVal
		}
		//  如果stringBuilder的长度不为0，说明已将找到了顺子
		if sb.Len() > 0 {
			break
		}
	}

	//  如果有顺子且flag为false，说明要在头部或者尾部插入组成顺子
	if sb.Len() > 0 && insertFlag == false {
		if hand[0] == 'A' {
			sb.WriteString("T")
		} else {
			hand = sb.String()
			sb.Reset()
			sb.WriteString(FaceName[FaceRank[hand[0:1]]+1])
			sb.WriteString(hand)
		}
	} else if sb.Len() == 0 && hand[0] == 'A' {
		//  A5432的特殊情况
		hand = hand[1:]
		if strings.Compare(hand, "543") == 0 ||
			strings.Compare(hand, "542") == 0 ||
			strings.Compare(hand, "532") == 0 ||
			strings.Compare(hand, "432") == 0 {
			sb.WriteString("5432A")
		}
	}
	return sb.String(), sb.Len() > 0
}

func (c *counterG) maybeIsStraightByFaceCount(count FaceCount) (string, bool) {
	sb.Reset()
	//  获取无重复子串
	for i := 14; i >= 2; i-- {
		v := count[i]
		if v == 0 {
			continue
		}
		sb.WriteString(FaceName[i])
	}

	if sb.Len() < 4 {
		return "", false
	}
	return c.maybeIsStraightByNoDuplicate(sb.String(), false)
}

func (c *counterG) equalJudgeHighCard(count1, count2 FaceCount) int {
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

//  可优化，同花可不必排序
func (c *counterG) equalJudgeFlush(hand1, hand2 string) int {
	if len(hand1) == 4 {
		for i := 14; i >= 2; i++ {
			face := FaceName[i]
			faceNext := FaceName[i-1]
			if !strings.Contains(hand1, face) && !strings.Contains(hand1, faceNext) {
				hand1 += face
				break
			}
		}
	}
	if len(hand2) == 4 {
		for i := 14; i >= 2; i++ {
			face := FaceName[i]
			faceNext := FaceName[i-1]
			if !strings.Contains(hand1, face) && !strings.Contains(hand1, faceNext) {
				hand2 += face
			}
		}
	}
	hand1, hand2 = sort(hand1), sort(hand2)

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

//func (c *counterG) equalJudgePair(count1, count2 [15]int, rank int) int {
//	map1 := c.getFaceCountMap(count1)
//	map2 := c.getFaceCountMap(count2)
//	var rst int
//	var isEqual bool
//	switch rank {
//	case HandRank["一对"]:
//		c1, c2 := FaceRank[map1[2]], FaceRank[map2[2]]
//		if rst, isEqual = max(c1, c2); isEqual {
//			hand1, hand2 := map1[1], map2[1]
//			for i := 0; i < 3; i++ {
//				c1 = FaceRank[hand1[i:i+1]]
//				c2 = FaceRank[hand2[i:i+1]]
//				if c1 > c2 {
//					rst = 1
//					break
//				} else if c1 < c2 {
//					rst = 2
//					break
//				}
//			}
//		}
//
//	case HandRank["两对"]:
//		c1s, c2s := map1[2], map2[2]
//		c1, c2 := FaceRank[c1s[0:1]], FaceRank[c2s[0:1]]
//
//		if rst, isEqual = max(c1, c2); isEqual {
//			c1, c2 = FaceRank[c1s[1:2]], FaceRank[c2s[1:2]]
//			if rst, isEqual = max(c1, c2); isEqual {
//				//  可能有3个两对
//				var c3, c4 int
//				if len(c1s) == 3 {
//					c3 = FaceRank[c1s[2:3]]
//				}
//				c4 = FaceRank[map1[1][0:1]]
//				if c3 > c4 {
//					c1 = c3
//				} else {
//					c1 = c4
//				}
//
//				if len(c2s) == 3 {
//					c3 = FaceRank[c2s[2:3]]
//				}
//				c4 = FaceRank[map2[1][0:1]]
//				if c3 > c4 {
//					c2 = c3
//				} else {
//					c2 = c4
//				}
//
//				rst, _ = max(c1, c2)
//			}
//		}
//	case HandRank["三条"]:
//		c1, c2 := FaceRank[(map1[3])[0:1]], FaceRank[(map2[3])[0:1]]
//
//		if rst, isEqual = max(c1, c2); isEqual {
//			hand1, hand2 := map1[1], map2[1]
//			for i := 0; i < 2; i++ {
//				c1 = FaceRank[hand1[i:i+1]]
//				c2 = FaceRank[hand2[i:i+1]]
//				if c1 > c2 {
//					rst = 1
//					break
//				} else if c1 < c2 {
//					rst = 2
//					break
//				}
//			}
//		}
//	case HandRank["葫芦"]:
//		c1, c2 := FaceRank[map1[3][0:1]], FaceRank[map2[3][0:1]]
//
//		if rst, isEqual = max(c1, c2); isEqual {
//			if len(map1[2]) == 0 {
//				c1 = FaceRank[map1[3][1:2]]
//			} else {
//				c1 = FaceRank[map1[2][0:1]]
//			}
//			if len(map2[2]) == 0 {
//				c2 = FaceRank[map2[3][1:2]]
//			} else {
//				c2 = FaceRank[map2[2][0:1]]
//			}
//			rst, _ = max(c1, c2)
//		}
//
//	case HandRank["四条"]:
//		c1, c2 := FaceRank[map1[4][0:1]], FaceRank[map2[4][0:1]]
//
//		if rst, isEqual = max(c1, c2); isEqual {
//			c1, c2 = FaceRank[map1[1][0:1]], FaceRank[map2[1][0:1]]
//			rst, _ = max(c1, c2)
//		}
//	}
//	return rst
//}
