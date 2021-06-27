package casino

import (
	"fmt"
)

type Dealer struct {
}

func (d *Dealer) Deal(hand1 string, hand2 string) ([]string, []string) {
	return d.DealHand(d.Sort(hand1)), d.DealHand(d.Sort(hand2))
}

func (d *Dealer) DealHand(handStr string) []string {
	var hands []string
	//  如果有7张牌，获取7张牌所有可能的牌组合
	if len(handStr) == 14 {
		hands = SevenToFive(handStr)
	} else {
		hands = append(hands, handStr)
	}
	return hands
}

func (*Dealer) Sort(hand string) string {
	l := len(hand)

	val := []byte(hand)
	for i := 2; i < l; i += 2 {
		for v := 0; v < i; v += 2 {
			if FaceRank[string(val[v])] < FaceRank[string(val[i])] {
				val[v], val[i] = val[i], val[v]
				val[v+1], val[i+1] = val[i+1], val[v+1]
			}
		}
	}

	return string(val)
}

// SevenToFive 7选5 使用穷举法给出排列组合，对于鬼牌，判断4张牌可能组成的最大值
func SevenToFive(hand string) []string {
	c1, c2, c3, c4, c5, c6, c7 := hand[0:2], hand[2:4], hand[4:6], hand[6:8], hand[8:10], hand[10:12], hand[12:14]
	var hands []string
	var format string

	//  有鬼牌时鬼牌必选，从剩下的6张牌选4张，一共有15种可能
	if hand[0] == 'X' {
		format = "%s%s%s%s"
		hands = append(hands, fmt.Sprintf(format, c2, c3, c4, c5), fmt.Sprintf(format, c2, c3, c4, c6), fmt.Sprintf(format, c2, c3, c4, c7), fmt.Sprintf(format, c2, c3, c5, c6), fmt.Sprintf(format, c2, c3, c5, c7), fmt.Sprintf(format, c2, c3, c6, c7), fmt.Sprintf(format, c2, c4, c5, c6), fmt.Sprintf(format, c2, c4, c5, c7), fmt.Sprintf(format, c2, c4, c6, c7), fmt.Sprintf(format, c2, c5, c6, c7), fmt.Sprintf(format, c3, c4, c5, c6), fmt.Sprintf(format, c3, c4, c5, c7), fmt.Sprintf(format, c3, c4, c6, c7), fmt.Sprintf(format, c3, c5, c6, c7), fmt.Sprintf(format, c4, c5, c6, c7))
	} else {
		//  无鬼牌时从7张牌中选5张，一共有21种可能
		format = "%s%s%s%s%s"
		hands = append(hands, fmt.Sprintf(format, c1, c2, c3, c4, c5), fmt.Sprintf(format, c1, c2, c3, c4, c6), fmt.Sprintf(format, c1, c2, c3, c4, c7), fmt.Sprintf(format, c1, c2, c3, c5, c6), fmt.Sprintf(format, c1, c2, c3, c5, c7), fmt.Sprintf(format, c1, c2, c3, c6, c7), fmt.Sprintf(format, c1, c2, c4, c5, c6), fmt.Sprintf(format, c1, c2, c4, c5, c7), fmt.Sprintf(format, c1, c2, c4, c6, c7), fmt.Sprintf(format, c1, c2, c5, c6, c7), fmt.Sprintf(format, c1, c3, c4, c5, c6), fmt.Sprintf(format, c1, c3, c4, c5, c7), fmt.Sprintf(format, c1, c3, c4, c6, c7), fmt.Sprintf(format, c1, c3, c5, c6, c7), fmt.Sprintf(format, c1, c4, c5, c6, c7), fmt.Sprintf(format, c2, c3, c4, c5, c6), fmt.Sprintf(format, c2, c3, c4, c5, c7), fmt.Sprintf(format, c2, c3, c4, c6, c7), fmt.Sprintf(format, c2, c3, c5, c6, c7), fmt.Sprintf(format, c2, c4, c5, c6, c7), fmt.Sprintf(format, c3, c4, c5, c6, c7))
	}

	return hands
}
