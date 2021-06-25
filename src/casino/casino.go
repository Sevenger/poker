package casino

type Casino struct {
	Dealer
}

func (c *Casino) start(hand1, hand2 string) (result int) {
	c.Dealer.Deal(hand1, hand2)
	return 0
}
