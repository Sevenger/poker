package casino

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
