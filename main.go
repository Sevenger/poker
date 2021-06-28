package main

import (
	"fmt"
	"poker/src/casino"
)

func main() {

	c := casino.Casino{}
	rst := c.Start("6s5h4c3s2c", "As2h3s4c5s")
	rst2 := c.Start("Ac9d6h3dTc", "2h6d8d7sJh")
	fmt.Println(rst, rst2)
}
