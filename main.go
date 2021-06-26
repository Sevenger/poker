package main

import (
	"fmt"
	"poker/src/casino"
)

func main() {

	counter := casino.Counter{}
	deal := casino.Dealer{}

	hs := counter.Count(deal.Sort("AdKsQsJs9s"))
	fmt.Printf("hs %+v\n", hs)
}
