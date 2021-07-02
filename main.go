package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "poker/src"
	"poker/src/casino"
	"time"
)

func main() {
	startTime := time.Now()
	spendTimeFiveHand := testFiveHand("./input/match_result.json")
	spendTimeSevenHand1 := testFiveHand("./input/seven_cards_with_ghost.json")
	spendTimeSevenHand2 := testFiveHand("./input/seven_cards_with_ghost.result.json")
	spendTimeAll := time.Since(startTime)

	fmt.Printf("FiveHand Spend:%v\n"+
		"SevenHand1 Spend:%v\n"+
		"SevenHand2 Spend:%v\n"+
		"All Spend:%v",
		spendTimeFiveHand, spendTimeSevenHand1, spendTimeSevenHand2, spendTimeAll)

}

func testFiveHand(filePath string) time.Duration {
	var matches Matches
	c := casino.Casino{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}

	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		if c.Start(v.Hand1, v.Hand2) != v.Result {
			panic("Result not equal")
		}
	}
	return time.Since(startTime)
}
