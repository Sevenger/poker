package main

import (
	"poker/src"
	"poker/src/casino"
	"testing"
)

func TestGetHandType(t *testing.T) {
	useCases := []struct {
		name string
		hand string
		rst  string
	}{
		{"1", "AsKsQsJsTs", "皇家同花顺"},
		{"2", "AhKhQhJhTh", "皇家同花顺"},
		{"3", "2s3s4s5sAs", "同花顺"},
		{"9", "QsQhQdQcJh", "四条"},
		{"9", "2hAsAhAdAc", "四条"},
		{"10", "AsAhAcJsJc", "葫芦"},
		{"10", "2s2cAsAhAc", "葫芦"},
		{"10", "3s3cAsAhAc", "葫芦"},
		{"12", "As5s6s8sTs", "同花"},
		{"12", "ThKh9h8h3h", "同花"},
		{"12", "6s5s4s2sAs", "同花"},
		{"3", "AhKsQsJsTs", "顺子"},
		{"4", "AsKsQsJsTh", "顺子"},
		{"5", "2s3h4s5sAs", "顺子"},
		{"6", "AsKhQcJsTc", "顺子"},
		{"6", "2s3cAsAhAc", "三条"},
		{"6", "AsAhAcKs2c", "三条"},
		{"6", "AsAhQcQsTc", "两对"},
		{"6", "KsKhJcJsTc", "两对"},
		{"6", "6s5hJc6sJc", "两对"},
		{"7", "5d6dJcJh7d", "一对"},
		{"8", "Js7cKdKh3c", "一对"},
		{"8", "Js7cKd4h3c", "高牌"},
		{"8", "2s3c4d5h7c", "高牌"},
	}

	for _, use := range useCases {
		var hs casino.HandStruct
		counter := casino.Counter{}
		deal := casino.Dealer{}
		t.Run(use.name, func(t *testing.T) {
			hs = *counter.Count(deal.Sort(use.hand))
			if src.HandName[hs.HandType] != use.rst {
				t.Fatalf("牌型不对 预期牌型: %v 输出牌型: %v", use.rst, src.HandName[hs.HandType])
			}
		})
	}
}
