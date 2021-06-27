package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"poker/src/casino"
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
	file, err := ioutil.ReadFile("C:\\Users\\Arcwi\\go\\src\\poker\\input\\match_result.json")
	if err != nil {
		t.Fatal("Read file error", err)
	}
	if err := json.Unmarshal(file, &matches); err != nil {
		t.Fatal("Read json error", err)
	}

	c := casino.Casino{}
	for i, v := range matches.MatchSlice {
		t.Run(fmt.Sprintf("Match %d", i), func(t *testing.T) {
			rst := c.Start(v.Hand1, v.Hand2)
			if rst != v.Result {
				t.Fatalf("ID: %d h1: %v, h2: %v. 预期结果: %d 输出结果: %d",
					i, v.Hand1, v.Hand2, v.Result, rst)
			}
		})
	}
}
