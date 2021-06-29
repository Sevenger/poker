package casino

import (
	"strconv"
	"strings"
)

//counter 用于计算牌的牌型
type counter struct{}

type countRst struct {
	Hand     []string
	IsGhost  bool
	HandRank int
}

// QuickCount 计算牌型切片中最大的牌型
func (c *counter) QuickCount(hands []string) *countRst {
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

	return &countRst{Hand: newHands, IsGhost: isGhost, HandRank: maxType}
}

//isTongHua 判断是否是同花色
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

//getHandRank 得到5手牌的牌型
func (c *counter) getHandRank(hand string) int {
	code := c.getHandFaceCountTable(c.getHandFaceCount(hand))
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

//getHostHandRank 得到四手牌的牌型
func (c *counter) getHostHandRank(hand string) int {
	count := c.getHandFaceCount(hand)
	code := c.getHandFaceCountTable(count)
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

//hasFlush 顺子必然是连续的
func (*counter) hasFlush(hand string) bool {
	var rst = true

	//  通过前后数字相减差值判断是否连续
	last := FaceRank[hand[0:1]]
	for i := 2; i < len(hand); i += 2 {
		val := FaceRank[hand[i:i+1]]
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

//canBeFlush 判读四手牌是否是连续的
func (*counter) canBeFlush(hand string) bool {
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

	//  2345A特殊判断
	if hand[0] == 'A' {
		sb := strings.Builder{}
		for i := 0; i < len(hand); i += 2 {
			sb.WriteString(hand[i : i+1])
		}
		handFaces := sb.String()
		if handFaces == "A432" || handFaces == "A532" || handFaces == "A542" || handFaces == "A543" {
			rst = true
		}
	}
	return rst
}

//isRoyalFlush 判断五手牌是否可能为皇家同花顺
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

//canBeRoyalFlush 判断四手牌是否可能为皇家同花顺
func (*counter) canBeRoyalFlush(hand string) bool {
	var rst bool
	sb := strings.Builder{}
	for i := 0; i < len(hand); i += 2 {
		sb.WriteString(hand[i : i+1])
	}
	handFaces := sb.String()
	if handFaces == "AKQJ" || handFaces == "AKQT" || handFaces == "AKJT" || handFaces == "AQJT" || handFaces == "KQJT" {
		rst = true
	}
	return rst
}

//getHandFaceCount 计算手牌中每种牌出现的次数
func (*counter) getHandFaceCount(hand string) [15]int {
	//  一共有12种牌，最小牌在map中值为2，最大为14，为了方便计算，数组长度为15
	count := [15]int{}
	for i := 0; i < len(hand); i += 2 {
		count[FaceRank[string(hand[i])]] += 1
	}

	return count
}

//getHandFaceCountTable 计算出现1次至4次的牌的情况
func (*counter) getHandFaceCountTable(count [15]int) string {
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
	//return fmt.Sprintf("%d%d%d%d", card1, card2, card3, card4)
}
