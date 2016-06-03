package main

import "fmt"

func main() {
	tree := newBTree(3)

	tree.Insert([2,3,4])

	fmt.Println(tree)
}
