package casino

import "strings"

// Casino 可优化：1、字符串拼接改为strings.Builder 2、使用goroutine计算牌型
type Casino struct {
	dealer  //  发牌员，处理牌型
	counter //  算派员，计算牌型
	judge   //  裁判，判断胜负
}

func (c *Casino) Start(hand1, hand2 string) int {
	hands1, hands2 := c.Deal(hand1, hand2)
	handRst1, handRst2 := c.QuickCount(hands1), c.QuickCount(hands2)
	return c.ResultJudge(handRst1, handRst2)
}

func WriteString(sb *strings.Builder, args ...string) {
	for _, v := range args {
		sb.WriteString(v)
	}
}

func Sort(hand string) string {
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
