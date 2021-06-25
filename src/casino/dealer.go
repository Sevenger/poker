package casino

import (
	"poker/src"
)

type Dealer struct{}

func (d *Dealer) Deal(hand1 string, hand2 string) {
	d.Sort(hand1)
	d.Sort(hand2)

	if len(hand1)/2 == 7 {

	} else {
		(&Counter{}).Count(hand1)
	}

	if len(hand2)/2 == 7 {

	}
}

func (*Dealer) Sort(hand string) string {
	l := len(hand)

	val := []byte(hand)
	for i := 2; i < l; i += 2 {
		for v := 0; v < i; v += 2 {
			if src.Face[string(val[v])] < src.Face[string(val[i])] {
				val[v], val[i] = val[i], val[v]
				val[v+1], val[i+1] = val[i+1], val[v+1]
			}
		}
	}

	return string(val)
}

// SevenToFive 7选5
func SevenToFive(hand string) string {
	//var hands []string
	//runes := []rune(hand)
	//if hand[0] == 'X' {
	//	hands = []string{
	//
	//	}
	//} else {
	//	hands = []string{
	//
	//	}
	//}

	return ""
}
