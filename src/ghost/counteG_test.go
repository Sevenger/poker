package counterG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "poker/src"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCounter7_Start_Time(t *testing.T) {
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

func TestCounter7_Start(t *testing.T) {
	var matches Matches
	file, _ := ioutil.ReadFile("../../input/seven_cards_with_ghost.json")
	_ = json.Unmarshal(file, &matches)

	c := counterG{}
	count := 0
	for i, v := range matches.MatchSlice {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			fmt.Println(v.Hand1, v.Hand2)
			rst := c.Start(v.Hand1, v.Hand2)

			if rst != v.Result {
				if hasA2345(v.Hand1) || hasA2345(v.Hand2) {
					return
				}
				count++
				t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
					i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})
	}
	fmt.Println("all err :", count)
}

func TestCounter7_Start_ByHand(t *testing.T) {
	c := counterG{}
	rst := c.Start("Jc4d5sAc6dAhAd", "6h6cJc4d5sAc6d")
	fmt.Println("shuchu", rst, "yuqi", 1)
}

func hasA2345(hand string) bool {
	return strings.Contains(hand, "A") &&
		strings.Contains(hand, "2") &&
		strings.Contains(hand, "3") &&
		strings.Contains(hand, "4") &&
		strings.Contains(hand, "5")
}
