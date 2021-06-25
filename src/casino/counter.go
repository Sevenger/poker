package casino

import "C"
import (
	"poker/src"
	"strings"
)

//  算牌

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
	hs.HandType = c.GetHandType(&hs)
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

func (c *Counter) GetHandType(hand *HandStruct) int {
	var tp = 0

	//  连续必有顺子
	if c.HasFlush(hand) {
		//  同花则为皇家同花顺或同花顺
		if hand.IsTongHua {
			if c.IsRoyalFlush(hand.Hand, hand.IsTongHua) {
				tp = src.HandRank["皇家同花顺"]
			} else {
				tp = src.HandRank["同花顺"]
			}
		} else {
			//  否则为顺子
			tp = src.HandRank["顺子"]
		}
	} else if hand.IsTongHua {
		//  不连续却同花色必为同花
		tp = src.HandRank["同花"]
	} else {

	}

	return tp
}

// HasFlush 顺子必然是连续的
func (*Counter) HasFlush(hand *HandStruct) bool {
	var rst = true

	last := src.Face[string(hand.Hand[0])]
	for i := 2; i < len(hand.Hand)-2; i += 2 {
		val := src.Face[string(hand.Hand[i])]
		if last-1 != val {
			rst = false
			break
		}
		last = val
	}

	//  A2345在德州扑克里是最小顺子
	if strings.Contains(hand.Hand, "A") &&
		strings.Contains(hand.Hand, "4") &&
		strings.Contains(hand.Hand, "3") &&
		strings.Contains(hand.Hand, "2") {
		length := len(hand.Hand) / 2
		if length == 5 {
			strings.Contains(hand.Hand, "5")
			rst = true
		} else if length == 4 {
			rst = true
		}
	}

	return rst
}

func (*Counter) IsRoyalFlush(hand string, isTongHua bool) bool {
	var rst bool
	if hand[0] == 'A' && hand[2] == 'K' && hand[4] == 'Q' && hand[6] == 'J' && hand[8] == 'T' && isTongHua == true {
		rst = true
	}

	return rst
}

func (*Counter) IsFourOfKind() {}

func (*Counter) IsFullHouse() {}

func (*Counter) IsThreeOfKind() {}

func (*Counter) IsTwoPairs() {}

func (*Counter) IsOnePair() {}

func (*Counter) IsHighCard(hand *HandStruct) {
	//  每一个点数都不同

}
