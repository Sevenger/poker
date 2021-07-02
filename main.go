package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "poker/src"
	"poker/src/fivehand"
	"poker/src/sevenhand"
	"time"
)

// 结果比较在main_test.go里
func main() {
	fiveHandSpend := testFiveHandsSpend("./input/five_cards.json")
	sevenHandSpend := testSevenHandsSpend("./input/seven_cards.json")
	ghostHandSpend := testGhostHandsSpend("./input/seven_cards_with_ghost.json")

	fmt.Printf("五手牌耗时：%v\n"+
		"七手牌耗时：%v\n"+
		"癞子牌耗时：%v\n",
		fiveHandSpend, sevenHandSpend, ghostHandSpend)
}

func testFiveHandsSpend(filePath string) time.Duration {
	var matches Matches
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := fivehand.Counter{}
	var rst int

	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		if rst = counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			fmt.Printf("Resuel not equal.Hand1:%s Hand2:%s Want:%d Output:%d",
				v.Hand1, v.Hand2, v.Result, rst,
			)
		}
	}
	return time.Since(startTime)
}

func testSevenHandsSpend(filePath string) time.Duration {
	var matches Matches
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := sevenhand.Counter{}
	var rst int

	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		if rst = counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			fmt.Printf("Resuel not equal.Hand1:%s Hand2:%s Want:%d Output:%d",
				v.Hand1, v.Hand2, v.Result, rst,
			)
		}
	}
	return time.Since(startTime)
}

func testGhostHandsSpend(filePath string) time.Duration {
	var matches Matches
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := sevenhand.Counter{}
	var rst int

	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		if rst = counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			fmt.Printf("Resuel not equal.Hand1:%s Hand2:%s Want:%d Output:%d",
				v.Hand1, v.Hand2, v.Result, rst,
			)
		}
	}
	return time.Since(startTime)
}
