package main

import (
	"fmt"
	"poker/src/casino"
)

func main() {
	c := casino.Casino{}
	v := c.Start("5dThTsTdXn7s3s", "Ah8c5dThTsTdXn")
	fmt.Println(v)
}
