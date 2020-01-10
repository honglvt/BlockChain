package main

import (
	"fmt"
	"strconv"
)
import "../block"

func main() {
	fmt.Print("helloWorld")
	bc := block.NewBlockChain()
	bc.AddBlock("Send1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for index, block := range bc.Blocks {
		fmt.Printf("Block %d Data is %s \n",index, block.Data)
		fmt.Printf("Block %d Hash is %x \n",index, block.Hash)
	}

	for _,bc := range bc.Blocks {
		pow :=block.NewProofOfWork(bc)
		fmt.Printf("Pow :%s \n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
