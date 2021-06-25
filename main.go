package main

import (
	"fmt"
	"poker/src/casino"
)

func main() {

	//(&casino.Counter{}).Count("Kc9c3s9sQh3c5c", "8s6cKc9c3s9sQh")

	counter := casino.Counter{}
	deal := casino.Dealer{}

	hs := counter.Count(deal.Sort("AdKsQsJsTs"))

	fmt.Println(counter.HasFlush(hs))
	fmt.Println("th", counter.IsTongHua("AsKsQsJsTs"))
	fmt.Println("th", counter.IsTongHua("QsQhQdQcJh"))

	fmt.Printf("%+v", hs)
}
