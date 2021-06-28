package main

import (
	"encoding/json"
	"io/ioutil"
	"poker/src/casino"
	"strconv"
	"testing"
)

type Matches struct {
	MatchSlice []Match `json:"matches"`
}

type Match struct {
	Hand1  string `json:"alice"`
	Hand2  string `json:"bob"`
	Result int    `json:"result"`
}

func TestFiveCard(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("./input/match_result.json")
	_ = json.Unmarshal(file, &matches)

	c := casino.Casino{}
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rst := c.Start(v.Hand1, v.Hand2)

			if rst != v.Result {
				t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
					i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})
	}
}

func TestSevenCard(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("./input/seven_cards_with_ghost.json")
	_ = json.Unmarshal(file, &matches)

	c := casino.Casino{}
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rst := c.Start(v.Hand1, v.Hand2)

			if rst != v.Result {
				t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
					i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})
	}
}

func TestSevenCard2(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("./input/seven_cards_with_ghost.result.json")
	_ = json.Unmarshal(file, &matches)

	c := casino.Casino{}
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rst := c.Start(v.Hand1, v.Hand2)

			if rst != v.Result {
				t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
					i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})
	}
}
