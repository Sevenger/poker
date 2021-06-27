package casino

import (
	"fmt"
	"strings"
)

type Counter struct {
}

// QuickCount 计算牌型切片中最大的牌型
func (c *Counter) QuickCount(hands []string) *CountRst {
	var newHands []string
	isGhost := len(hands[0]) == 4*2
	maxType := 0

	var rank int
	for _, v := range hands {
		if isGhost {
			rank = c.GetHostHandType(v)
		} else {
			rank = c.GetHandType(v)
		}
		if rank > maxType {
			maxType = rank
			newHands = newHands[0:0]
			newHands = append(newHands, v)
		} else if rank == maxType {
			newHands = append(newHands, v)
		}
	}

	return &CountRst{Hand: newHands, IsGhost: isGhost, HandType: maxType}
}

func (*Counter) IsTongHua(hand string) bool {
	rst := true

	key := hand[1]
	for i := 3; i < len(hand); i += 2 {
		if hand[i] != key {
			rst = false
			break
		}
	}

	return rst
}

func (c *Counter) GetHandType(hand string) int {
	code := c.GetHandFaceCountInfo(c.GetHandFaceCount(hand))
	tp := FiveHandCount[code]

	if tp == 1 {
		hasFlush := c.HasFlush(hand)
		isTongHua := c.IsTongHua(hand)
		if hasFlush {
			if isTongHua {
				if c.IsRoyalFlush(hand) {
					tp = HandRank["皇家同花顺"]
				} else {
					tp = HandRank["同花顺"]
				}
			} else {
				tp = HandRank["顺子"]
			}
		} else if isTongHua {
			tp = HandRank["同花"]
		} else {
			tp = HandRank["高牌"]
		}
	}

	return tp
}

func (c *Counter) GetHostHandType(hand string) int {
	count := c.GetHandFaceCount(hand)
	code := c.GetHandFaceCountInfo(count)
	tp := ForHandCount[code]

	if tp == 1 {
		canBeFlush := c.CanBeFlush(hand)
		isTongHua := c.IsTongHua(hand)
		if canBeFlush {
			if isTongHua {
				if c.CanBeRoyalFlush(hand) {
					tp = HandRank["皇家同花顺"]
				} else {
					tp = HandRank["同花顺"]
				}
			} else {
				tp = HandRank["顺子"]
			}
		} else if isTongHua {
			tp = HandRank["同花"]
		} else {
			tp = HandRank["一对"]
		}
	}
	return tp
}

// HasFlush 顺子必然是连续的
func (*Counter) HasFlush(hand string) bool {
	var rst = true

	last := FaceRank[string(hand[0])]
	for i := 2; i < len(hand); i += 2 {
		val := FaceRank[string(hand[i])]
		if last-1 != val {
			rst = false
			break
		}
		last = val
	}

	//  A2345在德州扑克里是最小顺子，由于事先对手牌进行了排序，A总是出现在第一位，所以特殊判断
	if strings.Contains(hand, "A") &&
		strings.Contains(hand, "4") &&
		strings.Contains(hand, "3") &&
		strings.Contains(hand, "2") {
		length := len(hand) / 2
		if length == 5 {
			strings.Contains(hand, "5")
			rst = true
		} else if length == 4 {
			rst = true
		}
	}

	return rst
}

func (*Counter) CanBeFlush(hand string) bool {
	var rst = true
	var flag bool

	last := FaceRank[string(hand[0])]
	for i := 2; i < len(hand); i += 2 {
		val := FaceRank[string(hand[i])]
		if last-1 != val && !flag {
			flag = true
		} else {
			rst = false
			break
		}
		last = val
	}

	if hand[0] == 'A' {
		if hand == "A234" || hand == "A235" || hand == "A245" || hand == "A345" {
			rst = true
		}
	}
	return rst
}

func (*Counter) IsRoyalFlush(hand string) bool {
	var rst bool
	if hand[0] == 'A' &&
		hand[2] == 'K' &&
		hand[4] == 'Q' &&
		hand[6] == 'J' &&
		hand[8] == 'T' {
		rst = true
	}

	return rst
}

func (*Counter) CanBeRoyalFlush(hand string) bool {
	var rst bool
	str := fmt.Sprintf("%s%s%s%s", hand[0:1], hand[2:3], hand[4:5], hand[6:7])
	if str == "AKQJ" || str == "AKQT" || str == "AKJT" || str == "AQJT" || str == "KQJT" {
		rst = true
	}
	return rst
}

// GetHandFaceCount 计算手牌中每种牌出现的次数
func (*Counter) GetHandFaceCount(hand string) [15]int {
	//  一共有12种牌，最小牌在map中值为2，最大为14，为了方便计算，数组长度为15
	count := [15]int{}
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[string(hand[i])]] += 1
	}

	return count
}

// GetHandFaceCountInfo 计算每种牌出现的次数结果`中同点数牌的情况
func (*Counter) GetHandFaceCountInfo(count [15]int) string {
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

	return fmt.Sprintf("%d%d%d%d", card1, card2, card3, card4)
}
