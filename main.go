package main

import (
	"block"
	"fmt"
)

func main() {
	blockObj := block.Block{}
	newer := blockObj.CreateFirstBlock()
	fmt.Println(newer)
	for i := 0; i < 2000000; i++ {
		bd := block.BlockData{}
		newer = blockObj.New(newer, bd.New(string(i)))
	}
}
