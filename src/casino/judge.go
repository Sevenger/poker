package casino

import (
	"fmt"
	"strings"
)

//  比牌

type judge struct{}

func (j *judge) ResultJudge(countRst1, countRst2 *CountRst) int {
	rank1, rank2 := countRst1.HandType, countRst2.HandType
	rst := j.quickJudge(rank1, rank2)
	//  如果是平局，需要先计算出每一方的最大牌，再根据最大牌比较。皇家同花顺由于默认平局，无需判断
	if rst == 0 && rank1 != HandRank["皇家同花顺"] {
		//  如果是鬼牌则填充牌
		if countRst1.IsGhost {
			countRst1.Hand = insertGhostHands(countRst1.Hand, countRst1.HandType)
		}
		if countRst2.IsGhost {
			countRst2.Hand = insertGhostHands(countRst2.Hand, countRst2.HandType)
		}
		rst = j.equalJudge(countRst1.Hand, countRst2.Hand, rank1)
	}
	return rst
}

func (j *judge) quickJudge(handType1, handType2 int) int {
	if handType1 > handType2 {
		return 1
	} else if handType1 < handType2 {
		return 2
	} else {
		return 0
	}
}

func (j *judge) equalJudge(hands1, hands2 []string, handRank int) int {
	rst := 0
	if handRank != HandRank["皇家同花顺"] {
		i1, _ := j.getBestHand(hands1, handRank)
		i2, _ := j.getBestHand(hands2, handRank)
		hand1, hand2 := hands1[i1], hands2[i2]
		if hand1 != hand2 {
			maxIndex, isEqual := j.getBestHand([]string{hand1, hand2}, handRank)
			rst = maxIndex + 1
			if isEqual {
				rst = 0
			}
		}
	}
	return rst
}

//insertGhostHands 手牌必须是排序的
func insertGhostHands(hands []string, handRank int) []string {
	var newHands []string
	for _, v := range hands {
		var hand string
		switch handRank {
		case HandRank["一对"]: //  牌型为 XYZW
			hand = fmt.Sprintf("%s%s", v[0:2], v)

		case HandRank["三条"]: //  牌型为 XXYZ|YXXZ|YZXX
			if v[0] == v[2] {
				hand = fmt.Sprintf("%s%s", v[0:2], v)
			} else if v[2] == v[4] {
				hand = fmt.Sprintf("%s%s%s%s", v[0:2], v[2:4], v[2:6], v[6:8])
			} else if v[4] == v[6] {
				hand = fmt.Sprintf("%s%s", v, v[6:8])
			}

		case HandRank["葫芦"]: //  牌型为 XXYY
			hand = fmt.Sprintf("%s%s", v[0:2], v)

		case HandRank["四条"]: //  牌型为 XXXY|YXXX|XXXX
			//  首尾相同即XXXX型
			if v[0:1] == v[6:7] {
				hand = fmt.Sprintf("%s%s", v, "As")
			} else if v[0:1] == v[2:3] {
				hand = fmt.Sprintf("%s%s", v, v[6:8])
			} else {
				hand = fmt.Sprintf("%s%s", v[0:2], v)
			}

		case HandRank["同花"]: //  牌型为 XYZW
			//  从A-2按顺序中找出一个手牌中不存在的牌
			for _, k := range Faces {
				if !strings.Contains(hand, v) {
					hand = fmt.Sprintf("%s%s%s", k, v[0:1], v)
					break
				}
			}

		case HandRank["皇家同花顺"]: //  牌型为 XYZW
			fallthrough
		case HandRank["同花顺"]: //  牌型为 XYZW
			//  判断是否往中间插入
			last := FaceRank[v[0:1]]
			for i := 2; i < len(v); i += 2 {
				cur := FaceRank[v[i:i+1]]
				if last-cur == 2 {
					//  todo test this code
					hand = fmt.Sprintf("%s%s%s%s", v[0:i], FaceName[last+1], v[1:2], v[i:])
					break
				}
				last = cur
			}

			//  如果hand长度为0说明需要在头或尾插入
			if len(hand) == 0 {
				//  除非开头是A，否则始终往头部插入
				if v[0] == 'A' {
					//  开头为A时判断A5432牌型
					handFaces := fmt.Sprintf("%s%s%s%s", v[0:1], v[2:3], v[4:5], v[6:7])
					if handFaces == "A432" || handFaces == "A532" || handFaces == "A542" || handFaces == "A543" {
						if handRank == FaceRank["同花顺"] {
							hand = fmt.Sprintf("5s4s3s2sAs")
						} else {
							hand = fmt.Sprintf("5s4s3s2sAh")
						}
					} else {
						hand = fmt.Sprintf("%s%s%s", v, "T", v[1:2])
					}
				} else {
					hand = fmt.Sprintf("%s%s%s", FaceName[FaceRank[v[0:1]]+1], v[1:2], v)
				}
			}
		}

		newHands = append(newHands, hand)
	}

	return newHands
}

// getBestHand 返回牌点数最大的牌型，返回在切片中的下标，相同时返回0,不比较皇家同花顺
func (j *judge) getBestHand(hands []string, handRank int) (int, bool) {
	if len(hands) == 1 {
		return 0, false
	}

	var max = hands[0]
	var maxIndex int
	var isEqual bool
	var cur string
	maxScoreTable := getHandScoreTable(max)
	var curScoreTable [5]string

	for i := 1; i < len(hands); i++ {
		cur = hands[i]
		switch handRank {
		//  todo  可能有A5432的情况
		case HandRank["顺子"]: //  顺子和同花顺只需要比较头牌
			fallthrough
		case HandRank["同花顺"]:
			rst := whoIsMax(max[0:2], cur[0:2], 2)
			if rst == 2 {
				max = cur
				maxIndex = i
				isEqual = false
			} else if rst == 0 {
				isEqual = true
			}

		case HandRank["同花"]: //  同花和高牌需要按顺序比较牌点
			fallthrough
		case HandRank["高牌"]:
			rst := whoIsMax(max, cur, 2)
			if rst == 2 {
				max = cur
				maxIndex = i
				isEqual = false
			} else if rst == 0 {
				isEqual = true
			}

		default: //  先比较重复牌，再按顺序比较牌点
			curScoreTable = getHandScoreTable(cur)
			var bit1, bit2 int
			switch handRank {
			case HandRank["两对"]:
				fallthrough
			case HandRank["一对"]:
				bit1, bit2 = 2, 1
			case HandRank["3条"]:
				bit1, bit2 = 3, 1
			case HandRank["葫芦"]:
				bit1, bit2 = 3, 2
			case HandRank["4条"]:
				bit1, bit2 = 4, 1
			}

			rst := whoIsMax(maxScoreTable[bit1]+maxScoreTable[bit2],
				curScoreTable[bit1]+curScoreTable[bit2], 1)

			if rst == 2 {
				max = cur
				maxIndex = i
				isEqual = false
			} else if rst == 0 {
				isEqual = true
			}
		}
	}

	return maxIndex, isEqual
}

//  递归比较牌面
func whoIsMax(s1, s2 string, step int) int {
	if len(s1) == 0 {
		return 0
	}

	v1, v2 := FaceRank[s1[0:1]], FaceRank[s2[0:1]]
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return 2
	} else {
		return whoIsMax(s1[step:], s2[step:], step)
	}
}

func getHandScoreTable(hand string) [5]string {
	//  一共有12种牌，最小牌在map中值为2，最大为14，为了方便计算，数组长度为15
	dist := [15]int{}
	for i := 0; i < len(hand); i += 2 {
		dist[FaceRank[string(hand[i])]] += 1
	}

	var table [5]string
	for i, v := range dist {
		if v == 0 {
			continue
		}
		face := FaceName[i]
		table[v] = face + table[v]
	}
	return table
}
