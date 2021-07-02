package fivehand

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "poker/src"
	"strconv"
	"testing"
	"time"
)

//  测试7手牌耗费时间
func TestFiveHand_Spend_Time(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/five_cards.json")
	_ = json.Unmarshal(file, &matches)

	c := Counter{}
	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime := time.Since(startTime)
	fmt.Println("Spend:", endTime)
}

//  测试7手牌判断是否正确
func TestFiveHand_Result(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/five_cards.json")
	_ = json.Unmarshal(file, &matches)

	c := Counter{}
	errCount := 0
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rst := c.Start(v.Hand1, v.Hand2)
			if rst != v.Result {
				errCount++
				t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
					i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})
	}
	fmt.Println("all err :", errCount)
}

func TestFiveHand_ByHand(t *testing.T) {
	c := Counter{}
	hand1, hand2 := "6s5h4c3s2c", "As2h3s4c5s"
	result := 1
	rst := c.Start(hand1, hand2)
	if rst != result {
		t.Fatalf("h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
			hand1, hand2, result, rst)
	}
}
