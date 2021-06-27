package casino

type Casino struct {
	Dealer
	Counter
	Judge
}

func NewCasino() *Casino {
	casino := &Casino{}
	return casino
}

func (c *Casino) Start(hand1, hand2 string) int {
	hands1, hands2 := c.Deal(hand1, hand2)
	handRst1, handRst2 := c.QuickCount(hands1), c.QuickCount(hands2)

	return c.ResultJudge(handRst1, handRst2)
}

type CountRst struct {
	Hand      []string
	IsFlush   bool
	IsTongHua bool
	IsGhost   bool
	HandType  int
}
