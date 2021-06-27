package main

import (
	"fmt"
	"poker/src/casino"
)

func main() {
	c := casino.NewCasino()
	rst := c.Start("5d6dJcJh7d7dXn", "Js7cKdKh3c")
	rst2 := c.Start("AsKsQsJsTs", "QsQhQdQcJh")
	fmt.Println(rst, rst2)
}
