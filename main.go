package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"poker/src/casino"
	"time"
)

type Matches struct {
	MatchSlice []Match `json:"matches"`
}

type Match struct {
	Hand1  string `json:"alice"`
	Hand2  string `json:"bob"`
	Result int    `json:"result"`
}

func main() {
	c := casino.Casino{}
	var matches Matches
	var file []byte
	file, _ = ioutil.ReadFile("./input/match_result.json")
	_ = json.Unmarshal(file, &matches)
	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime1 := time.Since(startTime)

	file, _ = ioutil.ReadFile("./input/seven_cards_with_ghost.json")
	_ = json.Unmarshal(file, &matches)
	startTime = time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime2 := time.Since(startTime)

	file, _ = ioutil.ReadFile("./input/seven_cards_with_ghost.result.json")
	_ = json.Unmarshal(file, &matches)
	startTime = time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime3 := time.Since(startTime)

	fmt.Printf("FiveHand Spend:%v\nSevenHand1 Spend:%v\nSevenHand2 Spend:%v\n", endTime1, endTime2, endTime3)
}
