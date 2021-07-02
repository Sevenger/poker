package counterG

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
func TestCounterG_SevenHand_Spend_Time(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/seven_cards_with_ghost.json")
	_ = json.Unmarshal(file, &matches)

	c := counterG{}
	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime := time.Since(startTime)
	fmt.Println("Spend:", endTime)
}

//  测试7手牌判断是否正确
func TestCounterG_SevenHand_Result(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/seven_cards_with_ghost.json")
	_ = json.Unmarshal(file, &matches)

	c := counterG{}
	errCount := 0
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Println(v.Hand1, v.Hand2)
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

//  测试鬼牌耗费时间
func TestCounterG_GhostHand_Spend_Time(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/seven_cards_with_ghost.result.json")
	_ = json.Unmarshal(file, &matches)

	c := counterG{}
	startTime := time.Now()
	for _, v := range matches.MatchSlice {
		c.Start(v.Hand1, v.Hand2)
	}
	endTime := time.Since(startTime)
	fmt.Println("Spend:", endTime)
}

//  鬼牌判断是否正确
func TestCounterG_GhostHand_Result(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/seven_cards_with_ghost.result.json")
	_ = json.Unmarshal(file, &matches)

	c := counterG{}
	errCount := 0
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Println(v.Hand1, v.Hand2)
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

func TestCounterG_ByHand(t *testing.T) {
	c := counterG{}
	hand1, hand2 := "XnAs9c7dTs8sKc", "Jc4hXnAs9c7dTs"
	result := 0
	rst := c.Start(hand1, hand2)
	if rst != result {
		t.Fatalf("h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
			hand1, hand2, result, rst)
	}
}
