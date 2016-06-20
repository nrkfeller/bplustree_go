package main

import "fmt"
import "github.com/Workiva/go-datastructures/slice"

func main() {
	s := slice.Int64Slice{3, 6, 1, 0, -1}
	s.Sort()
	fmt.Println(s)
}
