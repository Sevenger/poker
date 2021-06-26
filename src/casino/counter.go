package casino

import (
	"fmt"
	"poker/src"
	"strings"
)

type Counter struct{}

type CounterGroup struct {
	Counter1 Counter
	Counter2 Counter
}

type HandStruct struct {
	Hand      string
	IsTongHua bool
	IsGhost   bool
	HandType  int
}

func (c *Counter) Count(hand string) *HandStruct {
	var hs = HandStruct{
		Hand:      hand,
		IsTongHua: false,
		IsGhost:   false,
		HandType:  0,
	}

	//  todo  算鬼牌
	if hand[0] == 'X' {
		hs.IsGhost = true
	}

	hs.IsTongHua = c.IsTongHua(hs.Hand)
	hs.HandType = c.GetHandType(hs.Hand, hs.IsTongHua)
	return &hs
}

func (*Counter) MultiCount(hands []string) *HandStruct {
	var hs HandStruct
	return &hs
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

func (c *Counter) GetHandType(hand string, isTongHua bool) int {
	var tp = 0

	//  连续必有顺子
	if c.HasFlush(hand) {
		//  同花则为皇家同花顺或同花顺
		if isTongHua {
			if c.IsRoyalFlush(hand) {
				tp = src.HandRank["皇家同花顺"]
			} else {
				tp = src.HandRank["同花顺"]
			}
		} else {
			//  否则为顺子
			tp = src.HandRank["顺子"]
		}
	} else if isTongHua {
		//  不连续却同花色必为同花
		tp = src.HandRank["同花"]
	} else {
		//  根据四条、葫芦、三条、两对、一对、高牌牌型具有的牌点数出现次数的特征
		//  提前创建map，算出牌点数出现次数，得到牌型
		count := c.GetHandFaceCount(hand)
		code := c.GetHandFaceCountInfo(count)
		tp = src.HandCount[code]
	}

	return tp
}

// HasFlush 顺子必然是连续的
func (*Counter) HasFlush(hand string) bool {
	var rst = true

	last := src.Face[string(hand[0])]
	for i := 2; i < len(hand); i += 2 {
		val := src.Face[string(hand[i])]
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

// GetHandFaceCount 计算手牌中每种牌出现的次数
func (*Counter) GetHandFaceCount(hand string) [15]int {
	//  一共有12种牌，最小牌在map中值为2，最大为14，为了方便计算，数组长度为15
	count := [15]int{}
	for i := 0; i < len(hand); i += 2 {
		count[src.Face[string(hand[i])]] += 1
	}

	return count
}

// GetHandFaceCountInfo 计算`每种牌出现的次数结果`中同点数牌的情况
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
