package main

import (
	"encoding/json"
	"io/ioutil"
	. "poker/src"
	"poker/src/fivehand"
	"poker/src/sevenhand"
	"testing"
)

func TestFiveCards_Result(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("./input/five_cards.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := fivehand.Counter{}

	for i, v := range matches.MatchSlice {
		if rst := counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
				i, v.Hand1, v.Hand2, v.Result, rst)
		}
	}
}

func TestSevenCards_Result(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("./input/seven_cards.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := sevenhand.Counter{}

	for i, v := range matches.MatchSlice {
		if rst := counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
				i, v.Hand1, v.Hand2, v.Result, rst)
		}
	}
}

func TestSevenCardsWithGhost_Result(t *testing.T) {
	var matches Matches
	file, err := ioutil.ReadFile("./input/seven_cards_with_ghost.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(file, &matches); err != nil {
		panic(err)
	}
	counter := sevenhand.Counter{}

	for i, v := range matches.MatchSlice {
		if rst := counter.Start(v.Hand1, v.Hand2); rst != v.Result {
			t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
				i, v.Hand1, v.Hand2, v.Result, rst)
		}
	}
}
