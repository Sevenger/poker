package casino

import "time"

type Casino struct {
	reader  Reader
	counter Counter
	Timer   time.Timer
}

func NewCasino(timer time.Timer) Casino {
	return Casino{Timer: timer}
}

func (c *Casino) StartCasino(filePath string) {
	c.reader.ReadFile(filePath)

	for range c.Timer.C {
		c.Timer.Stop()
	}
}
