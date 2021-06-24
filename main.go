package main

import (
	"poker/src/casino"
	"time"
)

func main() {
	timer := time.NewTimer(time.Nanosecond)
	casino.StartCasino("./input/match.json")

}
