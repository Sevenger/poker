package casino

import (
	"fmt"
	"poker/src"
	"strings"
)

//  比牌

type Judge struct {
}

func (j *Judge) ResultJudge(countRst1, countRst2 *CountRst) int {
	rank1, rank2 := countRst1.HandType, countRst2.HandType
	rst := j.QuickJudge(rank1, rank2)
	fmt.Printf("左边牌型: %v 右边牌型: %v\n", src.HandName[countRst1.HandType], src.HandName[countRst2.HandType])
	//  如果是平局，需要先计算出每一方的最大牌，再根据最大牌比较
	if rst == 0 && rank1 != src.HandRank["皇家同花顺"] {
		//  如果是鬼牌则填充牌
		if countRst1.IsGhost {
			countRst1.Hand = InsertGhostHands(countRst1.Hand, countRst1.HandType)
		}
		if countRst2.IsGhost {
			countRst2.Hand = InsertGhostHands(countRst2.Hand, countRst2.HandType)
		}
		rst = j.EqualJudge(countRst1.Hand, countRst2.Hand, rank1)
	}
	return rst
}

func (j *Judge) QuickJudge(handType1, handType2 int) int {
	if handType1 > handType2 {
		return 1
	} else if handType1 < handType2 {
		return 2
	} else {
		return 0
	}
}

func (j *Judge) EqualJudge(hands1, hands2 []string, handRank int) int {
	rst := 0
	if handRank != src.HandRank["皇家同花顺"] {
		hand1, _ := j.GetBestHand(hands1, handRank)
		hand2, _ := j.GetBestHand(hands2, handRank)
		if hand1 != hand2 {
			_, rst = j.GetBestHand([]string{hand1, hand2}, handRank)
		}
	}
	return rst
}

func InsertGhostHands(hands []string, handRank int) []string {
	var newHands []string
	for _, v := range hands {
		var hand string
		if handRank == src.HandRank["一对"] {
			hand = fmt.Sprintf("%s%s", v[0:2], v)
		} else if handRank == src.HandRank["三条"] {
			if v[0] == v[2] {
				hand = fmt.Sprintf("%s%s", v[0:2], v)
			} else if v[2] == v[4] {
				hand = fmt.Sprintf("%s%s%s%s", v[0:2], v[2:4], v[2:6], v[6:8])
			} else if v[4] == v[6] {
				hand = fmt.Sprintf("%s%s", v, v[6:8])
			}
		} else if handRank == src.HandRank["葫芦"] {
			hand = fmt.Sprintf("%s%s", v[0:2], v)
		} else if handRank == src.HandRank["四条"] {
			//  首尾相同XXXX型
			if v[0:1] == v[6:7] {
				hand = fmt.Sprintf("%s%s", v, "As")
			} else if v[0:1] == v[2:3] {
				//  1、2相同XXXY型
				hand = fmt.Sprintf("%s%s", v, v[6:8])
			} else {
				//  1、2不同XYYY型
				hand = fmt.Sprintf("%s%s", v[0:2], v)
			}
		} else if handRank == src.HandRank["同花"] {
			for _, k := range []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"} {
				if !strings.Contains(hand, v) {
					hand = fmt.Sprintf("%s%s%s", k, v[0:1], v)
					break
				}
			}
		} else if handRank == src.HandRank["同花顺"] || handRank == src.HandRank["顺子"] {
			var i int
			last := src.FaceRank[hand[0:1]]

			for i = 2; i < len(hand); i += 2 {
				if last-src.FaceRank[hand[i:i+1]] != 1 {
					hand = fmt.Sprintf("%s%s%s%s", v[0:i], src.FaceName[last+1], v[1:2], v[i+2:])
					break
				}
			}
			//  如果hand长度为0说明需要在头或尾插入
			if len(hand) == 0 {
				//  除非开头是A，否则始终往头部插入
				if hand[0] == 'A' {
					hand = fmt.Sprintf("%s%s%s", v, "T", v[1:2])
				} else {
					hand = fmt.Sprintf("%s%s%s", src.FaceName[src.FaceRank[v[0:1]]+1], v[1:2], v)
				}
			}

			//  todo A2345，缺少2345任意数的情况
		}
		newHands = append(newHands, hand)
	}

	return newHands
}

// GetBestHand 返回牌点数最大的牌型，返回在切片中的下标，相同时返回0,不比较皇家同花顺
func (j *Judge) GetBestHand(hands []string, handRank int) (string, int) {
	var max = hands[0]
	var cur string

	dealer := Dealer{}
	for i := 1; i < len(hands); i++ {
		//cur = hands[i]
		cur = dealer.Sort(hands[i])
		switch handRank {
		//  按顺序比较牌点数
		case src.HandRank["顺子"]:
			fallthrough
		case src.HandRank["同花顺"]:
			fallthrough
		case src.HandRank["同花"]:
			fallthrough
		case src.HandRank["高牌"]:
			rst := whoIsMax(max, cur)
			if rst == 2 {
				max = cur
			}
		case src.HandRank["一对"]:

		case src.HandRank["两对"]:
		case src.HandRank["3条"]:
		case src.HandRank["葫芦"]:
		case src.HandRank["4条"]:
		}
	}

	return max, 0
}

func whoIsMax(s1, s2 string) int {
	if len(s1) == 0 {
		return 0
	}

	v1, v2 := src.FaceRank[s1[0:1]], src.FaceRank[s2[0:1]]
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return 2
	} else {
		return whoIsMax(s1[2:], s2[2:])
	}
}
