package casino

import (
	"strings"
)

//dealer 用于发牌
type dealer struct{}

func (d *dealer) Deal(hand1 string, hand2 string) ([]string, []string) {
	return d.dealHand(hand1), d.dealHand(hand2)
}

func (d *dealer) dealHand(handStr string) []string {
	handStr = sort(handStr)
	var hands []string
	//  如果有7张牌，获取7张牌所有可能的牌组合
	if len(handStr) == 7*2 {
		hands = sevenToFive(handStr)
	} else {
		hands = append(hands, handStr)
	}
	return hands
}

// sevenToFive 7选5 使用排列组合给出所有组合，对于鬼牌，判断4张牌可能组成的最大值
func sevenToFive(hand string) []string {
	var hands []string

	//  有鬼牌时鬼牌必选，从剩下的6张牌选4张，一共有C4/6=15种可能
	if hand[0] == 'X' {
		hands = make([]string, 0, 15)
		combine([]string{hand[2:4], hand[4:6], hand[6:8], hand[8:10], hand[10:12], hand[12:14]}, 4, &hands)
	} else {
		//  无鬼牌时从7张牌中选5张，一共有C5/7=21种可能
		hands = make([]string, 0, 21)
		combine([]string{hand[0:2], hand[2:4], hand[4:6], hand[6:8], hand[8:10], hand[10:12], hand[12:14]}, 5, &hands)
	}

	return hands
}

func combine(arr []string, combineLen int, hands *[]string) {
	arrLen := len(arr) - combineLen
	for i := 0; i <= arrLen; i++ {
		result := make([]string, combineLen)
		result[0] = arr[i]
		doProcess(arr, result, i, 1, len(arr), combineLen, hands)
	}
}

func doProcess(a []string, result []string, rawIndex int, curIndex int, rawLen int, combineLen int, hands *[]string) {
	choice := rawLen - rawIndex + curIndex - combineLen
	var tResult []string
	for i := 0; i < choice; i++ {
		if i != 0 {
			tResult := make([]string, combineLen)
			copyArr(result, tResult)
		} else {
			tResult = result
		}
		tResult[curIndex] = a[i+1+rawIndex]

		if curIndex+1 == combineLen {
			sb := strings.Builder{}
			for _, v := range tResult {
				sb.WriteString(v)
			}
			*hands = append(*hands, sb.String())
			continue
		} else {
			doProcess(a, tResult, rawIndex+i+1, curIndex+1, rawLen, combineLen, hands)
		}
	}
}

func copyArr(rawArr []string, target []string) {
	for i := 0; i < len(rawArr); i++ {
		target[i] = rawArr[i]
	}
}
