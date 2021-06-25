package casino

import "poker/src"

//  比牌

type Judge struct {
}

func QuickJudge(handType1, handType2 string) int {
	rank1, rank2 := src.HandRank[handType1], src.HandRank[handType2]
	if rank1 > rank2 {
		return 1
	} else if rank1 < rank2 {
		return 2
	} else {
		return 0
	}
}

func EqualJudge(hand1, hand2 string, handType string) {
}

func GhostJudge() {}
