package casino

import (
	"fmt"
)

type counter struct {
}

type CountRst struct {
	Hand     []string
	IsGhost  bool
	HandType int
}

// QuickCount 计算牌型切片中最大的牌型
func (c *counter) QuickCount(hands []string) *CountRst {
	var newHands []string
	isGhost := len(hands[0]) == 4*2
	maxType := 0

	var rank int
	for _, v := range hands {
		if isGhost {
			rank = c.getHostHandRank(v)
		} else {
			rank = c.getHandRank(v)
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

func (*counter) isTongHua(hand string) bool {
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

func (c *counter) getHandRank(hand string) int {
	code := c.getHandFaceCountInfo(c.getHandFaceCount(hand))
	tp := FiveHandCount[code]

	if tp == 1 {
		hasFlush := c.hasFlush(hand)
		isTongHua := c.isTongHua(hand)
		if hasFlush {
			if isTongHua {
				if c.isRoyalFlush(hand) {
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

func (c *counter) getHostHandRank(hand string) int {
	count := c.getHandFaceCount(hand)
	code := c.getHandFaceCountInfo(count)
	tp := ForHandCount[code]

	if tp == 1 {
		canBeFlush := c.canBeFlush(hand)
		isTongHua := c.isTongHua(hand)
		if canBeFlush {
			if isTongHua {
				if c.canBeRoyalFlush(hand) {
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

// hasFlush 顺子必然是连续的
func (*counter) hasFlush(hand string) bool {
	var rst = true

	//  通过前后数字相减差值判断是否连续
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
	if hand[0] == 'A' && hand[2] == '5' && hand[4] == '4' && hand[6] == '3' && hand[8] == '2' {
		rst = true
	}

	return rst
}

func (*counter) canBeFlush(hand string) bool {
	var rst = true
	var flag bool

	//  通过前后数字相减差值判断连续，允许出现一次数字不连续且可补位
	last := FaceRank[string(hand[0])]
	for i := 2; i < len(hand); i += 2 {
		val := FaceRank[string(hand[i])]

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

	//  2345A特殊判断
	if hand[0] == 'A' {
		handFaces := fmt.Sprintf("%s%s%s%s", hand[0:1], hand[2:3], hand[4:5], hand[6:7])
		fmt.Println(handFaces)
		if handFaces == "A432" || handFaces == "A532" || handFaces == "A542" || handFaces == "A543" {
			rst = true
		}
	}
	return rst
}

func (*counter) isRoyalFlush(hand string) bool {
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

func (*counter) canBeRoyalFlush(hand string) bool {
	var rst bool
	handFaces := fmt.Sprintf("%s%s%s%s", hand[0:1], hand[2:3], hand[4:5], hand[6:7])
	if handFaces == "AKQJ" || handFaces == "AKQT" || handFaces == "AKJT" || handFaces == "AQJT" || handFaces == "KQJT" {
		rst = true
	}
	return rst
}

// getHandFaceCount 计算手牌中每种牌出现的次数
func (*counter) getHandFaceCount(hand string) [15]int {
	//  一共有12种牌，最小牌在map中值为2，最大为14，为了方便计算，数组长度为15
	count := [15]int{}
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[string(hand[i])]] += 1
	}

	return count
}

// getHandFaceCountInfo 计算每种牌出现的次数结果`中同点数牌的情况
func (*counter) getHandFaceCountInfo(count [15]int) string {
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
