package main

import (
	"fmt"
	"log"

	"github.com/Workiva/go-datastructures/bitarray"
)

func main() {
	ba := bitarray.NewBitArray(10)
	err := ba.SetBit(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ba.GetBit(5))

	err = ba.ClearBit(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ba.GetBit(5))
}
