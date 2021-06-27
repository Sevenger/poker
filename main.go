package main

import (
	"fmt"
	"poker/src/casino"
)

func main() {
	c := casino.NewCasino()
	rst := c.Start("AsAhAcJsTc", "As2h3s4c5s")
	rst2 := c.Start("AsKsQsJsTs", "QsQhQdQcJh")
	fmt.Println(rst, rst2)

}
